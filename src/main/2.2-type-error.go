package main

import "errors"
import "fmt"

func main() {
    err := errors.New("emit macho dwarf: elf header corrupted")
    if err != nil {
        fmt.Println(err)
    }

}
