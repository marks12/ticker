# ticker
Переодический запуск функций по заданному расписанию

## Использование

```golang

// Удаляем устаревшие данные
func clearOldData(ctx context.Context) (err error) {

    err = common.Ticker(ctx, TickerClearCode, func(tick ticker.Tick) (err error) {
        err = clearOldData(ctx)
        return err
    }, time.Hour * 24, 0)
    
    return
}
	
```
