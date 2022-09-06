package utils

import (
	"bufio"
	"log"
	"os"
)

func FoldExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil && f.IsDir() {
		return true, nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Printf("mkdir failed![%v]\n", err)
		} else {
			return true, nil
		}
	}
	return false, err
}

func FileExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil && f.Size() > 0 {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FileReadN(file *os.File, N int64) (string, error) {
	defer file.Close()
	brd := bufio.NewReader(file)
	buf := make([]byte, N)
	_, bufErr := brd.Read(buf)
	if bufErr != nil {
		log.Println("Error when read file: ", bufErr)
		return "", bufErr
	}
	return string(buf), nil
}

func DeleteCodes(paths ...string) {
	for _, path := range paths {
		if path != "" {
			os.Remove(path)
		}
	}
}
