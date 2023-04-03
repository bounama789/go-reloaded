package main

import (
	"fmt"
	"goreloaded/lib"
	"os"
)

var cont = new(string)


func main() {

	args := os.Args

	if len(args) == 3 {
		filename := args[1]
		outfilename := args[2]
		content, _ := os.ReadFile(filename)
		*cont = string(content)
		var output string

		file, _ := os.Create(outfilename)
		erro := os.Truncate(outfilename, 0)

		if erro != nil {
			panic(erro)
		}
		output = lib.Process(*cont)
		_, err := file.WriteString(output)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("USAGE: go run . <input file> <output file>")
		return
	}
}
