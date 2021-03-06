[![Go Report Card](https://goreportcard.com/badge/github.com/le-michael/simple)](https://goreportcard.com/report/github.com/le-michael/simple) [![Build Status](https://travis-ci.org/le-michael/simple.svg?branch=master)](https://travis-ci.org/le-michael/simple) [![Coverage Status](https://coveralls.io/repos/github/le-michael/simple/badge.svg?branch=master)](https://coveralls.io/github/le-michael/simple?branch=master)


# simple
A compilation of simple data-structures written in go

## Why?
Golang is a great language but it's missing some basic data structures. I made this package so I won't have to have to reimplement all these data structures.

The code is also simple to read for anyone trying to get themselves familiar with Go. 
## Usage

### Stack 
<details>
    
<summary>Example</summary>
<br>

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
        stack.Pop()
    }
    // 9 8 7 6 5 4 3 2 1 0
}
```
</details>

### Queue 
<details>
    
<summary>Example</summary>
<br>
    
```Go
package main

import (
    "fmt"

    "github.com/le-michael/simple"
)

func main() {
    queue := simple.NewQueue()

    for i := 0; i < 10; i++ {
        queue.Push(i)
    }

    for !queue.Empty() {
        fmt.Println(queue.Front())
        queue.Pop()
    }
    // 0 1 2 3 4 5 6 7 8 9
}
```
</details>
