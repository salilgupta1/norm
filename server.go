package main

import (
    "net/http"
    "log"
    "os"

    "gopkg.in/robfig/cron.v2"
    _ "github.com/joho/godotenv/autoload"
    "github.com/abhinavdahiya/go-messenger-bot"
    "github.com/salilgupta1/reminder-bot/reminderbot"
)

var token = os.Getenv("token")
var secret = os.Getenv("secret")
var vtoken = os.Getenv("vtoken")

var Bot = &mbotapi.BotAPI{
        Token:       token,
        AppSecret:   secret,
        VerifyToken: vtoken,
        Debug:       true,
        Client:      &http.Client{},
    }

func main() {

    c := cron.New()
    // Europe/Iceland is same as GMT (which is 0 hours off of UTC)
    // Run job every hour in GMT time
    c.AddFunc("TZ=Europe/Iceland 0 0 * * * *", func() { reminderbot.SendReminders(Bot) })
    c.Start()

    callbacks, mux := Bot.SetWebhook("/webhook")
    go http.ListenAndServe(":3000", mux)

    for callback := range callbacks {
        log.Printf("[%#v] ", callback)
        if callback.IsPostback() {
            log.Printf("Callback is postback")

            botResponse := reminderbot.SaveUserResponse(callback)
            msg := mbotapi.NewMessage(botResponse)
            Bot.Send(callback.Sender, msg, mbotapi.RegularNotif)
        }
    }
}
