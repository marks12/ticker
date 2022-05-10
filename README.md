# ticker
Переодический запуск функций по заданному расписанию

## Использование

```go

// Удаляем устаревшие данные
func clearOldData(ctx context.Context) (err error) {

    err = ticker.Ticker(ctx, TickerClearCode, func(tick ticker.Tick) (err error) {
        err = clearOldData(ctx)
        return err
    }, time.Hour * 24, 0)
    
    return
}

//...

// remove old data
func clearOldData()  {
    // some func body
}
	
```
