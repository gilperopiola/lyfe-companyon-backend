package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/mailjet/mailjet-apiv3-go"
)

func initCron() {
	gocron.Every(1).Day().At("07:00").Do(sendDailyMail)
	gocron.Every(5).Seconds().Do(sendDailyMail)

	<-gocron.Start()
}

func sendDailyMail() {
	//todo config
	mailjetClient := mailjet.NewMailjetClient("82a01557701a0dd9f319df7c84418785", "151f081f1debe5f691ac6b0cc6caa8eb")

	task := &Task{}
	params := &SearchParameters{
		Filter:           "",
		FilterTagID:      1,
		FilterImportance: 1,
		ShowPrivate:      true,
		Limit:            1000,
		Offset:           0,
	}

	dailies, _ := task.Search(params)

	log.Println(len(dailies))

	dailyElements := ""
	for i, daily := range dailies {
		color := "#e4e4e4"
		if i%2 == 0 {
			color = "#c3c3c3"
		}

		dailyElements += `<p style="background-color: ` + color + `; padding: 8px; margin: 0;">` + daily.Name + `</p>`
	}

	html := `<html>
	<head>
	  <title>Dailies for ` + time.Now().Format("06/01/02") + `</title>
	</head>
	<body>` + dailyElements + `
	  <p style="background-color: black; margin: 0; font-size: 8px">~</p>
	  <br>
	</body>
	</html>
	`

	log.Println(dailyElements)

	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "lucfran2005@hotmail.com",
				Name:  "Franco",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "ferra.main@gmail.com",
					Name:  "Franco",
				},
			},
			Subject:  "Keonda perro",
			TextPart: "Acá tamo",
			HTMLPart: html,
			CustomID: "Daily",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data: %+v\n", res)
}
