package test

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/xiashura/anime-notification-cli/model"
)

func TestNoti(t *testing.T) {

	dat, err := ioutil.ReadFile("./notification.json")
	if err != nil {
		t.Error("can't find file error:")
	}

	var ongoings model.Ongoings
	err = json.Unmarshal(dat, &ongoings)

	if err != nil {
		t.Error("can't parse file")
	}

	for _, ongoing := range ongoings {
		re := ongoing.NextEpisodeAt.UnixNano() - time.Now().UnixNano()
		if re > 0 {
			t.Error("fail test time")
		}
	}
}

func TestWorker(t *testing.T) {

	dat, err := ioutil.ReadFile("./notification.json")
	if err != nil {
		t.Error("can't find file error:")
	}

	var ongoings model.Ongoings
	err = json.Unmarshal(dat, &ongoings)

	if err != nil {
		t.Error("can't parse file")
	}
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
				t.Logf("\nDoneTimer=[%v], Index=[%v]\nAnimeTime=[%v], AnimeName=[%v]\n",
					time.String(), index, ongoings[index].NextEpisodeAt.String(), ongoings[index].Anime.Name)
				wg.Done()
			}
		}(i)
	}
	wg.Wait()
}
