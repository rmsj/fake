# Fake

Fake is a Go package that generates fake data for you to help with development.
It tris to give a simple API to access anywhere.

[![codecov](https://codecov.io/gh/rmsj/fake/branch/main/graph/badge.svg?token=W3IGROKH1R)](https://codecov.io/gh/rmsj/fake)
[![Go Report Card](https://goreportcard.com/badge/github.com/rmsj/fake)](https://goreportcard.com/report/github.com/rmsj/fake)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/fake/blob/main/LICENSE)
[![GoDoc](https://godoc.org/github.com/rmsj/fake?status.svg)](https://godoc.org/github.com/rmsj/fake)
[![Go.Dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/rmsj/fake/v0.0.3?tab=doc)

Coming from a PHP background, I like the way [PHP Faker](https://github.com/fzaninotto/Faker) works and how 
easy it is to generate data such as first and last names, emails, domains, texts, etc. to bootstrap databases for 
development, stress-test persistence layer, etc. So this Faker is heavily based on the same API provided by [PHP Faker](https://github.com/fzaninotto/Faker).

# Table of Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Data Factory](#data-factory)
- [Adding Providers](#adding-providers)
- [It's a WIP](#work-in-progress)

## Installation

You can simply add the package `github.com/rmsj/fake` to you import section and run `go mod tidy` 

It's still early days in development so API might change. 
To upgrade, or downgrade the dependency, run go get:

```sh
go get github.com/rmsj/fake@v0.0.5
```

Using the appropriate version number you want.

## Basic Usage

Fake API, as of version 0.0.5 has data generation capability  and `providers` for person, internet, text, lotem ipsum text and DNA sequences. 

- PersonProvider
  - for first and last names, gender, etc.
- InternetProvider
  - for email, domain names, user names, urls, etc.
- TextProvider 
  - for "real sentences" randomized from Alice in Wonder Land.
- DNAProvider
  - for generation of random fake DNA sequences of `n` length.
- CompanyProvider
  - for generation of random fake company data.
- ImageProvider
  - uses [Lorem Flickr](https://loremflickr.com) to generate random images.
- Lorem
  - generates lorem ipsum texts.

```go
package main

import (
	"fmt"
	"github.com/rmsj/fake"
)

func main() {

	f, err := fake.New()
	if err != nil {
		panic(err)
	}

	// print random first name
	fmt.Println(f.FirstName())

	// prints random email
	fmt.Println(f.Email())
}
```

## Data Factory

To create multiple data fo fill in databases for development for example, you can use the `Factory` function.
The `Factory` function requires a specific function type (`type Builder func() interface{}`) and an `int` as second parameter -
being the amount of types the operation should be repeated.

```go
package main

import (
	"fmt"
	"github.com/rmsj/fake"
)

type user struct {
	firstName string
	lastName  string
	email     string
}

func main() {

	f, err := fake.New()
	if err != nil {
		panic(err)
	}

	builder := func() interface{} {
		return user{
			firstName: f.FirstName(),
			lastName:  f.LastName(),
			email:     f.Email(),
		}
	}

	users := f.Factory(builder, 10)
	for _, v := range users {
		u, ok := v.(user)
		if !ok {
			panic("this should not happen")
		}

		// you can use the user value as normal
		fmt.Println(u.firstName)
		fmt.Println(u.email)
	}
	
	// or for more predictable values, with deterministic mode
    f.Determistix(42)

    sameUsers := f.Factory(builder, 10)
    for _, v := range sameUsers {
      u, ok := v.(user)
      if !ok {
        panic("this should not happen")
      }
  
      // all the generated users should have the same name and email
      fmt.Println(u.firstName)
      fmt.Println(u.email)
    }
	
}
```

## Adding Your Own Providers

The number of providers will grow over time and the idea is that you can change a specific `provider` by implementing 
the required interface with a different set of data - to have more control of what data is generated, change language, etc.

So you could implement the `PersonProvider` interface to have `portuguese` names, for example, and the [basic usabe](#basic-usage) example above 
would look like:

```go
package pt_provider


type PortuguesePersonProvider struct {}


func (p PortuguesePersonProvider) FirstNames() []string {
	return []string{"Manoel", "Pedro"}
}

// rest of implementation for all required methods on interface PersonProvider...
```

```go
package main

import (
	"fmt"
	"github.com/rmsj/fake"

	"github.com/user/project/pt_provider"
)

func main() {

	f, err := fake.New()
	if err != nil {
		panic(err)
	}
	
	f.SetPersonProvider(pt_provider)

	// print random first name from your list of names
	fmt.Println(f.FirstName())
}
```

## Work In Progress

The package is a work in progress as I'm slowly adding more `providers` and related data generation from them.