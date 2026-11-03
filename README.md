# DyJSON

[![License][badge-1-img]][badge-1-link]
[![go.dev][badge-2-img]][badge-2-link]
[![Go Report Card][badge-3-img]][badge-3-link]

A small and simple package to help you handle dynamic and unknown JSON
elegantly.

## Why

DyJSON lets you handle dynamic JSON in a different way. It is perfect
for cases where you have to deal differently based on a field type or
structure, e.g. call a function recursivelly if it's an object or array
until you find a string or number value (something similar to what I
needed).

If, while dealing with your JSON, you know the specific value path, or
just want to validate it, I suggest you to use [Gabs][1].

## Example

```go
func main() {

    const json = `
        {
            "key": {
                "key": [
                    "test",
                    1,
                    true
                ]
            }
        }
    `

    handle(dyjson.ParseString(json))

    // The output is:
    //   found string: test
    //   found number: 1.00
}

func handle(v *dyjson.JSONValue) {
    switch {
    case v.IsArray():
        for _, av := range v.Array() {
            handle(av)
        }
    case v.IsObject():
        handle(v.Object()["key"])
    case v.IsString():
        fmt.Printf("found string: %s\n", v.String())
    case v.IsNumber():
        fmt.Printf("found number: %.2f\n", v.Number())
    }
}
```

## Source code

The main repository is on [SourceHut][source-srht]. Mirrors exist on
[Codeberg][source-codeberg] and [GitHub][source-github] for convenience, but
they might occasionally be outdated.

## License and contribution

All original code is licensed under the
[Mozilla Public License 2.0][license-code].

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you shall be licensed under the same license as
the material to which it is contributed, without any additional terms or
conditions.

### License history

Starting with v2.0.0, this project is licensed under MPL-2.0.

[1]: https://github.com/Jeffail/gabs

[badge-1-img]: https://img.shields.io/badge/license-MPL--2.0-green?style=flat-square
[badge-1-link]: ./LICENSE
[badge-2-img]: https://img.shields.io/badge/go.dev-reference-007d9c?style=flat-square&logo=go&logoColor=white
[badge-2-link]: https://pkg.go.dev/github.com/paemuri/dyjson/v2
[badge-3-img]: https://goreportcard.com/badge/github.com/paemuri/dyjson?style=flat-square
[badge-3-link]: https://goreportcard.com/report/github.com/paemuri/dyjson

[source-srht]: https://git.sr.ht/~paemuri/dyjson
[source-codeberg]: https://codeberg.org/paemuri/dyjson
[source-github]: https://github.com/paemuri/dyjson

[license-code]: ./LICENSES/MPL-2.0.txt
