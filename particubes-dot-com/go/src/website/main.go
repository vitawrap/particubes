package main

import (
	"calvados"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// preprocessor function
func CreateRedirectionsFromFrontmatter(c *calvados.Calvados) error {
	var aliases map[string]string = make(map[string]string)
	var err error = nil
	// generate aliases map
	aliases, err = aliasParseMarkdownFiles("/www")
	if err != nil {
		return err
	}
	for alias, canonical := range aliases {
		c.SetRedirection(alias, canonical)
	}
	return nil
}

// beta form handling route
func HandleBetaForm(c *gin.Context, calva *calvados.Calvados) {
	fmt.Println("💥 HandleBetaForm 💥")
	c.JSON(http.StatusOK, gin.H{})
}

//
func main() {
	config := calvados.NewConfig("en", "Particubes", false)

	c := calvados.WithConfig(config)

	c.AddTemplateDir("/www/_templates")

	c.SetRedirection("/get", "https://itunes.apple.com/app/id1299143207?mt=8")

	c.SetRedirection("/download-mac-alpha", "https://download.particubes.com/Particubes-0.0.5.dmg")

	c.SetRedirection("/discord", "https://discord.gg/NbpdAkv")
	c.SetRedirection("/discord-fr", "https://discord.gg/xVSqdJu")

	c.AddPreprocessorFunc(CreateRedirectionsFromFrontmatter)

	c.AddCustomRoute(calvados.NewCustomRoute("POST", "/beta/form", HandleBetaForm))

	c.Run(":80")
}
