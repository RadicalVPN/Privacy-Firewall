package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func filter(line string) string {
	if strings.HasPrefix(line, "#") {
		return ""
	}

	if strings.HasPrefix(line, "!") {
		return ""
	}

	return line
}

func main() {
	recommended_sources := []string{
		"https://cdn.jsdelivr.net/gh/badmojr/1Hosts@master/Lite/hosts.txt", // https://o0.pages.dev
		"https://adguardteam.github.io/AdGuardSDNSFilter/Filters/filter.txt",
		"https://easylist.to/easylist/easylist.txt", //https://easylist.to
		"https://easylist.to/easylist/easyprivacy.txt",
		"https://small.oisd.nl/", //https://oisd.nl/downloads
		"http://sbc.io/hosts/hosts", //https://github.com/StevenBlack/hosts
	}


	source_slice := []string{}
	for _, source := range recommended_sources {
		fmt.Println("Loading source: " + source)

		res, err := http.Get(source)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		rawBody, err := io.ReadAll(res.Body)
		if err != nil {
		   log.Fatalln(err)
		}

		body := string(rawBody)

		for _, line := range strings.Split(strings.TrimRight(body, "\n"), "\n") {
			parsed := filter(line)
			if parsed == "" {
				continue
			}

			source_slice = append(source_slice, line)
		 }
	}

	f, err := os.OpenFile("../lists/hosts.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f.Write([]byte(strings.Join(source_slice, "\n")))
	f.Close()
}