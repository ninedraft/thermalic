package podcast

type Podcast struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EpisodesNum uint64 `json:"series_num"`
}
