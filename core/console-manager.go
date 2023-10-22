package core

import (
	"fmt"
	"os"
	"os/exec"
)

func CollectOption(limit int) (int, error) {
	fmt.Print("Enter with option number: ")
	var option int
	_, err := fmt.Scan(&option)
	if err != nil {
		fmt.Println("An error occurred while trying to get the selected option!")
		fmt.Println("Error: ", err)
		return 0, err
	}
	if option > limit {
		fmt.Println("The chosen option is invalid")
		return 0, fmt.Errorf("the chosen option [%d] is invalid", option)
	}
	return option, nil
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func ExitApp() {
	fmt.Println("Exiting application...")
	os.Exit(-1)
}
