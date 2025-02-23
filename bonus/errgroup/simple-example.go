package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	urls := []string{
		"https://www.golang.org",
		"https://google.com",
		"https://x.com",
	}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Printf("Status da URL %s: %s\n", url, resp.Status)
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		fmt.Println("Erro ao processar as URLs:", err)
		return
	}

	fmt.Println("Todas as URLs foram processadas com sucesso")
}
