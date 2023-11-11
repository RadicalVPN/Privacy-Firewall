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

func download(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(rawBody)
}

func create_blocklist(alias string, url string) {
	source_slice := []string{}

	for _, line := range strings.Split(strings.TrimRight(download(url), "\n"), "\n") {
		parsed := filter(line)
		if parsed == "" {
			continue
		}

		source_slice = append(source_slice, line)
	}

	f, err := os.OpenFile("../lists/"+alias+".txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f.Write([]byte(strings.Join(source_slice, "\n")))
	defer f.Close()
}

func main() {
	recommended_list := []string{
		"https://cdn.jsdelivr.net/gh/badmojr/1Hosts@master/Lite/hosts.txt", // https://o0.pages.dev
		"https://adguardteam.github.io/AdGuardSDNSFilter/Filters/filter.txt",
		"https://easylist.to/easylist/easylist.txt", // https://easylist.to
		"https://easylist.to/easylist/easyprivacy.txt",
		"https://small.oisd.nl/",    // https://oisd.nl/downloads
		"http://sbc.io/hosts/hosts", // https://github.com/StevenBlack/hosts
	}

	for _, source := range recommended_list {
		fmt.Println("Loading recommended source: " + source)
		create_blocklist("basic", source)
	}

	comprehensive_list := []string{
		"https://cdn.jsdelivr.net/gh/badmojr/1Hosts@master/Pro/hosts.txt",               // https://o0.pages.dev
		"https://www.github.developerdan.com/hosts/lists/ads-and-tracking-extended.txt", // https://github.com/lightswitch05/hosts
		"https://gitlab.com/hagezi/mirror/-/raw/main/dns-blocklists/domains/multi.txt",  // https://github.com/hagezi/dns-blocklists#normal,
		"https://big.oisd.nl/", // https://oisd.nl/downloads
	}

	//merge recommended and comprehensive lists
	comprehensive_list = append(comprehensive_list, recommended_list...)

	for _, source := range comprehensive_list {
		fmt.Println("Loading comprehensive source: " + source)
		create_blocklist("comprehensive", source)
	}

	agressive_list := []string{
		"https://cdn.jsdelivr.net/gh/badmojr/1Hosts@master/Xtra/hosts.txt",                 // https://o0.pages.dev
		"https://www.github.developerdan.com/hosts/lists/tracking-aggressive-extended.txt", // https://github.com/lightswitch05/hosts
		"https://gitlab.com/hagezi/mirror/-/raw/main/dns-blocklists/domains/ultimate.txt",  // https://github.com/hagezi/dns-blocklists#ultimate
	}

	//merge recommended and comprehensive lists
	agressive_list = append(agressive_list, comprehensive_list...)

	for _, source := range agressive_list {
		fmt.Println("Loading agressive source: " + source)
		create_blocklist("agressive", source)
	}
}
