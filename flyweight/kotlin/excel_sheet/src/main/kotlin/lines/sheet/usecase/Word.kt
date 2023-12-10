package lines.sheet.usecase

import lines.sheet.domain.Text

class WordFactory {
    companion object {
        fun GetWord(type: WordType) {

        }
    }
}

enum class WordType {
    A, B, C, D, E
}

class WordA : Text {
    override fun Draw() {
        TODO("Not yet implemented")
    }
}

class WordB : Text {
    override fun Draw() {
        TODO("Not yet implemented")
    }
}

class WordC : Text {
    override fun Draw() {
        TODO("Not yet implemented")
    }
}

class WordD : Text {
    override fun Draw() {
        TODO("Not yet implemented")
    }
}

class WordE : Text {
    override fun Draw() {
        TODO("Not yet implemented")
    }
}