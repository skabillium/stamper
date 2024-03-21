package main

type License struct {
	Name      string
	Aliases   []string
	HasParams bool
	Body      string
}

var Licenses = map[string]License{
	"agpl-3.0": {
		Name:      "agpl-3.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Apl3Body,
	},
	"apache-2.0": {
		Name:      "apache-2.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Apache2Body,
	},
	"bsd-2-clause": {
		Name:      "bsd-2-clause",
		Aliases:   []string{},
		HasParams: true,
		Body:      Bsd2ClauseBody,
	},
	"bsd-3-clause": {
		Name:      "bsd-3-clause",
		Aliases:   []string{},
		HasParams: true,
		Body:      Bsd3ClauseBody,
	},
	"bsl-1.0": {
		Name:      "bsl-1.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Bsl1Body,
	},
	"cc0-1.0": {
		Name:      "cc0-1.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Cc01Body,
	},
	"epl-2.0": {
		Name:      "epl-2.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Epl2Body,
	},
	"gpl-2.0": {
		Name:      "gpl-2.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Gpl2Body,
	},
	"gpl-3.0": {
		Name:      "gpl-3.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Gpl3Body,
	},
	"lgpl-2.1": {
		Name:      "lgpl-2.1",
		Aliases:   []string{},
		HasParams: false,
		Body:      Lgpl2Body,
	},
	"mit": {
		Name:      "mit",
		Aliases:   []string{},
		HasParams: true,
		Body:      MitBody,
	},
	"mpl-2.0": {
		Name:      "mpl-2.0",
		Aliases:   []string{},
		HasParams: false,
		Body:      Mpl2Body,
	},
	"unlicense": {
		Name:      "unlicense",
		Aliases:   []string{},
		HasParams: false,
		Body:      UnlicenseBody,
	},
}
