package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unpack(t *testing.T) {
	tests := []struct {
		Name           string
		Input          string
		ExpectedResult string
		ExpectedErr    error
	}{
		{
			Name:           "Success",
			Input:          "a4bc2d5eqwe\\45qwe\\4\\5",
			ExpectedResult: "aaaabccdddddeqwe44444qwe45",
			ExpectedErr:    nil,
		},
		{
			Name:           "Only digits",
			Input:          "45",
			ExpectedResult: "",
			ExpectedErr:    ErrIncorrectString,
		},
		{
			Name:           "Empty input",
			Input:          "",
			ExpectedResult: "",
			ExpectedErr:    nil,
		},
		{
			Name:           "Uncompleted screen",
			Input:          "asdf\\",
			ExpectedResult: "",
			ExpectedErr:    ErrIncorrectString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res, err := unpack(tt.Input)
			assert.Equal(t, tt.ExpectedResult, res)
			assert.ErrorIs(t, err, tt.ExpectedErr)
		})
	}
}
