# EsbFireplace - Go

The [FIREPLACEv1.0 protocol](https://github.com/luxedo/esb/blob/main/doc/FIREPLACEv1.0.md)
allows the use of the `esb` tooling for solving Advent of Code problems.
This is an implementation of FIREPLACEv1.0 for [go](https://go.dev/).

Check [esb](https://github.com/luxedo/esb) for more information.

## Installation

The package can be installed by adding `esb_fireplace-go` to your list of dependencies in `go.mod`:

```mod
require github.com/luxedo/esb_fireplace-go v1.0.0
```

or by running:

```bash
go get github.com/luxedo/esb_fireplace-go
```

## Usage

Create your solution functions and add `esb_fireplace.V1Run` to `main`.

```go
package main

import (
  "github.com/luxedo/esb_fireplace-go"
)

func SolvePt1(inputData string, args []string) (interface{}, error) {
	return 25, nil
}


func SolvePt2(inputData string, args []string) (interface{}, error) {
	return "December", nil
}

func main() {
  // 🎅🎄❄️☃️🎁🦌
  // Bright christmas lights HERE
  esb_fireplace.V1Run(SolvePt1, SolvePt2)
}
```

Running can be done with `go run`, but this library is meant to be used with [esb](https://github.com/luxedo/esb).

```bash
# You can do this...
go run . --part 1 < inputData.txt

# But instead do this:
esb run --year 2023 --day 1 --lang rust --part 1
```

The docs can be found at <https://pkg.go.dev/github.com/luxedo/esb_fireplace-go/>.

## License

> ESB - Script your way to rescue Christmas as part of the ElfScript Brigade team.
> Copyright (C) 2024 Luiz Eduardo Amaral <luizamaral306@gmail.com>
>
> This program is free software: you can redistribute it and/or modify
> it under the terms of the GNU General Public License as published by
> the Free Software Foundation, either version 3 of the License, or
> (at your option) any later version.
> This program is distributed in the hope that it will be useful,
> but WITHOUT ANY WARRANTY; without even the implied warranty of
> MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
> GNU General Public License for more details.
> You should have received a copy of the GNU General Public License
> along with this program. If not, see <http://www.gnu.org/licenses/>.
