package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/example/processor/count"
	"github.com/example/processor/read"
)

func main() {
	if len(os.Args) == 2 {
		res, err := process(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(res)
	} else {
		fmt.Println("no source specified")
	}
}

func process(s string) (string, error) {
	if s[0:7] == "http://" || s[0:8] == "https://" {
		res, err := read.FromWeb(s)
		if err != nil {
			return "", err
		}
		defer res.Close()

		n, err := count.FromReader(res)
		if err != nil {
			return "", err
		}

		return strconv.Itoa(n), nil
	} else {
		res, err := read.FromFile(s)
		if err != nil {
			return "", err
		}
		defer res.Close()

		n, err := count.FromReader(res)
		if err != nil {
			return "", err
		}

		return strconv.Itoa(n), nil
	}
}
