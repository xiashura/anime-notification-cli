package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/xiashura/anime-notification-cli/model"
	"github.com/xiashura/anime-notification-cli/sender"
)

func TestSenderNoti(t *testing.T) {

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
		sender.Post(ongoing)
	}
}
