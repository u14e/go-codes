package client

import (
	"ch14/series"
	"testing"
)


func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
}
