package main

import (
	"fmt"
	"os"

	bubble "github.com/charmbracelet/bubbletea"
	"github.com/ysrckr/currency-converter-cli/internal/tea"
)

func main() {
	p := bubble.NewProgram(tea.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
