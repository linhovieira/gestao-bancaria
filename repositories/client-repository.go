package repositories

import (
	"fmt"
	mdl "github.com/linhovieira/gestao-bancaria/models"
	"slices"
)

type ClientRepository struct {
	clients []mdl.Client
}

func (cr *ClientRepository) Add() (mdl.Client, error) {
	fmt.Print("Enter with name: ")
	var name = ""
	var _, err = fmt.Scan(&name)
	if err != nil {
		fmt.Println("You enter with invalid name!!!")
		panic(err)
	}
	fmt.Print("Enter with cpf: ")
	var cpf = ""
	_, err = fmt.Scan(&cpf)
	if err != nil {
		fmt.Println("You enter with invalid cpf!!!")
		panic(err)
	}
	client := mdl.Client{}
	client.SetName(name)
	client.SetCpf(cpf)
	cr.clients = append(cr.clients, client)
	return client, nil
}

func (cr *ClientRepository) Search() (mdl.Client, error) {
	var cpf = ""
	fmt.Print("Enter with cpf: ")
	var _, err = fmt.Scan(&cpf)
	if err != nil {
		fmt.Println("You enter with invalid cpf!!!")
		panic(err)
	}
	for i := range cr.clients {
		if cr.clients[i].Cpf() == cpf {
			return cr.clients[i], err
		}
	}
	return mdl.Client{}, err
}

func (cr *ClientRepository) Delete() error {
	var cpf = ""
	fmt.Print("Enter with cpf: ")
	var _, err = fmt.Scan(&cpf)
	if err != nil {
		fmt.Println("You enter with invalid cpf!!!")
		panic(err)
	}
	index := slices.IndexFunc(cr.clients, func(client mdl.Client) bool { return client.Cpf() == cpf })
	if index > -1 {
		cr.clients = append(cr.clients[:index], cr.clients[index+1:]...)
	}
	return err
}
