package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str " + s)

	if i, err := strconv.Atoi(s); err == nil {
		t.Log(10 + i)
	}
}