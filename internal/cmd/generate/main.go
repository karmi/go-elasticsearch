package main

import (
	"github.com/karmi/go-elasticsearch/internal/cmd/generate/commands"
	_ "github.com/karmi/go-elasticsearch/internal/cmd/generate/commands/gensource"
	_ "github.com/karmi/go-elasticsearch/internal/cmd/generate/commands/gentests"
)

func main() {
	commands.Execute()
}
