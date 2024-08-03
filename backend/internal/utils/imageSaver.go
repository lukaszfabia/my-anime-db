package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type UploadDir string

const Avatar UploadDir = "avatars"
const AnimesImg UploadDir = "animes"
const CharactersImg UploadDir = "characters"
const ActorsImg UploadDir = "actors"

type Placeholders string

const DefaultImage Placeholders = "def.png"
const DefaultConent Placeholders = "content.jpg"

type SaverProps struct {
	Dir         UploadDir
	Placeholder Placeholders
	KeyToImg    string
	Filename    string
}

/*
c - gin's context

filename - e.g username for user, title for anime

keyToImg - json's key name e.g picUrl, defined in models
*/
func SaveImage(c *gin.Context, props SaverProps) string {
	err := godotenv.Load()
	if err != nil {
		panic("Could't find .env file!")
	}

	var host string = fmt.Sprintf("http://%s:%s", os.Getenv("HOST"), os.Getenv("API_PORT"))
	var defaultsPath string = fmt.Sprintf("%s/upload/placeholders/", host)
	var placeholderFile string = fmt.Sprintf("%s%s", defaultsPath, props.Placeholder)

	// get file from request
	image, err := c.FormFile(props.KeyToImg)

	if err != nil || image == nil {
		log.Println("There was problem with file or no image in form")
		return placeholderFile
	}

	// get extension
	var extension string = strings.Split(image.Filename, ".")[1]

	relativePath := fmt.Sprintf("/upload/%s/%s.%s", props.Dir, props.Filename, extension)

	path := fmt.Sprintf(".%s", relativePath)

	if err := c.SaveUploadedFile(image, path); err != nil {
		return placeholderFile
	}

	var picURL string = fmt.Sprintf("%s%s", host, relativePath)

	return picURL
}
