package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

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

func ReadConfigsAndContent(configDirectory string) (Config, []WorkExperienceItem, []ProjectItem, error) {
	var errs []error
	config := Config{}
	err := ReadContentFromJSON(fmt.Sprintf("%s/%s", configDirectory, "configs/site-config.json"), &config)
	if err != nil {
		errs = append(errs, fmt.Errorf("Could not read site config: %s", err))
	}

	workExperienceItems := []WorkExperienceItem{}
	err = ReadContentFromJSON(fmt.Sprintf("%s/%s", configDirectory, "configs/work-experience-items.json"), &workExperienceItems)
	if err != nil {
		errs = append(errs, fmt.Errorf("Could not read work experience items config: %s", err))
	}

	projectItems := []ProjectItem{}
	err = ReadContentFromJSON(fmt.Sprintf("%s/%s", configDirectory, "configs/project-items.json"), &projectItems)
	if err != nil {
		errs = append(errs, fmt.Errorf("Could not read project items config: %s", err))
	}

	return config, workExperienceItems, projectItems, errors.Join(errs...)
}

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

	config, workExperienceItems, projectItems, err := ReadConfigsAndContent(path)

	if err != nil {
		slog.Default().Error("Initial configuration read failed", "err", err)
		os.Exit(1)
	}

	pageData := PageData{
		PersonName:          config.PersonName,
		NextNavbarAction:    config.NextNavbarAction,
		WorkExperienceItems: workExperienceItems,
		ProjectItems:        projectItems,
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", pageData)
	})

	r.GET("/navbar/:navbarAction", func(c *gin.Context) {
		wanted := c.Param("navbarAction")
		var next string
		if wanted == "show" {
			next = "hide"
		} else {
			next = "show"
		}
		sessionPageData := pageData
		sessionPageData.NextNavbarAction = next
		c.HTML(http.StatusOK, "navbar.htmx", sessionPageData)
	})

	fmt.Printf("\n*****\nStarting on http://localhost:%d\n*****\n\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
