package com.printer

object SingletonPrinter {
    private val printer = SamsungPrinter()

    fun getPrinterInstance() : Printer {
        return printer
    }
}

class SamsungPrinter : Printer {
    override fun PutPaperIn() {
        print("Put Paper In!")
    }

    override fun Print() {
        print("Print!")
    }
}

interface Printer {
    fun PutPaperIn()
    fun Print()
}