package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func Generate(args string, filename string) string {
	src, err := os.ReadFile("asciiart/" + filename + ".txt")
	if err != nil {
		fmt.Println("Bad request", err.Error())
		os.Exit(1)
	}

	var ans string
	data := strings.Split(args, "\\n")
	strarr := strings.Split(string(src), "\n")

	for k := range data {
		if data[k] == "" {
			ans = ans + "\n"
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(data[k]); j++ {
				num := data[k][j] - 32
				line := int(num)*9 + i
				ans = ans + strings.TrimSuffix(strarr[line], "\r")
			}
			ans = ans + "\n"
		}
	}
	return ans
}
