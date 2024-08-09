package utils

import (
	"fmt"
	"log"
	"net/url"
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
Saves image in/on filesystem/server

params:

	c: *gin.Context - gin context, used to get raw file from request
	props: SaverProps - struct with properties needed to save image
	e.g SaverProps{
		Dir:         utils.AnimesImg, // directory where image will be saved
		Placeholder: utils.DefaultConent, // placeholder, used when something goes wrong
		KeyToImg:    "picUrl", // key to get image from request
		Filename:    body.Title, // new name of a file
	}

returns:

	string - path to saved image
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

	props.Filename = strings.ReplaceAll(props.Filename, ".", "")
	props.Filename = strings.ReplaceAll(props.Filename, " ", "_")
	props.Filename = strings.ToLower(props.Filename)

	relativePath := fmt.Sprintf("/upload/%s/%s.%s", props.Dir, props.Filename, extension)

	path := fmt.Sprintf(".%s", relativePath)

	if err := c.SaveUploadedFile(image, path); err != nil {
		return placeholderFile
	}

	var picURL string = fmt.Sprintf("%s%s", host, relativePath)

	return picURL
}

/*
Remove image from filesystem/server

params:

	filepath: string - path to image on server comes from e.g user.picUrl
*/
func RemoveImage(filepath string) {
	log.Println("path", filepath)

	parsedUrl, err := url.Parse(filepath)

	if err != nil {
		log.Println(err)
		return
	}

	relativePath := fmt.Sprintf(".%s", parsedUrl.Path)

	elems := strings.Split(relativePath, "/")

	filename := elems[len(elems)-1]

	if filename == "def.png" || filename == "content.jpg" {
		log.Println("can't remove placeholder!")
		return
	}

	if err := os.Remove(relativePath); err != nil {
		log.Printf("%s not found\n", relativePath)
	} else {
		log.Println("file successfully removed")
	}
}
