package heap

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	return Slots{}
}
