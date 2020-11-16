package main

func main() {
	foo([]byte{0x01, 0x02, 0x03, 0x04})
	bar([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06})
}

func foo(b []byte) {
	b[0] = 0 // Found IsInBounds
	b[1] = 1 // Found IsInBounds
	b[2] = 2 // Found IsInBounds
	b[3] = 3 // Found IsInBounds
}

func bar(b []byte) {
	_ = b[5] // Found IsInBounds
	b[0] = 0
	b[1] = 1
	b[2] = 2
	b[3] = 3
	b[4] = 4
	b[5] = 5
}
