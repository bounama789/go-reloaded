package main

import (
	"os"
	"piscine"
	"strconv"
	"strings"
)

func main() {

	const (
		hex = "0123456789ABCDEF"
		bin = "01"
	)

	args := os.Args
	options := []string{"(low)", "(up)", "(bin)", "(hex)", "(cap)"}

	if len(args) > 1 {
		filename := args[1]
		content, _ := os.ReadFile(filename)
		words := strings.Split(string(content), " ")
		var ln string

		for i, word := range words {
			if word == "\n" {
				ln = strings.Join(words[:i], " ")
				var option string

				for y, l := range ln {
					if l == '(' {
						option = getOption(ln,y)
						println(option)
					}

					if option == options[3] {
						ln = convertHex(ln,y)
					}
					
				}
			}
			print(ln)
		}

	}

}

func getOption(s string,index int) string {
	for i := index; i < len(s); i++ {
		if s[i] == ')' {
			return s[index:i]
		}
	}
	return ""
}

func convertHex(s string, idx int) string{
	base := "0123456789ABCDEF"

	var hex string
	for i := idx-2; i > 0; i-- {
		if s[i] == ' ' {
			hex = s[i:idx-1]
			s =s[:i] + strconv.Itoa(piscine.AtoiBase(hex, base))+ s[:idx-1]
		}
	}
	return s
}