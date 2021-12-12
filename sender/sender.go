package sender

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/xiashura/anime-notification-cli/model"
)

//This func use cli utils notify-send
func Post(ongoing model.Ongoing) {

	urlImage := "https://shikimori.one" + ongoing.Anime.Image.Original

	pathImage := getImageAndSave(urlImage, ongoing.Anime.ID)

	cmd := exec.Command("notify-send", "-i", pathImage, fmt.Sprintf(`%v Вышла серия номер %d`, ongoing.Anime.Russian, ongoing.NextEpisode))

	cmd.Run()
}

func getImageAndSave(url string, id int) string {
	dirname, _ := os.UserHomeDir()

	img, _ := os.Create(dirname + "/.cache/" + fmt.Sprint(id) + ".png")
	defer img.Close()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	b, err := io.Copy(img, resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)

	return dirname + "/.cache/" + fmt.Sprint(id) + ".png"
}
