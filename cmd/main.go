package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const DefaultLicenseName = "LICENSE"

func main() {
	var name string
	flag.StringVar(&name, "name", "", "Full name to be put in license")

	flag.Parse()

	kind := strings.ToLower(os.Args[len(os.Args)-1])

	switch kind {
	case "list", "ls":
		fmt.Println("List of supported licenses")
		for _, l := range Licenses {
			fmt.Printf("- %s\n", l.Name)
		}
	default:
		license, ok := Licenses[kind]
		if !ok {
			fmt.Printf("License '%s' not supported\n", kind)
			return
		}

		if license.HasParams {
			if name == "" {
				fmt.Println("License", kind, "requires name")
				return
			}

			rep := strings.NewReplacer("[year]", fmt.Sprint(time.Now().Year()), "[fullname]", name)
			license.Body = rep.Replace(license.Body)

			strings.NewReplacer()
		}

		file, err := os.Create(DefaultLicenseName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		file.WriteString(license.Body)
		fmt.Println("Created file", DefaultLicenseName)
	}

}
