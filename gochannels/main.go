package main

import (
	"fmt"
	"math/rand"
)

func broadcast(c chan int) {
	// infinite loop to create a random numbers
	for {
		//generate a random number 0-999 and put it into the channel
		c <- rand.Intn(3)
	}
}

func generateAccountNumber(accountNumberChannel chan int) {
	var accountNumber int
	accountNumber = 30000001

	for {
		if accountNumber > 30000005 {
			close(accountNumberChannel)
			return
		}
		accountNumberChannel <- accountNumber
		accountNumber += 1
	}

}

func main() {

	//numbersStation := make(chan int)
	//
	//// execute broadcast in a separate thread
	//
	//go broadcast(numbersStation)
	//
	//// retrieve values from the channel
	//
	//for num := range numbersStation {
	//	// delay for artistic effect only
	//	time.Sleep(3 * time.Millisecond)
	//	fmt.Printf("%d ", num)
	//}

	accountNumberChannel := make(chan int)

	go generateAccountNumber(accountNumberChannel)

	newCustomers := []string{"SMITH", "SINGH", "JONES", "LOPEZ", "CLARK", "ALLEN"}
	for _, newCustomer := range newCustomers {
		accnum, ok := <-accountNumberChannel
		if !ok {
			fmt.Printf("%s: “No number available\n", newCustomer)
		} else {
			fmt.Printf("%s: %d\n", newCustomer, accnum)
		}
	}

	//fmt.Printf("SMITH: %d\n", <-accountNumberChannel)
	//fmt.Printf("SINGH: %d\n", <-accountNumberChannel)
	//fmt.Printf("JONES: %d\n", <-accountNumberChannel)
	//fmt.Printf("LOPEZ: %d\n", <-accountNumberChannel)
	//fmt.Printf("CLARK: %d\n", <-accountNumberChannel)

}
