package com.printer

import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    repeat(50_000) {
        launch {
            val printerInstance = SingletonPrinter.getPrinterInstance()

            print("Memory Address is $printerInstance \r\n")
        }
        launch {
            val printerInstance = SingletonPrinter.getPrinterInstance()

            print("Memory Address is $printerInstance \r\n")
        }
        launch {
            val printerInstance = SingletonPrinter.getPrinterInstance()

            print("Memory Address is $printerInstance \n")
        }
        launch {
            val printInstance = SingletonPrinter2.getPrintInstance()

            print("Memory Address is $printInstance \n")
        }
        launch {
            val printInstance = SingletonPrinter2.getPrintInstance()

            print("Memory Address is $printInstance \n")
        }
        launch {
            val printInstance = SingletonPrinter2.getPrintInstance()

            print("Memory Address is $printInstance \n")
        }
    }
}