package main

/*
	Buffer em channels são os "slots" destes. O padrão de um channel é buffer 1,
	ou seja, apenas 1 valor pode ser atribuído ao channel de uma vez. Ao inserir mais
	de um buffer é possível adicionar múltiplos valores que irão ser consumidos em FIFO.
	Caso não possua uma concorrência suficiente para consumir todos os valores rapidamente,
	adicionar um número maior de buffer não trará benefícios e consumirá memória
*/

func main() {
	ch := make(chan string, 2) // buffer de 2 slots
	ch <- "Hello"
	/*
		Se o buffer possuir o default (1) a linha a seguir irá gerar um deadlock pois
		a segunda mensagem nunca será lida até a primeira ser lida, o que nunca acontecerá
	*/
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
