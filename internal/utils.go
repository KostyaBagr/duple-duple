// File represents some common utils for package

package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
