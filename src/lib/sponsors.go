package lib

type Tier struct {
	Spons    []Sponsor
	TierName string
	CSSName  string
}

type Sponsor struct {
	HREF   string
	Source string
	ALT    string
	Width  string
}

type Sponsors struct {
	AllTiers []Tier
}

func GetSponsors() Sponsors {
	title := []Sponsor{
		{
			HREF:   "https://www.cyber.nj.gov/",
			Source: "NJCCIC_logo.png",
			ALT:    "NJCCIC",
			Width:  "200",
		},
	}

	ciso := []Sponsor{
		// {
		// 	HREF:   "https://www.merck.com/",
		// 	Source: "Merck.png",
		// 	ALT:    "Merck",
		// 	Width:  "200",
		// },
	}

	manager := []Sponsor{
		{
			HREF:   "https://www.blackwood.ai/",
			Source: "Blackwood_Horizontal.png",
			ALT:    "Blackwood",
			Width:  "200",
		},
		{
			HREF:   "https://www.splunk.com/",
			Source: "Splunk_No_Bkg.png",
			ALT:    "Splunk",
			Width:  "200",
		},
	}

	networkEng := []Sponsor{
		{
			HREF:   "https://www.adp.com/",
			Source: "ADP_Logo.png",
			ALT:    "ADP",
			Width:  "200",
		},
	}

	analyst := []Sponsor{
		{
			HREF:   "https://aws.amazon.com/",
			Source: "AWS_Logo.png",
			ALT:    "AWS",
			Width:  "200",
		},
		{
			HREF:   "https://futureconevents.com/",
			Source: "FC_Logo_Red_Background.png",
			ALT:    "Futurecon",
			Width:  "200",
		},
		{
			HREF:   "https://www.paloaltonetworks.com/",
			Source: "palo_alto_logo.png",
			ALT:    "Palo Alto",
			Width:  "200",
		},
		{
			HREF:   "https://www.crowdstrike.com/en-us/",
			Source: "Crowdstrike_Logo.png",
			ALT:    "CrowdStrike",
			Width:  "200",
		},
		{
			HREF:   "https://www.isaca.org/",
			Source: "ISACA_Logo.jpg",
			ALT:    "ISACA",
			Width:  "200",
		},
		{
			HREF:   "https://github.com/",
			Source: "github_logo.png",
			ALT:    "GitHub",
			Width:  "200",
		},
		{
			HREF:   "https://www.tekstream.com/",
			Source: "TekStream.png",
			ALT:    "TekStrem",
			Width:  "200",
		},
		// {
		// 	HREF:   "https://letsdefend.io/",
		// 	Source: "LetsDefendLogo.png",
		// 	ALT:    "Let's Defend",
		// 	Width:  "200",
		// },
	}

	tiers := []Tier{
		{
			Spons:    title,
			TierName: "Title Sponsor",
			CSSName:  "sponsor-title",
		},

		{
			Spons:    ciso,
			TierName: "CISO Sponsor",
			CSSName:  "sponsor-ciso",
		},
		{
			Spons:    networkEng,
			TierName: "Network Engineer Sponsor",
			CSSName:  "sponsor-network",
		},
		{
			Spons:    manager,
			TierName: "Security Manager Sponsor",
			CSSName:  "sponsor-manager",
		},
		{
			Spons:    analyst,
			TierName: "Security Analyst Sponsors",
			CSSName:  "sponsor-analyst",
		},
	}

	return Sponsors{
		AllTiers: tiers,
	}
}
