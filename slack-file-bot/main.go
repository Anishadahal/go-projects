package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3879019227168-3879459446579-3lkT5xm6XdjM31QZNJ8Rdnih")
	os.Setenv("CHANNEL_ID", "C03RV0K75DE")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channel := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"./perfectSquare.py", "abc.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channel,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s, URL: %s \n", file.Name, file.URL)
	}
}
