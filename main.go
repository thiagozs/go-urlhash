package main

import (
	"flag"
	"fmt"
	"go-urlhash/modules"
	"os"
	"strings"
)

type arrStr []string

func (ar *arrStr) String() string {
	return fmt.Sprint(*ar)
}

func (ar *arrStr) Set(value string) error {
	rr := strings.Split(value, " ")
	for _, v := range rr {
		*ar = append(*ar, strings.TrimSpace(v))
	}
	return nil
}

func init() {
	flag.Var(&urlArr, "urls", "Array of urls for process, you can use with multiple args with -urls x -urls y -urls z OR all urls inside a double quotes, ex: \"http://url.com https://niceurl.com ...\"")
}

var (
	urlArr    arrStr
	totalProc = flag.Int("parallel", 10, "Total of number parallel process")
	verApp    = flag.Bool("version", false, "prints current program version")

	//nolinter check
	version = "0.0.0"
)

func main() {

	flag.Parse()

	if *verApp {
		fmt.Println("Version : ", version)
		os.Exit(0)
	}

	for _, k := range flag.Args() {
		if flag.NFlag() == 0 {
			if !strings.HasPrefix(k, "-") {
				urlArr = append(urlArr, k)
			}
		} else {
			urlArr = append(urlArr, k)
		}
	}

	pp := modules.NewParallel(*totalProc, urlArr)
	_, err := pp.GetData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
