package lib

import (
	"net/http"
	"text/template"
)

// Segment Number refers to how one wishes to break down
// cards
const (
	SEGMENT_NUM = 1
)

type SpeakerStruct struct {
	Pic      string
	Name     string
	Position string
	Talk     string
	Time     string
}

// Copy and Paste Add
// {
// 	Pic: "",
// 	Name: "",
// 	Position: "",
// 	Role: "",
// 	LinkedIn: "",
// },

func GetSpeakers() []SpeakerStruct {
	return []SpeakerStruct{

		{
			Pic:      "JamesPerry.jpeg",
			Name:     "James Perry",
			Position: "Security Lead at AWS",
			Talk:     "Cultivating Cybersecurity Culture: Navigating New Threats in the Cloud",
			Time:     "",
		},
		{
			Pic:      "personPlaceholder.png",
			Name:     "Guest FBI Speaker",
			Position: "",
			Talk:     "",
			Time:     "",
		},
		{
			Pic:      "Patrick_Park.jpg",
			Name:     "Patrick Park",
			Position: "Board Director - ISACA New Jersey Chapter",
			Talk:     "Cyberattacks in Corporations: Myth and Reality",
			Time:     "",
		},

		/*
					{
				Pic:      "JamesPerry.jpeg",
				Name:     "James Perry",
				Position: "Security Lead at AWS",
				Talk:     "Cultivating Cybersecurity Culture: Navigating New Threats in the Cloud",
				Time:     "",
			},
			{
				Pic:      "personPlaceholder.png",
				Name:     "Guest FBI Speaker",
				Position: "[REDACTED]",
				Talk:     "[REDACTED]",
				Time:     "",
			},
			{
				Pic:      "Patrick_Park.jpg",
				Name:     "Patrick Park",
				Position: "Director of Information security at Milbank",
				Talk:     "",
				Time:     "",
			},
		*/
	}
}

// --------------------------------------------

// 1- 5 inclusive
// All on the same line
// 6 - 12 inclusive
// Being equal or bottom line has more than top line
// >12
// More Dynamic Grid like in Recognitions

// Handles Speaker Tab and writes to website
func Speaker(w http.ResponseWriter, tpl *template.Template) {
	speakers := GetSpeakers()

	tpl.ExecuteTemplate(w, "speaker-start", nil)

	tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)

	for _, speaker := range speakers {
		tpl.ExecuteTemplate(w, "speaker-card", speaker)
	}

	tpl.ExecuteTemplate(w, "speaker-div-end", nil)

	tpl.ExecuteTemplate(w, "speaker-div-end", nil)

}
