package controllers

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"gin/controllers/pkg/collector"
	"gin/controllers/pkg/config"
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
	"time"
)

func ColUrl(url string) []string{
	AllUrl := make([]string, 1)
	conf := config.NewConfig()
	// define and parse command line flags
	conf.Url = url
	conf.Depth = 5
	conf.Outdir = ""
	conf.Cookie = ""
	conf.AuthHeader = ""
	conf.Scope = "subs"
	conf.Wayback = false
	conf.Plain = false
	conf.Runlinkfinder = false
	conf.IncludeAll = true
	// which data to include in output?

	// if -v is given, just display version number and exit
	if conf.DisplayVersion {
		fmt.Println(conf.Version)
		os.Exit(1)
	}

	// set up the bools
	if conf.IncludeJS || conf.IncludeSubs || conf.IncludeURLs || conf.IncludeForms || conf.IncludeRobots || conf.IncludeSitemap {
		conf.IncludeAll = false
	}

	au := aurora.NewAurora(!conf.Plain)


	stdout := bufio.NewWriter(os.Stdout)

	c := collector.NewCollector(&conf,au, stdout)

	urls := make(chan string, 1)
	var reqsMade []*http.Request
	var crawlErr error
	var wg sync.WaitGroup

	if conf.Url != "" {
		urls <- conf.Url
		close(urls)
	} else {
		ch := readStdin()
		go func() {
			//translate stdin channel to domains channel
			for u := range ch {
				urls <- u
			}
			close(urls)
		}()
	}

	// flush to stdout periodically
	t := time.NewTicker(time.Millisecond * 500)
	defer t.Stop()
	go func() {
		for {
			select {
			case <-t.C:
				stdout.Flush()
			}
		}
	}()

	for u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			// url set but does not include schema
			if !strings.Contains(url, "://") && url != "" {
				url = "http://" + url
			}
			reqsMade, crawlErr = c.Crawl(url)

			// Report errors and flush requests to files as we go
			if crawlErr != nil {
				writeErrAndFlush(stdout, crawlErr.Error(), au)
			}
			if conf.Outdir != "" {
				_, err := os.Stat(conf.Outdir)
				if os.IsNotExist(err) {
					errDir := os.MkdirAll(conf.Outdir, 0755)
					if errDir != nil {
						writeErrAndFlush(stdout, errDir.Error(), au)
					}
				}

				err = printRequestsToRandomFiles(reqsMade, conf.Outdir)
				if err != nil {
					writeErrAndFlush(stdout, err.Error(), au)
				}
			}

		}(u)
	}

	wg.Wait()

	// just in case anything is still in buffer
	stdout.Flush()

	Allurl := collector.ReturnAllUrl()
	fmt.Println("此时AllUrl的值",AllUrl)
	collector.Destory()
	fmt.Println("销毁后AllUrl的值",AllUrl)
	defer func() {
		AllUrl = nil
		fmt.Println(collector.ReturnAllUrl())
		fmt.Println("defer after...:",AllUrl)
	}()

	return Allurl
}

func readStdin() <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			url := strings.ToLower(sc.Text())
			if url != "" {
				lines <- url
			}
		}
	}()
	return lines
}

func writeErrAndFlush(w *bufio.Writer, msg string, au aurora.Aurora) {
	s := fmt.Sprintln(au.BrightRed("[error]"), msg)
	w.Write([]byte(s))
	w.Flush()
}

func printRequestsToRandomFiles(rs []*http.Request, dir string) error {
	for _, r := range rs {
		if r == nil {
			// Skip requests that were malformed
			continue
		}
		raw, err := httputil.DumpRequest(r, true)
		if err != nil {
			return err
		}

		uuid, _ := uuid.NewRandom()
		if dir[len(dir)-1:] != "/" {
			dir = dir + "/"
		}

		err = ioutil.WriteFile(dir+"hakrawler_"+uuid.String()+".req", []byte(raw), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
