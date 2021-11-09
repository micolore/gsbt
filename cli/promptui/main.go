package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func ExecSelectUi() {
	prompt := promptui.Select{
		Label: "Select Init Type",
		Items: []string{"init_license", "update_service_code", "verify_license"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose %q\n", result)

}
