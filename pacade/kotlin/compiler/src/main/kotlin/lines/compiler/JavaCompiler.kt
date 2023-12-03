package lines.compiler

import lines.builder.CompilerBuilder

class JavaCompiler(var compiler: CompilerBuilder) : lines.domain.Compiler {
    override fun Compile() {
        compiler.node.Process()

        compiler.parser.Process()

        compiler.scanner.Process()
    }
}