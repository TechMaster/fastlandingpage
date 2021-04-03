package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	fs := &fasthttp.FS{
		// Path to directory to serve.
		Root: "./",

		// Generate index pages if client requests directory contents.
		GenerateIndexPages: false,

		IndexNames: []string{"index.html", "index.htm"},

		// Enable transparent compression to save network traffic.
		Compress: true,
	}

	// Create request handler for serving static files.
	h := fs.NewRequestHandler()

	// Start the server.
	if err := fasthttp.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
