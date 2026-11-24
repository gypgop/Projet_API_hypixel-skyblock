package models

type Perk struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Minister    bool   `json:"minister"`
}

type Candidate struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Perks []Perk `json:"perks"`
	Votes int    `json:"votes"`
}

type Mayor struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Perks []Perk `json:"perks"`
}

type Minister struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Perk Perk   `json:"perk"`
}

type Election struct {
	Year       int         `json:"year"`
	Candidates []Candidate `json:"candidates"`
}

type Current struct {
	Year       int         `json:"year"`
	Candidates []Candidate `json:"candidates"`
}

type ElectionData struct {
	Mayor    Mayor    `json:"mayor"`
	Minister Minister `json:"minister"`
	Current  Current  `json:"current"`
}

//-----------------------------Mojang UUID----------------------------------\\

type PlayerUUId struct {
	UUID string `json:"id"`
}
