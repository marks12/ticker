package ticker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Tick - DTO user operation tick data
type Tick struct {
	// function context
	Ctx context.Context
	// user operation code
	Code string
	// user operation tick period
	Period time.Duration
	// current tick counter
	Counter int
}

// Единственный инстанс, хранящий список запущенных функций для однократного запуска
var tickerInstance *sync.Map
var once sync.Once

func getInstance() *sync.Map {
	once.Do(func() {
		tickerInstance = &sync.Map{}
	})

	return tickerInstance
}

// New - создание тикера, выполняющего функцию с некоторой периодичностью
// times - количество циклов которые отработает тикер
// Если times=0 тикер работает бесконечно
func New(ctx context.Context, code string, function func(tick Tick) error, period time.Duration, times int) (err error) {

	inst := getInstance()

	if existsPeriod, ok := inst.Load(code); ok == true {
		err = fmt.Errorf("Метод с кодом %s уже запущен. Вызывается с периодом %+v ", code, existsPeriod)
		return
	} else {
		inst.Store(code, period)
	}

	var count = 0

	// канал управления запуском итераций
	ch := make(chan int)

	go func() {

		// бесконечно циклим
		for {

			// выбираем входящее данные по каналам
			select {

			// если есть данные из управляющего канала, получаем их и запускаем целевую функцию
			case <-ch:

				go func() {

					if count > 0 && count == times {
						return
					}

					time.Sleep(period)

					// даже если прошел таймаут
					// смотрим не пришел ли нам cancel
					// если пришел, то никаких выполнений, на выход
					select {
					case <-ctx.Done():
						return
					default:
						_ = function(Tick{
							Ctx:     ctx,
							Code:    code,
							Period:  period,
							Counter: count,
						})

						count += 1
						ch <- 1
					}
				}()

				if count > 0 && count == times {
					return
				}

			//	если вдруг приходит отмена, немедленно сворачиваемся не дожидаясь следующей итерации
			case <-ctx.Done():
				fmt.Printf("Cancel func %+v\n", code)
				return
			}
		}
	}()

	// инициируем запуск тикера
	ch <- 1

	return
}
