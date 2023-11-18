package main

type LinkData struct {
	URL     string
	Text    string
	SVGData SVGData
}

type PortfolioItem struct {
	Name     string
	Platform string
	Links    []LinkData
	Video    string
	Details  []ItemDetails
}

var portfolioItems = []PortfolioItem{
	{
		Name:     "Infinigolf (Unity)",
		Platform: "Android",
		Links: []LinkData{
			{
				URL:     "https://puttehi.itch.io/infinigolf",
				Text:    "Infinigolf (itch.io)",
				SVGData: SVGItchIo,
			},
		},
		Video: "https://puttehi.github.io/mp4/infinigolf.mp4",
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
		Name:     "Routa Engine Tutorial",
		Platform: "Cross-platform",
		Links: []LinkData{
			{
				URL:     "https://gitlab.dclabra.fi/wiki/s/HJrXEPW58",
				Text:    "Written tutorial: Part 1 (CodiMD @KAMK DCLabra)",
				SVGData: SVGText,
			},
			{
				URL:     "https://gitlab.dclabra.fi/wiki/s/rkHV83ccI",
				Text:    "Written tutorial: Part 1 (CodiMD @KAMK DCLabra)",
				SVGData: SVGText,
			},
		},
		Video: "https://puttehi.github.io/mp4/routa_tutorials.mp4",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "Routa Engine is a game engine developed by students, for students in-house @ KAMK. No one knew how to use it though.",
				ListItems: []string{
					"3-man team",
					"Developed a single platformer game",
					"Wrote a tutorial on Routa basics for first-year students",
				},
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"More advanced C++",
					"C++ build systems",
					"Entity Component Systems (ECS)",
					"Technical writing",
					"Learning a new, big codebase",
					"Fixing a lot of low-level game engine bugs",
					"Tilemap creation (Tiled editor)",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"C++",
					"Markdown",
				},
			},
		},
	},
}
