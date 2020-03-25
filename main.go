package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)
//writer from: http://networkbit.ch/golang-column-print/

func main() {
	writer := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	writer.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer writer.Flush()

	_, _ = fmt.Fprintf(writer, "\n %s\t\t%s\t%s\t%s\t", "Spaceline", "Days", "Trip Type", "Price")
	_, _ = fmt.Fprintf(writer, "\n %s\t\t%s\t%s\t%s\t%s", "----", "----", "----", "----", "\n")

	for i := 0; i < 10; i++ {
		speed := getSpeed()
		tripType := getTripType()
		_, _ = fmt.Fprintf(writer, "%s\t\t%d\t%s\t$%d\t\n", getSpaceline(),
			calculateDays(speed),
			tripType,
			getPrice(speed, tripType))
	}
}

func getPrice(speed int, tripType string) int {
	rand.Seed(time.Now().UnixNano())
	minPrice := 36
	maxPrice := 50
	if speed > 23 {
		minPrice = 42
	} else {
		maxPrice = 42
	}
	price := rand.Intn(maxPrice - minPrice + 1) + minPrice
	if tripType == "round trip" {
		return price * 2
	} else {
		return price
	}
}

func getTripType() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(2) +1
	switch num {
		case 1: return "one way"
		case 2: return "round trip"
		default: return "bad number"
	}
}

func calculateDays(speed int) int {
	distance := 62100000

	return distance/speed/86400
}

func getSpeed() int {
	rand.Seed(time.Now().UnixNano())
	//Randomly choose the speed the ship will travel, from 16 to 30 km/s.
	min := 16
	max := 30
	speed := rand.Intn(max - min + 1) + min
	return speed

}

func getSpaceline() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(3) + 1

	switch num {
		case 1: return "Space Adventures"
		case 2: return "Space X"
		case 3: return "Virgin Galactic"
		default: return "badNumber"
	}
}
