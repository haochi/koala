package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type ReadMe struct {
	KoalaExample string
	RepoUrl string
}

func BuildReadMe() {
	readMeTemplate, err := ioutil.ReadFile("./README_TEMPLATE.md")
	checkErr(err)

	basicKoalaExample, err := ioutil.ReadFile("../examples/basic/main.go")
	checkErr(err)

	readMe := ReadMe{
		KoalaExample: string(basicKoalaExample),
		RepoUrl: "github.com/haochi/koala",
	}
	tmpl, err := template.New("README").Parse(string(readMeTemplate))
	checkErr(err)

	readMeFile, err := os.Create("../README.md")
	checkErr(err)
	defer readMeFile.Close()

	err = tmpl.Execute(readMeFile, readMe)
	checkErr(err)
}
