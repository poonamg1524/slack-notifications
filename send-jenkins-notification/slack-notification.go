package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"

	// "github.com/joho/godotenv"
	"encoding/json"
	"net/http"
)

type jenkinsBuild struct {
	BuildURL    string `json:"buildurl"`   
	BuildResult string `json:"buildresult"`
	BuildNumber string `json:"buildnumber"`
	JobName     string `json:"jobname"`
}

// HandlerFunction
func sendSlackMessage(w http.ResponseWriter, r *http.Request) {
	// err := godotenv.Load("E:/Web Tech/Go/slack-jenkins-go/.env")
	// if err != nil {
	// 	fmt.Println("Error loading .env:", err)
	// }

	// Get the command-line arguments passed to the script
	args := os.Args[1:]

	// Ensure you have the right number of arguments
	if len(args) < 4 {
		fmt.Println("Insufficient arguments. Expected: URL, build result, build number, job name.")
		return
	}

	fmt.Println(args)

	// // token := os.Getenv("SLACK_BOT_TOKEN") // this will take token from jenkins environment variable
	// token := os.Getenv("SLACK_BOT_TOKEN")
	// if token == "" {
	// 	fmt.Println("Slack Bot Token is missing!")
	// 	return
	// }
	// api := slack.New(token)

	// //Introductory Message Components
	// preText := "*Hello! Your Jenkins Build has finished!*"
	// jenkinsURL := "*Build URL:* " + args[0]
	// buildResult := "*" + args[1] + "*"
	// buildNumber := "*" + args[2] + "*"
	// jobName := "*" + args[3] + "*"

	// //Adding ❌ if the build fails and ✅ if build is success
	// if buildResult == "*SUCCESS*" {
	// 	buildResult = buildResult + " :white_check_mark:"
	// } else {
	// 	buildResult = buildResult + " :x:"
	// }

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Sent Slack Message</h1>")

	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		fmt.Println("Slack Bot Token is missing!")
		return
	}

	build := jenkinsBuild{}
	err0 := json.NewDecoder(r.Body).Decode(&build)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}

	api := slack.New(token)

	preText := "*Hello! Your Jenkins Build has finished!*"
	jenkinsURL := "*Build URL:* " + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	// preText := "*Hello! Your Jenkins Build has finished!*"
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

	// send msg func returns channel, timestamp
	_, _, _, errSendMsg := api.SendMessage(
		"C08Q213RFAR", // SlackChannelID
		msg,
	)

	if errSendMsg != nil {
		fmt.Printf("%s\n", errSendMsg)
	}
}

func main() {
	http.HandleFunc("/sendSlackMessage", sendSlackMessage)
	http.ListenAndServe(":8091", nil)
}
