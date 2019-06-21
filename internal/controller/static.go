package controller

import (
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type static struct {
	prefix string
}

func (c static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := c.prefix + r.URL.Path
	modTime := time.Now().UTC()
	stat, err := os.Stat(path)
	if err != nil {
		http.Error(w, "Not Found : "+r.URL.Path, 404)
		return
	}
	modTime = stat.ModTime()
	// 304
	modHdr := r.Header.Get("If-Modified-Since")
	w.Header().Set("Last-Modified", modTime.UTC().Format(time.RFC1123))
	hdrModTime, err := time.Parse(time.RFC1123, modHdr)
	if err == nil && modHdr != "" && modTime.Unix() <= hdrModTime.Unix() {
		w.WriteHeader(304)
		return
	}
	// Serve file
	file := []byte("")
	ext := filepath.Ext(path)
	mimetype := mime.TypeByExtension(ext)
	file, err = ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if mimetype == "" {
		mimetype = http.DetectContentType(file)
	}
	w.Header().Set("Content-Type", mimetype)
	w.Write(file)
}
