# ticker
Периодический запуск функций по заданному расписанию

## Использование

```go
package main

import (
	"context"
	"fmt"
	"github.com/marks12/ticker"
	"time"
)

func InitTicker() {

	err := ticker.New(context.Background(), "test", func(tick ticker.Tick) error {

		fmt.Printf("tick: %+v\n", tick)
		return nil

	}, time.Second * 3, 0)

	if err != nil {
		fmt.Println("Can`t start ticker. Err: ", err)
	}
}

func main() {
	
	InitTicker()
	
	for {
		fmt.Println("text ticker")
		time.Sleep(time.Second)
    }
}
```

```shell
$ go run ./main.go 
text ticker
text ticker
text ticker
tick: {Ctx:context.Background Code:test Period:3s Counter:0}
text ticker
text ticker
text ticker
tick: {Ctx:context.Background Code:test Period:3s Counter:1}
text ticker
text ticker
text ticker
tick: {Ctx:context.Background Code:test Period:3s Counter:2}
text ticker
text ticker
text ticker
tick: {Ctx:context.Background Code:test Period:3s Counter:3}
text ticker

```
