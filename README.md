# Env

[![GoDoc](https://godoc.org/github.com/rrandall91/env?status.svg)](https://godoc.org/github.com/rrandall91/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/rrandall91/env)](https://goreportcard.com/report/github.com/rrandall91/env)
[![Maintainability](https://api.codeclimate.com/v1/badges/93cdc0bb6b1aacda2551/maintainability)](https://codeclimate.com/github/rrandall91/env/maintainability)
[![Tests](https://github.com/rrandall91/env/actions/workflows/test.yml/badge.svg)](https://github.com/rrandall91/env/actions/workflows/test.yml)
[![Test Coverage](https://api.codeclimate.com/v1/badges/93cdc0bb6b1aacda2551/test_coverage)](https://codeclimate.com/github/rrandall91/env/test_coverage)

This is a minimalistic library for accessing environment variables in Go.

## Features

* Simple methods for creating and retrieving environment variables
* Allows a fallback value if an environment variable is not set

## Usage

```go
package main

import (
    "github.com/rrandall91/env"
)

func main() {
    // Create a new environment variable
    env.Set("MY_ENV_VAR", "my value")

    // Get the value of an environment variable or a fallback value if it is not set
    env.Get("MY_ENV_VAR", "fallback value")

    // Get the value of an environment variable with a fallback
    env.GetWithDefault("MY_ENV_VAR", "default value")
}
```

## License

Copyright (c) 2022-present [Rashaad Randall](https://github.com/rrandall91). Env is free and open-source software licensed under the [GNU Affero General Public License](https://github.com/rrandall91/env/blob/master/LICENSE).