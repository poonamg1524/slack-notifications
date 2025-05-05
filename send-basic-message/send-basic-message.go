package main

import (
	"fmt"
	"os"
	// "time"
	"github.com/slack-go/slack"
	"github.com/joho/godotenv"
)

//creating a global variable
// var SlackChannelID = "C08Q213RFAR"

func main() {

	err := godotenv.Load("E:/Web Tech/Go/slack-jenkins-go/.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	token := os.Getenv("SLACK_BOT_TOKEN")
	fmt.Println("Token:", os.Getenv("SLACK_BOT_TOKEN"))

	api := slack.New(token) //getting slack token
	

	channelID, timeStamp, err := api.PostMessage(
		"C08Q213RFAR", //SlackChannelID ,
		slack.MsgOptionText("Hello World", false),
	)

	if err != nil {
		fmt.Printf("%s\v", err)
		return
	}

	fmt.Printf("Message sent successfully to channel %s at %s",channelID, timeStamp)
		
	


}