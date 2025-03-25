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

		// {
		// 	Pic:      "personPlaceholder.png",
		// 	Name:     "James Perry",
		// 	Position: "Security Lead at AWS",
		// 	Talk:     "TBD",
		// 	Time:     "",
		// },
		{
			Pic:      "AnthonyHarvey.jpg",
			Name:     "Anthony Harvey",
			Position: "Sr. Security Solutions Architect at AWS",
			Talk:     "From the Trenches to the Clouds: Active Incident Response and the Future of Security",
			Time:     "",
		},
		{
			Pic:      "KenFishkin.jpg",
			Name:     "Ken Fishkin",
			Position: "President of ISC2",
			Talk:     "When Cybersecurity meets Stranger Things!  Learn How Being a Dungeon Master Can Improve Your Cybersecurity Program.",
			Time:     "",
		},
		{
			Pic:      "KrishLatha.jpg",
			Name:     "Krish Latha",
			Position: "Lead STAT Analyst (Global Security) at ADP",
			Talk:     "Application development security",
			Time:     "",
		},
		{
			Pic:      "JasonSylvia.jpg",
			Name:     "Jason Sylvia",
			Position: "Sr. Dir. Information Security (Global Security) at ADP",
			Talk:     "Application development security",
			Time:     "",
		},
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

	// tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)

	for _, speaker := range speakers {
		tpl.ExecuteTemplate(w, "speaker-card", speaker)
	}

	// tpl.ExecuteTemplate(w, "speaker-div-end", nil)

	// tpl.ExecuteTemplate(w, "speaker-div-end", nil)

}
