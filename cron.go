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
	sendWeeklyMail()

	//you have to take out 3 hours to get the real Argentina time
	gocron.Every(1).Day().At("10:00").Do(sendDailyMail)
	gocron.Every(1).Sunday().At("13:00").Do(sendWeeklyMail)

	<-gocron.Start()
}

func createMailRow(text string, backgroundColor string, foregroundColor string, isTitle bool) string {
	if isTitle {
		return `<p style='font-size: 32px; font-weight: bold; text-align: center; background-color: ` + backgroundColor + `; color: ` + foregroundColor + `; padding: 12px; margin: 0;'>` + text + `</p>`
	}

	return `<p style='font-size: 16px; background-color: ` + backgroundColor + `; color: ` + foregroundColor + `; padding: 8px; margin: 0;'>` + text + `</p>`
}

func getRowColor(i int) string {
	if i%2 == 0 {
		return "#292929"
	}
	return "black"
}

func sendDailyMail() {

	/*------------------------*/
	/* PART 1: INFO RETRIEVAL */
	/*------------------------*/

	// get dailies

	task := &Task{}
	params := &SearchParameters{
		FilterTagID:      1,
		FilterImportance: 1,
		ShowPrivate:      true,
		Limit:            1000,
		Offset:           0,
	}

	periodicalsExpiringToday, _ := connect.GetPeriodicalsExpiringToday()
	periodicalsDoneYesterday, _ := connect.GetPeriodicalsDoneYesterday()

	dailies, _ := task.Search(params)

	// get doing

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

	//get entities

	problemEntities, _ := connect.GetEntitiesOfKind("Problems")
	axiomEntities, _ := connect.GetEntitiesOfKind("Supuestos")
	errorEntities, _ := connect.GetEntitiesOfKind("Errores")
	knowledgeEntities, _ := connect.GetEntitiesOfKind("Knowledge")

	//get money & transactions

	money, _ := connect.GetMoneyAmount()
	transactions, _ := connect.GetDayTransactions()

	/*------------------------*/
	/*  PART 2: PREPARATION   */
	/*------------------------*/

	// tasks

	dailyElements := ""
	for i, daily := range dailies {
		elapsed := frutils.ToString(frutils.GetDaysBetween(time.Now(), daily.DateCreated))
		dailyElements += createMailRow(daily.Name+" > "+elapsed, getRowColor(i), "white", false)
	}

	doingElements := ""
	for i, taskDoing := range doing {
		elapsed := frutils.ToString(frutils.GetDaysBetween(time.Now(), taskDoing.DateCreated))
		doingElements += createMailRow(taskDoing.Name+" > "+elapsed, getRowColor(i), "white", false)
	}

	doneYesterdayElements := ""
	for i, taskDone := range doneYesterday {
		elapsed := frutils.ToString(frutils.GetDaysBetween(time.Now(), taskDone.DateCreated))
		doneYesterdayElements += createMailRow(taskDone.Name+" > "+elapsed, getRowColor(i), "white", false)
	}

	addedYesterdayElements := ""
	for i, taskAdded := range addedYesterday {
		elapsed := frutils.ToString(frutils.GetDaysBetween(time.Now(), taskAdded.DateCreated))
		addedYesterdayElements += createMailRow(taskAdded.Name+" > "+elapsed, getRowColor(i), "white", false)
	}

	// periodicals

	periodicalsTodayElements := ""
	for i, periodical := range periodicalsExpiringToday {
		periodicalsTodayElements += createMailRow(periodical.Name, getRowColor(i), "white", false)
	}

	periodicalsYesterdayElements := ""
	for i, periodical := range periodicalsDoneYesterday {
		periodicalsYesterdayElements += createMailRow(periodical.Name, getRowColor(i), "white", false)
	}

	// entities

	entitiesElements := ""

	problemIndex := frutils.GetRandomInt(0, len(problemEntities)-1)
	entitiesElements += createMailRow(problemEntities[problemIndex].Name, getRowColor(0), "white", false)

	axiomIndex := frutils.GetRandomInt(0, len(axiomEntities)-1)
	entitiesElements += createMailRow(axiomEntities[axiomIndex].Name, getRowColor(1), "white", false)

	errorIndex := frutils.GetRandomInt(0, len(errorEntities)-1)
	entitiesElements += createMailRow(errorEntities[errorIndex].Name, getRowColor(2), "white", false)

	knowledgeIndex := frutils.GetRandomInt(0, len(knowledgeEntities)-1)
	entitiesElements += createMailRow(knowledgeEntities[knowledgeIndex].Name, getRowColor(3), "white", false)

	// transactions

	transactionElements := ""
	for i, transaction := range transactions {
		transactionElements += createMailRow(transaction.Name+" > "+frutils.ToString(transaction.Amount), getRowColor(i), "white", false)
	}

	/*------------------------*/
	/*  PART 3: MAIL SENDING  */
	/*------------------------*/

	subject := "Daily - " + time.Now().Weekday().String() + " " + time.Now().Format("06/01/02")

	html := `
	<html>
		<body> ` +

		createMailRow("DAILY - "+frutils.ToString(money)+"$", "#511480", "white", true) + dailyElements +
		createMailRow("DOING", "#511480", "white", true) + doingElements +
		createMailRow("DONE / ARCHIVED YESTERDAY", "#511480", "white", true) + doneYesterdayElements +
		createMailRow("ADDED YESTERDAY", "#511480", "white", true) + addedYesterdayElements +
		createMailRow("PERIODICALS TO DO TODAY", "#b9c217", "white", true) + periodicalsTodayElements +
		createMailRow("PERIODICALS DONE YESTERDAY", "#b9c217", "white", true) + periodicalsYesterdayElements +
		createMailRow("TRANSACTIONS MADE YESTERDAY", "#b9c217", "white", true) + transactionElements +
		createMailRow("PROB / SUPP / ERR / KNOW", "#511480", "white", true) + entitiesElements + `

			<p style='background-color: black; margin: 0; font-size: 8px'>~</p>
			<br>
		</body>
	</html>`

	status, response := connect.SendMail("ferra.main@gmail.com", subject, "", strings.Replace(strings.Replace(html, "\n", "", -1), "\t", "", -1))
	log.Println(frutils.ToString(status) + ": " + response)
}

func sendWeeklyMail() {

	/*------------------------*/
	/* PART 1: INFO RETRIEVAL */
	/*------------------------*/

	// weeklies

	task := &Task{}
	params := &SearchParameters{
		FilterTagID:      2,
		FilterImportance: 1,
		ShowPrivate:      true,
		Limit:            1000,
		Offset:           0,
	}

	weeklies, _ := task.Search(params)

	// done & archived last week

	weekAgo := time.Now().Add(-24 * 7 * time.Hour)
	doneAndArchived, _ := task.GetDoneAndArchivedSince(weekAgo)

	// transactions

	transactions, _ := connect.GetWeekTransactions()

	/*------------------------*/
	/*  PART 2: PREPARATION   */
	/*------------------------*/

	// weekly

	weeklyElements := ""

	for i, weeklyTask := range weeklies {
		weeklyElements += createMailRow(weeklyTask.Name, getRowColor(i), "white", false)
	}

	// done & archived last week

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

	// transactions

	transactionElements := ""
	for i, transaction := range transactions {
		transactionElements += createMailRow(transaction.Name+" > "+frutils.ToString(transaction.Amount), getRowColor(i), "white", false)
	}

	/*------------------------*/
	/*  PART 3: MAIL SENDING  */
	/*------------------------*/

	subject := "Weekly - " + time.Now().Weekday().String() + " " + time.Now().Format("06/01/02")

	html := `
	<html>
		<body> ` +

		createMailRow("TO DO THIS WEEK", "#511480", "white", true) + weeklyElements +
		createMailRow("DONE THIS WEEK", "#b9c217", "white", true) + doneElements +
		createMailRow("ARCHIVED THIS WEEK", "#b9c217", "white", true) + archivedElements +
		createMailRow("TRANSACTED THIS WEEK", "#511480", "white", true) + transactionElements + `

		<p style='background-color: black; margin: 0; font-size: 8px'>~</p>
		<br>
		</body>
	</html>`

	status, response := connect.SendMail("ferra.main@gmail.com", subject, "", strings.Replace(strings.Replace(html, "\n", "", -1), "\t", "", -1))
	log.Println(frutils.ToString(status) + ": " + response)
}
