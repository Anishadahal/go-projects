package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println() //print empty line
	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3879019227168-3852351358997-VsRuYXlz7sSX37sDS4LrHIaW")                                         //OAUTH TOKEN
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03R58LM0KU-3855225322067-8c5ff2430e861195d7e8e38e1bbba6b77f3d5a30a7b7d1f9bf9596802bf59c89") //SOCKET TOKEN

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	fmt.Println(bot)

	go printCommandEvents(bot.CommandEvents()) //GO ROUTINE

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB calculator",
		Examples:    []string{"My YOB is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != err {
				log.Fatal(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)

		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
