![go-blablacar logo](https://github.com/Midneit/go-blablacar/blob/master/logo/logo.png?raw=true)
go-blablacar is a Go client library for accessing the [BlaBlaCar API](https://support.blablacar.com/hc/en-gb/sections/360004167199-Documentation-BlaBlaCar-API).

[![Go Reference](https://pkg.go.dev/badge/github.com/Midneit/go-blablacar.svg)](https://pkg.go.dev/github.com/Midneit/go-blablacar)
[![Lint](https://github.com/Midneit/go-blablacar/actions/workflows/lint.yml/badge.svg)](https://github.com/Midneit/go-blablacar/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Midneit/go-blablacar)](https://goreportcard.com/report/github.com/Midneit/go-blablacar)

----

## Installation
```bash
go get github.com/Midneit/go-blablacar
```

## Example
```go
package main

import (
	"context"
	"github.com/Midneit/go-blablacar/blablacar"
	"log"
)

func main() {
	token := "<TOKEN>"
	client := blablacar.NewClient(token)

	search, err := client.Search(context.Background(), &blablacar.SearchRequest{
		FromCoordinate: "55.755826,37.6173",
		ToCoordinate:   "57.626074,39.88447",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(search)
}
```