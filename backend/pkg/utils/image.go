package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/* TODO: sciezka w bazie powinna byc /path/to/image.jpg reszta bedzie dodawana na froncie api path +sciezka
jesli user nie ma zdjecia to ustawiamy na pusty string czy cos a avatar bedzie jako piersza litera username
wymyslic jak ogranac nazewnictwo plików
nazwa musi byc unikalna a pole nie moze byc mutowalne
uzyc tego do generowania nazwy pliku : "github.com/google/uuid"

placeholdery do wyjebania z kodu
ten struct do wyjebania
przyjac upload dir / key to img albo surowy plik / username jako skladowa nazwy pliku
file name genrowac w funkcji save image
*/

type UploadDir string

const Avatar UploadDir = "avatars"
const AnimesImg UploadDir = "animes"
const CharactersImg UploadDir = "characters"
const ActorsImg UploadDir = "actors"
const PostsImg UploadDir = "posts"

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
func SaveImage(c *gin.Context, uploadDir UploadDir, keyToImg string) *string {
	// get extension
	file, err := c.FormFile(keyToImg)
	if err != nil {
		return nil
	}

	filename := uuid.New()
	var extension string = path.Ext(file.Filename) // its like a .jpg

	pathToFile := path.Join("/", "upload", string(uploadDir), fmt.Sprintf("%s%s", filename, extension))

	if err := c.SaveUploadedFile(file, "."+pathToFile); err != nil {
		return nil
	}

	log.Println("file saved to", pathToFile)
	return &pathToFile
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

func UpdateImage(c *gin.Context, oldPath string, uploadDir UploadDir, keyToImg string) *string {

	RemoveImage(oldPath)
	return SaveImage(c, uploadDir, keyToImg)
}
