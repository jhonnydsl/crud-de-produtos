package services

import (
	"fmt"

	"github.com/jhonnydsl/crud-go/model"
)

func AdicionarProduto(nome string, Nid int, preco float64, _produto map[int]model.Produto) {
	novoProduto := model.Produto{
		Nome:  nome,
		Id:    Nid,
		Preco: preco,
	}
	if _, existe := _produto[Nid]; existe {
		fmt.Println("Ja existe um produto com esse Id")
		return
	}
	_produto[Nid] = novoProduto
}

func ListarProduto(_produto map[int]model.Produto) {
	for _, valor := range _produto {
		fmt.Printf("ID: %d | nome: %s | preço: %.2f\n",valor.Id, valor.Nome, valor.Preco)
	}	
}

func AtualizaProduto(nome string, id int, preco float64, _produto map[int]model.Produto) {
	produtoAtualizado := model.Produto {
		Nome: nome,
		Id: id,
		Preco : preco,
	}
	if _, existe := _produto[id]; !existe {
		fmt.Println("Não existe esse produto para ser atualizado!")
		return
	}
	_produto[id] = produtoAtualizado
}

func DeletaProduto(id int, _produto map[int]model.Produto) {
	if _, existe := _produto[id]; !existe {
		fmt.Println("Esse produto informado não existe")
		return
	}
	delete(_produto, id)
}
