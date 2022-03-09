package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

type Downloader struct {
	concurrency int
}

func NewDownloader(concurrency int) *Downloader {
	return &Downloader{concurrency: concurrency}
}

// Download 自动选择下载策略
func (d *Downloader) Download(strURL, filename string) error {
	if filename == "" {
		filename = path.Base(strURL)
	}

	resp, err := http.Head(strURL)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return d.multiDownload(strURL, filename, int(resp.ContentLength))
	}

	return d.singleDownload(strURL, filename)
}

// multiDownload 多段下载
func (d *Downloader) multiDownload(strURL, filename string, contentLen int) error {
	log.Println("multiDownload")
	partSize := contentLen / d.concurrency

	// 创建部分文件的存放目录
	partDir := d.getPartDir(filename)
	err := os.Mkdir(partDir, 0777)
	if err != nil {
		return err
	}
	defer os.RemoveAll(partDir)

	var wg sync.WaitGroup
	wg.Add(d.concurrency)

	rangeStart := 0

	for i := 0; i < d.concurrency; i++ {
		// 并发请求
		go func(i, rangeStart int) {
			defer wg.Done()

			rangeEnd := rangeStart + partSize
			// 最后一部分，总长度不能超过 ContentLength
			if i == d.concurrency-1 {
				rangeEnd = contentLen - 1
			}

			d.downloadPartial(strURL, filename, rangeStart, rangeEnd, i)

		}(i, rangeStart)

		rangeStart += partSize + 1
	}

	wg.Wait()

	// 合并文件
	err = d.merge(filename)
	if err != nil {
		return err
	}

	return nil
}

// singleDownload 整体下载
func (d *Downloader) singleDownload(strURL, filename string) error {
	resp, err := http.Get(strURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(f, resp.Body, buf)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// downloadPartial 分段下载
func (d *Downloader) downloadPartial(strURL, filename string, rangeStart, rangeEnd, i int) {
	if rangeStart >= rangeEnd {
		return
	}

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	flags := os.O_CREATE | os.O_WRONLY
	partFile, err := os.OpenFile(d.getPartFilename(filename, i), flags, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer partFile.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(partFile, resp.Body, buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}

// getPartDir 部分文件存放的目录
func (d *Downloader) getPartDir(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

// getPartFilename 构造部分文件的名字
func (d *Downloader) getPartFilename(filename string, partNum int) string {
	partDir := d.getPartDir(filename)
	return fmt.Sprintf("%s/%s-%d", partDir, filename, partNum)
}

// merge 合并文件
func (d *Downloader) merge(filename string) error {
	destFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for i := 0; i < d.concurrency; i++ {
		partFileName := d.getPartFilename(filename, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(destFile, partFile)
		if err != nil {
			return err
		}
		partFile.Close()
		err = os.Remove(partFileName)
		if err != nil {
			return err
		}
	}

	return nil
}
