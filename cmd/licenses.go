package main

type License struct {
	Name string
	Aliases []string
	HasParams bool
	Body string
}

var Licenses = map[string]License{
	"agpl-3.0": {
		Name: "GNU Affero General Public License v3.0",
		Aliases: []string{},
		HasParams: false,
		Body: Apl3Body,
	},
	"apache-2.0": {
		Name: "Apache License 2.0",
		Aliases: []string{},
		HasParams: false,
		Body: Apache2Body,
	},
	"bsd-2-clause": {
		Name: "BSD 2-Clause \"Simplified\" License",
		Aliases: []string{},
		HasParams: true,
		Body: Bsd2ClauseBody,
	},
	"bsd-3-clause": {
		Name: "BSD 3-Clause \"New\" or \"Revised\" License",
		Aliases: []string{},
		HasParams: true,
		Body: Bsd3ClauseBody,
	},
	"bsl-1.0": {
		Name: "Boost Software License 1.0",
		Aliases: []string{},
		HasParams: false,
		Body: Bsl1Body,
	},
	"cc0-1.0": {
		Name: "Creative Commons Zero v1.0 Universal",
		Aliases: []string{},
		HasParams: false,
		Body: Cc01Body,
	},
	"epl-2.0": {
		Name: "Eclipse Public License 2.0",
		Aliases: []string{},
		HasParams: false,
		Body: Epl2Body,
	},
	"gpl-2.0": {
		Name: "GNU General Public License v2.0",
		Aliases: []string{},
		HasParams: false,
		Body: Gpl2Body,
	},
	"gpl-3.0": {
		Name: "GNU General Public License v3.0",
		Aliases: []string{},
		HasParams: false,
		Body: Gpl3Body,
	},
	"lgpl-2.1": {
		Name: "GNU Lesser General Public License v2.1",
		Aliases: []string{},
		HasParams: false,
		Body: Lgpl2Body,
	},
	"mit": {
		Name: "MIT License",
		Aliases: []string{},
		HasParams: true,
		Body: MitBody,
	},
	"mpl-2.0": {
		Name: "Mozilla Public License 2.0",
		Aliases: []string{},
		HasParams: false,
		Body: Mpl2Body,
	},
	"unlicense": {
		Name: "The Unlicense",
		Aliases: []string{},
		HasParams: false,
		Body: UnlicenseBody,
	},
}
