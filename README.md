# Sengkala

**Sengkala** is a Go package for generating _sengkala_ (Javanese [chronogram](https://en.wikipedia.org/wiki/Chronogram)).

It will generate two types of _sengkala_:
- **Surya Sengkala**, a Gregorian (sun) calendar based chronogram
- **Candra Sengkala**, a Javanese (moon) calendar based chronogram

## Usage Example

```go
package main

import (
    "fmt"

    "github.com/matriphe/sengkala"
)

func main()  {
    // get sengkala for year 2020
    s := sengkala.FromYear("2020")
    // s.SetYear("2021")

    // get surya sengkala
    suryaSengkala := s.GetSuryaSengkala()
    // get year
    fmt.Println(suryaSengkala.GetYear())
    // get sentence
    fmt.Println(suryaSengkala.GetSengkala())
    // get meaning for every word
    // fmt.Println(suryaSengkala.GetMeaning())
    
    // get candra sengkala
    candraSengkala := s.GetCandraSengkala()
    // get year
    fmt.Println(candraSengkala.GetYear())
    // get sentence
    fmt.Println(candraSengkala.GetSengkala())
    // get meaning for every word
    // fmt.Println(candraSengkala.GetMeaning())
}
```
### Results
```text
2020
Talingan Langit Tanpa Nembah
1934
Keblat Naut Kori Semedi
```

## Resource
- [Membuat Sengkalan](https://lantip.xyz/2020/05/membuat-sengkalan/), a blog post in Bahasa Indonesia (Indonesian) by [Lantip](https://github.com/lantip/sengkalan)

## License

Sengkala is distributed using [MIT License](LICENSE).