package helpers

import (
	"fmt"
	"time"
)

// Marshaler represent marshal interface
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// JSONTime represent custom time format
type JSONTime time.Time

// MarshalJSON represent masrshal json
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.000Z"))
	return []byte(stamp), nil
}
