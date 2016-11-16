package reminderbot

import (
    "time"
    "fmt"

    "gopkg.in/pg.v5"
)

type Response struct {
    ID              int             `json:"id"`
    RecipientID     int64           `json:"recipient_id"`
    HabitID         int             `json:"habit_id"`
    Response        string          `sql:", null" json:"response"`
    SentAt          time.Time       `json:"sent_at"`
    RespondedAt     time.Time       `sql:", null" json:"responded_at"`
}

type Habit struct {
    ID              int               `json:"id"`
    RecipientID     int64             `json:"recipient_id"`
    Content         string            `json:"content"`
    Frequency       string            `json:"frequency"`
    Responses       []*Response       `json:"responses"`
    CreatedAt       time.Time         `json:"created_at"`
    UpdatedAt       time.Time         `sql:", null" json:"updated_at"`
}

type Schedule struct {
    ID              int             `json:"id"`
    HabitID         int             `json:"habit_id"`
    Time            int             `json:"time"`
}

func DBSetup() {
    options := &pg.Options {
        User: "salilgupta",
        Database: "norm",
        Password:"",
        Addr:"localhost:5432",
    }
    db := pg.Connect(options)
    defer db.Close()
    err := db.CreateTable(&Habit{}, nil)
    if err != nil {
        fmt.Println(err)
    }

    err = db.CreateTable(&Response{}, nil)
    if err != nil {
        fmt.Println(err)
    }
}

/*
Habit has many responses
A response has one habit
*/