package main

import "strings"

func GetVars(s string) ([]string, []string, []string) {
	lines := strings.Split(s, "\n")

	key := []string{}
	val := []string{}
	typ := []string{}

	for i := 0; i < len(lines); i++ {
		codes := strings.Split(lines[i], " ")

		if codes[0] == "var" {
			key = append(key, codes[2])
			val = append(val, RemoveFirstAndLast(codes[3]))
			typ = append(typ, codes[1])
		}
	}

	return key, val, typ
}
