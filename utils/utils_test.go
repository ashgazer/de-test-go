package utils

import (
	"reflect"
	"testing"
)


func TestSomeShit (t *testing.T) {
	want := []int{9, 26, 31, 32, 38, 41, 46, 53, 56, 58, 79}
	got := GetFiles("../data")

	if assert := reflect.DeepEqual(want, got);  !assert {
		t.Errorf("got %v want %v", got, want)
	}
}