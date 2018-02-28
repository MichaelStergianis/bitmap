package bitmap

// Bitmap ...
// The bitmap data structure
type Bitmap struct {
	Bits     []uint32
	Size     uint32
	Buckets  uint32
	HashFunc func([]uint32) uint32
}

//////////////////// Bitwise operations ////////////////////

// pow2
// raises 2 to the power of n using bitwise operations
func pow2(n uint32) uint32 {
	var k uint32
	k = 1 << n
	return k
}

// log2
// returns the ceiling of log base 2 of n
// Does so by returning the first non zero bit of the integer plus one
func log2(n uint32) uint32 {
	for i := uint32(31); i > 0; i-- {
		p := pow2(i)
		if ((n & p) >> i) == 1 {
			return (i + 1)
		}
	}
	return 0
}

//////////////////// Hashing Functions ////////////////////

// PhiHash ...
// Uses the golden ratio to hash the numbers into buckets.
func PhiHash(group []uint32) uint32 {
	return 0
}

// MHash ...
// Uses bitwise complement to vary the buckets silghtly.
// For all odd indices, takes complement and does exclusive or with result
func MHash(group []uint32) uint32 {
	var res uint32
	res = 0
	for i, n := range group {
		if i%2 == 1 {
			res = res ^ (^n)
		} else {
			res = res ^ n
		}
	}
	return res
}

// OrHash ...
// Takes the bitwise or of both constituents of a [2]int slice and returns the result
func OrHash(group []uint32) uint32 {
	var res uint32
	res = 0
	for _, n := range group {
		res = res | n
	}
	return res
}

//////////////////// Bitmap Functions ////////////////////

// NewBitmap ...
// Creates a bitmap out of the counts prepared earlier
func NewBitmap(length uint32, hashFunc string) *Bitmap {
	hashes := map[string]func([]uint32) uint32{"OrHash": OrHash, "MHash": MHash, "PhiHash": PhiHash}
	var b *Bitmap
	b = new(Bitmap)
	nBits := pow2(log2(length))
	nInts := (nBits / 32)

	// initialize all of the bits
	b.Size = nInts
	b.Buckets = length
	b.Bits = make([]uint32, nInts, nInts)

	// set function
	if hashes[hashFunc] == nil {
		panic("The hashing function " + hashFunc + " is not supported.")
	}
	b.HashFunc = hashes[hashFunc]

	return b
}

// getUnhashed ...
// Looks up the bit in the bitmap. Returns either 0 or 1.
func (b *Bitmap) getUnhashed(index uint32) uint32 {
	// want to look up the ith bit
	// each int is 32 bits
	i0 := uint32(index / 32)
	i1 := uint32(index % 32)
	return (b.Bits[i0] & pow2(i1)) >> i1
}

// Get ...
// Looks up the corresponding bit in the bitmap by taking the hash of the array. Returns either 0 or 1.
func (b *Bitmap) Get(items []uint32) uint32 {
	idx := b.HashFunc(items) % b.Buckets
	return b.getUnhashed(idx)
}

// Set ...
// Sets the index of the Bitmap to 1.
func (b *Bitmap) Set(index uint32) {
	i0 := uint32(index / 32)
	i1 := uint32(index % 32)
	b.Bits[i0] = b.Bits[i0] | pow2(i1)
}

// Populate ...
// populates a Bitmap using counts
func (b *Bitmap) Populate(count []uint32, support uint32) {
	for i, v := range count {
		if v >= support {
			b.Set(uint32(i))
		}
	}
}
