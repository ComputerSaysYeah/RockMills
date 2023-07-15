package api

import "testing"

func TestParseSquare(t *testing.T) {
	if ParseSquare("GG").IsValid() {
		t.Fatal()
	}
}
