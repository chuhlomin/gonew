package main

import "testing"

func TestFormatOutput(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"", ""},
		{"text", "result=text"},
		{"%", "result=%"},
		{"some\ntext", "result<<OUTPUT\nsome\ntext\nOUTPUT"},
		{"\n", "result<<OUTPUT\n\n\nOUTPUT"},
		{"\r", "result=\r"},
	}

	for _, tt := range tests {
		actual := formatOutput("result", tt.in)
		if actual != tt.expected {
			t.Errorf("formatOutput(%q) got: %q, want: %q.", tt.in, actual, tt.expected)
		}
	}
}

func TestDoSomething(t *testing.T) {
	tests := []struct {
		key      string
		expected string
	}{
		{"", ""},
	}

	for _, tt := range tests {
		actual, err := doSomething(tt.key)
		if err != nil {
			t.Errorf("doSomething(%q) got error: %q, want: %q.", tt.key, err, tt.expected)
		}
		if actual != tt.expected {
			t.Errorf("doSomething(%q) got: %q, want: %q.", tt.key, actual, tt.expected)
		}
	}
}
