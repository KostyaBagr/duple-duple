// File represents some common utils for package

package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
	"slices"
	"time"
)

// Returns current date time in RFC3339 fomat
func CurrentDateTimeRFC3339() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}

// Converts file to slice of bytes
func ConvertFileToBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("No such file %v", filePath)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("Can't get stat of file")
	}

	bfile := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bfile)
	if err != nil && err != io.EOF {
		log.Print(err)
		return nil, fmt.Errorf("Unexpected error %v", err)
	}
	return bfile, nil

}

// Checks that object is not nil
func IsEmpty(object interface{}) bool {

	switch object {
	case nil:
		return true
	case "":
		return true
	case false:
		return true
	}

	if reflect.ValueOf(object).Kind() == reflect.Struct {
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		} else {
			return false
		}
	}
	return false
}

// TODO: fix it. Now it can only create a dir
// Check that file or dir exists
func PathExists(path string, create bool) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		if create {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				fmt.Println(err)
				return false, err
			}
			return true, nil
		}
		return false, errors.New("File does not exist")
	}
	return true, nil
}

// Splits fullPath by "/" and returns the last element of slice
func GetFileNameFromPath(filePath string) (string, error) {

	_, err := PathExists(filePath, false)
	if err != nil {
		return "", errors.New("File does not exist")
	}

	res := regexp.MustCompile("/").Split(filePath, -1)
	return res[len(res)-1], nil

}

// Returns file size in bytes
func FileSize(path string) (int64, error) {
	file, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return file.Size(), nil
}

// Checks that one slice is a subslice of another
func SliceIsSubSlice[T comparable](mainSlice, subSlice []T) bool {
	for _, item := range subSlice {
		if !slices.Contains(mainSlice, item) {
			return false
		}
	}
	return true
}
