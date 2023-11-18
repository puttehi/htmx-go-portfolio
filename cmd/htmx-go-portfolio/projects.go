package main

type PortfolioItem struct {
	Name       string
	Platform   string
	SVGData    SVGData
	SourceURL  string
	SourceText string
	Video      string
	Details    []ItemDetails
}

var portfolioItems = []PortfolioItem{
	{
		Name:       "Infinigolf (Unity)",
		Platform:   "Android",
		SVGData:    SVGItchIo,
		SourceURL:  "https://puttehi.itch.io/infinigolf",
		SourceText: "Infinigolf (itch.io)",
		Video:      "https://puttehi.github.io/mp4/infinigolf.mp4",
		Details: []ItemDetails{
			{
				Header: "General",
				Text:   "Prototype developed solo in 3 months for a game engineering course.",
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"Algorithms and data structures (procedural generation)",
					"Advanced cameras and directing (Cinemachine)",
					"Data tweening (DOTween)",
					"3D modeling (DOTween)",
					"Node-based shader progamming (Shader Graph)",
					"Object pooling (LeanPool)",
					"Query building (LINQ)",
					"Market research (Mobile)",
					"Publishing (Store requirements)",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"C# (Unity)",
					"GLSL",
					"(Shader Graph)",
				},
			},
		},
	},
	{
		Name:       "Project 2",
		Platform:   "Platform",
		SVGData:    SVGCode,
		SourceURL:  "https://www.gitlab.com/",
		SourceText: "group-name/repository-name (GitLab)",
		Video:      "",
		Details: []ItemDetails{
			{
				Header: "Header 1",
				Text:   "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
			},
			{
				Header: "Header 2",
				Text:   "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
			},
			{
				Header: "Header 3",
				Text:   "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
			},
			{
				Header: "Header 4",
				Text:   "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
			},
		},
	},
}
