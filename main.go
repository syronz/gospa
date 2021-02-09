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
var confContent = `#it can use for another directory
port = "8080"
public_dir = "."
index_file = "index.html"`

// SPAHandler Serve from a public directory with specific index
type SPAHandler struct {
	Port      string `toml:"port"`
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
	var createConfig bool
	flag.StringVar(&configFile, "config", "spa.toml", "name of config file, by default it is spa.toml")
	flag.BoolVar(&createConfig, "init", false, "create the config file")
	flag.Parse()

	if createConfig {
		f, err := os.Create("spa.toml")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, errWriting := f.WriteString(confContent)
		if errWriting != nil {
			log.Fatal(err)
		}

		// err := ioutil.WriteFile("spa.toml", confContent, 0755)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		fmt.Println("spa.toml created successfully, now you can run 'gospa'")
		return
	}

	tomlContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("There is no file for config, run 'gospa -init' for create spa.toml")
		// log.Fatal(err)
		return
	}

	// var conf tomlConfig
	var spa SPAHandler
	if _, err := toml.Decode(string(tomlContent), &spa); err != nil {
		log.Fatal(err)
	}
	fmt.Println(spa)

	http.HandleFunc("/", spa.ServeHTTP)

	err = http.ListenAndServe(":"+spa.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
