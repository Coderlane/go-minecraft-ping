package mcclient

// User represents a minecraft user
type User struct {
	Name string `json:"name" firebase:"name"`
	UUID string `json:"id" firebase:"uuid"`
}

// StatusVersion represents the server version
type StatusVersion struct {
	Name     string `json:"name" firebase:"name"`
	Protocol int    `json:"protocol" firebase:"protocol"`
}

// StatusPlayers represents the players logged in to the server
type StatusPlayers struct {
	Max    int    `json:"max" firebase:"max"`
	Online int    `json:"online" firebase:"online"`
	Users  []User `json:"sample" firebase:"users"`
}

// StatusDescription represents the message of the day reported by the server
type StatusDescription struct {
	Text string `json:"text" firebase:"text"`
}

// StatusResponse is the response to a status request to the server
type StatusResponse struct {
	Version     StatusVersion     `json:"version" firebase:"version"`
	Players     StatusPlayers     `json:"players" firebase:"players"`
	Description StatusDescription `json:"description" firebase:"description"`
	Favicon     string            `json:"favicon" firebase:"favicon"`
}
