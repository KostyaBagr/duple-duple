// File represents some common utils for package

package internal

import "time"

// Returns current date time in RFC3339 fomat
func CurrentDateTimeRFC3339() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}
