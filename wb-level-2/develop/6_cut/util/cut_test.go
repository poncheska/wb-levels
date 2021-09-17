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
		in      string
		out     string
		ok      bool
		sets    *Settings
	}{
		{
			"ok test",
			"test_1_in.txt",
			"test_1_out.txt",
			true,
			NewSettings("1","   ",false),
		},
		{
			"only separated",
			"test_1_in.txt",
			"test_1_seponly_out.txt",
			true,
			NewSettings("1","   ",true),
		},
		{
			"custom delimiter",
			"test_2_in.txt",
			"test_2_out.txt",
			true,
			NewSettings("0,1",",",false),
		},
		{
			"invalid fields",
			"test_2_in.txt",
			"",
			false,
			NewSettings("0,a",",",false),
		},
	}


	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := Cut(fileReader(v.in), v.sets)
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