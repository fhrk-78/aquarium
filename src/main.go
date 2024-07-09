package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"slices"
)

// ファイルを修正する
func writetmp(target string) {
	rfs, errr := os.ReadFile(target)
	if errr != nil {
		fmt.Println("[!ERR] Read: " + errr.Error())
		return
	}

	wfs, erro := os.Create("tmp.aqua")
	if erro != nil {
		fmt.Println("[!ERR] Open: " + erro.Error())
		return
	}
	_, errw := wfs.Write([]byte(FileNewlineCharConvert(string(rfs))))
	if errw != nil {
		fmt.Println("[!ERR] Write: " + errw.Error())
		return
	}
	defer func() {
		errc := wfs.Close()
		if errc != nil {
			fmt.Println("[!ERR] Close: " + errc.Error())
		}
	}()
}

// Aquaで実行する
func aqua() string {
	abp, err := filepath.Abs(`tmp.aqua`)
	if err != nil {
		fmt.Println("[!ERR] Path: " + err.Error())
	} else {
		fmt.Println("[INFO] aqua " + "\"" + abp + "\" --yes")
	}

	aq := exec.Command("powershell", "aqua", "\""+abp+"\"", "--yes")

	var aqout bytes.Buffer
	var aqerr bytes.Buffer
	aq.Stdout = &aqout
	aq.Stderr = &aqerr

	aq.Start()
	err = aq.Wait()

	if err != nil {
		fmt.Println("[!ERR] Aqua: " + aqerr.String())
	}

	return regexp.MustCompile(`\x1b\[[0-9;]*m`).ReplaceAllString(aqout.String(), "")
}

func serve(host string, port string) {
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
			abp, err := filepath.Abs(path.Join("pages", RemoveFirstAndLast(cvals[i])))
			if err != nil {
				fmt.Println("[!ERR] " + err.Error())
			}
			if err != nil {
				fmt.Println("[!ERR] " + abp + " : " + err.Error())
			}
		}
	}

	// Start Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(routings_pattern, r.RequestURI) && r.Method == "GET" {
			raddress := slices.Index(routings_pattern, r.RequestURI)
			writetmp(path.Join("pages", routings_path[raddress]))
			bytes := []byte(aqua())
			_, err := w.Write(bytes)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("[ 500] " + r.RequestURI + " : " + r.Method)
			} else {
				fmt.Println("[ 200] " + r.RequestURI + " : " + r.Method)
			}
		} else if slices.Contains(routings_pattern, "404") {
			raddress := slices.Index(routings_pattern, "404")
			writetmp(path.Join("pages", routings_path[raddress]))
			bytes := []byte(aqua())
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

	http.ListenAndServe(host+":"+port, nil)
}

func build() {
	fmt.Println("This function is under developping.")
}

func main() {
	if len(os.Args) == 1 {
		serve("localhost", "8000")
	} else {
		switch os.Args[1] {
		case "serve":
		case "dev":
			serve(os.Args[2], os.Args[3])
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
