package print_console

import "fmt"

type car struct {
	color   string
	mileage int
}

func MainSlice() {
	cars := []car{
		{
			color:   "red",
			mileage: 5000,
		},
		{
			color:   "green",
			mileage: 10_000,
		},
		{
			color:   "blue",
			mileage: 7000,
		},
	}

	carPtr := cars[0]     // red, 5000
	carPtr.mileage += 100 // red, 5100

	cars = append(cars, car{color: "yellow", mileage: 15_000}) // ...cars,  {yellow, 15000}
	carPtr.mileage += 50                                       // 5050

	fmt.Println(cars[0].mileage, cars[0].color) // red, 5050
	fmt.Println(carPtr.mileage, carPtr.color)   // red, 5050
}
