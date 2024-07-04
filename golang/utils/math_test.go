package utils

import "testing"

func Test_Add(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Fatal("error")
	}
	t.Log("success")
}
