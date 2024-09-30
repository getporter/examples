package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", Say)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Say(w http.ResponseWriter, r *http.Request) {
	var msg = "Whale Hello There!"
	if defaultMsg, ok := os.LookupEnv("DEFAULT_MESSAGE"); ok {
		msg = defaultMsg
	}
	userMsg := r.FormValue("msg")
	if userMsg != "" {
		msg = userMsg
	}
	cowsay := exec.Command("cowsay", msg)

	buf := bytes.Buffer{}
	cowsay.Stdout = &buf
	err := cowsay.Start()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Oops, we couldn't get the whale started, sorry!"))
		return
	}

	err = cowsay.Wait()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Oops, the whale ran outta steam, sorry!"))
		return
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Oops, the whale is a bit tongue tied, sorry!"))
		return
	}
}
