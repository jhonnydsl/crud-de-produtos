package main

import (
	"github.com/jhonnydsl/crud-go/model"
	"github.com/jhonnydsl/crud-go/services"
)

func main() {
	var produtos = make(map[int]model.Produto)

	services.AdicionarProduto("redmi", 1, 990.00, produtos)
	services.AdicionarProduto("Notebook", 2, 1350.00, produtos)
	services.AtualizaProduto("Notebook", 2, 1500.00, produtos)
	services.ListarProduto(produtos)

	services.AdicionarProduto("TV", 3, 2000.00, produtos)
	services.DeletaProduto(3, produtos)
	services.ListarProduto(produtos)
}