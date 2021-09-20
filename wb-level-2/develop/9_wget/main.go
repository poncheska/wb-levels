package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

var urlPattern = "(href=\"(.*?)\"|src=\"(.*?)\")"

func main() {
	out := flag.String("o", "./out", "output directory")
	url := os.Args[len(os.Args)-1]

	err := Wget(url, *out)
	if err != nil {
		fmt.Println("wget: ", err.Error())
	}
}

//Wget ...
func Wget(u, outDir string) error {
	pu, err := url.Parse(u)
	if err != nil {
		return err
	}

	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	site, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outDir+"/index.html", site, 0666)
	if err != nil {
		return err
	}

	rx := regexp.MustCompile(urlPattern)

	us := rx.FindAllSubmatch(site, -1)
	var relPath string
	chr := NewUniqueChecker()
	counter := 0
	for _, v := range us {
		if len(v[2]) != 0 {
			if v[2][0] == '/' {
				pp := url.URL{
					Scheme: pu.Scheme,
					Host:   pu.Host,
					Path:   string(v[2]),
				}
				relPath = pp.String()
			} else {
				relPath = string(v[2])
			}
		}
		if !chr.Check(relPath) {
			continue
		}
		counter++
		fmt.Printf("download in file%v.txt from %v \n", counter, relPath)
		err = downloadFile(fmt.Sprintf("%v%v%v%v", outDir, "/file", counter, ".txt"), relPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//UniqueChecker ...
type UniqueChecker struct {
	m map[string]struct{}
}

//NewUniqueChecker ...
func NewUniqueChecker() *UniqueChecker {
	return &UniqueChecker{m: make(map[string]struct{})}
}

//Check ...
func (c *UniqueChecker) Check(s string) bool {
	if _, ok := c.m[s]; ok {
		return false
	}
	c.m[s] = struct{}{}
	return true
}
