package service

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Service_Singleton(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {

			fmt.Printf("Address : %v \r\n", NewGetPrinterSingleton_Way3())

			wg.Done()
		}()
	}

	wg.Wait()
}

func Test_Service_Singleton2(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {

			fmt.Printf("Address : %v \r\n", NewGetPrintSingleton_Way2())

			wg.Done()
		}()
	}

	wg.Wait()
}

func Test_Service_Singleton1(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {

			fmt.Printf("Address : %v \r\n", NewGetPrinterNotSingleton_Way1())

			wg.Done()
		}()
	}

	wg.Wait()
}
