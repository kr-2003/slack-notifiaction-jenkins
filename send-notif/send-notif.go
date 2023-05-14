package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main(){
	args := os.Args[1:]
	fmt.Println(args)

	api := slack.New("xoxb-5254573433526-5261523930851-fvzxZ5zFI0A4lYMkqzhKEVUp")
	preText := "*Hello! Your Jenkins build is finished!*"
	jenkinsURL := "*Build URL: *" + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*"{
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText + "\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)
	jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)
	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		jenkinsBuildDetailsSection,
	)
	_, _, _, err := api.SendMessage(
		"C058CRTRLUQ",
		msg,
	)
	if err!=nil {
		fmt.Printf("%s\n", err)
	}else {
		fmt.Println("Message sent successfully!!")
	}
} 
