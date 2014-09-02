package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/davecheney/profile"
)

var (
	addr         = flag.String("addr", ":3000", "http service address")
	assets       = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTemplate *template.Template
)

func defaultAssetPath() string {
	p, err := build.Default.Import("sockettome/public", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	homeTemplate.Execute(w, req.Host)
}

func assetHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "public/"+req.URL.Path[1:])
}

func main() {
	if len(os.Getenv("PROFILE")) > 0 {
		cfg := profile.Config{
			CPUProfile:     true,
			MemProfile:     true,
			ProfilePath:    "tmp", // store profiles in current directory
			NoShutdownHook: false, // do not hook SIGINT
		}
		defer profile.Start(&cfg).Stop()
	}
	flag.Parse()
	homeTemplate = template.Must(template.ParseFiles(filepath.Join(*assets, "index.html")))
	game := Game{}

	go hub.run()
	go game.Run()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", websocketHandler)
	http.HandleFunc("/assets/", assetHandler)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	} else {
		log.Println("Listening on: ", *addr)
	}
}
