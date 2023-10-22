package core

import "fmt"

func ShowBanner() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("\t\tBank Management")
	fmt.Println("-------------------------------------------------")
}

func ShowMainMenu() (int, error) {
	fmt.Println("- Main Menu")
	fmt.Println("-------------------------------------------------")
	fmt.Println("- [1] - Clients")
	fmt.Println("- [2] - Accounts")
	fmt.Println("- [3] - Exit")
	fmt.Println("-------------------------------------------------")
	return CollectOption(3)
}

func ShowClientMenu() (int, error) {
	fmt.Println("- Client Menu")
	fmt.Println("-------------------------------------------------")
	fmt.Println("- [1] - Add")
	fmt.Println("- [2] - View")
	fmt.Println("- [3] - Remove")
	fmt.Println("- [4] - Back")
	fmt.Println("-------------------------------------------------")
	return CollectOption(4)
}

func ShowAccountMenu() (int, error) {
	fmt.Println("- Account Menu")
	fmt.Println("-------------------------------------------------")
	fmt.Println("- [1] - Add")
	fmt.Println("- [2] - View")
	fmt.Println("- [3] - Remove")
	fmt.Println("- [4] - Back")
	fmt.Println("-------------------------------------------------")
	return CollectOption(4)
}
