package main

import (
	"bufio"
	"fmt"
	"os"
)

//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

// days of the week as constants
const (
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
	Sunday    = "Sunday"
)

const (
	Admin      = "Admin"
	Manager    = "Manager"
	Contractor = "Contractor"
	Member     = "Member"
	Guest      = "Guest"
)

func accessGranted() {
	fmt.Println("Access granted")
}

func accessRevoked() {
	fmt.Println("Access denied")
}

func isWeekDay(day string) bool {
	return !isWeekend(day)
}

func isWeekend(day string) bool {
	return day == Saturday || day == Sunday
}
func main() {
	var today, role string

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("What day is it? (Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)")
		reader.Scan()
		today = reader.Text()
		fmt.Println("What is your role? (Admin, Manager, Contractor, Member, Guest)")
		reader.Scan()
		role = reader.Text()

		if role == Admin || role == Manager {
			accessGranted()
		} else if role == Contractor && isWeekend(today) {
			accessGranted()
		} else if role == Member && isWeekDay(today) {
			accessGranted()
		} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
			accessGranted()
		} else if role == "none" {
			exit()
			break
		} else {
			accessRevoked()
		}
	}

}

func exit() {
	fmt.Println("Exiting program")
}
