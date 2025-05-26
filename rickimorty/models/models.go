package models

type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

//

type ResponseCharacters struct {
	Info    Info                `json:"info"`
	Results []ResultsCharacters `json:"results"`
}

type ResultsCharacters struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Status   string         `json:"status"`
	Species  string         `json:"species"`
	Type     string         `json:"type"`
	Gender   string         `json:"gender"`
	Origin   OriginLocation `json:"origin"`
	Location OriginLocation `json:"location"`
	Image    string         `json:"image"`
	Episode  []string       `json:"episode"`
}

type OriginLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//

type ResponseEpisodes struct {
	Info    Info              `json:"info"`
	Results []ResultsEpisodes `json:"results"`
}

type ResultsEpisodes struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	AirDate    string   `json:"air_date"`
	Episode    string   `json:"episode"`
	Characters []string `json:"characters"`
}

//

type CharacterName struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type Character struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	Species      string `json:"species"`
	Type         string `json:"type"`
	Gender       string `json:"gender"`
	OriginName   string `json:"originName"`
	LocationName string `json:"locationName"`
	ImageUrl     string `json:"imageUrl"`
}

type Episode struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	EpisodeCode string `json:"episodeCode"`
}

// Response
type StandardResponse struct {
	Message string `json:"message,omitempty"`
}
