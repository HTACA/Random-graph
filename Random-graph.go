package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var files []string

func main() {
	imageDir := `./list`
	
	var err error
	files, err = getImageFiles(imageDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No image files found in the directory.")
		return
	}

	HttpService()

}

func getImageFiles(dir string) ([]string, error) {
	var imageFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isImageFile(path) {
			imageFiles = append(imageFiles, path)
		}
		return nil
	})

	return imageFiles, err
}

func isImageFile(file string) bool {
	ext := filepath.Ext(file)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return true
	}
	return false
}

func getRandomImage() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(files))
	return files[randomIndex]
}

func HttpService() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		randomImage := getRandomImage()
		http.ServeFile(w, r, randomImage)
	})
	port := ":8080"
	fmt.Println("http service start,post is", port)
	http.ListenAndServe(port, nil)
}
