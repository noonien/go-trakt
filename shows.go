package trakt

import "time"

type Show struct {
	Ids struct {
		Imdb   string `json:"imdb"`
		Slug   string `json:"slug"`
		Tmdb   int    `json:"tmdb"`
		Trakt  int    `json:"trakt"`
		Tvdb   int    `json:"tvdb"`
		Tvrage int    `json:"tvrage"`
	} `json:"ids"`

	Status                string   `json:"status"`
	Title                 string   `json:"title"`
	Year                  int      `json:"year"`
	Genres                []string `json:"genres"`
	Overview              string   `json:"overview"`
	Country               string   `json:"country"`
	Language              string   `json:"language"`
	Trailer               string   `json:"trailer"`
	Network               string   `json:"network"`
	Homepage              string   `json:"homepage"`
	AvailableTranslations []string `json:"available_translations"`
	Certification         string   `json:"certification"`

	FirstAired    time.Time `json:"first_aired"`
	AiredEpisodes int       `json:"aired_episodes"`
	Runtime       int       `json:"runtime"`
	Airs          struct {
		Day      string `json:"day"`
		Time     string `json:"time"`
		Timezone string `json:"timezone"`
	} `json:"airs"`

	Images struct {
		Banner   Images `json:"banner"`
		Clearart Images `json:"clearart"`
		Fanart   Images `json:"fanart"`
		Logo     Images `json:"logo"`
		Poster   Images `json:"poster"`
		Thumb    Images `json:"thumb"`
	} `json:"images"`

	Rating float64 `json:"rating"`
	Votes  int     `json:"votes"`

	UpdatedAt time.Time `json:"updated_at"`
}
