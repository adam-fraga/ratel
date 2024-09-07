package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintSuccessMsg(text string) {
	printer := color.New(color.FgGreen).SprintfFunc()
	msg := fmt.Sprintf("%s", text)
	fmt.Println(printer(msg))
}

func PrintInfoMsg(text string) {
	printer := color.New(color.FgBlue).SprintfFunc()
	msg := fmt.Sprintf("%s", text)
	fmt.Println(printer(msg))
}

func PrintCyanInfoMsg(text string) {
	printer := color.New(color.FgCyan).SprintfFunc()
	msg := fmt.Sprintf("%s", text)
	fmt.Println(printer(msg))
}

func PrintMagentaInfoMsg(text string) {
	printer := color.New(color.FgMagenta).SprintfFunc()
	msg := fmt.Sprintf("%s", text)
	fmt.Println(printer(msg))
}

func PrintWarningMsg(text string) {
	printer := color.New(color.FgYellow).SprintfFunc()
	msg := fmt.Sprintf("%s", text)
	fmt.Println(printer(msg))
}

func PrintErrorMsg(err string) {
	printer := color.New(color.FgRed).SprintfFunc()
	fmt.Println(printer(err))
}
