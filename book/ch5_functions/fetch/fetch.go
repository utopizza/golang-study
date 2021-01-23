package ch5

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	// get response by url
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	// get filename by url path
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	// create local file by filename
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	// copy response to file
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}
