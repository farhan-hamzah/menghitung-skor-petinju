package main

import "fmt"

func skorDanPemenang(n int) {
    var skorA, skorB int

    for i := 0; i < n; i++ {
        var a1, a2, a3, b1, b2, b3 int
        fmt.Scan(&a1, &a2, &a3, &b1, &b2, &b3)

        skorA += a1 + a2 + a3
        skorB += b1 + b2 + b3
    }

    var pemenang string
    if skorA > skorB {
        pemenang = "A"
    } else if skorB > skorA {
        pemenang = "B"
    } else {
        pemenang = "0"
    }

    fmt.Println(skorA, skorB, pemenang)
}

func main() {
    var n int
    fmt.Scan(&n)
    skorDanPemenang(n)
}
