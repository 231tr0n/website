package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("build"))
	fmt.Println("Starting fileserver and serving 'build' folder on port 8080.")
	err := http.ListenAndServe(":8080", fs)
	if err != nil {
		fmt.Println(err)
	}
}
