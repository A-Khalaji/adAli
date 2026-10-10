package models

import "time"

type ScheduleChunk struct {
	ID        string             `json:"id"`
	StartTime time.Time          `json:"start_time"`
	EndTime   time.Time          `json:"end_time"`
	Programs  []ScheduledProgram `json:"programs"`
}

type ScheduledProgram struct {
	ProgramID uint  `json:"program_id"`
	Limit     int64 `json:"limit"`
	Used      int64 `json:"used"`
}