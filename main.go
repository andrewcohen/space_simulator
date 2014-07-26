package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"text/template"
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
	//cfg := profile.Config{
	//CPUProfile:     true,
	//MemProfile:     true,
	//ProfilePath:    "profiles", // store profiles in current directory
	//NoShutdownHook: false,      // do not hook SIGINT
	//}
	//defer profile.Start(&cfg).Stop()
	runtime.GOMAXPROCS(4)
	flag.Parse()
	homeTemplate = template.Must(template.ParseFiles(filepath.Join(*assets, "index.html")))
	go hub.run()
	go game.run()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/assets/", assetHandler)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
