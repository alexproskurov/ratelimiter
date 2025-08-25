# ratelimiter

A Go package providing implementations of rate limiting algorithms such as **Leaky Bucket** and more to come.  
This package is designed to be used as a framework for experimenting with different rate limiting strategies.

## Features

- `RateLimiter` interface for extensibility
- Leaky Bucket algorithm implementation
- Easy to extend with additional algorithms (e.g., Token Bucket, Fixed Window, Sliding Window)

## Installation

```bash
go get github.com/alexproskurov/ratelimiter
```

## Usage

Example using the Leaky Bucket rate limiter:

```go
package main

import (
    "fmt"
    "time"

    "github.com/alexproskurov/ratelimiter/ratelimiter"
)

func main() {
    rl := ratelimiter.NewLeakyBucket(5, 200 * time.Millisecond)

    for i := 0; i < 10; i++ {
        if rl.Allow() {
            fmt.Println("Request", i, "allowed")
        } else {
            fmt.Println("Request", i, "denied")
        }
        time.Sleep(100 * time.Millisecond)
    }
}
```

## Roadmap

- [x] Leaky Bucket
- [ ] Token Bucket
- [ ] Fixed Window
- [ ] Sliding Window

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.
