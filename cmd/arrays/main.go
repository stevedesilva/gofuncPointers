package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("..... ARRAYS")
	arrays()
	fmt.Println("..... SLICES")
	slices()
	fmt.Println("..... MAPS")
	maps()
	fmt.Println("..... STRUCTS")
	structs()
}

func arrays() {
	nums := [...]int{1, 2, 3}

	fmt.Printf("\nOriginal nums : %p - nums %v \n", &nums, nums)
	s.Show("nums orig >", nums)
	// whole array will be copied
	incr(nums)
	fmt.Printf("\nOriginal nums : %p - nums %v \n", &nums, nums)

	incrByPtr(&nums)
	fmt.Printf("\nOriginal nums : %p - nums %v \n", &nums, nums)
}

func incr(nums [3]int) {
	// copy of array - disconnect original array
	fmt.Printf("\tincr nums : %p\n", &nums)
	for i := range nums {
		nums[i]++
		fmt.Printf("\nincr.nums[%d] : %p\n", i, &nums[i])
	}
	fmt.Println("\tincr nums : ", nums)
}

func incrByPtr(nums *[3]int) {
	// copy of the pointer value
	fmt.Printf("\tincrByPtr nums : %p\n", &nums)
	for i := range nums {
		// go automatically derefs
		nums[i]++ // same as (*nums)[i]++
	}
	fmt.Println("\tincrByPtr nums : ", nums)
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}
	fmt.Printf("slices >\t : v %p \t, addr %p\n", dirs, &dirs)
	s.Show("slices >", dirs)
	up(dirs)
	fmt.Printf("slices <\t :  v %p \t, addr %p\n", dirs, &dirs)
	// s.Show("slices <", dirs)

	upPtr(&dirs)
	fmt.Printf("slices <\t :  v %p \t, addr %p\n", dirs, &dirs)
	// s.Show("slices <", dirs)
	fmt.Println(dirs)

}

func up(list []string) {
	// slice header(pointing to backing array) will have a local copied :. you can still change the slice values
	// s.Show("up >", list)
	fmt.Printf("up >\t : v %p \t, addr %p\n", list, &list)
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}
	list = append(list, "STEVE")
	// s.Show("up append <", list)
	fmt.Printf("up <\t : v %p \t, addr %p\n", list, &list)
}

func upPtr(list *[]string) {
	fmt.Printf("upPtr >\t : v %p \t, addr %p\n", list, &list)
	lv := *list
	fmt.Printf("lv >\t : v %p \t, addr %p\n", lv, &lv)
	// s.Show("upPtr >", *list)

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "STEVE")
	lv = *list
	fmt.Printf("lv <\t : v %p \t, addr %p\n", lv, &lv)
	fmt.Printf("append <\t : v %p \t, addr %p\n", *list, &*list)
	// s.Show("upPtr append <", *list)
	fmt.Printf("upPtr <\t : v %p \t, addr %p\n", list, &list)
}

func maps() {
	confused := map[string]int{"two": 1, "one": 2}
	fix(confused)
	fmt.Println(confused)

	// can't find the address of a map element because behind the scenes go can change the map element address
	// &confused["one"] = 2

}

func fix(m map[string]int) {
	s.Show("fix >", m)
	// map value (like a slice) is already a pointer
	// local copy of map(stored at different location) will store the same pointer value
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

type house struct {
	name  string
	rooms int
}

func structs() {

	myHouse := house{name: "My House", rooms: 5}
	addRooms(myHouse)
	fmt.Printf("structs()\t : %p %v\n", &myHouse, myHouse)

	addRoomsPtr(&myHouse)
	fmt.Printf("structs()\t : %p %v\n", &myHouse, myHouse)
}

func addRooms(h house) {
	// the struct is a copy
	h.rooms++
	fmt.Printf("addRooms()\t : %p %v\n", &h, h)
}

func addRoomsPtr(h *house) {
	// the struct is a copy
	h.rooms++ // go automatically defs (*h).rooms++
	fmt.Printf("addRoomsPtr()\t : %p %v\n", h, h)
	// structs stored in contiguous blocks (string value 16 bytes, int values 8 bytes)
	fmt.Printf("&h.name\t : %p %v\n", h, &h.name)
	fmt.Printf("&h.room\t : %p %v\n", h, &h.rooms)

}
