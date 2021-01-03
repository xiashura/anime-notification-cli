package main

import (
	"flag"
	"log"

	"github.com/xiashura/anime-notification-cli/notification"
	"github.com/xiashura/anime-notification-cli/parse"
)

func main() {

	urlUser := flag.String("url", "", "")
	flag.Parse()
	if *urlUser == "" {
		log.Fatal("url can't be empty pleas set url !")
	}

	ongoings, err := parse.GetUserOngoing(*urlUser)
	if err != nil {
		log.Fatal("fatal get ongoing")
	}

	if len(ongoings) == 0 {
		log.Fatal("ongoing can't be empty")
	}

	notification.Worker(ongoings)

}
