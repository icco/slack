package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/slack-go/slack"
	"go.uber.org/zap"
)

func main() {
	token, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		fmt.Println("Missing SLACK_TOKEN in environment")
		os.Exit(1)
	}
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	api := slack.New(
		token,
		slack.OptionDebug(true),
		slack.OptionLog(sugar),
	)
}
