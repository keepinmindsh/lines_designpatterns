package com.printer

class SingletonPrinter2 private constructor() {
    companion object {
        private var printer: Printer? = null

        fun getPrintInstance(): Printer {
            return printer ?: synchronized(this) {
                printer ?: SamsungPrinter()
            }
        }
    }
}