package main

import "fmt"

// Flags is the data type for storing and retrieving flag binary numbers.
type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopBack
	FlagPointToPoint
	FlagMulticast
)

// IsUp tests whether the given flags object is turned up.
func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

// TurnDown turns down the given flags object.
func TurnDown(v *Flags) {
	*v &^= FlagUp
}

// SetBroadcast turns on the broadcast flag of the given flags object.
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

// IsCast tests whether the given flags object has a turned on any cast flag.
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))

	i1 := intelligent{V: 4, Z: 98}
	fmt.Println(i1.multiply())
}

type intelligent struct {
	V int
	Z int
}

func (i intelligent) multiply() int {
	return i.Z * i.V
}
