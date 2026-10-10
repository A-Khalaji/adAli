
package cache

import (
	"encoding/json"
	"fmt"

	"adAli/internal/database"
	"adAli/internal/models"
)

func GetScheduleChunk(id string) (*models.ScheduleChunk, error) {
	data, err := database.Redis.Get(
		database.Ctx,
		fmt.Sprintf("schedule:chunk:%s", id),
	).Bytes()
	if err != nil {
		return nil, err
	}

	var chunk models.ScheduleChunk
	if err := json.Unmarshal(data, &chunk); err != nil {
		return nil, err
	}

	return &chunk, nil
}

func SetScheduleChunk(chunk models.ScheduleChunk) error {
	data, err := json.Marshal(chunk)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("schedule:chunk:%s", chunk.ID)

	if err := database.Redis.Set(database.Ctx, key, data, 0).Err(); err != nil {
		return err
	}

	return database.Redis.SAdd(database.Ctx, "schedule:chunks", chunk.ID).Err()
}

func SetScheduleChunkIfNotExists(chunk models.ScheduleChunk) error {
	data, err := json.Marshal(chunk)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("schedule:chunk:%s", chunk.ID)

	if err := database.Redis.SetNX(database.Ctx, key, data, 0).Err(); err != nil {
		return err
	}

	return database.Redis.SAdd(database.Ctx, "schedule:chunks", chunk.ID).Err()
}