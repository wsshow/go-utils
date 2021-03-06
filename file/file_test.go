package file

import (
	"strings"
	"testing"
)

const filePath = "./test.txt"

func TestIsPathExist(t *testing.T) {
	logLn(IsPathExist("./file.go"))
	logLn(IsPathExist("./file"))
	logLn(IsPathExist("D:\\GoProjects"))
	logLn(IsPathExist("D:\\Go"))
	logLn(IsPathExist("D:\\Go\\go-utils"))
}

func TestWriteInfoToFile(t *testing.T) {
	if !WriteInfoToFile(filePath, "hello golang\n123 456 789\n,./ &*(") {
		t.Error()
	}
}

func TestReadAll(t *testing.T) {
	if bs := ReadAll(filePath); bs != nil {
		logLn(string(bs))
	} else {
		t.Error()
	}
}

func TestReadFileToSlice(t *testing.T) {
	handleFunc := func(s string) (string, bool) {
		ss := strings.Split(s, " ")
		return ss[0], true
	}
	ss := ReadFileToSlice(filePath, handleFunc)
	if ss != nil {
		logLn(ss)
	} else {
		t.Error()
	}
}

func TestWriteSliceToFile(t *testing.T) {
	if !WriteSliceToFile(filePath, []string{"123", ",./", "iop"}) {
		t.Error()
	}
}

func TestRemoveFile(t *testing.T) {
	RemoveFile(filePath)
	if IsPathExist(filePath) {
		t.Error()
	}
}

func TestCreateFile(t *testing.T) {
	CreateFile(filePath)
	if !IsPathExist(filePath) {
		t.Error()
	}
}

func TestRenameFile(t *testing.T) {
	newPath := "./testNew.txt"
	RenameFile(filePath, newPath)
	if !IsPathExist(newPath) {
		t.Error()
	}
}

func TestGetFileNameWithoutSuffix(t *testing.T) {
	CreateFile(filePath)
	s := GetFileNameWithoutSuffix(filePath)
	if s != "test" {
		t.Error()
	}
}

func TestGetFileStat(t *testing.T) {
	stat := GetFileStat(filePath)
	if stat != nil {
		logLn(stat.Name(), stat.Size(), stat.IsDir(), stat.Mode(), stat.ModTime(), stat.Sys())
	} else {
		t.Error()
	}
}

func TestGetPrevDir(t *testing.T) {
	ex := "/dev/sda/123/"
	s := GetPrevDir(ex)
	logLn(s)
}
