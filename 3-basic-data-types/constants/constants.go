package main

// Weekday is the data type which holds the index of a weeday.
type Weekday uint

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Flags is the data type which holds the amount of net capability
type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point linke
	FlagMulticast                      // supports multicast access capability
)

// IsUp determines whether the set flags is "Flag Up"
func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

// TurnDown is
func TurnDown(v *Flags) {
	*v &^= FlagUp
}

// SetBroadcast is
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func main() {

}
