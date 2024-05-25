package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintSuccessMsg(text string, action string) {
	printer := color.New(color.FgGreen).SprintfFunc()
	msg := fmt.Sprintf("%s %s successfully !", text, action)
	fmt.Println(printer(msg))
}

func PrintErrorMsg(text string, action string) {
	printer := color.New(color.FgRed).SprintfFunc()
	msg := fmt.Sprintf("%s %s !", text, action)
	fmt.Println(printer(msg))
}

func PrintInfoMsg(text string, action string) {
	printer := color.New(color.FgBlue).SprintfFunc()
	msg := fmt.Sprintf("%s %s !", text, action)
	fmt.Println(printer(msg))
}

func PrintWarningMsg(text string, action string) {
	printer := color.New(color.FgYellow).SprintfFunc()
	msg := fmt.Sprintf("%s %s !", text, action)
	fmt.Println(printer(msg))
}
