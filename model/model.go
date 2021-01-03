package model

import "time"

type Ongoings []Ongoing

type Ongoing struct {
	NextEpisode   int       `json:"next_episode"`
	NextEpisodeAt time.Time `json:"next_episode_at"`
	Duration      int       `json:"duration"`
	Anime         struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Russian string `json:"russian"`
		Image   struct {
			Original string `json:"original"`
			Preview  string `json:"preview"`
			X96      string `json:"x96"`
			X48      string `json:"x48"`
		} `json:"image"`
		URL           string      `json:"url"`
		Kind          string      `json:"kind"`
		Score         string      `json:"score"`
		Status        string      `json:"status"`
		Episodes      int         `json:"episodes"`
		EpisodesAired int         `json:"episodes_aired"`
		AiredOn       string      `json:"aired_on"`
		ReleasedOn    interface{} `json:"released_on"`
	} `json:"anime"`
}

//type User struct {}
