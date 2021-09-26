# Faker

Faker is a Go package that generates fake data for you to help with development.
It tris to give a simple API to access anywhere.

Coming from a PHP background, I like the way [PHP Faker](https://github.com/fzaninotto/Faker) works and how 
easy it is to generate data such as first and last names, emails, domains, texts, etc. to bootstrap databases for 
development, stress-test persistence layer, etc. So this Faker is heavily based on the same API provided by [PHP Faker](https://github.com/fzaninotto/Faker).

# Table of Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Adding Providers](#adding-providers)
- [It's a WIP](#work-in-progress)

## Installation

You can simply add the package `github.com/rmsj/faker` to you import section and run `go mod tidy` 

It's still early days in development so API might change. 
To upgrade, or downgrade the dependency, run go get:

```sh
go get github.com/rmsj/faker@v0.0.1
```

Using the appropriate version number you want.

## Basic Usage

Faker API, as of version 0.0.1 has only two `provider` to help with generation of fake data.

- Person - for names 
- Internet - for email, domain, urls, etc.


```go
package main

import (
	"fmt"
	"github.com/rmsj/faker"
	"github.com/rmsj/faker/provider"
)

func main() {
	pp := provider.NewEnglishPersonProvider()
	ip := provider.NewEnglishInternetProvider()

	f, err := faker.New(pp, ip)
	if err != nil {
		panic(err)
	}

	// print random first name
	fmt.Println(f.FirstName())

	// prints random email
	fmt.Println(f.Email())
}
```

## Adding Providers

The number of providers will grow over time and the idea is that you can change a specific `faker` provider by implementing 
the required interface with a different set of data - to have more control, change language, etc.

So you could implement the `PersonProvider` interface to have `portuguese` names, for example, and the basic example above 
would look like:

```go
package pt_provider


type PortuguesePersonProvider struct {
	//
}

func NewPortuguesePersonProvider() PortuguesePersonProvider {
	// construct code here
}

func (p PortuguesePersonProvider) FirstNames() []string {
	return []string{"Manoel", "João"}
}

// rest of implementation ...
```

```go
package main

import (
	"fmt"
	"github.com/rmsj/faker"
	"github.com/rmsj/faker/provider"

	"github.com/your-name-or-company/pt_provider"
)

func main() {
	pp := pt_provider.NewPortuguesePersonProvider()
	ip := provider.NewEnglishInternetProvider()

	f, err := faker.New(pp, ip)
	if err != nil {
		panic(err)
	}

	// print random first name
	fmt.Println(f.FirstName())
}
```

## Work In Progress

The package is a work in progress as I'm slowly adding more `fakers` and `providers` to it.