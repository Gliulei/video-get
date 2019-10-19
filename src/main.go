package main

import (
	"net/http"
	"io/ioutil"
	"os"
	"log"
	"strconv"
)

func main() {
	download()
}

//解析图片url
func download() {
	url := "https://zao.momocdn.com/tmpvideo/3E/DA/3EDAB17F-E24D-6B80-528B-6D3E6313A1F820191018.mp4";
	getHeader(url);
	//save(url,"video_test")

}

func getHeader(url string) {
	response, err := http.Head(url)
	if err != nil {
		log.Println("get url failed:", err)
		return
	}

	cnt := 10;
	data := response.Header
	length, err := strconv.Atoi(data["Content-Length"][0]);
	avg_length := length / cnt;
	for i := 0; i < cnt; i++ {
		start := i*avg_length;
		if i == (cnt - 1) {
			end := length - start;

			log.Println("start:", start)
			log.Println("end:", end)
		} else {
			end := start + avg_length

			log.Println("start:", start)
			log.Println("end:", end)
		}
	}

	log.Println("header:", avg_length)
	log.Println("content-length:", data["Content-Length"])

}

//下载video
func save(url string, name string) {
	log.Println(url)

	req, err := http.NewRequest("GET",url,nil)
	req.Header.Add("Range:" , "0-10")
	response, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Println("get url failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", url, err)
		return
	}
	filename := "/Users/MOMO/my/go/video-get/movie/" + name + ".mp4";
	video, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}

	defer video.Close()
	video.Write(data)
	log.Println("finish")
}
