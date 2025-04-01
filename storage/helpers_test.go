package storage

import (
	"fmt"
	"testing"
	"time"
)

func TestReplaceStringWithDate(t *testing.T) {
	nowtime := time.Now()
	// Format the date
	year := fmt.Sprintf("%d", nowtime.Year())
	month := fmt.Sprintf("%02d", nowtime.Month())
	monthAbbrev := nowtime.Format("Jan")
	day := fmt.Sprintf("%02d", nowtime.Day())

	testCases := []struct {
		name     string
		test     string
		expected string
	}{{
		name:     "yyyy-MM-dd",
		test:     "filepath/filename-yyyy-MM-dd.md",
		expected: fmt.Sprintf("filepath/filename-%s-%s-%s.md", year, month, day),
	},
		{
			name:     "MMM-yyyy",
			test:     "filepath/filename-MMM-yyyy.md",
			expected: fmt.Sprintf("filepath/filename-%s-%s.md", monthAbbrev, year),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			replacedString := ReplaceStringWithDate(tc.test)
			if replacedString != tc.expected {
				t.Errorf("expected %s, received %s", tc.expected, replacedString)
			}
		})
	}
}
