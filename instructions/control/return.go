package control

import "github.com/jvm-in-go/instructions/base"

type RETUREN struct {
	// return void from method
	base.NoOperandsInstruction
}

type ARETURN struct {
	// return reference from method
	base.NoOperandsInstruction
}

type DRETURN struct {
	// return double from method
	base.NoOperandsInstruction
}

type FRETURN struct {
	// return float from method
	base.NoOperandsInstruction
}

type IRETURN struct {
	// return int from method
	base.NoOperandsInstruction
}

type LRETURN struct {
	// return long from method
	base.NoOperandsInstruction
}
