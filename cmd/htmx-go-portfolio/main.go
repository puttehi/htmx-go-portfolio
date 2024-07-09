package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

type PageData struct {
	NextNavbarAction string // TODO
}

const webRoot = "web/public"
const port = 3000

func main() {
	r := gin.Default()

	// Root directory is current working directory
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	absoluteWebRoot := filepath.Join(path, webRoot)
	r.Use(static.ServeRoot("/", absoluteWebRoot))

	fmt.Printf("\n\n*****\nServing web root (/) from %s\n", absoluteWebRoot)
	fmt.Printf("\nStarting on http://localhost:%d\n*****\n\n", port)

	r.Run(fmt.Sprintf(":%d", port))
}
