# DyJSON

[![License][badge-1-img]][badge-1-link]
[![go.dev][badge-2-img]][badge-2-link]
[![Travis CI][badge-3-img]][badge-3-link]
[![Codecov.io][badge-4-img]][badge-4-link]
[![Go Report Card][badge-5-img]][badge-5-link]

A small and simple package to help you handle dynamic and unknown JSON
elegantly.

## Why

DyJSON lets you handle dynamic JSON in a different way. It is perfect for cases
where you have to deal differently based on a field type or structure, e.g. call
a function recursivelly if it's an object or array until you find a string or
number value (something similar to what I needed).

If, while dealing with your JSON, you know the specific value path, or just want
to validate it, I suggest you to use [Gabs][1].

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

## License

This project code is in the public domain. See the [LICENSE file][2].

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you shall be in the public domain, without any
additional terms or conditions.

[1]: https://github.com/Jeffail/gabs
[2]: ./LICENSE

[badge-1-img]: https://img.shields.io/github/license/Nhanderu/dyjson?style=flat-square
[badge-1-link]: https://github.com/Nhanderu/dyjson/blob/master/LICENSE
[badge-2-img]: https://img.shields.io/badge/go.dev-reference-007d9c?style=flat-square&logo=go&logoColor=white
[badge-2-link]: https://pkg.go.dev/github.com/Nhanderu/dyjson
[badge-3-img]: https://img.shields.io/travis/Nhanderu/dyjson?style=flat-square
[badge-3-link]: https://travis-ci.org/Nhanderu/dyjson
[badge-4-img]: https://img.shields.io/codecov/c/gh/Nhanderu/dyjson?style=flat-square
[badge-4-link]: https://codecov.io/gh/Nhanderu/dyjson
[badge-5-img]: https://goreportcard.com/badge/github.com/Nhanderu/dyjson?style=flat-square
[badge-5-link]: https://goreportcard.com/report/github.com/Nhanderu/dyjson
