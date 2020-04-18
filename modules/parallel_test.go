package modules_test

import (
	"fmt"
	"go-urlhash/modules"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMD5(t *testing.T) {

	pp := modules.NewParallel(0, []string{})

	tests := []struct {
		value    string
		md5      string
		expected bool
	}{
		{value: "my string", md5: "2ba81a47c5512d9e23c435c1f29373cb", expected: true},
		{value: "aaa", md5: "47bce5c74f589f4867dbd57e9ca9f808", expected: true},
		{value: "", md5: "", expected: true},
	}

	for i, test := range tests {
		actual := pp.GetMD5(test.value)

		if strings.Compare(test.md5, actual) != 0 && !test.expected {
			t.Errorf("Case %d, unexpected equality result: expected '%v', received '%v'", i, test.md5, actual)
		}
	}

}

func TestURL(t *testing.T) {
	pp := modules.NewParallel(0, []string{})

	tests := []struct {
		url      string
		expected bool
	}{
		{url: "http://sss.bbb.com", expected: true},
		{url: "http://bbb.com", expected: true},
		{url: "https://sss.bbb.com", expected: true},
		{url: "https://bbb.com", expected: true},
		{url: "https://sss.bbb", expected: true},
		{url: "//www.sss.bbb", expected: true},
		{url: "://sss.bbb", expected: false},
		{url: "www.sss.bbb", expected: false},
	}

	for i, test := range tests {
		actual := pp.CheckURL(test.url)
		//t.Logf("%s %v\n", test.url, actual)
		if actual != test.expected {
			t.Errorf("Case %d, unexpected equality result: expected '%v', received '%v'", i, test.expected, actual)
		}
	}

}

func TestFetchURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello there")
	}))
	defer ts.Close()

	pp := modules.NewParallel(0, []string{})

	resp, err := pp.FetchURL(ts.URL)
	if err != nil {
		t.Errorf("In this case of FetchURL, we got: %s", err.Error())
	}

	if resp == "" {
		t.Error("In this case of FetchURL, we got a blank response")
	}

}

func TestGetData(t *testing.T) {

	//TODO: improve this test
	pp := modules.NewParallel(2, []string{"https://google.com",
		"https://adjust.com", "https://yandex.com",
		"https://facebook.com", "https://twitter.com",
		"https://reddit.com/r/funny", "https://reddit.com/r/notfunny",
		"https://yahoo.com", "https://baroquemusiclibrary.com"})

	results, err := pp.GetData()
	if err != nil && len(results) == 0 {
		t.Error("Error on get data")
	}
}
