package trakt

import "time"

type Episode struct {
	Ids struct {
		Imdb   string `json:"imdb"`
		Tmdb   int    `json:"tmdb"`
		Trakt  int    `json:"trakt"`
		Tvdb   int    `json:"tvdb"`
		Tvrage int    `json:"tvrage"`
	} `json:"ids"`

	Season                int       `json:"season"`
	Number                int       `json:"number"`
	Title                 string    `json:"title"`
	Overview              string    `json:"overview"`
	FirstAired            time.Time `json:"first_aired"`
	AvailableTranslations []string  `json:"available_translations"`

	Images struct {
		Screenshot Images `json:"screenshot"`
	} `json:"images"`

	Rating float64 `json:"rating"`
	Votes  int     `json:"votes"`

	UpdatedAt string `json:"updated_at"`
}
