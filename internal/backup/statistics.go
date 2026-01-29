// Module for formatting dump statistics. Time taken, storagies, etc
package backup

import (
	"fmt"
	"time"

	"github.com/KostyaBagr/duple-duple/internal/utils"
)

type DumpFileStats struct {
	timeStart time.Time
	timeEnd   time.Time

	filePath string
	fileSize float64
	Storages []string
	Dbms     string
}

func (d *DumpFileStats) String() string {
	timeStart := d.timeStart.Format("2006-01-02 15:04:05")
	timeEnd := d.timeEnd.Format("2006-01-02 15:04:05")

	text := fmt.Sprintf(
		`
		DBMS: %s;
		Start time: %s;
		End time: %s;
		File path: %s;
		File size: %.5f mb;
		Storagies: %s; 
		`, d.Dbms, timeStart, timeEnd, d.filePath, d.fileSize, d.Storages,
	)
	return text
}

func (d *DumpFileStats) startTime() {
	d.timeStart = time.Now()
}

func (d *DumpFileStats) endTime() {
	d.timeEnd = time.Now()
}

// Compute and returns file size in MB
func (d *DumpFileStats) calcFileSize(path string) error {
	size, err := utils.FileSize(path)
	if err != nil {
		return err
	}

	d.fileSize = float64(size) / float64(1024) / float64(1024)
	return nil
}
