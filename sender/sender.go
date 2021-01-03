package sender

import (
	"fmt"
	"io"
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

	img, _ := os.Create("../tmp/" + fmt.Sprint(id) + ".png")
	defer img.Close()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// toPng([]byte(resp.Body.()))
	b, _ := io.Copy(img, resp.Body)
	fmt.Println(b)
	// pwd, _ := os.Getwd()
	//test
	pwd := "/home/xiashura/Projects/go/src/github.com/xiashura/anime-notification-cli"
	return pwd + "/tmp/" + fmt.Sprint(id) + ".png"
}
