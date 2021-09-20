package util

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// тесты запускать из директории utils(или поменять testDataPath)

const testDataPath = "test_data/"

func TestGrep(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		in      string
		out     string
		ok      bool
		sets    *Settings
	}{
		{
			"ok test with flags: -n -C 1 -F",
			"qww",
			"test_1_in.txt",
			"test_1_out.txt",
			true,
			NewSettings(0, 0, 1, false, false, false, true, true),
		},
		{
			"ok test with flags: -n -C 1 -F -i",
			"qww",
			"test_1_in.txt",
			"test_1_inv_out.txt",
			true,
			NewSettings(0, 0, 1, false, false, true, true, true),
		},
		{
			"ok test with flags: -c",
			"qww",
			"test_2_in.txt",
			"test_2_out.txt",
			true,
			NewSettings(0, 0, 0, true, false, false, false, false),
		},
		{
			"ok test with flags: -c -i",
			"qww",
			"test_2_in.txt",
			"test_2_ign_out.txt",
			true,
			NewSettings(0, 0, 0, true, true, false, false, false),
		},
		{
			"ok regex test with flags: -B 1",
			"y(a|b|c)",
			"test_3_in.txt",
			"test_3_out.txt",
			true,
			NewSettings(0, 1, 0, false, false, false, false, false),
		},
		{
			"ok fix test with flags: -F -A 1",
			"y(a|b|c)",
			"test_3_in.txt",
			"test_3_fix_out.txt",
			true,
			NewSettings(1, 0, 0, false, false, false, true, false),
		},
		{
			"invalid regex",
			"y(a|b|c",
			"test_3_in.txt",
			"",
			false,
			NewSettings(0, 0, 0, false, false, false, false, false),
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := Grep(fileReader(v.in), v.pattern, v.sets)
			if v.ok {
				assert.Nil(t, err)
				assert.Equal(t, fileString(v.out), res)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func fileReader(path string) io.Reader {
	reader, err := os.Open(testDataPath + path)
	if err != nil {
		panic(err)
	}
	return reader
}

func fileString(path string) string {
	reader, err := ioutil.ReadFile(testDataPath + path)
	if err != nil {
		panic(err)
	}
	return string(reader)
}
