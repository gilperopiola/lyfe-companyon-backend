package main

import (
	"log"
	"strings"
	"time"

	"github.com/gilperopiola/frutils"
	connect "github.com/gilperopiola/gilperopiola-ms-connect"
	"github.com/jasonlvhit/gocron"
)

func initCron() {
	sendDailyMail()

	//you have to take out 3 hours to get the real Argentina time
	gocron.Every(1).Day().At("10:00").Do(sendDailyMail)
	gocron.Every(1).Sunday().At("13:00").Do(sendWeeklyDoneMail)

	<-gocron.Start()
}

func createMailRow(text string, backgroundColor string, foregroundColor string, isTitle bool) string {
	if isTitle {
		return `<p style='font-size: 32px; font-weight: bold; text-align: center; background-color: ` + backgroundColor + `; color: ` + foregroundColor + `; padding: 12px; margin: 0;'>` + text + `</p>`
	}

	return `<p style='font-size: 16px; background-color: ` + backgroundColor + `; color: ` + foregroundColor + `; padding: 8px; margin: 0;'>` + text + `</p>`
}

func sendDailyMail() {
	//get dailies
	task := &Task{}
	params := &SearchParameters{
		FilterTagID:      1,
		FilterImportance: 1,
		ShowPrivate:      true,
		Limit:            1000,
		Offset:           0,
	}

	dailies, _ := task.Search(params)

	//get doing
	params.FilterTagID = 0
	allTasks, _ := task.Search(params)
	doing := []*Task{}
	for _, task := range allTasks {
		if task.Status == Doing {
			doing = append(doing, task)
		}
	}

	//get done and added
	doneYesterday, _ := task.GetDoneAndArchivedSince(time.Now().AddDate(0, 0, -1))
	addedYesterday, _ := task.GetAddedSince(time.Now().AddDate(0, 0, -1))

	//prepare elements
	dailyElements := ""
	for _, daily := range dailies {
		dailyElements += createMailRow(daily.Name, "black", "white", false)
	}

	doingElements := ""
	for _, taskDoing := range doing {
		doingElements += createMailRow(taskDoing.Name, "black", "white", false)
	}

	doneYesterdayElements := ""
	for _, taskDone := range doneYesterday {
		doneYesterdayElements += createMailRow(taskDone.Name, "black", "white", false)
	}

	addedYesterdayElements := ""
	for _, taskAdded := range addedYesterday {
		addedYesterdayElements += createMailRow(taskAdded.Name, "black", "white", false)
	}

	//send mail
	subject := "Daily - " + time.Now().Weekday().String() + " " + time.Now().Format("06/01/02")

	html := `
	<html>
		<body> ` +

		createMailRow("DAILY", "#511480", "white", true) + dailyElements +
		createMailRow("DOING", "#511480", "white", true) + doingElements +
		createMailRow("DONE / ARCHIVED YESTERDAY", "#511480", "white", true) + doneYesterdayElements +
		createMailRow("ADDED YESTERDAY", "#511480", "white", true) + addedYesterdayElements + `

			<p style='background-color: black; margin: 0; font-size: 8px'>~</p>
			<br>
		</body>
	</html>`

	status, response := connect.SendMail("ferra.main@gmail.com", subject, "", strings.Replace(strings.Replace(html, "\n", "", -1), "\t", "", -1))
	log.Println(frutils.ToString(status) + ": " + response)
}

func sendWeeklyDoneMail() {
	task := &Task{}
	weekAgo := time.Now().Add(-24 * 7 * time.Hour)
	doneAndArchived, _ := task.GetDoneAndArchivedSince(weekAgo)

	doneElements := ""
	archivedElements := ""

	for _, task := range doneAndArchived {
		if task.Status == Done {
			doneElements += createMailRow(task.Name, "gray", "black", false)
		}

		if task.Status == Done {
			archivedElements += createMailRow(task.Name, "gray", "black", false)
		}
	}

	//send mail
	subject := "Weekly - " + time.Now().Weekday().String() + " " + time.Now().Format("06/01/02")

	html := `
	<html>
		<body> ` +

		createMailRow("DONE", "black", "white", true) + doneElements +
		createMailRow("ARCHIVED", "black", "white", true) + archivedElements + `

		<p style="background-color: black; margin: 0; font-size: 8px">~</p>
		<br>
		</body>
	</html>`

	status, response := connect.SendMail("ferra.main@gmail.com", subject, "", strings.Replace(strings.Replace(html, "\n", "", -1), "\t", "", -1))
	log.Println(frutils.ToString(status) + ": " + response)
}
