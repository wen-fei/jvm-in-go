package rtda

import "github.com/jvm-in-go/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
