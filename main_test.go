package main

import "testing"

func TestAlwaysFail(t *testing.T) {
	t.Fatal("testing out CI")
}
