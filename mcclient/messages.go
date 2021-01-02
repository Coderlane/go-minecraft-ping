package mcclient

type User struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type StatusVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type StatusPlayers struct {
	Max    int    `json:"max"`
	Online int    `json:"online"`
	Sample []User `json:"sample"`
}

type StatusDescription struct {
	Text string `json:"text"`
}

type StatusResponse struct {
	Version     StatusVersion     `json:"version"`
	Players     StatusPlayers     `json:"players"`
	Description StatusDescription `json:"description"`
	Favicon     string            `json:"favicon"`
}
