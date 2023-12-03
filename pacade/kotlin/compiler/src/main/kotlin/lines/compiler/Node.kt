package lines.compiler

import lines.domain.Element

class Node : Element {
    override fun Process() {
        println("Node is processing")
    }
}