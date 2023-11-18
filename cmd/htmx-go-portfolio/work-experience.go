package main

type WorkExperienceItem struct {
	RoleTitle string
	Company   string
	Date      string
	Details   []ItemDetails
}

var workExperienceItems = []WorkExperienceItem{
	{
		RoleTitle: "DevOps Engineer",
		Company:   "Critical Force",
		Date:      "02/2022 - ... (current)",
		Details: []ItemDetails{
			{
				Header: `Migrating an "on-prem" datacenter to Google Cloud`,
				Text:   "Cloud architecture, networking, tunnels, migrations, ...",
			},
			{
				Header: "Evaluating and designing infrastructure for a monitoring platform",
				Text:   "Different types of collector agents, metric backends, log backends, backups, cloud storage options, response times, historical data, ...",
			},
			{
				Header: "Designing and building a GitOps platform over its skeleton (GKE)",
				Text:   "GKE Autopilot, Shared VPC, private clusters. Kubernetes administration: networking, volumes, ingress-nginx, external-dns, cert-manager, kustomize overlays, namespacing, security, migrations, ...",
			},
			{
				Header: "Maintaining company infrastructure",
				Text:   "Tooling and service upgrades, maintaining and improving internal tooling and services, modularizing, improving git hygiene, ...",
			},
			{
				Header: "Helping others grow while doing so myself",
				Text:   "Assisting others while learning from them myself",
			},
		},
	},
	{
		RoleTitle: "Junior DevOps Engineer",
		Company:   "", // Critical Force, template won't print empties
		Date:      "01/2022 - 02/2023 (~14 months)",
		Details: []ItemDetails{
			{
				Header: "Broadening expertise to other infrastructure",
				Text:   "Amazon Web Services, game backends, game servers, Linux servers, internal tools and services, ...",
			},
			{
				Header: "Designing and building analytics platforms",
				Text:   "Automation-first, git-driven, reusable modular platforms to support many products",
			},
			{
				Header: "Improving the DevOps culture",
				Text:   "Automating and documenting manual processes, empowering developers, embracing git, improving Infrastructure as Code, ...",
			},
			{
				Header: "Maintaining and optimizing the larger tech stack as a whole",
				Text:   "Terraform/Kubernets/Python/Go upgrades, maintaining and improving internal services, right-sizing workloads, improving traffic balancing, migrations to modern patterns, ...",
			},
			{
				Header: "Helping others grow while doing so myself",
				Text:   "Learning never stops; Guiding the local university towards modern tech solutions, onboarding new joins to tech and company practices, ...",
			},
		},
	},
	{
		RoleTitle: "TechOps Intern",
		Company:   "", // Critical Force, template won't print empties
		Date:      "06/2021 - 12/2021 (~6 months)",
		Details: []ItemDetails{
			{
				Header: "Google Kubernetes Engine (GKE) migrations",
				Text:   "Migrating unknown complexity towards known simplicity of dedicated servers, monitoring/alerting, ...",
			},
			{
				Header: "Learning the analytics backend",
				Text:   "Building data and CI pipelines for stakeholders, maintaining internal and external services, patching data, ...",
			},
			{
				Header: "Maintaining the tech stack",
				Text:   "Terraform upgrades, Kubernetes upgrades, Python upgrades, ...",
			},
			{
				Header: "Growing to an engineer from an engineering student",
				Text:   "Studying, researching, asking lots and lots of questions to absorb knowledge from our awesome seniors.",
			},
		},
	},
	{
		RoleTitle: "Cashier",
		Company:   "Neste Sotkamo",
		Date:      "09/2013 - 08/2019 (~6 years)",
		Details: []ItemDetails{
			{
				Header: "Customer service & experience",
				Text:   "Face-to-face, by phone and by email. Only happy customers!",
			},
			{
				Header: "Social media marketing",
				Text: `Building the company Facebook account. Designing and implementing marketing campaigns to Facebook.
                Improving social media and web presence.`,
			},
			{
				Header: "Website maintenance",
				Text:   "Keeping the company website running steady and its information up to date.",
			},
			{
				Header: "Troubleshooting and maintenance",
				Text:   "Ranging from connectivity issues to malfunctioning burger toasters.",
			},
		},
	},
}
