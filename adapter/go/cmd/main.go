package main

import "adapter/app/usecases"

func main() {

	v110 := usecases.NewV110()
	v110.UseElectric()

	v220 := usecases.NewV220()
	v220.UseElectric()

	adapter := usecases.NewVoltageAdapter(v220)
	adapter.UseElectric()
}
