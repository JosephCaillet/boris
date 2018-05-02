package main

import "github.com/JosephCaillet/boris/musicorganizer"

func main() {
	musicorganizer.LoadConfiguration("conf.hjson")
	musicorganizer.Reorganize()
}
