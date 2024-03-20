/*
Let's make a console app to split bill

Input: comma seperated price/item, price/items
Another Input: Number of people
Output: bill per person
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type itemType struct {
	name  string
	price float64
}

type Items []itemType

func splitBill(items, noOfPeople string) float64 {
	/*
		- split string by comma ','
		- for each item, split by '/' and make a struct of price and item name
	*/

	subItems := strings.Split(items, ",") // ["momo/150", "fanta/100"]
	var parsedItems Items
	for _, item := range subItems {
		item := strings.Split(item, "/") // ["momo", "150"]
		name, price := item[0], item[1]

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Error reading input: ", err)
		}

		var newItem = itemType{
			name:  name,
			price: priceFloat,
		}

		parsedItems = append(parsedItems, newItem)
	}

	fmt.Println("Items: ", parsedItems)
	return 5.5

}
func main() {
	// var totalBill float64

	var items, noOfPeople string

	fmt.Println("Enter the items and price")
	fmt.Println("For eg: momo/150, fanta/100")
	_, err := fmt.Scan(&items)

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	for {
		fmt.Println("Enter the number of people")
		_, err = fmt.Scan(&noOfPeople)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			fmt.Println("Please enter a valid input.")
			continue
		}
		break
	}

	splitBill(items, noOfPeople)

	fmt.Println("Items: ", items)
	fmt.Scanln("Press enter to exit...")
	fmt.Scanln()
	fmt.Println("Exiting...")
}
