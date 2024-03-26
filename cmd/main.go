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

func ListLicenses() {
	message := "List of supported licenses and their commands: \n"
	for k, l := range Licenses {
		message += fmt.Sprintf("- \"%s\" %s\n", k, l.Name)
	}

	fmt.Println(message)
}

func main() {
	var nameLong string
	var nameShort string
	var licenseTypeLong string
	var licenseTypeShort string

	flag.StringVar(&nameLong, "name", "", "Full name for license")
	flag.StringVar(&nameShort, "n", "", "Short alias for name")
	flag.StringVar(&licenseTypeLong, "license", "", "License type to generate")
	flag.StringVar(&licenseTypeShort, "l", "", "Short alias for license type")
	flag.Parse()

	licenseType := licenseTypeLong
	if licenseType == "" {
		licenseType = licenseTypeShort
	}

	name := nameLong
	if name == "" {
		name = nameShort
	}

	// This could be either a command ("list") or a license path ("path/to/license")
	arg := flag.Arg(0)
	switch strings.ToLower(arg) {
	case "list", "ls":
		ListLicenses()
		return
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
