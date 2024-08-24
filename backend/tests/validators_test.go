package tests

// import (
// 	"api/pkg/validators"
// 	"testing"
// )

// func TestIsPasswordSecure(t *testing.T) {
// 	got := validators.IsSecurePassword(",jsdhf@#sdfjSDkjd")
// 	want := false

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	got = validators.IsSecurePassword("job@#$sdfbasdASDFSF1234")
// 	want = true

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	got = validators.IsSecurePassword("LLLLLL")
// 	want = false

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	got = validators.IsSecurePassword("La1!")
// 	want = false

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	got = validators.IsSecurePassword("!!!!!!aB1")
// 	want = true

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	got = validators.IsSecurePassword("")
// 	want = false

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}
// }

// func TestIsValidUsername(t *testing.T) {
// 	// Test case 1: Valid username
// 	got := validators.IsUsernameValid("john!doe")
// 	want := false
// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	// Test case 2: Invalid username with special characters
// 	got = validators.IsUsernameValid("john@doe")
// 	want = false
// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	// Test case 3: Invalid username with whitespace
// 	got = validators.IsUsernameValid("    ")
// 	want = false
// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	// Test case 4: Invalid username with length less than 3
// 	got = validators.IsUsernameValid("jd")
// 	want = false
// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	got = validators.IsUsernameValid("luSkasz2003")
// 	want = true

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// 	got = validators.IsUsernameValid("lukasz")
// 	want = true

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}

// }
