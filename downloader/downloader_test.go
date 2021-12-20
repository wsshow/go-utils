package downloader

import (
	"runtime"
	"testing"
)

func TestNewDownloader(t *testing.T) {
	dl := NewDownloader(runtime.NumCPU())
	err := dl.Download("https://studygolang.com/dl/golang/go1.17.5.src.tar.gz", "go1.17.5.src.tar.gz")
	if err != nil {
		t.Error(err)
		return
	}
}
