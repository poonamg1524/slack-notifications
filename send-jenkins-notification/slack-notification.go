package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	// Get the Slack token from the Jenkins environment variables
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		fmt.Println("Slack Bot Token is missing!")
		return
	}
	api := slack.New(token)

	// Get the command-line arguments passed to the script
	args := os.Args[1:]
	fmt.Println(args)

	// Ensure you have the right number of arguments
	if len(args) < 4 {
		fmt.Println("Insufficient arguments. Expected: URL, build result, build number, job name.")
		return
	}

	// Introductory message components
	preText := "*Hello! Your Jenkins Build has finished!*"
	jenkinsURL := "*Build URL:* " + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	// Modify build result to include success or failure emoji
	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	// Slack message components
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

	// Send the message to Slack channel
	_, _, _, errSendMsg := api.SendMessage(
		"C08Q213RFAR", // SlackChannelID,
		msg,
	)

	if errSendMsg != nil {
		fmt.Printf("%s\n", errSendMsg)
	}
}
