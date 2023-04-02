package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case a, ok := <-ch: //читаем из канала, канал конкурентно безопасен, т.к. содержит мьютексы внтури себя
			if !ok {
				continue
			}
			fmt.Println(a) // пишем в stdout
		case <-ctx.Done():
			a, ok := <-ch //так как в канале могут остаться данные, считываем и выводим их
			if ok {
				fmt.Println(a)
			}
			wg.Done() //убавляем счетчик и завершаем работу воркера
			return
		}
	}
}

func main() {
	var N uint                                                                //инициализируем тип N
	N = 100                                                                   // присваиваем N начение
	ch := make(chan int, 3)                                                   // создаем буферизованный канал, можно и небуферизованный
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT) //контекст отлавливающий прерывание "Ctrl + C"
	var wg sync.WaitGroup                                                     // создаем WaitGroup для того, чтобы полностью вычитать канал + чтобы все горутины доработали
	var i uint
	for i = 0; i < N; i++ {
		wg.Add(1)               //для каждого воркера +1 в wg
		go worker(ch, ctx, &wg) //запускаем N воркеров
	}

	var iterating_flg = true         // флаг, когда можно писать в канал.
	for k := 0; iterating_flg; k++ { // пишем в канал значения
		select {
		case <-ctx.Done(): //Если контекст исполнился, закрываем его, закрываем канал, завершаем работу
			cancel()              // Закрываем контекст
			close(ch)             // Закрываем канал. (кто пишет - тот и закрывает)
			iterating_flg = false // в закрытый канал писать нельзя
		default:
			ch <- k //Пишем в канал
		}
	}
	wg.Wait() //ждем закрытия всех воркеров, потому что часть может не успеть доработать
}
