package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	filepath := os.Getenv("FILE_SERVER_CONFIG_PATH")
	if filepath == "" {
		panic("filepath is empty")
	}
	_, err := os.Stat(filepath)
	if err != nil {
		panic(fmt.Errorf("read file info error: %v", err))
	}
	passBytes := make([]byte, 64)
	_, err = io.ReadFull(rand.Reader, passBytes)
	if err != nil {
		panic(err)
	}
	password := base64.RawURLEncoding.EncodeToString(passBytes)
	fmt.Println("generate new password", password)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /clash/{id}", func(w http.ResponseWriter, r *http.Request) {
		pass := r.PathValue("id")
		if pass == "" || pass != password {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		out, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Println("read file error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(out)
		if err != nil {
			fmt.Println("write body error", err)
		}
	})
	port := ":8080"
	fmt.Printf("starting the http server, port: %s\n", port)
	panic(http.ListenAndServe(port, mux))
}
