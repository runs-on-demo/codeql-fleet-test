package main

import (
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Deliberate command injection for the CodeQL test.
	out, _ := exec.Command("sh", "-c", "ls "+r.URL.Query().Get("dir")).Output()
	_, _ = w.Write(out)
}

func main() {
	http.HandleFunc("/ls", handler)
	_ = http.ListenAndServe(":8080", nil)
}
