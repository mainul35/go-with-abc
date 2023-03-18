package main

import "fmt"

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
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
	Sunday    = 7
)

const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Access granted")
}

func accessRevoked() {
	fmt.Println("Access denied")
}

func weekDay(day int) bool {
	return day < 6
}
func main() {

	today, role := Sunday, Contractor

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekDay(today) {
		accessGranted()
	} else if role == Member && weekDay(today) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessRevoked()
	}
}
