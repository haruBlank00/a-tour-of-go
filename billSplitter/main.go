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

func extractItems(items string) []itemType {
	subItems := strings.Split(items, ",") // ["momo/150", "fanta/100"]

	var parsedItems Items
	var total float64
	for _, item := range subItems {
		item := strings.Split(item, "/") // ["momo", "150"]
		if len(item) != 2 {
			fmt.Println("Invalid item format:", item)
			continue
		}

		name, price := item[0], item[1]

		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Error reading input: ", err, name, price)
		}

		var newItem = itemType{
			name:  strings.TrimSpace(name),
			price: priceFloat,
		}

		total += newItem.price

		parsedItems = append(parsedItems, newItem)
	}
	return parsedItems
}

func sum(items []itemType) float64 {
	var total float64

	for _, item := range items {
		total += item.price
	}
	return total
}

func split(total, noOfPeople float64) float64 {
	return total / noOfPeople
}

func main() {

	var inputItems, inputNoOfPeople = "", "0"

	for {
		fmt.Println("Enter the items and price")
		fmt.Print("For eg: momo/150, fanta/100: ")
		_, err := fmt.Scan(&inputItems)

		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}
		break
	}

	for {
		fmt.Println("\nEnter the number of people")
		_, err := fmt.Scan(&inputNoOfPeople)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			fmt.Println("Please enter a valid input.")
			continue
		}
		break
	}

	noOfPeople, err := strconv.ParseFloat(inputNoOfPeople, 64)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		fmt.Println("Press enter to exit...")
		fmt.Scanln()
		fmt.Println("Exiting...")
	}
	itemsStruct := extractItems(inputItems)
	totalBill := sum(itemsStruct)
	perPerson := split(totalBill, noOfPeople)

	fmt.Println("Total bill: ", totalBill)
	fmt.Println("Bill per person: ", perPerson)

	fmt.Println("Press enter to exit...")
	fmt.Scanln()
	fmt.Println("Exiting...")
}
