package main

import (
	"testing"
)

func TestLabelUpperCase(t *testing.T) {
	testCases := []struct {
		name          string
		expectSuccess bool
	}{
		{
			name:          "existing pod",
			expectSuccess: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			labelValue, err := uppercasePodLabel()
			if err != nil && test.expectSuccess {
				t.Fatalf("uxpected %v", err)
			} else if err == nil && !test.expectSuccess {
				t.Fatalf("expect")
			} else if labelValue != "bob" && !test.expectSuccess {
				t.Fatalf("label %v", labelValue)
			}
		})
	}
}
