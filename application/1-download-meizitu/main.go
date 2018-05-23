package main

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	SavePath  string
	MinWidth  int
	MinHeight int
	Overwrite bool
	MaxPage   int
	StartPage int
}

func NewConfig(savePath string, minWidth, minHeight, maxPage, startPage int, overwrite bool) *Config {
	return &Config{
		savePath,
		minWidth,
		minHeight,
		overwrite,
		maxPage,
		startPage,
	}
}

const (
	PAGE_URL          string = "http://www.meizitu.com/a/sifang_5_%d.html"
	IMAGE_LIST_LINKS  string = "http://www.meizitu.com/a/[0-9]+.html"
	IMAGE_IMAGE_LINKS string = "http://mm.chinasareview.com/wp-content/uploads/[^\\.]+\\.(jpg|png|gif)"
)

type Webpage struct {
	Config *Config
}

func NewWebpage(config *Config) *Webpage {
	return &Webpage{Config: config}
}

func (self *Webpage) ParsePage(url string) []string {
	offset := self.Config.StartPage + self.Config.MaxPage
	var urls []string
	for curPage := self.Config.StartPage; curPage < offset; curPage++ {
		urls = append(urls, fmt.Sprintf(url, curPage))
	}
	return urls
}

func (self *Webpage) Get(url string) (body string) {
	resp, ok := http.Get(url)
	if nil != ok {
		return ""
	}
	defer resp.Body.Close()
	str, ok := ioutil.ReadAll(resp.Body)
	if ok != nil {
		return ""
	}
	return string(str)
}

func (self *Webpage) ParseUrl(url, pattern string) (links []string) {
	fmt.Println("Parse url ==>", url)
	body := self.Get(url)
	if "" == body {
		return []string{}
	}
	reg := regexp.MustCompile(pattern)
	return reg.FindAllString(body, -1)
}

func (self *Webpage) GetSaveName(url string) string {
	paths := strings.Split(url, "/")
	len := len(paths)
	fileName := self.Config.SavePath + paths[len-4] + paths[len-3] + paths[len-2] + paths[len-1]
	return fileName
}

func (self *Webpage) Download(urls []string) {
	for _, url := range urls {
		fmt.Println("Start download image from url ==>", url)
		fileName := self.GetSaveName(url)
		if self.FileExist(fileName) && !self.Config.Overwrite {
			fmt.Println("Image already exists, skip download ==>", url)
			continue
		}
		body := self.Get(url)
		if "" == body {
			continue
		}
		if !self.CheckSize(body, self.GetExt(url)) {
			fmt.Println("Image size too small, skip download ==>", url)
			continue
		}
		if !self.SaveImage(body, fileName) {
			fmt.Println("Save image failed ==>", url)
		}
	}
}

func (self *Webpage) SaveImage(body, name string) bool {
	f, ok := os.Create(name)
	if ok != nil {
		fmt.Println("open file error")
		return false
	}
	defer f.Close()
	if _, err := f.WriteString(body); err == nil {
		return true
	}
	return false
}

func (self *Webpage) GetExt(url string) string {
	if url == "" {
		return ""
	}
	temp := strings.Split(url, ".")
	return temp[len(temp)-1]
}

func (self *Webpage) CheckSize(body, ext string) bool {
	if self.Config.MinWidth <= 0 && self.Config.MinHeight <= 0 {
		return true
	}
	var iImage image.Image
	var ok error = errors.New("Unknow image type")
	switch ext {
	case "jpg":
		iImage, ok = jpeg.Decode(strings.NewReader(body))
	case "png":
		iImage, ok = png.Decode(strings.NewReader(body))
	case "gif":
		iImage, ok = gif.Decode(strings.NewReader(body))
	default:
		fmt.Println("Unknow image format")
		return false
	}
	if ok == nil {
		rect := iImage.Bounds()
		if self.Config.MinWidth <= rect.Max.X && self.Config.MinHeight <= rect.Max.Y {
			return true
		}
	}
	return false
}

func (self *Webpage) FileExist(name string) bool {
	if _, ok := os.Stat(name); ok == nil {
		return true
	}
	return false
}

func (self *Webpage) RunTask() {
	urls := self.ParsePage(PAGE_URL)
	sum := 0
	l := len(urls)
	c := make(chan int, l)
	for _, url := range urls {
		go func(url string) {
			links := self.ParseUrl(url, IMAGE_LIST_LINKS)
			for _, v := range links {
				uris := self.ParseUrl(v, IMAGE_IMAGE_LINKS)
				self.Download(uris)
			}
			c <- 1
		}(url)
	}
forEnd:
	for {
		select {
		case <-c:
			sum++
			if sum == l {
				break forEnd
			}
		}
	}
}

func main() {
	os.Mkdir("E:/girls", 0777)
	config := NewConfig(
		"E:/girls/",
		400,
		400,
		12,
		20,
		false,
	)

	webpage := NewWebpage(config)
	webpage.RunTask()

	fmt.Println("done!")
}
