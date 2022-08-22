package main

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	got_slice := encode("Survivors Guilt",5)
	got := strings.Join(got_slice, " ")	
	want := "xzwanatwx lznqy" 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test2(t *testing.T) {
	got_slice := encode("Untreated Trauma",5)
	got := strings.Join(got_slice, " ")	
	want := "zsywjfyji ywfzrf" 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test3(t *testing.T) {
	got_slice := encode("Occupational Hazard",5)
	got := strings.Join(got_slice, " ")	
	want := "thhzufyntsfq mfefwi" 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test4(t *testing.T) {
	got_slice := encode("Beyond Bulletproof",5)
	got := strings.Join(got_slice, " ")	
	want := "gjdtsi gzqqjyuwttk" 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test5(t *testing.T) {
	got_slice := encode("Gangland Landlord",5)
	got := strings.Join(got_slice, " ")	
	want := "lfslqfsi qfsiqtwi" 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

