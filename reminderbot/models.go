package reminderbot

import (
    "time"
)

type Response struct {
    ID              int64           `json:"id"`
    RecipientID     int64           `json:"recipient_id"`
    HabitID         int64           `json:"habit_id"`
    Response        string          `json:"response"`
    SentAt          time.Time       `json:"sent_at"`
    RespondedAt     time.Time       `json:"responded_at"`
}

type Habit struct {
    ID              int64             `json:"id"`
    RecipientID     int64             `json:"recipient_id"`
    Content         string            `json:"content"`
    Frequency       string            `json:"frequency"`
    TimeOfDay       string            // not sure what to do with this one
    Responses       []*Response       `json:"responses"`
    CreatedAt       time.Time         `json:"created_at"`
    UpdatedAt       time.Time         `json:"updated_at"`
}


/*
Habit has many responses
A response has one habit
*/