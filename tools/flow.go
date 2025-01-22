package tools

import (
	"reflect"
	"sort"
)

// var non_overlap_paths [][]int

func Flow(paths [][]string) map[int][][]string {
	var allcompatiblepaths [][]int

	// Make a map so we can store the paths with an index id.
	pathsmap := make(map[int][]string)
	for a, b := range paths {
		// Remove start and end from path.
		pathsmap[a] = b[1 : len(b)-1]
		// fmt.Println("Path ID:", a, ":", pathsmap[a]) // Prints Path ID and path contents.
	}

	var keys []int
	for key := range pathsmap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	// We have to order these otherwise the order of operations doesn't always start with index 0 and will skip checking some paths.

	for _, keya := range keys {
		path1 := pathsmap[keya]

		var compatiblepaths []int
		compatiblepaths = append(compatiblepaths, keya)

		checkrooms := []string{}
		checkrooms = append(checkrooms, path1...)
		// fmt.Println("Check these rooms A:", checkrooms) // Prints the rooms that we are checking against.

		for _, keyb := range keys {
			path2 := pathsmap[keyb]
			if keya != keyb {
				// fmt.Println("a:", keya, "b:", keyb) // Prints the Path ID.
				// fmt.Println("PATHS", checkrooms, pathsmap[keyb]) // Prints the rooms we are checking against, and what we are checking.
				if NoOverlap(checkrooms, path2) {
					// fmt.Println("Non-overlapping paths:", path1, path2) // Prints if the paths are non-overlapping, and the contents of the paths
					compatiblepaths = append(compatiblepaths, keyb)
					checkrooms = append(checkrooms, path2...)
					// fmt.Println("Check these rooms B:", checkrooms) // Prints the rooms that we are checking against.
				} /* else {
					fmt.Println("Paths overlapped") // Prints if the paths overlap.
				} */
			}
		}
		allcompatiblepaths = append(allcompatiblepaths, compatiblepaths)
		// fmt.Println("END OF PATHS") // Prints when we have checked all overlapping paths from chosen starting path.
	}

	allcompatiblepaths = RemoveDuplicates(allcompatiblepaths)
	// fmt.Println("FINALPATHS", allcompatiblepaths) // Prints all the compatible path combinations, but only Path ID.

	result := make(map[int][][]string)

	for i, a := range allcompatiblepaths {
		var combo [][]string
		// fmt.Print("Compatible paths combo:")
		for _, b := range a {
			combo = append(combo, pathsmap[b])
			// fmt.Print(pathsmap[b]) // Prints the compatible path combinations.
		}
		result[i] = combo
		// fmt.Println()
	}
	// fmt.Println(result)
	return result
}

func NoOverlap(checkrooms, path2 []string) bool {
	var no_overlap bool
	for _, room1 := range checkrooms {

		no_overlap = true
		for _, room2 := range path2 {
			// fmt.Println("ROOMS:", room1, room2)
			if room1 == room2 {
				no_overlap = false
				// fmt.Println(no_overlap)
				return false
			}
			// fmt.Println(no_overlap)
		}
	}
	return no_overlap
}

func RemoveDuplicates(allcompatiblepaths [][]int) [][]int {
	var list [][]int
	for _, compatiblepath := range allcompatiblepaths {
		sort.Ints(compatiblepath)
		if !Contains(list, compatiblepath) {
			list = append(list, compatiblepath)
		}
	}
	return list
}

func Contains(list [][]int, compatiblepath []int) bool {
	for _, item := range list {
		if reflect.DeepEqual(item, compatiblepath) {
			return true
		}
	}
	return false
}
