package main

import (
	"fmt"
	"sync"
	"time"
)

// 5 философов, 5 вилок, 5 тарелок

// Philosopher Структура в которой хранится информация о них
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// Теперь нам нужен слайс философов
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// определим некоторые переменные
var hunger = 3                  // столько раз философы будут кушать
var eatTime = 1 * time.Second   // время для трапизы
var thinkTime = 3 * time.Second // время на подумать
var sleepTime = 1 * time.Second // время для сна

func main() {
	fmt.Println("Проблема обедающих философов")
	fmt.Println("----------------------------")
	fmt.Println("Стол пустой.")
	dine()
	fmt.Println("Стол пустой.")
}

func dine() {

	// 1 группа для трапизы
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// 2 группа для посадки философов за стол
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// создадим карту для 5 вилок
	forks := make(map[int]*sync.Mutex)
	// инициализируем карту ключами и значениями
	for i := range philosophers {
		forks[i] = &sync.Mutex{}
	}

	// философы начинаю кушать
	for i := range philosophers {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// философ садиться за стол
	fmt.Printf("%s садиться за стол\n", philosopher.name)
	seated.Done()

	seated.Wait()
	// кушаем три раза
	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\tФилософ %s взял правую вилку.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\tФилософ %s взял левую вилку.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\tФилософ %s взял левую вилку.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\tФилософ %s взял правую вилку.\n", philosopher.name)
		}

		fmt.Printf("\tФилософа У %s обе вилки, он кушает...\n", philosopher.name)
		time.Sleep(eatTime)
		fmt.Printf("\tФилософ %s поел и теперь думает...\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("Философ %s положил обе вилки.\n", philosopher.name)

		fmt.Printf("Философ %s покинул стол.\n", philosopher.name)
	}
}
