package libs

import "math/rand"

var Buffers []*GiBObject

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

type GiBObject struct {
	B [GiB]byte
}

func WasteMemory(gib int) {
	Buffers = make([]*GiBObject, 0, gib)
	for gib > 0 {
		o := new(GiBObject)
		rand.Read(o.B[:])
		Buffers = append(Buffers, o)
		gib -= 1
	}
}
