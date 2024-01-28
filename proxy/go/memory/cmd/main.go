package main

import (
	"fmt"
	"memory/app/loader"
	"memory/domain"
)

func main() {

	objectMetaLeader := loader.MustNewObjectLoader(domain.META)

	for i := 0; i < 10; i++ {
		load := objectMetaLeader.Load().(string)

		fmt.Print(load)
	}

	fmt.Println()

	objectRealLeader := loader.MustNewObjectLoader(domain.REAL)

	for i := 0; i < 100; i++ {
		load := objectRealLeader.Load()

		fmt.Println(load)
	}
}
