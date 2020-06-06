package infoCollection

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func judgeDirStatus(target, dirstr string) bool {
	path := target + dirstr
	path = strings.TrimSpace(path)
	res, err := http.Head(path)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 || res.StatusCode == 403 {
		fmt.Printf("%s \t %v \n", path, res.StatusCode)
		return true
	} else {
		if res.StatusCode == 404 {

		} else {
			fmt.Printf("%s \t %v \n", path, res.StatusCode)
		}
		return false
	}
}
func DirScan(dirs []string, target string) []string {
	fmt.Println("start dirscan", target)
	var AllDirStr []string
	AllDirStr = ReadFileDir("controllers/infoCollection/dir.dir", AllDirStr)
	for _, dirStr := range AllDirStr[:] {
		if judgeDirStatus(target, dirStr) {
			dirs = append(dirs, target+dirStr)
		}
	}
	fmt.Println("dirs: ", dirs)
	return dirs
}

func ReadFileDir(filname string, AllDirStr []string) []string {
	inputFile, inputError := os.Open(filname)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return AllDirStr
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		AllDirStr = append(AllDirStr, inputString)
		if readerError == io.EOF {
			return AllDirStr
		}
	}
}
