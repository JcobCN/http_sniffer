package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir("C:\\"))
	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}