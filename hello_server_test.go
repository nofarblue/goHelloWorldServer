package main
import (
    "strings"
    "time"
    "testing"
)


func TestGreetingSpecificJohn(t *testing.T) {
	time.Sleep(8 * time.Second)
	greeting := CreateGreeting("John")
	if greeting != "Hello, John\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificDemo(t *testing.T) {
	greeting := CreateGreeting("Demo")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

/* func TestShowFailure(t *testing.T) {
	greeting := CreateGreeting("Demo1")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
} */



func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	if greeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}

// Test case for whitespace-only names
func TestGreeting_WhitespaceOnly(t *testing.T) {
	testCases := []string{
		" ",
		"   ",
		"\t",
		"\t\t",
		" \t ",
		"\n",
		"\r\n",
	}
	
	for _, input := range testCases {
		greeting := CreateGreeting(input)
		if greeting != "Hello, Guest\n" {
			t.Errorf("Greeting for whitespace input %q was incorrect, got: %s, want: %s.", input, greeting, "Hello, Guest\n")
		}
	}
}

// Test case for long names (>100 characters)
func TestGreeting_LongName(t *testing.T) {
	longName := strings.Repeat("a", 150) // 150 characters
	greeting := CreateGreeting(longName)
	expectedName := strings.Repeat("a", 100) // Should be truncated to 100
	expected := "Hello, " + expectedName + "\n"
	
	if greeting != expected {
		t.Errorf("Greeting for long name was incorrect, got length: %d, want length: %d", len(greeting), len(expected))
	}
	
	// Verify the name portion is exactly 100 characters
	nameOnly := strings.TrimPrefix(strings.TrimSuffix(greeting, "\n"), "Hello, ")
	if len(nameOnly) != 100 {
		t.Errorf("Name was not truncated properly, got length: %d, want: 100", len(nameOnly))
	}
}

// Test case for names with newlines and control characters
func TestGreeting_NewlineInjection(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"John\nAdmin", "Hello, JohnAdmin\n"}, // Newline removed
		{"Jane\rDoe", "Hello, JaneDoe\n"},     // Carriage return removed
		{"Bob\tSmith", "Hello, BobSmith\n"},   // Tab removed
		{"Alice\x00", "Hello, Alice\n"},       // Null character removed
		{"Test\x01\x02", "Hello, Test\n"},     // Control characters removed
	}
	
	for _, tc := range testCases {
		greeting := CreateGreeting(tc.input)
		if greeting != tc.expected {
			t.Errorf("Greeting for input %q was incorrect, got: %q, want: %q", tc.input, greeting, tc.expected)
		}
	}
}

// Test case for names with special symbols
func TestGreeting_SpecialSymbols(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Jane!@#", "Hello, Jane!@#\n"},
		{"User$%^", "Hello, User$%^\n"},
		{"Test&*()", "Hello, Test&*()\n"},
		{"Name-_+=", "Hello, Name-_+=\n"},
		{"User[]{}|", "Hello, User[]{}|\n"},
	}
	
	for _, tc := range testCases {
		greeting := CreateGreeting(tc.input)
		if greeting != tc.expected {
			t.Errorf("Greeting for input %q was incorrect, got: %q, want: %q", tc.input, greeting, tc.expected)
		}
	}
}

// Test case for names with leading/trailing whitespace
func TestGreeting_TrimWhitespace(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{" John ", "Hello, John\n"},
		{"\tJane\t", "Hello, Jane\n"},
		{" \n Bob \r ", "Hello, Bob\n"},
		{"  Alice  ", "Hello, Alice\n"},
	}
	
	for _, tc := range testCases {
		greeting := CreateGreeting(tc.input)
		if greeting != tc.expected {
			t.Errorf("Greeting for input %q was incorrect, got: %q, want: %q", tc.input, greeting, tc.expected)
		}
	}
}

// Test case for sanitization helper functions
func TestSanitizeControlChars(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"normal", "normal"},
		{"with\nnewline", "withnewline"},
		{"with\ttab", "withtab"},
		{"with\rcarriage", "withcarriage"},
		{"with\x00null", "withnull"},
		{"", ""},
	}
	
	for _, tc := range testCases {
		result := sanitizeControlChars(tc.input)
		if result != tc.expected {
			t.Errorf("sanitizeControlChars for input %q was incorrect, got: %q, want: %q", tc.input, result, tc.expected)
		}
	}
}

func TestSanitizeForLogging(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "<empty>"},
		{"normal", "normal"},
		{"with\nnewline", "withnewline"},
		{"with\ttab", "withtab"},
	}
	
	for _, tc := range testCases {
		result := sanitizeForLogging(tc.input)
		if result != tc.expected {
			t.Errorf("sanitizeForLogging for input %q was incorrect, got: %q, want: %q", tc.input, result, tc.expected)
		}
	}
}
 


