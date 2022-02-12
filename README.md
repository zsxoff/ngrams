# ngrams

> Create sequence of N items from a given string.

## Usage example

```go
import (
    "fmt"
    "github.com/zsxoff/ngrams"
)

func main() {
    s := "Hello!"

    for i, ngram := range ngrams.Ngrams(s, 2) {
        fmt.Println(i, ngram)
    }
}

// Output:
// 0 He
// 1 el
// 2 ll
// 3 lo
// 4 o!
```

Also you can use `Unigrams(line string)`, `Bigrams(line string)` and `Trigrams(line string)` like

```go
s := "Hello!"

for i, ngram := range ngrams.Trigrams(s) {
    fmt.Println(i, ngram)
}

// Output:
// 0 Hel
// 1 ell
// 2 llo
// 3 lo!
```

---

## License

[![License: Unlicense](https://img.shields.io/badge/License-Unlicense-green.svg?style=flat-square)](https://unlicense.org/)

This project is licensed under the terms of the [Unlicense](https://unlicense.org/) (see [LICENSE](<https://github.com/zsxoff/ngrams/blob/master/LICENSE>) file).
