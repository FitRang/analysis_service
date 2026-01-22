package geminiapi

type Profile struct {
	ID           string  `json:"id"`
	FullName     string  `json:"fullName"`
	Username     string  `json:"username"`
	AccessStatus string  `json:"accessStatus"`
	ProfileURL   *string `json:"profileUrl,omitempty"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

type Dossier struct {
	ID              string   `json:"id"`
	Username        string   `json:"username"`
	FaceType        string   `json:"faceType"`
	SkinTone        SkinTone `json:"skinTone"`
	BodyType        string   `json:"bodyType"`
	Gender          Gender   `json:"gender"`
	PreferredColors []string `json:"preferredColors,omitempty"`
	DislikedColors  []string `json:"dislikedColors,omitempty"`
	Viewers         []string `json:"viewers"`
	Height          *string  `json:"height,omitempty"`
	Weight          *string  `json:"weight,omitempty"`
	CreatedAt       string   `json:"createdAt"`
	UpdatedAt       string   `json:"updatedAt"`
}
type SkinTone string

const (
	SkinTonePale  SkinTone = "pale"
	SkinToneLight SkinTone = "light"
	SkinToneOlive SkinTone = "olive"
	SkinToneDark  SkinTone = "dark"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)
