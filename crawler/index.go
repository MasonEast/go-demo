package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	imgUrlChan chan string
	// 使用等待组完成多个任务的同步，等待组可以保证在并发环境中完成指定数量的任务
	wg sync.WaitGroup
	taskChan chan string
	imgRegx = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	filePath = "/Users/manbang/Desktop/go/img"
)

func HandleError(err error, where string) {
	if err != nil {
			fmt.Println(where, err)
	}
}

// 任务统计协程
func CheckOk() {
	var count int
	for {
		url := <- taskChan
		fmt.Printf("%s 完成了爬取任务，目前共完成任务数：%d\n", url, count)
		count++

		if count == 20 {
			// 关闭通道，关闭后依然可以让没读取数据的通道读取数据，关闭主要防止死循环
			close(imgUrlChan)
			break
		}
	}

	close(taskChan)
	wg.Done()
}

func main() {
	imgUrlChan = make(chan string, 1000)
	taskChan = make(chan string, 10)
	// 1. 找到目标网站的图片下载url
	for i := 1; i < 11; i++ {
		wg.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}

	// 2. 统计爬取图片下载url任务，任务完成后关闭管道
	wg.Add(1)
	go CheckOk()

	// 3. 通过第一步的url下载图片到本地
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}

func getImgUrls(url string) {

	var urls []string

	resp, err := http.Get(url)
	HandleError(err, "http.get getImgUrls")

	defer resp.Body.Close()

	// 获取网页body内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll getImgUrls")

	pageStr := string(pageBytes)

	// 从网页内容中匹配出图片url
	reg := regexp.MustCompile(imgRegx)
	data := reg.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d张图片url\n", len(data))

	// 将图片url加入到切片中
	for _, v := range data {
		url := v[0]
		urls = append(urls, url)
	}

	// 通过通道将 查图片url协程中的url传给下载协程
	for _, url := range urls {
		imgUrlChan <- url
	}

	// 将url传给任务统计协程
	taskChan <- url

	wg.Done()

}

func DownloadImg() {
	// 通过range循环读出通道中的url
	for url := range imgUrlChan {
		filename := GetFilenameFromUrl(url)

		resp, err := http.Get(url)
		HandleError(err, "http.get DownloadImg")
		defer resp.Body.Close()

		bytes, err := ioutil.ReadAll(resp.Body)
		HandleError(err, "ioutil.ReadAll DownloadImg")

		path := filePath + "/" + filename

	  //向指定位置写文件，0666是指文件的权限是具有读写权限
		err = ioutil.WriteFile(path, bytes, 0666)

		if err != nil {
			fmt.Printf("%s 下载失败\n", filename)
		} else {
			fmt.Printf("图片下载成功, 存放路径：%s\n", path)
		}

	}
}

func GetFilenameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex + 1:]

  // 使用时间戳，避免命名重复
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

