package notification

import (
	"fmt"
	"sync"
	"time"

	"github.com/xiashura/anime-notification-cli/model"
	"github.com/xiashura/anime-notification-cli/sender"
)

func Worker(ongoings model.Ongoings) {
	var wg sync.WaitGroup
	noti := []time.Timer{}
	notichan := make(chan time.Time, len(ongoings))
	for i, ongoing := range ongoings {

		setTimers := time.NewTimer(time.Duration(ongoing.NextEpisodeAt.UnixNano() - time.Now().UnixNano()))

		noti = append(noti, *setTimers)
		wg.Add(2)
		go func(index int) {
			a := <-noti[index].C
			notichan <- a
			wg.Done()
		}(i)
		go func(index int) {
			select {
			case time := <-notichan:
				fmt.Printf("Done timer %v index %v animeID %v \n", time, index, ongoings[index].Anime.Name)
				sender.Post(ongoings[index])
				wg.Done()
			}
		}(i)
	}
	wg.Wait()
}
