package core

import (
	"fmt"
	repo "github.com/linhovieira/gestao-bancaria/repositories"
)

var clientRepository = new(repo.ClientRepository)

func RunApp() {
	for {
		ClearScreen()
		ShowBanner()
		selectedMainMenuOption, _ := ShowMainMenu()
		switch selectedMainMenuOption {
		case 1:
			ClearScreen()
			ShowBanner()
			var selectedOption, _ = ShowClientMenu()
			if selectedOption > 0 {
				ClearScreen()
				ShowBanner()
				executeActionClient(selectedOption)
			}
		case 2:
			ClearScreen()
			ShowBanner()
			var selectedOption, _ = ShowAccountMenu()
			if selectedOption > 0 {
				ClearScreen()
				ShowBanner()
				fmt.Println("This content is in the development phase!")
			}
		case 3:
			ExitApp()
		}
	}
}

func executeActionClient(option int) {
	switch option {
	case 1:
		fmt.Println("Starting to register a new client...")
		var client, _ = clientRepository.Add()
		fmt.Println("-----------------------------------------")
		fmt.Println("New Client registered!!!")
		fmt.Println("-----------------------------------------")
		fmt.Printf("Name: %s\r\nCpf: %s\r\n", client.Name(), client.Cpf())
		fmt.Println("-----------------------------------------")
	case 2:
		fmt.Println("Starting to search clients...")
		var client, _ = clientRepository.Search()
		fmt.Println("Client founded!!!")
		fmt.Printf("Name: %s\r\nCpf: %s\r\n", client.Name(), client.Cpf())
	case 3:
		fmt.Println("Starting to remove a client...")
		_ = clientRepository.Delete()
		fmt.Println("Client removed!!!")
	default:
		return
	}
}
