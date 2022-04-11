package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "*****************************************************\n")
	fmt.Fprintf(w, "███████████  ████████████      ████████      ████████\n")
	fmt.Fprintf(w, "███████████  ███████████████   █████████    █████████\n")
	fmt.Fprintf(w, "   █████        ████   █████     ████████  ████████  \n")
	fmt.Fprintf(w, "   █████        ███████████      ████  ███ ███ ████  \n")
	fmt.Fprintf(w, "   █████        ███████████      ████  ███████ ████  \n")
	fmt.Fprintf(w, "   █████        ████   █████     ████   █████  ████  \n")
	fmt.Fprintf(w, "███████████  ███████████████   ██████    ███   ██████\n")
	fmt.Fprintf(w, "███████████  ████████████      ██████     █    ██████\n")
	fmt.Fprintf(w, "*****************************************************\n")
	fmt.Fprintf(w, "Hello World!\n")
	fmt.Fprintf(w, "And thanks for all the fish!\n")
	fmt.Fprintf(w, "And thanks for all the fish!\n")
}
