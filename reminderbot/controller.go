package reminderbot

import (
    "log"
    "encoding/json"
    "time"

    "github.com/abhinavdahiya/go-messenger-bot"
    "gopkg.in/pg.v5"
)

type PostBackPayload struct {
    ResponseID int
    UserResponse string
}

var Options = &pg.Options {
    User: "salilgupta",
    Database: "norm",
    Password:"",
    Addr:"localhost:5432",
}


func getHabits() []Habit {
    db := pg.Connect(Options)
    defer db.Close()
    var habits []Habit

    err := db.Model(&habits).Select()
    if err != nil {
        panic(err.Error())
    }

    return habits
}

func createResponse(habit Habit) int64 {
    db := pg.Connect(Options)
    defer db.Close()

    response := Response{
                    RecipientID: habit.RecipientID,
                    HabitID: habit.ID,
                    Response: "No Response",
                    SentAt: time.Now(),
                }

    id, err := db.Insert(&response)
    if err != nil {
        panic(err)
    }
    return id
}

func generateTemplate(responseID int64) mbotapi.GenericTemplate {
    generic := mbotapi.NewGenericTemplate()
    element := mbotapi.Element{
                        Title: "Reminder",
                        Subtitle: habit.Content,
                    }

    yesPayload, _ := json.Marshal(PostBackPayload{ResponseID: responseID, UserResponse: "Yes"})
    noPayload, _ := json.Marshal(PostBackPayload{ResponseID: responseID, UserResponse: "No"})

    yesButton := mbotapi.NewPostbackButton("Yes", string(yesPayload))
    noButton := mbotapi.NewPostbackButton("No", string(noPayload))

    element.AddButton(yesButton, noButton)
    generic.AddElement(element)
    return generic
}

func SendReminders(bot *mbotapi.BotAPI) {
    habits := getHabits()
    for _, habit := range habits {
        log.Printf("[%#v] ", habit)
        go func (habit Habit) {
            responseID := createResponse(habit)
            generic := generateTemplate(responseID)
            user := mbotapi.NewUserFromID(habit.RecipientID)
            bot.Send(user, generic, mbotapi.RegularNotif)
        }(habit)
    }
}

func SaveUserResponse(callback mbotapi.Callback) string {
    db := pg.Connect(Options)
    defer db.Close()

    var payload PostBackPayload
    json.Unmarshal([]byte(callback.Postback.Payload), &payload)

    response := Response{
                    ID: payload.ResponseID,
                    Response:   payload.UserResponse,
                    RespondedAt: time.Unix(callback.Timestamp, 0),
                }

    err := db.Update(&response)

    if err != nil {
        log.Printf("[%s]", err)
    }

    log.Printf("[%#v]", payload)
    if response.Response == "No" {
        return "Oh poop ... "
    }
    return "Yea! You go cowboy!"
}