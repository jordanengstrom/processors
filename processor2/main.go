package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/abadojack/whatlanggo"
	"github.com/example/processor/count"
	"github.com/example/processor/read"
)

func main() {
	if len(os.Args) == 3 {
		countOrLang := os.Args[1]
		resourcePath := os.Args[2]

		res, err := process(countOrLang, resourcePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(res)
	} else {
		fmt.Println("expected command and resource")
		os.Exit(1)
	}
}

func process(cmd, rp string) (string, error) {
	if rp[0:7] == "http://" || rp[0:8] == "https://" {
		res, err := read.FromWeb(rp)
		if err != nil {
			return "", err
		}
		defer res.Close()

		if cmd == "count" {
			n, err := count.FromReader(res)
			if err != nil {
				return "", err
			}
			return strconv.Itoa(n), nil
		} else if cmd == "lang" {
			l, err := detect(res)

			if err != nil {
				return "", err
			}

			return l, nil
		} else {
			return "", errors.New("unknown command")
		}

	} else {
		res, err := read.FromFile(rp)
		if err != nil {
			return "", err
		}
		defer res.Close()

		if cmd == "count" {
			n, err := count.FromReader(res)
			if err != nil {
				return "", err
			}

			return strconv.Itoa(n), nil
		} else if cmd == "lang" {
			l, err := detect(res)

			if err != nil {
				return "", err
			}

			return l, nil
		} else {
			return "", errors.New("unknown command")
		}
	}
}

func detect(r io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return "", err
	}

	rStr := buf.String()
	info := whatlanggo.Detect(rStr)

	lang := info.Lang.String()
	return lang, nil
}
