package lines

import lines.builder.CompilerBuilder
import lines.compiler.JavaCompiler
import lines.compiler.Node
import lines.compiler.Parser
import lines.compiler.Scanner
import lines.domain.Compiler

fun main() {
    val compiler : Compiler = JavaCompiler(CompilerBuilder(
        Node(),
        Parser(),
        Scanner(),
    ))

    compiler.Compile()
}
