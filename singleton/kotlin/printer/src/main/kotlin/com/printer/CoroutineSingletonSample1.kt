package com.printer

import kotlinx.coroutines.*
import java.util.Random

// Concurrently executes both sections
suspend fun doWorld() = coroutineScope { // this: CoroutineScope
    launch {
        delay(2000L)
        println("World ${Random().nextInt()}")
    }
    println("Hello")
}