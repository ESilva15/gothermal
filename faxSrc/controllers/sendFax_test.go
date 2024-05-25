package controllers

import (
	"testing"
	"unicode/utf8"
)

func TestSignMessage_NormalUserName(t *testing.T) {
	// Arrange
	user := "Alice"
	msg := "Hello, world!"
	tm := "2024-05-31 15:02:23"

	expected := `
+----------------------------------------------+
| User: Alice                                  |
| Time: 2024-05-31 15:02:23                    |
+----------------------------------------------+`
	truncation := "\n\n+- END ---------------------------------- END -+"
	expectedOutput := expected + "\n" + msg + truncation

	// Act
	result := signMessage(msg, user, tm)

	// Assert
	if result != expectedOutput {
		t.Errorf("Expected:\n%v\nGot:\n%v", expectedOutput, result)
	}
}

func TestSignMessage_VeryLongUserName(t *testing.T) {
	// Arrange

	user := "Alice Cuja Mãe Deu Um Nome Demasiado Longo"
	msg := "Hello, world!"
	tm := "2024-05-31 15:02:23"

	expected := `
+----------------------------------------------+
| User: Alice Cuja Mãe Deu Um Nome Demasiado...|
| Time: 2024-05-31 15:02:23                    |
+----------------------------------------------+`
	truncation := "\n\n+- END ---------------------------------- END -+"
	expectedOutput := expected + "\n" + msg + truncation

	// Act
	result := signMessage(msg, user, tm)

	// Assert
	if result != expectedOutput {
		t.Errorf("Expected:\n%v\nGot:\n%v", expectedOutput, result)
	}
}

func TestSignMessage_UserNameMatchesAvailableSpace(t *testing.T) {
	// Arrange

	user := "Alice Cuja Mãe Deu Um Nome Demasiado Longo"
	msg := "Hello, world!"
	tm := "2024-05-31 15:02:23"
	// Alice Cuja Mãe Deu Um Nome Demasiado...
	// umnomeextrasuperduperlongoesperoquei...
	expected := `
+----------------------------------------------+
| User: Alice Cuja Mãe Deu Um Nome Demasiado...|
| Time: 2024-05-31 15:02:23                    |
+----------------------------------------------+`
	truncation := "\n\n+- END ---------------------------------- END -+"
	expectedOutput := expected + "\n" + msg + truncation

	// Act
	result := signMessage(msg, user, tm)

	// Assert
	if result != expectedOutput {
		t.Errorf("Expected:\n%v\nGot:\n%v", expectedOutput, result)
	}
}

func TestTruncateString_StringLargerThenLim(t *testing.T) {
	// Arrange
	user := "umnomeextrasuperduperlongoesperoqueistochegue"
	// Alice Cuja Mãe Deu Um Nome Demasiado...
	// umnomeextrasuperduperlongoesperoquei...

	expected := "umnomeextrasuperduperlongoesperoquei..."

	// Act
	result := truncateString(user, 39)

	// Assert
	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}

func TestTruncateString_StringSmallerThenLim(t *testing.T) {
	// Arrange
	user := "umnome"
	// Alice Cuja Mãe Deu Um Nome Demasiado...
	// umnomeextrasuperduperlongoesperoquei...

	expected := "umnome                                 "

	// Act
	result := truncateString(user, 39)

	// Assert
	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}

func TestTruncateString_StringEqualsThenLim(t *testing.T) {
	// Arrange
	user := "umnomeextrasuperduperlongoesperoqueisto"
	// Alice Cuja Mãe Deu Um Nome Demasiado...
	// umnomeextrasuperduperlongoesperoquei...

	expected := "umnomeextrasuperduperlongoesperoqueisto"

	// Act
	result := truncateString(user, 39)

	// Assert
	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n[%v]", expected, result)
	}
}

// This is the right method to use to get the number of characters
// in a string
func TestRuneCount(t *testing.T) {
	// Arrange
	str := "Mãe"
	expected := 3

	// Act
	result := utf8.RuneCountInString(str)

	// Assert
	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}

func TestSubstr_Correct(t *testing.T) {
	// Arrange
	str := "Mãe"
	expected := "M"

	// Act
	result := substr(str, 0, 1)

	// Assert
	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}
