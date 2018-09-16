package character

type Species struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}

type Short struct {
	UID              string `json:"uid"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	YearOfBirth      int    `json:"yearOfBirth"`
	AlternateReality bool   `json:"alternateReality"`
}

type Full struct {
	UID                UID       `json:"uid"`
	Name               string    `json:"name"`
	Gender             string    `json:"gender"`
	YearOfBirth        int       `json:"yearOfBirth"`
	MonthOfBirth       int       `json:"monthOfBirth"`
	DayOfBirth         int       `json:"dayOfBirth"`
	Hologram           bool      `json:"hologram"`
	FictionalCharacter bool      `json:"fictionalCharacter"`
	Mirror             bool      `json:"mirror"`
	AlternateReality   bool      `json:"alternateReality"`
	CharacterSpecies   []Species `json:"characterSpecies"`
	Titles             []struct {
		UID            string `json:"uid"`
		Name           string `json:"name"`
		MilitaryRank   bool   `json:"militaryRank"`
		FleetRank      bool   `json:"fleetRank"`
		ReligiousTitle bool   `json:"religiousTitle"`
		Position       bool   `json:"position"`
		Mirror         bool   `json:"mirror"`
	} `json:"titles"`
}
