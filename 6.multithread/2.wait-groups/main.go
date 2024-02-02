package main

import (
	"fmt"
	"sync"
	"time"
)

// wait groups são métodos para que seja possível esperar que um conjunto de goroutines termine de executar.

func task(name string, w *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s Running!\n", i, name)
		time.Sleep(1 * time.Second)
		/*
			Ao inserir um .Done() é possível que o WaitGroup saiba que a goroutine terminou de executar.
			Isso adiciona +1 na contagem do WaitGroup que, ao chegar no valor estipolado na hora de sua instanciação,
			para de esperar e permite a executação do restante do código.
		*/
		w.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{} // Instancia um novo WaitGroup
	waitGroup.Add(25)             // Define que o WaitGroup deve esperar 25 goroutines terminarem de executar

	go task("A", &waitGroup) // Informa o ponteiro do waitgroup para que a goroutine possa executar o .Done(). É preciso informar o ponteiro se não será feito uma cópia
	go task("B", &waitGroup)

	go func(name string, w *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s Running!\n", i, name)
			time.Sleep(1 * time.Second)
			w.Done() // Adiciona +1 ao contador do wait gorup
		}
	}("anonymous", &waitGroup)

	waitGroup.Wait() // Faz com que o WaitGroup aguarde para que todas as contagens definidas (25) sejam feitas usando o .Done()
}
