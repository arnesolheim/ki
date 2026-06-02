package main

import (
	"testing"

	"github.com/pocketbase/pocketbase/tools/types"
)

func mustDateTime(s string) types.DateTime {
	dt, err := types.ParseDateTime(s)
	if err != nil {
		panic(err)
	}
	return dt
}

func TestFormatDate(t *testing.T) {
	tests := []struct {
		input    types.DateTime
		expected string
	}{
		{mustDateTime("2026-04-21T10:00:00Z"), "21. april 2026"},
		{mustDateTime("2026-01-01T00:00:00Z"), "1. januar 2026"},
		{mustDateTime("2026-06-02T12:00:00Z"), "2. juni 2026"},
		{mustDateTime("2026-12-31T23:59:59Z"), "31. desember 2026"},
		{types.DateTime{}, "-"},
	}

	for _, tt := range tests {
		got := formatDate(tt.input)
		if got != tt.expected {
			t.Errorf("formatDate(%v) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}
