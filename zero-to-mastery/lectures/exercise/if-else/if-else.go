//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//

package main

import "fmt"

// Days of the week
//goland:noinspection GoUnusedConst
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

//goland:noinspection GoBoolExpressions
func main() {
	// The day and role. Change these to check your work.
	today, role := Tuesday, Guest

	if role == Admin || role == Manager {
		//* Access at any time: Admin, Manager
		accessGranted()
		return
	} else if role == Contractor {
		//* Access weekends: Contractor
		if today >= 5 {
			accessGranted()
			return
		}
	} else if role == Member {
		//* Access weekdays: Member
		if today <= 4 {
			accessGranted()
			return
		}
	} else if today == Monday || today == Wednesday || today == Friday {
		//* Access Mondays, Wednesdays, and Fridays: Guest
		accessGranted()
		return
	}

	accessDenied()
}
