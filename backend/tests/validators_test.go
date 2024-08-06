package tests

import (
	"api/pkg/middleware"
	"testing"
)

func TestIsPasswordSecure(t *testing.T) {
	got := middleware.IsSecurePassword(",jsdhf@#sdfjSDkjd")
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = middleware.IsSecurePassword("job@#$sdfbasdASDFSF1234")
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = middleware.IsSecurePassword("LLLLLL")
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = middleware.IsSecurePassword("La1!")
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = middleware.IsSecurePassword("!!!!!!aB1")
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsValidUsername(t *testing.T) {
	// Test case 1: Valid username
	got := middleware.IsUsernameValid("john!doe")
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	// Test case 2: Invalid username with special characters
	got = middleware.IsUsernameValid("john@doe")
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	// Test case 3: Invalid username with whitespace
	got = middleware.IsUsernameValid("    ")
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	// Test case 4: Invalid username with length less than 3
	got = middleware.IsUsernameValid("jd")
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = middleware.IsUsernameValid("luSkasz2003")
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = middleware.IsUsernameValid("lukasz")
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}
