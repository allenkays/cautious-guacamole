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

		fmt.Println("")
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3836070680288-3814928018628-wDviqNkQ14RP9QD9wVnJi57F")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03PPQ8LER4-3836210009888-1851aae53a3d5db541e41e82b28a840f3da10ba3e169ec6123f9c594f043ce3a")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Example:     "My YOB is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
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
