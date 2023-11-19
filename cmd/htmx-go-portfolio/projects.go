package main

type LinkData struct {
	URL     string
	Text    string
	SVGData SVGData
}

type ProjectItem struct {
	Name     string
	Platform string
	Links    []LinkData
	Video    string
	Details  []ItemDetails
}

var projectItems = []ProjectItem{
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
		Video: "assets/mp4/infinigolf.mp4",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "Mobile game development in the form of an \"infinite runner\" -style minigolf game.",
				ListItems: []string{
					"Prototype developed solo in about 3 months",
				},
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
		Name:     "www.puttehi.eu infrastructure and service automation",
		Platform: "Linux",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/puttehi-cloud",
				Text:    "puttehi/puttehi-cloud (GitHub)",
				SVGData: SVGCode,
			},
		},
		Video: "",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "My personal cloud server configuration.",
				ListItems: []string{
					"Development services behind authentication",
					"Single command deployments",
				},
			},
			{
				Header: "Acquired (rather, used) expertise",
				ListItems: []string{
					"Infrastructure automation (Terraform)",
					"Service automation (Ansible)",
					"Reverse proxy (nginx)",
					"Cloud services (Linode)",
					"Domain ownership (GoDaddy)",
					"Remote workstations (automated personalization)",
					"Authentication layers (Vouch)",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"Terraform (HCL)",
					"Ansible (YAML)",
					"Nginx configuration",
				},
			},
		},
	},
	{
		Name:     "Python application template with full CI/CD",
		Platform: "GitHub",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/slacktube",
				Text:    "puttehi/slacktube (GitHub)",
				SVGData: SVGCode,
			},
			{
				URL:     "https://slacktube.readthedocs.io/en/latest/",
				Text:    "Documentation (readthedocs.io)",
				SVGData: SVGText,
			},
			{
				URL:     "https://github.com/fedejaure/cookiecutter-modern-pypackage/issues/409",
				Text:    "fedejaure/cookiecutter-modern-pypackage improvements (GitHub)",
				SVGData: SVGCode,
			},
		},
		Video: "",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "Initially a side-project to dump YouTube links from a #jukebox Slack channel to a YouTube playlist for easy listening. Turns out automation was much more interesting.",
				ListItems: []string{
					"GitHub automation",
					"Modern Python package automation",
				},
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"GitHub Actions (CI/CD)",
					"Dependabot (GitHub)",
					"Sphinx documentation generation",
					"Read The Docs automation",
					"Python Package Index (pip) automation",
					"Modernizing old Python project templates",
					"Git hooks (pre-commit)",
					"Python environments (Poetry)",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"GitHub Actions",
					"Python",
				},
			},
		},
	},
	{
		Name:     "Rocket League replay uploader",
		Platform: "Linux, Windows",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/bcsync",
				Text:    "puttehi/bcsync (GitHub)",
				SVGData: SVGCode,
			},
			{
				URL:     "https://ballchasing.com/",
				Text:    "Replay database (ballchasing.com)",
				SVGData: SVGText,
			},
			{
				URL:     "https://store.steampowered.com/app/252950/Rocket_League/",
				Text:    "Rocket League in Steam (Steam Store)",
				SVGData: SVGText,
			},
		},
		Video: "",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "When you cannot get Bakkesmod to work on Linux and do it for you, you do it yourself.",
				ListItems: []string{
					"CLI tool",
					"Keep track of your statistics and share your game with your friends",
					"Just needs the (free) ballchasing.com API token",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"Python",
					"Bash",
					"GNU Make",
				},
			},
		},
	},
	{
		Name:     "Simplechat",
		Platform: "Windows",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/qt-chat",
				Text:    "puttehi/qt-chat (GitHub)",
				SVGData: SVGCode,
			},
		},
		Video: "assets/mp4/chat.mp4",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "Modern Windows application development in the form of a server/client chat application.",
				ListItems: []string{
					"Prototype developed solo in 1-2 months",
					"Chat commands (/me, /slap, /time)",
					"Server timestamps with optional client timestamps",
					"Persistent chat history with optional saving on clients",
				},
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"Qt framework (C++)",
					"Websockets (QtWebsocket)",
					"qmake and GNU Make",
					"Messaging protocols",
					"Sharing server/client code",
					"UI/UX design",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"C++",
					"GNU Make",
				},
			},
		},
	},
	{
		Name:     "React applications",
		Platform: "Android, iOS, Web",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/mobile-exercises",
				Text:    "puttehi/mobile-exercises (GitHub)",
				SVGData: SVGCode,
			},
		},
		Video: "assets/mp4/react.mp4",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "Modern mobile and web app development in the form of over a dozen small apps. Developed for a summertime course.",
				ListItems: []string{
					"Apps developed solo in 2-3 months",
					"Over a dozen different apps",
				},
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"React",
					"React Native",
					"Expo",
					"Ads",
					"Native toasts",
					"PWAs",
					"Cross-app traversal",
					"Google products (Maps, Places, etc.)",
					"Web APIs, fetching and parsing)",
					"NPM",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"Javascript",
				},
			},
		},
	},
	{
		Name:     "Neuroevolution inside a game clone",
		Platform: "Web",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/dj-machine-learning",
				Text:    "puttehi/dj-machine-learning (GitHub)",
				SVGData: SVGCode,
			},
		},
		Video: "assets/mp4/machine_learning.mp4",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "Playing around with AI by cloning a classic mobile game and making it play itself better than any human ever could.",
				ListItems: []string{
					"Prototype developed in about a month",
					"Trained model surpasses scores of 100 000+",
				},
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"Neuroevolution and machine learning basics",
					"AI modeling",
					"Reward functions",
					"Tensorflow.js",
					"Real-time web development",
					"p5 framework",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"Javascript",
				},
			},
		},
	},
	{
		Name:     "Easily extendable Discord bot",
		Platform: "Linux, Windows",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/discord-spaghetti-bot/tree/feat-mongo",
				Text:    "puttehi/discord-spaghetti-bot (GitHub)",
				SVGData: SVGCode,
			},
		},
		Video: "",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "A Discord bot built for my friend groups private Discord server. Initial idea was to support the release of New World with Discord-integration.",
				ListItems: []string{
					"Easily extendible command palette with a permission system",
					"New World API integrations",
					"Web scraping with parameter-based navigation and automatic screenshot threads",
				},
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"Discord API",
					"Chat-bot development (commands, user groups, admin tools)",
					"External APIs (New World World Status)",
					"Web scraping (Puppeteer)",
					"Screenshot generation",
					"Local caching",
					"Online caching and data storage (MongoDB)",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"Javascript",
					"Bash",
				},
			},
		},
	},
	{
		Name:     "Discord lunch bot",
		Platform: "Linux",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/lounaat-info",
				Text:    "puttehi/lounaat-info (GitHub)",
				SVGData: SVGCode,
			},
			{
				URL:     "https://www.lounaat.info/",
				Text:    "Lunch menus (lounaat.info)",
				SVGData: SVGText,
			},
		},
		Video: "",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "A Discord bot built for my friend groups private Discord server. Announces local lunch menus daily.",
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"Python dependency management (Poetry)",
					"Discord API",
					"External APIs (Lounaat.info)",
					"Web scraping (BeautifulSoup)",
					"Parsing HTML trees",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"Python",
				},
			},
		},
	},
	{
		Name:     "Bitburner scripts",
		Platform: "Windows/Linux (Steam)",
		Links: []LinkData{
			{
				URL:     "https://github.com/puttehi/bitburner-scripts/tree/BN1",
				Text:    "puttehi/bitburner-scripts BitNode 1 (GitHub)",
				SVGData: SVGCode,
			},
			{
				URL:     "https://store.steampowered.com/app/1812820/Bitburner/",
				Text:    "Bitburner in Steam (Steam Store)",
				SVGData: SVGText,
			},
		},
		Video: "",
		Details: []ItemDetails{
			{
				Header: "Overview",
				Text:   "Bitburner is an \"idle game\" played through programming. The goal of the game is to create the best botnet you can.",
			},
			{
				Header: "Acquired expertise",
				ListItems: []string{
					"JSDoc",
					"Typescript",
					"Distributed (local) systems i.e. a botnet",
					"Migrating from Javascript to TypeScript",
					"Code editor integration in-game through game API (VSCode)",
					"Reverse-engineering a custom botnet UI",
					"Optimization",
				},
			},
			{
				Header: "Languages",
				ListItems: []string{
					"Javascript",
					"Typescript",
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
				Text:    "Written tutorial: Part 2 (CodiMD @KAMK DCLabra)",
				SVGData: SVGText,
			},
		},
		Video: "assets/mp4/routa_tutorials.mp4",
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
