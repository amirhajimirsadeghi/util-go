package money

import (
	"testing"
)

func TestFromCents(t *testing.T) {
	tests := []struct {
		Input  int
		Output string
	}{
		{
			Input:  -200000,
			Output: "-$2,000.00",
		},
		{
			Input:  -20000,
			Output: "-$200.00",
		},
		{
			Input:  -1000,
			Output: "-$10.00",
		},
		{
			Input:  -999,
			Output: "-$9.99",
		},
		{
			Input:  -99,
			Output: "-$0.99",
		},
		{
			Input:  0,
			Output: "$0.00",
		},
		{
			Input:  99,
			Output: "$0.99",
		},
		{
			Input:  999,
			Output: "$9.99",
		},
		{
			Input:  1000,
			Output: "$10.00",
		},
		{
			Input:  20000,
			Output: "$200.00",
		},
		{
			Input:  200000,
			Output: "$2,000.00",
		},
	}

	for _, test := range tests {
		m := FromCents(test.Input)
		out := m.String()
		if out != test.Output {
			t.Fatalf("Input: %d | Expected %s | Got %s", test.Input, test.Output, out)
		}
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
		err    bool
	}{
		{"$10", 1000, false},
		{"$10.23", 1023, false},
		{"$0.99", 99, false},
		{"$0", 0, false},
		{"$123456789", 12345678900, false},
		{"$0.1", 10, false},
		{"$0.01", 1, false},
		{"$0.009", 0, true},
		{"$10.00", 1000, false},
		{"$10.", 1000, false},
		{".$99", 99, false},
		{"-$10", -1000, false},
		{"-$0.99", -99, false},
		{"$abc", 0, true},
		{"$10.10.10", 0, true},
		{"$1,010.10", 101010, false},
		{"-$1,010.10", -101010, false},
		{"10", 1000, false},
		{"10.23", 1023, false},
		{"0.99", 99, false},
		{"0", 0, false},
		{"123456789", 12345678900, false},
		{"0.1", 10, false},
		{"0.01", 1, false},
		{"0.009", 0, true},
		{"10.00", 1000, false},
		{"10.", 1000, false},
		{".99", 99, false},
		{"-10", -1000, false},
		{"-0.99", -99, false},
		{"abc", 0, true},
		{"10.10.10", 0, true},
		{"1,010.10", 101010, false},
		{"-1,010.10", -101010, false},
	}

	for _, test := range tests {
		m, err := FromString(test.Input)
		if test.err && err == nil {
			t.Fatalf("Test case %s failed: expected an error but got none\n", test.Input)
			continue
		}
		if !test.err && err != nil {
			t.Fatalf("Test case %s failed: expected no error but got %s\n", test.Input, err)
			continue
		}

		if m.Cents() != test.Output {
			t.Fatalf("Test case %s failed: expected %d but got %d\n", test.Input, test.Output, m.Cents())
		}
	}
}
