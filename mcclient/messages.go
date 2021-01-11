package mcclient

// User represents a minecraft user
type User struct {
	Name string `json:"name" firestore:"name"`
	UUID string `json:"id" firestore:"uuid"`
}

// StatusVersion represents the server version
type StatusVersion struct {
	Name     string `json:"name" firestore:"name"`
	Protocol int    `json:"protocol" firestore:"protocol"`
}

// StatusPlayers represents the players logged in to the server
type StatusPlayers struct {
	Max    int    `json:"max" firestore:"max"`
	Online int    `json:"online" firestore:"online"`
	Users  []User `json:"sample" firestore:"users"`
}

// StatusDescription represents the message of the day reported by the server
type StatusDescription struct {
	Text string `json:"text" firestore:"text"`
}

// StatusResponse is the response to a status request to the server
type StatusResponse struct {
	Version     StatusVersion     `json:"version" firestore:"version"`
	Players     StatusPlayers     `json:"players" firestore:"players"`
	Description StatusDescription `json:"description" firestore:"description"`
	Favicon     string            `json:"favicon" firestore:"favicon"`
}
