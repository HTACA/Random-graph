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
	imageDir := `./list` // 赋值文件地址

	// 调用函数getImageFiles，获取图片的文件地址
	var err error
	files, err = getImageFiles(imageDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 如果没有找到图片文件
	if len(files) == 0 {
		fmt.Println("No image files found in the directory.")
		return
	}

	HttpService()

}

// 创建getImageFiles
func getImageFiles(dir string) ([]string, error) { // 输入文件夹地址，返回两个值，一个为文件地址，一个为可能的错误
	var imageFiles []string // 定义一个字符串切片，用来储存地址

	// 有三种遍历目录的方法，filepath.Walk，ioutil.ReadDir，os.File.Readdir，这里使用filepath.Walk方法获取文件列表
	//  dir为需要遍历的文件夹，第二个参数为一个函数，函数内接受path文件路径，info文件信息，err错误
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 先判断有没有错误
		if err != nil {
			return err
		}
		// IsDir若为false,则为文件，isImageFile检查是否为图片
		if !info.IsDir() && isImageFile(path) {
			// 通过切片方法加入数组
			imageFiles = append(imageFiles, path)
		}
		// 若上面正确，则 err = nil
		return nil
	})

	return imageFiles, err
}

// 创建isImageFile，接受一个文件地址
func isImageFile(file string) bool {
	// 将文件地址放到函数中，filepath.Ext方法返回一个文件的拓展名
	ext := filepath.Ext(file)
	// 拓展名为下面的列则为true
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return true
	}
	// 否则为false
	return false
}

func getRandomImage() string {
	// 生成随机数种子，根据时间来生成
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 指定随机的范围，生成随机的下标
	randomIndex := r.Intn(len(files))
	// 将下标给到数组，返回生成随机的数组内容
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
