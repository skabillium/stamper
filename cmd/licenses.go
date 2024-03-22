package main

type License struct {
	Name      string
	Aliases   []string
	HasParams bool
	Body      string
}

var Licenses = map[string]License{
	"agpl3": {
		Name:      "GNU Affero General Public License v3.0",
		HasParams: false,
		Body:      Apl3Body,
	},
	"apache2": {
		Name:      "Apache License 2.0",
		HasParams: false,
		Body:      Apache2Body,
	},
	"bsd2": {
		Name:      "BSD 2-Clause \"Simplified\" License",
		HasParams: true,
		Body:      Bsd2ClauseBody,
	},
	"bsd3": {
		Name:      "BSD 3-Clause \"New\" or \"Revised\" License",
		HasParams: true,
		Body:      Bsd3ClauseBody,
	},
	"bsl": {
		Name:      "Boost Software License 1.0",
		HasParams: false,
		Body:      Bsl1Body,
	},
	"cc": {
		Name:      "Creative Commons Zero v1.0 Universal",
		HasParams: false,
		Body:      Cc01Body,
	},
	"epl2": {
		Name:      "Eclipse Public License 2.0",
		HasParams: false,
		Body:      Epl2Body,
	},
	"gpl2": {
		Name:      "GNU General Public License v2.0",
		HasParams: false,
		Body:      Gpl2Body,
	},
	"gpl3": {
		Name:      "GNU General Public License v3.0",
		HasParams: false,
		Body:      Gpl3Body,
	},
	"lgpl2": {
		Name:      "GNU Lesser General Public License v2.1",
		HasParams: false,
		Body:      Lgpl2Body,
	},
	"mit": {
		Name:      "MIT License",
		HasParams: true,
		Body:      MitBody,
	},
	"mpl2": {
		Name:      "Mozilla Public License 2.0",
		HasParams: false,
		Body:      Mpl2Body,
	},
	"unlicense": {
		Name:      "The Unlicense",
		HasParams: false,
		Body:      UnlicenseBody,
	},
}
