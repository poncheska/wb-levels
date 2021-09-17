package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
		err  bool
	}{
		{
			"simple unpack",
			"a4bc2d5e",
			"aaaabccddddde",
			false,
		},
		{
			"nothing to unpack",
			"abcd",
			"abcd",
			false,
		},
		{
			"zero repeats",
			"abcc0d",
			"abcd",
			false,
		},
		{
			"several numbers in a row 1",
			"45",
			"",
			true,
		},
		{
			"several numbers in a row 1",
			"qwe45",
			"",
			true,
		},
		{
			"empty input",
			"",
			"",
			false,
		},
		{
			"escape number",
			"qwe\\4\\5",
			"qwe45",
			false,
		},
		{
			"repeat escaped number",
			"qwe\\45",
			"qwe44444",
			false,
		},
		{
			"repeat escaped slash",
			"qwe\\\\5",
			"qwe\\\\\\\\\\",
			false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := Unpack(v.in)
			if v.err {
				assert.NotNil(t, err)
			}else{
				assert.Nil(t, err)
				assert.Equal(t, v.out, res)
			}
		})
	}
}
