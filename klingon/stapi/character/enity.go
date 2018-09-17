package character

// Species represents a species of the current Character
type Species struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}

// Short represents brief structure of Character
type Short struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

// Full represents full structure of Character
type Full struct {
	UID              UID       `json:"uid"`
	Name             string    `json:"name"`
	Gender           string    `json:"gender"`
	CharacterSpecies []Species `json:"characterSpecies"`
}
