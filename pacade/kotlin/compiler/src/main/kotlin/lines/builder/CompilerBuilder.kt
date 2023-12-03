package lines.builder

import lines.domain.Element

data class CompilerBuilder (
    var parser: Element,
    var scanner: Element,
    var node: Element
)