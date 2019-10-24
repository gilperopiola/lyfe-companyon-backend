package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/mailjet/mailjet-apiv3-go"
)

func initCron() {
	sendDailyMail()

	//you have to take out 3 hours to get the real Argentina time
	gocron.Every(1).Day().At("10:00").Do(sendDailyMail)
	gocron.Every(1).Sunday().At("13:00").Do(sendWeeklyDoneMail)

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

	params.FilterTagID = 0

	allTasks, _ := task.Search(params)
	doing := []*Task{}
	for _, task := range allTasks {
		if task.Status == Doing {
			doing = append(doing, task)
		}
	}

	doneYesterday, _ := task.GetDoneAndArchivedSince(time.Now().AddDate(0, 0, -1))
	addedYesterday, _ := task.GetAddedSince(time.Now().AddDate(0, 0, -1))

	dailyElements := ""
	for i, daily := range dailies {
		color := "#e4e4e4"
		if i%2 == 0 {
			color = "#c3c3c3"
		}

		dailyElements += `<p style="background-color: ` + color + `; padding: 8px; margin: 0;">` + daily.Name + `</p>`
	}

	doingElements := ""
	for i, taskDoing := range doing {
		color := "#e4e4e4"
		if i%2 == 0 {
			color = "#c3c3c3"
		}

		doingElements += `<p style="background-color: ` + color + `; padding: 8px; margin: 0;">` + taskDoing.Name + `</p>`
	}

	doneYesterdayElements := ""
	for i, taskDone := range doneYesterday {
		color := "#e4e4e4"
		if i%2 == 0 {
			color = "#c3c3c3"
		}

		doneYesterdayElements += `<p style="background-color: ` + color + `; padding: 8px; margin: 0;">` + taskDone.Name + `</p>`
	}

	addedYesterdayElements := ""
	for i, taskAdded := range addedYesterday {
		color := "#e4e4e4"
		if i%2 == 0 {
			color = "#c3c3c3"
		}

		addedYesterdayElements += `<p style="background-color: ` + color + `; padding: 8px; margin: 0;">` + taskAdded.Name + `</p>`
	}

	html := `
	<html>
		<head>
	  		<title>Daily - ` + time.Now().Format("06/01/02") + `</title>
		</head>

		<body> 
			<p style="color: white; background-color: black; margin: 0; font-size: 14px; text-align: center; font-weight: bold;">DAILY</p>` +
		dailyElements + `
			<p style="color: white; background-color: black; margin: 0; font-size: 14px; text-align: center; font-weight: bold;">DOING</p>` +
		doingElements + `
			<p style="color: white; background-color: black; margin: 0; font-size: 14px; text-align: center; font-weight: bold;">DONE / ARCHIVED YESTERDAY</p>` +
		doneYesterdayElements + `
			<p style="color: white; background-color: black; margin: 0; font-size: 14px; text-align: center; font-weight: bold;">ADDED YESTERDAY</p>` +
		addedYesterdayElements + `

			<p style="background-color: black; margin: 0; font-size: 8px">~</p>
			<br>
		</body>
	</html>
	`

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
			Subject:  "Daily - " + time.Now().Format("06/01/02"),
			TextPart: "Que tengas buen día wachín!",
			HTMLPart: html,
			CustomID: "Daily",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Data: %+v\n", res)
}

func sendWeeklyDoneMail() {
	//todo config
	mailjetClient := mailjet.NewMailjetClient("82a01557701a0dd9f319df7c84418785", "151f081f1debe5f691ac6b0cc6caa8eb")

	task := &Task{}
	weekAgo := time.Now().Add(-24 * 7 * time.Hour)
	doneAndArchived, _ := task.GetDoneAndArchivedSince(weekAgo)

	doneElements := ""
	archivedElements := ""

	for i, task := range doneAndArchived {
		color := "#e4e4e4"
		if i%2 == 0 {
			color = "#c3c3c3"
		}

		if task.Status == Done {
			doneElements += `<p style="background-color: ` + color + `; padding: 8px; margin: 0;">` + task.Name + `</p>`
		}

		if task.Status == Done {
			archivedElements += `<p style="background-color: ` + color + `; padding: 8px; margin: 0;">` + task.Name + `</p>`
		}
	}

	html := `
	<html>
		<head>
	  		<title>Weekly - ` + weekAgo.Format("06/01/02") + ` - ` + time.Now().Format("06/01/02") + `</title>
		</head>
		<body> 
			<p style="color: white; background-color: black; margin: 0; font-size: 14px; text-align: center; font-weight: bold;">DONE</p>` +
		doneElements + `
		<p style="color: white; background-color: black; margin: 0; font-size: 14px; text-align: center; font-weight: bold;">ARCHIVED</p>` +
		archivedElements + `
			<p style="background-color: black; margin: 0; font-size: 8px">~</p>
			<br>
		</body>
	</html>
	`

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
			Subject:  `Weekly - ` + weekAgo.Format("06/01/02") + ` - ` + time.Now().Format("06/01/02"),
			TextPart: "Que tengas buena semana wachín!",
			HTMLPart: html,
			CustomID: "Weekly",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Data: %+v\n", res)
}
