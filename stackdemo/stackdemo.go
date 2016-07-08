package main

import (
    "fmt"
    "stack"
)

func main() {
    st := stack.New()
    st.push(1)
    val, err := st.peek()
    if err == nil {
        fmt.println(st.peek())
    }
}