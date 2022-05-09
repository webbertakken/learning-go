//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

type Server string
type Status int

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
func printServerStatuses(serverStatuses map[Server]Status) {
	fmt.Printf("\n%34v\n", "--- Server status ---")
	for server, status := range serverStatuses {
		fmt.Printf("%18v is ", server)
		switch status {
		case Online:
			fmt.Println("online")
		case Offline:
			fmt.Println("offline")
		case Maintenance:
			fmt.Println("in maintenance")
		case Retired:
			fmt.Println("retired")
		}
	}
}

func main() {
	servers := []Server{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	serverStatuses := make(map[Server]Status)

	//* Set all the server statuses to `Online` when creating the map
	for _, server := range servers {
		serverStatuses[server] = Online
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	printServerStatuses(serverStatuses)

	//  - change server status of `darkstar` to `Retired`
	serverStatuses["darkstar"] = Retired

	//  - change server status of `aiur` to `Offline`
	serverStatuses["aiur"] = Offline

	//  - call display server info function
	printServerStatuses(serverStatuses)

	//  - change server status of all servers to `Maintenance`
	for server := range serverStatuses {
		serverStatuses[server] = Maintenance
	}

	//  - call display server info function
	printServerStatuses(serverStatuses)
}
