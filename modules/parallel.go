package modules

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/sync/errgroup"
)

type Parallel struct {
	NumProc  int
	ListUrls []string
	ctx      context.Context
}

type Results struct {
	URL string
	MD5 string
}

// NewParallel return a new instance
func NewParallel(numProc int, listUrls []string) *Parallel {
	return &Parallel{
		NumProc:  numProc,
		ListUrls: listUrls,
		ctx:      context.Background(),
	}
}

// GetData return all processed parallel data
func (p *Parallel) GetData() ([]Results, error) {
	results := []Results{}
	limiter := make(chan struct{}, p.NumProc)

	g, _ := errgroup.WithContext(p.ctx)

	for i, tdata := range p.ListUrls {
		i, tdata := i, tdata // https://golang.org/doc/faq#closures_and_goroutines

		g.Go(func() error {
			return func(idx int, url string) error {
				defer func() {
					<-limiter
				}()
				limiter <- struct{}{}
				if !p.CheckURL(url) {
					return errors.New("Invalid URL")
				}
				resp, err := p.FetchURL(url)
				if err != nil {
					return err
				}
				fmt.Printf("%s %s\n", url, p.GetMD5(resp))
				results = append(results, Results{URL: url, MD5: p.GetMD5(resp)})

				return nil
			}(i, tdata)
		})
	}
	if err := g.Wait(); err != nil {
		return results, err
	}
	return results, nil
}

// GetMD5 return a md5 from string
func (p Parallel) GetMD5(data string) string {
	hasher := md5.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

// FetchURL get response data
func (p Parallel) FetchURL(url string) (string, error) {

	timeout := time.Duration(60 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// CheckURL return a valid url format
func (p *Parallel) CheckURL(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	return err == nil
}
