package main

import (
	"strconv"
	"strings"
)

// リテラルから型を取得する
func GetValtype(target string) string {
	// targetのint
	targeti, err := strconv.Atoi(target)

	// 小数点であるか
	targets := !strings.Contains(target, ".")

	if target == "true" || target == "false" {
		return "bool"
	} else if target[0:1] == "\"" && target[len(target)-1:] == "\"" {
		return "string"
	} else if targets && targeti > -1 {
		return "uint"
	} else if targets && targeti >= -2147483647 && targeti <= 2147483647 {
		return "int"
	} else if targets && err != nil {
		return "int64_t"
	} else if !targets {
		return "double"
	} else {
		return "unknown"
	}
}

// 文字リテラルのためのリテラル文字削除関数
func RemoveFirstAndLast(s string) string {
	runes := []rune(s)
	if len(runes) > 2 {
		return string(runes[1 : len(runes)-1])
	}
	return ""
}

func ReplacePathCharacter(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "_s_", "/"), "_u_", "_"), "_d_", ".")
}
