package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"slices"
)

// Aquaで実行する
func aqua(target string) string {
	aq := exec.Command("aqua", target, "--yes ")

	var aqout bytes.Buffer
	aq.Stdout = &aqout

	err := aq.Run()

	if err != nil {
		fmt.Println("[!ERR] " + aqout.String())
		os.Exit(1)
	}

	return regexp.MustCompile(`\x1b\[[0-9;]*m`).ReplaceAllString(aqout.String(), "")
}

func serve(prod bool, host string, port string) {
	// Read Config File
	routings_pattern := []string{}
	routings_path := []string{}

	rf, err := os.ReadFile(".aquarium")

	if err != nil {
		fmt.Println(err.Error())
	}

	ckeys, cvals, ctypes := GetVars(string(rf))
	for i := 0; i < len(ckeys); i++ {
		if ctypes[i] == "routing" {
			routings_path = append(routings_path, RemoveFirstAndLast(cvals[i]))
			routings_pattern = append(routings_pattern, ReplacePathCharacter(ckeys[i]))
			fmt.Println("[ LOG] Registered Routings: " + ckeys[i] + " : " + cvals[i])
		}
	}

	// Start Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(routings_pattern, r.RequestURI) && r.Method == "GET" {
			raddress := slices.Index(routings_pattern, r.RequestURI)
			bytes := []byte(aqua(path.Join("pages", routings_path[raddress])))
			_, err := w.Write(bytes)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("[ 500] " + r.RequestURI + " : " + r.Method)
			} else {
				fmt.Println("[ 200] " + r.RequestURI + " : " + r.Method)
			}
		} else if slices.Contains(routings_pattern, "404") {
			raddress := slices.Index(routings_pattern, "404")
			bytes := []byte(aqua(path.Join("pages", routings_path[raddress])))
			_, err := w.Write(bytes)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("[ 500] " + r.RequestURI + " : " + r.Method)
			} else {
				fmt.Println("[ 404] " + r.RequestURI + " : " + r.Method)
			}
		} else {
			bytes := []byte("404 Not Found")
			_, err := w.Write(bytes)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("[ 500] " + r.RequestURI + " : " + r.Method)
			} else {
				fmt.Println("[ 404] " + r.RequestURI + " : " + r.Method)
			}
		}
	})
	fmt.Println("[INFO] Server started on " + host + ":" + port + " !")
	fmt.Println(http.ListenAndServe(host+":"+port, nil))
}

func build() {
	fmt.Println("This function is under developping.")
}

func main() {
	if len(os.Args) == 1 {
		serve(false, "localhost", "8000")
	} else {
		switch os.Args[1] {
		case "prod":
		case "serve":
			serve(true, os.Args[2], os.Args[3])
			return
		case "dev":
			serve(false, os.Args[2], os.Args[3])
			return
		case "build":
			build()
			return
		default:
			fmt.Println("Unknown command, please read documents.")
			return
		}
	}
}
