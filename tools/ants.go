package tools

import (
	"fmt"
)

type Ant struct {
	Id       int
	Path     []string
	Travel   int
	Distance int
}

type PathCompare struct {
	Id    int
	Path  []string
	Queue int
}

var pathcompare = make(map[int]PathCompare)

func AntAssignToPath(optimalpath [][]string, ants int, steps int) string {
	antpath := make(map[int]Ant)

	// fill our ant info map with each ant (no other info at the moment)
	for l := 1; l <= ants; l++ {
		antpath[l] = Ant{
			Id: l,
		}
	}

	// create a map to keep track of the available paths and their queues.
	for id := range optimalpath {
		optimalpath[id] = optimalpath[id][1:]
		pathcompare[id] = PathCompare{
			Id:    id,
			Path:  optimalpath[id],
			Queue: len(optimalpath[id]),
		}
	}

	for l := 0; l < len(antpath); l++ {
		// for each ant, check which path has the shortest queue. Give ShortestPath func the info on our paths to compare the queue length.
		shortest := ShortestPath(pathcompare)
		// assign this information to the ant.
		antpath[l] = Ant{
			Id:       antpath[l].Id,
			Path:     shortest.Path,
			Travel:   0,
			Distance: len(shortest.Path),
		}
		// since we have added an ant to this queue, +1 to the queue for next time, we need to for loop to find the id again.
		for id := range pathcompare {
			if id == shortest.Id {
				pathcompare[id] = PathCompare{
					Id:    pathcompare[id].Id,
					Path:  pathcompare[id].Path,
					Queue: pathcompare[id].Queue + 1,
				}
			}
		}
	}
	return SendInTheAnts(antpath, pathcompare, steps, ants)
}

func ShortestPath(pathcompare map[int]PathCompare) PathCompare {
	var shortest PathCompare
	shortest = pathcompare[0]
	for id := range pathcompare {
		if pathcompare[id].Queue < shortest.Queue {
			shortest = pathcompare[id]
		}
	}
	return shortest
}

func SendInTheAnts(antpath map[int]Ant, pathcompare map[int]PathCompare, steps, ants int) string {
	// fmt.Println("STEPS:", steps, "\nPATHCOMPARE:", pathcompare, "\nANTPATH:", antpath)

	// make a map to store whether a room is free or busy.
	busyrooms := make(map[string]bool)

	// make a map to store whether a tunnel is free or busy.
	busytunnels := make(map[string]bool)

	// we need to know what the end room is, so create a variable to store it.
	var end string

	// loop through the paths and set all rooms to not busy.
	// also while we are here, grab the end room and save it to end variable.
	// also go through the paths and set all the tunnels to not busy.
	// if there is a tunnel that directly connects the start room with the end room (path length will be 1),
	// this is the main thing we care about for using the busytunnel map.

	for _, path := range pathcompare {
		for i, r := range path.Path {
			busyrooms[r] = false
			if i == len(path.Path)-1 {
				end = r
			}
			if i == 0 {
				busytunnels["start"+r] = false
			} else {
				busytunnels[path.Path[i-1]+r] = false
			}
		}
	}

	// to store what will be printed.
	var print string

	// for each step:
	for s := 1; s <= steps; s++ {
		// set all rooms to not busy.
		for r := range busyrooms {
			busyrooms[r] = false
		}
		// and set all tunnels to not busy.
		for t := range busytunnels {
			busytunnels[t] = false
		}
		// for each ant (starting with ant 1 always):
		for l := 1; l <= ants; l++ {
			// if the ant has not reached its destination:
			if antpath[l].Travel < antpath[l].Distance {
				// fmt.Println(antpath[l])
				// if the room the ant is moving to is not busy.
				if !busyrooms[antpath[l].Path[antpath[l].Travel]] && !busytunnels["start"+antpath[l].Path[antpath[l].Travel]] {
					// if the start-end tunnel has been used, it cannot be used again in this step, so set it to busy.
					if len(antpath[l].Path) == 1 {
						busytunnels["start"+antpath[l].Path[antpath[l].Travel]] = true
					}
					// create the print line: this create a line that includes the ant number and its destination in this step.
					print = print + fmt.Sprintf("L%v-%v ", antpath[l].Id, antpath[l].Path[antpath[l].Travel])

					// the room the ant now occupies is set to busy.
					busyrooms[antpath[l].Path[antpath[l].Travel]] = true
					// the ants travel distance is updated with plus 1, everything else is left as it was.
					antpath[l] = Ant{
						Id:       antpath[l].Id,
						Path:     antpath[l].Path,
						Travel:   antpath[l].Travel + 1,
						Distance: antpath[l].Distance,
					}
					// set the end room to false because it can hold infinite ants.
					busyrooms[end] = false
				}
			}
		}
		// until we have reached the last step, add a newline to print on for each step.
		if s != steps {
			print = print + "\n"
		}
	}
	return print
}
