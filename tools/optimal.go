package tools

import (
	"math"
)

func Optimal(paths [][]string, flowmap map[int][][]string, ants int) ([][]string, int) {
	// ! Another method:
	/* 	lengthmap := make(map[int][]int)
	   	var lengths []int

	   	for id, combo := range flowmap {
	   		lengths = nil
	   		for pathid := range combo {
	   			lengths = append(lengths, len(combo[pathid]))
	   		}
	   		lengthmap[id] = append(lengthmap[id], lengths...)
	   	}
	   	for id := range lengthmap {
	   		for L := 1; L <= ants; L++ {
	   			sort.Ints(lengthmap[id])
	   			lengthmap[id][0] += 1
	   			// fmt.Printf("L%v-%v\n", L, lengthmap[id])
	   		}
	   		fmt.Println("id:", id, "flowmap:", flowmap[id], "steps:", lengthmap[id][0])
	   	} */

	// this method uses an equation instead of iterating and updating.
	// this is the equation:
	// * ( number of ants / number of paths ) + ( number of rooms / number of paths ) = minimum steps !every value must be rounded up at stages of calculation!
	stepidmap := make(map[int]int)
	var counter float64
	counter = 0
	for id, combo := range flowmap {
		rooms := 0
		// get the number of rooms across all paths.
		for _, path := range combo {
			rooms += len(path)
		}
		// multistaged for more simple code.
		// * ( number of ants / number of paths )
		// every value must be rounded up at stages of calculation! that is what math.Ceil is doing. math.Ceil works on float64, not int, so we are converting to float64.
		ena := math.Ceil(float64(ants) / float64(len(combo)))
		// * ( number of rooms / number of paths )
		duo := math.Ceil(float64(rooms) / float64(len(combo)))

		// add them.
		// * ( number of ants / number of paths ) + ( number of rooms / number of paths )
		counter = math.Ceil(ena + duo)
		// counter is float64, we need int.
		steps := int(counter)

		// fmt.Println(flowmap[id], "steps", steps)
		// add our minimum steps value to the map to store all the path combos minimum step, and then we will compare and find the absolute optimal path!
		stepidmap[id] = steps
	}

	// fmt.Println(stepidmap)

	// get the combo with the lowest minimum steps.
	var minid int
	minvalue := stepidmap[0]

	// compares to each one, and finds the lowest.
	for id, steps := range stepidmap {
		if steps < minvalue {
			minvalue = steps
			minid = id
		}
	}

	var optimalpath [][]string
	// grab our start and end rooms again.
	start := paths[0][0]
	end := paths[0][len(paths[0])-1]

	for id := range flowmap {
		// if it is the combo with the lowest minimum steps.
		if id == minid {
			// assign this combo as our optimal combo.
			optimalpath = flowmap[id]
			// add on start and end rooms again.
			for i, path := range optimalpath {
				slc := []string{start}
				slc = append(slc, path...)
				slc = append(slc, end)
				optimalpath[i] = slc
			}
		}
	}
	return optimalpath, minvalue
}
