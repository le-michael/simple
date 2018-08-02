## simple
A compilation of simple data-structures written in go

## Why?
Golang is a great language but it's missing some basic data structures. I made this package so I won't have to have to reimplement all these data structures.

The code is also simple to read for anyone trying to get themselves familiar with Go. 
## Usage

#### Stack 

```Go
package main

import (
    "fmt"

    "github.com/le-michael/simple"
)

func main() {
    stack := simple.NewStack()

    for i := 0; i < 10; i++ {
        stack.Push(i)
    }

    for !stack.Empty() {
        fmt.Println(stack.Top())
        err := stack.Pop()
    }
    // 9 8 7 6 5 4 3 2 1 0
}
```

#### Queue

```Go
package main

import (
    "fmt"

    "github.com/le-michael/simple"
)

func main() {
    queue := simple.NewStack()

    for i := 0; i < 10; i++ {
        queue.Push(i)
    }

    for !queue.Empty() {
        fmt.Println(queue.Top())
        queue.Pop()
    }
    // 0 1 2 3 4 5 6 7 8 9
}
```