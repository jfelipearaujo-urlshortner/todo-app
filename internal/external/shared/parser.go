package shared

import (
	"time"
)

func ParseStringToTime(layout string, input *string) (*time.Time, error) {
	if input == nil {
		return nil, nil
	}

	t, err := time.Parse(layout, *input)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func FormatTime(layout string, time *time.Time) string {
	var output string

	if time != nil {
		output = time.Format(layout)
	}

	return output
}
