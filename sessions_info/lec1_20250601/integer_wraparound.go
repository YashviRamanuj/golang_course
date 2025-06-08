package main
import "fmt"
func main() {
	// var maxUint32 uint32 = 4294967295 // Max uint32 size
	// fmt.Println(maxUint32)

	// maxUint32++
	// fmt.Println(maxUint32)


	var maxUint32 uint64 = 4294967295 + 1
	fmt.Println(maxUint32)
}