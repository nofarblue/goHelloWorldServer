package main
import (
    "net/http"
    "net/http/httptest"
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

func TestHealthzEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthzHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := "OK\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
 


