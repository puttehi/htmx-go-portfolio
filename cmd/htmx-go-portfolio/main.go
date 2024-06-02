package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ItemDetails struct {
	Header    string   `json:"header"`
	Text      string   `json:"text,omitempty"`       // (optional)
	ListItems []string `json:"list_items,omitempty"` // (optional)
}

type PageData struct {
	PersonName          string
	NextNavbarAction    string
	WorkExperienceItems []WorkExperienceItem
	ProjectItems        []ProjectItem
}

const webTemplatesRoot = "web/templates"
const webVendorRoot = "web/vendor"
const webAssetsRoot = "web/assets"
const webCSSRoot = "web/css"
const webFaviconPath = "web/favicon.ico"
const port = 3000

func main() {
	r := gin.Default()

	// Root directory is current working directory
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Load templates
	r.LoadHTMLGlob(fmt.Sprintf("%s/%s/*.htm*", path, webTemplatesRoot))

	// Serve static content
	r.Static("/vendor", fmt.Sprintf("%s/%s", path, webVendorRoot))
	r.Static("/css", fmt.Sprintf("%s/%s", path, webCSSRoot))
	r.Static("/assets", fmt.Sprintf("%s/%s", path, webAssetsRoot))
	r.StaticFile("favicon.ico", fmt.Sprintf("%s/%s", path, webFaviconPath))

	workExperienceItems := []WorkExperienceItem{}
	err = ReadContentFromJSON[[]WorkExperienceItem](fmt.Sprintf("%s/%s", path, "configs/work-experience-items.json"), &workExperienceItems)
	if err != nil {
		slog.Default().Error("Could not read work experience items config", "err", err)
		os.Exit(1)
	}

	projectItems := []ProjectItem{}
	err = ReadContentFromJSON[[]ProjectItem](fmt.Sprintf("%s/%s", path, "configs/project-items.json"), &projectItems)
	if err != nil {
		slog.Default().Error("Could not read project items config", "err", err)
		os.Exit(1)
	}

	pageData := PageData{
		PersonName:          "Petteri Zitting",
		NextNavbarAction:    "show",
		WorkExperienceItems: workExperienceItems,
		ProjectItems:        projectItems,
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", pageData)
	})

	r.GET("/navbar/:navbarAction", func(c *gin.Context) {
		pageData.NextNavbarAction = "show"
		if c.Param("navbarAction") == "show" {
			pageData.NextNavbarAction = "hide"
		}
		c.HTML(http.StatusOK, "navbar.htmx", pageData)
	})

	fmt.Printf("\n*****\nStarting on http://localhost:%d\n*****\n\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
