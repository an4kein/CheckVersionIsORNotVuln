package main


// Author @anakein

import (
	"fmt"
	"strconv"
)

func CheckVersionVuln() {
	sessao1 := "retorno de uma funcao"
	cookie := 0 != 1337

	compare, _ := strconv.ParseInt(sessao1, 10, 64)
	newcompare := 0 != compare

	switch auto := true; auto {
	case auto == true:
		if !(cookie == false && newcompare == false) {
			fmt.Println("[+] Versao vulnerable")

		} else {
			fmt.Println("[-] not vulnerable")
		}
	case auto == true:
		fmt.Println("Error or (continue), SEI LAAAAH")
	}

	return
}

func main() {
	CheckVersionVuln()
}
