package main

func reverseBits(num uint32) uint32 {
	var newNum uint32
	for _ = range 32 {
		lastBit := num & 1
		newNum <<= 1
		num >>= 1
		newNum += lastBit
	}
	return newNum
}
