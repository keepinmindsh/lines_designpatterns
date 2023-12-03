package main

import (
	"compiler/app/builder"
	"compiler/app/compiler"
	"compiler/app/element"
)

func main() {
	compilerBuilder := builder.NewCompilerBuilder()

	compilerBuilder.Parser(element.NewParser())
	compilerBuilder.Scanner(element.NewScanner())
	compilerBuilder.ProgramNode(element.NewNode())

	compiler := compiler.NewCompiler(compilerBuilder)

	compiler.Compile()
}
