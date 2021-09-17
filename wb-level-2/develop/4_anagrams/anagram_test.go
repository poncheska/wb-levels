package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnagrams(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		res  map[string][]string
	}{
		{
			"all not anagrams",
			[]string{"йцу", "цук", "уке", "кен", "енг"},
			map[string][]string{},
		},
		{
			"anagrams",
			[]string{"Йцу", "уйЦ", "цук", "уке", "цйу", "кен", "енг", "уцй", "куц"},
			map[string][]string{
				"йцу": []string{"йцу", "уйц", "уцй", "цйу"},
				"цук": []string{"куц", "цук"},
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, v.res, Anagrams(v.in))
		})
	}
}
