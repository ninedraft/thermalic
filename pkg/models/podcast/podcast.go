package podcast

type Podcast struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SeriesNum   uint64 `json:"series_num"`
}
