package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("E:/Web Tech/Go/slack-jenkins-go/.env")
	if err != nil {
		fmt.Println("Error loading .env:", err)
	}
	args := os.Args[1:]
	fmt.Println(args)

	token := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(token)

	//Introductory Message Components
	preText := "*Hello! Your Jenkins Build has finished!*"
	jenkinsURL := "*Build URL:* " + args[0] 
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	//Adding ❌ if the build fails and ✅ if build is success

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL

	preTextField := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)

	jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)
	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		jenkinsBuildDetailsSection,
	)

	//send msg func returns channel,timestamp
	_, _, _, errSendMsg := api.SendMessage(
		"C08Q213RFAR", //SlackChannelID ,
		msg,
	)

	if errSendMsg != nil {
		fmt.Printf("%s\n", errSendMsg)
	}
}
