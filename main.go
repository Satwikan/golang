package main

import (
	"booking-app/customPackage"
	"fmt"
	"strings"
	"sync"
	"time"
)

type structureInGo struct {
	memberOne   string
	memberTwo   int
	memberThree float32
}

var wg = sync.WaitGroup{}

func main() {
	// to run externalFunc type
	// go run main.go helper.go
	// in command line
	externalFunc()
	customPackage.CustomPackageFunc()

	var object = structureInGo{
		memberOne:   "1",
		memberTwo:   2,
		memberThree: 3,
	}
	fmt.Println("object in go", object)

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// slice (vector)
	var bookings []string

	greetUsers(conferenceName)

	fmt.Println("we have total off", conferenceTickets, " tickets with", remainingTickets, " tickets still available")
	fmt.Println("Get your tickets here to attend")
	var userName string
	var email string
	var userTickets uint

	// infinite loop
	for {
		fmt.Println("what should i call you master?")
		fmt.Scan(&userName)

		fmt.Print("how should i contact you master (email)? ")
		fmt.Scan(&email)

		fmt.Print("how many tickets you want to book master? ")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Printf("we have only %v tickets remaining please re-enter your details\n", remainingTickets)
			continue
		}

		remainingTickets -= userTickets
		bookings = append(bookings, userName)

		// concurrent func (executes on seprate thread)
		// but the main thread won't wait for any child thread to execute
		// to avoid this wait group is used
		wg.Add(1)
		go sendTickets(userTickets, userName, email)
		fmt.Printf("%v ticket(s) remaning\n", remainingTickets)

		// syntactic sugar for slice declaration
		firstNames := []string{}
		for _, booking := range bookings {
			// split with " "
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("all bookings:\n%v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("our conference is booked out. come back next year")
			break
		}

	}

	city := "London"
	// switch in go
	switch city {
	case "New York":
		// bookinLondon()
	case "Singapore":
		// bookinSingapore()
	case "Mumbai", "Jaipur":
		// bookinJaipur()
	}

	var dictionary = make(map[string]int)
	dictionary["key-0"] = 0
	// wait for all wait group threads to finish
	wg.Wait()
}

func greetUsers(conferenceName string) {
	fmt.Printf("Welcome to %v, Master\n", conferenceName)
}

func sendTickets(userTickets uint, userName string, email string) {
	// simulating real world response time (10 seconds)
	time.Sleep(10 * time.Second)
	// save ticket as a string
	var ticket = fmt.Sprintf("you have booked %v ticket(s) by name of %v, you will receive your bill on %v\n", userTickets, userName, email)
	fmt.Println(ticket)
	// Inform wait group that this function is executed
	wg.Done()
}
