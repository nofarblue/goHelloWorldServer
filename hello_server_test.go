package main
import (

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

func TestGreeting_WhitespaceOnly(t *testing.T) {
	testCases := []string{
		"   ",           // spaces
		"\t",            // tab
		"\n",            // newline
		"\r",            // carriage return
		" \t\n\r ",      // mixed whitespace
	}
	
	for _, name := range testCases {
		greeting := CreateGreeting(name)
		if greeting != "Hello, Guest\n" {
			t.Errorf("Greeting for whitespace-only name '%s' was incorrect, got: %s, want: %s.", name, greeting, "Hello, Guest\n")
		}
	}
}

func TestGreeting_LongName(t *testing.T) {
	// Test name longer than 100 characters
	longName := "a"
	for i := 0; i < 120; i++ {
		longName += "a"
	}
	
	greeting := CreateGreeting(longName)
	// Should truncate to 100 characters
	expectedName := longName[:100]
	expected := "Hello, " + expectedName + "\n"
	
	if greeting != expected {
		t.Errorf("Greeting for long name was incorrect, got length %d, want length %d", len(greeting)-8-1, 100) // -8 for "Hello, " and -1 for "\n"
	}
}

func TestGreeting_NewlineInjection(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"John\nAdmin", "Hello, John Admin\n"},          // newline should be replaced with space
		{"Jane\r\nDoe", "Hello, Jane  Doe\n"},           // carriage return + newline
		{"Test\tUser", "Hello, Test User\n"},            // tab should be replaced with space
		{"User\x00", "Hello, User \n"},                  // null character
		{"Test\x01\x02", "Hello, Test  \n"},             // control characters
	}
	
	for _, tc := range testCases {
		greeting := CreateGreeting(tc.input)
		if greeting != tc.expected {
			t.Errorf("Greeting for input '%s' was incorrect, got: %s, want: %s.", tc.input, greeting, tc.expected)
		}
	}
}

func TestGreeting_SpecialSymbols(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Jane!@#", "Hello, Jane!@#\n"},
		{"User$%^&*()", "Hello, User$%^&*()\n"},
		{"Test-User_123", "Hello, Test-User_123\n"},
		{"João", "Hello, João\n"}, // Unicode characters should be preserved
	}
	
	for _, tc := range testCases {
		greeting := CreateGreeting(tc.input)
		if greeting != tc.expected {
			t.Errorf("Greeting for input '%s' was incorrect, got: %s, want: %s.", tc.input, greeting, tc.expected)
		}
	}
}

func TestGreeting_TrimWhitespace(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"  John  ", "Hello, John\n"},
		{"\tJane\t", "Hello, Jane\n"},
		{" \n Test \r ", "Hello, Test\n"},
	}
	
	for _, tc := range testCases {
		greeting := CreateGreeting(tc.input)
		if greeting != tc.expected {
			t.Errorf("Greeting for input '%s' was incorrect, got: %s, want: %s.", tc.input, greeting, tc.expected)
		}
	}
}

