package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackingString(t *testing.T) {
	tables := []struct {
		s   string
		res string
		err error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"", "", nil},
		{`qwe\\`, `qwe\`, nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
		{`q\410`, "q4444444444", nil},
		{"45", "", errors.New("invalid string")},
		{"0000", "", errors.New("invalid string")},
	}
	for _, table := range tables {
		totalS, totalErr := UnpackingString(table.s)

		assert.Equal(t, totalS, table.res)
		assert.Equal(t, totalErr, table.err)
	}
}
