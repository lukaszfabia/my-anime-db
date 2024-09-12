package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadDir string

const Avatar UploadDir = "avatars"
const AnimesImg UploadDir = "animes"
const CharactersImg UploadDir = "characters"
const ActorsImg UploadDir = "actors"
const PostsImg UploadDir = "posts"
const StudiosLogo UploadDir = "studios"

func SaveImage(c *gin.Context, uploadDir UploadDir, keyToImg string) *string {
	if !isFileInRequest(c, keyToImg) {
		return nil
	}

	file, _ := c.FormFile(keyToImg)

	filename := uuid.New()
	// get extension
	var extension string = path.Ext(file.Filename) // its like a .jpg

	pathToFile := buildNewPath(uploadDir, filename, extension)

	if err := c.SaveUploadedFile(file, fmt.Sprintf(".%s", *pathToFile)); err != nil {
		return nil
	}

	log.Println("file saved to", *pathToFile)
	return pathToFile
}

func RemoveImage(filepath *string) error {
	parsedUrl, err := url.Parse(*filepath)

	if err != nil {
		log.Println(err)
		return err
	}

	relativePath := fmt.Sprintf(".%s", parsedUrl.Path)

	if err := os.Remove(relativePath); err != nil {
		log.Printf("%s not found\n", relativePath)
		return err
	}

	log.Println("file removed")
	return nil
}

func UpdateImage(c *gin.Context, oldPath *string, uploadDir UploadDir, keyToImg string) *string {
	if !isFileInRequest(c, keyToImg) {
		return oldPath
	}

	file, _ := c.FormFile(keyToImg)

	if oldPath == nil {
		return SaveImage(c, uploadDir, keyToImg)
	}

	if err := os.Remove(fmt.Sprintf(".%s", *oldPath)); err != nil {
		return nil
	}

	var newExtension string = path.Ext(file.Filename) // its like a .jpg
	pathToFile := buildNewPath(uploadDir, uuid.New(), newExtension)

	if err := c.SaveUploadedFile(file, fmt.Sprintf(".%s", *pathToFile)); err != nil {
		return nil
	}
	return pathToFile
}

func isFileInRequest(c *gin.Context, key string) bool {
	_, err := c.FormFile(key)
	return err == nil
}

func buildNewPath(uploadDir UploadDir, filename uuid.UUID, extension string) *string {
	res := path.Join("/", "upload", string(uploadDir), fmt.Sprintf("%s%s", filename.String(), extension))
	return &res
}
