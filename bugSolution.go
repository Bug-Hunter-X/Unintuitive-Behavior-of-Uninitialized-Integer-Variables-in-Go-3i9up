```Go
package main

import (

    "fmt"
)

func main() {
    var i int
    fmt.Println("Uninitialized i:", i) // Output: 0
    fmt.Println("i == 0:", i == 0)     // Output: true
    fmt.Println("&i == nil:", &i == nil) // Output: false

    var j int = 0 // Explicit initialization
    fmt.Println("\nExplicitly initialized j:", j)
    fmt.Println("j == 0:", j == 0)     // Output: true
    fmt.Println("&j == nil:", &j == nil) // Output: false

    var k *int
    fmt.Println("\nInitialized pointer k:", k) // Output: <nil>
}
```