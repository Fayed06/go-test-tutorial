package lesson2

import (
	"testing"
)

func Test_Dedup_Empty(t *testing.T) {
	result := Dedup([]string{})
	if len(result) != 0 {
		t.Fail()
	}
}

func Test_Dedup_ManyEntries(t *testing.T) {
	result := Dedup([]string{"a", "b", "c"})
	if len(result) != 3 {
		t.Fail()
	}
}
