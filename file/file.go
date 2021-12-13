package file

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

const logTitle = "[file]"

func logLn(v ...interface{}) {
	log.Println(logTitle, v)
}

// IsPathExist 判断文件或文件夹路径是否存在
func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// WriteInfoToFile 将内容写入文件
func WriteInfoToFile(filePath string, content string) bool {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logLn("WriteInfoToFile OpenFile error:", err)
		return false
	}
	defer file.Close()
	_, err = io.WriteString(file, content)
	if err != nil {
		logLn("WriteInfoToFile WriteString error:", err)
		return false
	}
	return true
}

// ReadFileToSlice 将文件中的数据按行经过handleFunc处理后存入字符串切片
func ReadFileToSlice(filePath string, handleFunc func(string) (string, bool)) []string {
	f, err := os.Open(filePath)
	if err != nil {
		logLn("ReadFileToSlice open file err = [%s]", err.Error())
		return nil
	}
	defer f.Close()
	var fileList []string
	var s string
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		s = fileScanner.Text()
		if s, bOK := handleFunc(s); bOK {
			fileList = append(fileList, s)
		}
	}
	return fileList
}

// WriteSliceToFile 将切片写入文件并自动换行
func WriteSliceToFile(filePath string, contents []string) bool {
	s := strings.Join(contents, "\n")
	return WriteInfoToFile(filePath, s)
}

// RemoveFile 移除文件
func RemoveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		logLn("RemoveFile error:", err)
	}
}

// CreateFile 新建文件
func CreateFile(filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		logLn("CreateFile error:", err.Error())
	}
	f.Close()
}

// RenameFile 重命名文件
func RenameFile(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		logLn("RenameFile error:", err)
		return
	}
}

// GetFileNameWithoutSuffix 获取不含后缀的文件名称
func GetFileNameWithoutSuffix(filePath string) string {
	// 获取文件名带后缀
	filenameWithSuffix := path.Base(filePath)
	// 获取文件后缀
	fileSuffix := path.Ext(filenameWithSuffix)
	// 获取文件名
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

// ReadAll 读取文件全部内容到内存
func ReadAll(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		logLn("ReadAll error:", err)
		return nil
	}
	return data
}

// GetFileStat 获取文件基本信息
func GetFileStat(filePath string) os.FileInfo {
	fi, err := os.Stat(filePath)
	if err != nil {
		logLn("GetFileStat error:", err)
		return nil
	}
	return fi
}
