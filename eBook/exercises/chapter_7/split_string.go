package main

import (
        "fmt"
    )

    func main() {
            rawString := "Google"
                index := 3
                    sp1, sp2 := splitStringbyIndex(rawString, index)
                        fmt.Printf("The string %s split at position %d is: %s / %s\n", rawString, index, sp1, sp2)
                    }

                    func splitStringbyIndex(str string, i int) (sp1, sp2 string) {
                            rawStrSlice := []byte(str)
                                sp1 = string(rawStrSlice[:i])
                                    sp2 = string(rawStrSlice[i:])
                                        return
                                    }
