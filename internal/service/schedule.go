package service

import (
	"errors"
	"fmt"
	"time"

	"adAli/internal/cache"
	"adAli/internal/database"
	"adAli/internal/models"

	"github.com/redis/go-redis/v9"
)

const chunkDuration = 5 * time.Minute

func CreateSchedule(startTime, endTime time.Time) []models.ScheduleChunk {
	var chunks []models.ScheduleChunk

	for start := startTime; start.Before(endTime); start = start.Add(chunkDuration) {
		chunkEnd := start.Add(chunkDuration)

		if chunkEnd.After(endTime) {
			chunkEnd = endTime
		}

		chunks = append(chunks, models.ScheduleChunk{
			ID:        start.Format(time.RFC3339),
			StartTime: start,
			EndTime:   chunkEnd,
			Programs:  []models.ScheduledProgram{},
		})
	}

	return chunks
}

func InitializeSchedule(startTime, endTime time.Time) error {
	chunks := CreateSchedule(startTime, endTime)

	for _, chunk := range chunks {
		if err := cache.SetScheduleChunkIfNotExists(chunk); err != nil {
			return err
		}
	}

	return nil
}

func InitializeScheduleFor100Days() error {
	startTime := time.Now().Truncate(chunkDuration)
	endTime := startTime.Add(100 * 24 * time.Hour)

	return InitializeSchedule(startTime, endTime)
}

func CalculateProgramLimit(program models.Program, totalChunks int) int64 {
	if program.BidPrice <= 0 || totalChunks <= 0 {
		return 0
	}

	totalImpressions := program.Budget / program.BidPrice
	return int64(totalImpressions / float64(totalChunks))
}

func AddProgramToChunk(
	chunk *models.ScheduleChunk,
	program models.Program,
	limit int64,
) {
	for i := range chunk.Programs {
		if chunk.Programs[i].ProgramID == program.ID {
			chunk.Programs[i].Limit = limit
			return
		}
	}

	chunk.Programs = append(chunk.Programs, models.ScheduledProgram{
		ProgramID: program.ID,
		Limit:     limit,
		Used:      0,
	})
}

func AddProgramToSchedule(program models.Program) error {
	if program.StartDate == nil || program.EndDate == nil {
		return errors.New("program start_date and end_date are required")
	}

	if !program.EndDate.After(*program.StartDate) {
		return errors.New("end_date must be after start_date")
	}

	if program.Budget <= 0 || program.BidPrice <= 0 {
		return errors.New("budget and bid_price must be greater than zero")
	}

	chunkIDs, err := database.Redis.SMembers(
		database.Ctx,
		"schedule:chunks",
	).Result()
	if err != nil {
		return err
	}

	var matchingChunks []*models.ScheduleChunk

	for _, chunkID := range chunkIDs {
		chunk, err := cache.GetScheduleChunk(chunkID)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				if err := database.Redis.SRem(
					database.Ctx, "schedule:chunks", chunkID,
				).Err(); err != nil {
					return err
				}
				continue
			}
			return err
		}

		if chunk.EndTime.After(*program.StartDate) &&
			chunk.StartTime.Before(*program.EndDate) {
			matchingChunks = append(matchingChunks, chunk)
		}
	}

	if len(matchingChunks) == 0 {
		return fmt.Errorf(
			"no schedule chunks overlap program %d; check its dates and initialized schedule range",
			program.ID,
		)
	}

	limit := CalculateProgramLimit(program, len(matchingChunks))
	if limit <= 0 {
		return errors.New("calculated per-chunk impression limit is zero")
	}

	for _, chunk := range matchingChunks {
		AddProgramToChunk(chunk, program, limit)

		if err := cache.SetScheduleChunk(*chunk); err != nil {
			return err
		}
	}

	return nil
}
