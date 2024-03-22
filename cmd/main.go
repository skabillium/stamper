package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const DefaultLicenseName = "LICENSE"

func formatKey(key string) string {
	return strings.NewReplacer("-", "", "0", "", "1", "", ".", "", "clause", "").Replace(key)
}

func main() {
	var name string
	var licenseType string
	flag.StringVar(&name, "name", "", "Full name to be put in license")
	flag.StringVar(&licenseType, "license", "", "License type to generate")
	flag.Parse()

	// This could be either a command ("list") or a license path ("path/to/license")
	arg := flag.Arg(0)
	switch strings.ToLower(arg) {
	case "list", "ls":
		message := "List of supported licenses: \n"
		for _, l := range Licenses {
			message += fmt.Sprintf("- %s\n", l.Name)
		}

		fmt.Println(message)
	default:
		out := DefaultLicenseName
		if arg != "" {
			out = arg
		}

		key := formatKey(licenseType)
		license, ok := Licenses[key]
		if !ok {
			fmt.Printf("License '%s' not supported\n", licenseType)
			return
		}

		if license.HasParams {
			if name == "" {
				fmt.Printf("License '%s' requires full name \n", licenseType)
				return
			}

			rep := strings.NewReplacer("[year]", fmt.Sprint(time.Now().Year()), "[fullname]", name)
			license.Body = rep.Replace(license.Body)
		}

		file, err := os.Create(out)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		file.WriteString(license.Body)
		fmt.Println("Created file", out)
	}
}
