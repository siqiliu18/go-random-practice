package main

import (
	"fmt"
	"sort"
)

/**

Server names consist of an alphabetical host type (e.g. "sitebox") concatenated with the
server number. The availability of the server number follows the same rule as
part 1 (so "sitebox1", "sitebox2" are valid hostnames). Write a name tracking solution that
includes two operations - allocate(host) and deallocate(host). Allocate should reserve and
return the hostname with the lowest available ID and deallocate should release that hostname
for reuse.

server_number [5,3,2]
Example:
> tracker.allocate("sitebox") returns "sitebox1"
> tracker.allocate("apibox") returns "apibox1"
> tracker.allocate("sitebox") returns "sitebox2"
> tracker.allocate("apibox") returns "apibox2"
> tracker.deallocate("apibox1") returns null
> tracker.allocate("apibox") returns "apibox1"

this is unclear, after deallocate "apibox1", then allocate "sitbox", what is the output?

when deallocate, would server ID be down?
*/

var serverIDs = []int{}

func next_server_number(serverIDs []int) int {
	sort.SliceStable(serverIDs, func(i, j int) bool {
		return serverIDs[i] < serverIDs[j]
	})

	i := 0
	for ; i < len(serverIDs); i++ {
		if i+1 != serverIDs[i] {
			return i + 1
		}
	}
	return i + 1
}

func allocate(name string) string {

	return ""
}

func main() {
	fmt.Println(next_server_number([]int{5, 3, 1}))    // 2
	fmt.Println(next_server_number([]int{5, 4, 1, 2})) // 3
	fmt.Println(next_server_number([]int{3, 2, 1}))    // 4
	fmt.Println(next_server_number([]int{}))           // 1
	fmt.Println(next_server_number([]int{100}))        // 1

	// fmt.Println(allocate("A")) // A1
	// fmt.Println(allocate("B")) // B1
	// fmt.Println(allocate("A")) // A2
	// fmt.Println(allocate("A")) // A3
	// fmt.Println(allocate("A")) // A4
	// fmt.Println(allocate("B")) // B2
	// deallocate(A1)
	// allocate(A) // A1
}
