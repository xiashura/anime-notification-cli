package parse

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/xiashura/anime-notification-cli/model"
)

const urlOngoing = "https://shikimori.one/api/calendar"

func GetUserOngoing(url string) (model.Ongoings, error) {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, urlOngoing, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	var ongoings model.Ongoings
	json.NewDecoder(res.Body).Decode(&ongoings)

	animeIdsFollow := getUserAnimesIds(url)

	result := make(model.Ongoings, 0)
	for _, ongoing := range ongoings {
		if ongoing.Anime.Status != "ongoing" {
			continue
		}
		for _, idAnime := range animeIdsFollow {
			if ongoing.Anime.ID == idAnime {
				result = append(result, ongoing)
			}
		}
	}

	return result, nil
}

func getUserAnimesIds(url string) []int {
	var idList []int
	c := colly.NewCollector()

	c.OnHTML("tr", func(h *colly.HTMLElement) {
		idString := h.Attr("data-target_id")
		idInt, _ := strconv.Atoi(idString)
		if idInt != 0 && idInt > 0 {
			idList = append(idList, idInt)
		}
	})

	c.Visit(url)
	return idList
}
