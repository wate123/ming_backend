package util

import (
	"reflect"
	"testing"
	"time"
)

func Test_dateParser(t *testing.T) {
	type args struct {
		timeString string
	}
	tests := []struct {
		name  string
		args  args
		wantT time.Time
	}{
		// TODO: Add test cases.
		{"testCase1", args{"2016-10-09T00:00:00Z"},
			time.Date(2016, 10, 9, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := dateParser(tt.args.timeString); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("dateParser() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
