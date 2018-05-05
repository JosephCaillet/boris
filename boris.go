package main

import (
	"fmt"

	"github.com/JosephCaillet/boris/musicorganizer"
)

func main() {
	if err := musicorganizer.LoadConfiguration("conf.hjson"); err != nil {
		fmt.Printf("❌ failure loading configuration: %v", err)
		return
	}

	if err := musicorganizer.Reorganize(); err != nil {
		fmt.Printf("❌ failure reorganizing files: %v", err)
	}
}
