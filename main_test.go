package main

import "testing"

func TestWorkshopTitle(t *testing.T) {
	got := workshopTitle()
	if got != "Get Up and Go" {
		t.Errorf("want Get Up and Go, got %s", got)
	}
}

func BenchmarkWorkshopTitle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		workshopTitle()
	}
}
