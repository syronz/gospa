package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

/*
reload chrome
xdotool search --onlyvisible --class chromium-browser windowfocus key ctrl+r
*/

// SPAHandler Serve from a public directory with specific index
type SPAHandler struct {
	PublicDir string `toml:"public_dir"` // The directory from which to serve
	IndexFile string `toml:"index_file"` // The fallback/default file to serve
}

// Falls back to a supplied index (IndexFile) when either condition is true:
// (1) Request (file) path is not found
// (2) Request path is a directory
// Otherwise serves the requested file.
func (h *SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(h.PublicDir, filepath.Clean(r.URL.Path))

	if info, err := os.Stat(p); err != nil {
		http.ServeFile(w, r, filepath.Join(h.PublicDir, h.IndexFile))
		return
	} else if info.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.PublicDir, h.IndexFile))
		return
	}

	http.ServeFile(w, r, p)
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "spa.toml", "name of config file, by default it is spa.toml")
	flag.Parse()

	tomlContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// var conf tomlConfig
	var spa SPAHandler
	if _, err := toml.Decode(string(tomlContent), &spa); err != nil {
		log.Fatal(err)
	}
	fmt.Println(spa)

	http.HandleFunc("/", spa.ServeHTTP)

	http.ListenAndServe(":8090", nil)
}
