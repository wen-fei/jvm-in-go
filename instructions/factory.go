package instructions

import (
	"fmt"
	"github.com/jvm-in-go/instructions/base"
	"github.com/jvm-in-go/instructions/constants"
)

var (
	nop         = &constants.NOP{}
	aconst_null = &constants.ACONST_NULL{}
	dconst_0    = &constants.DCONST_0{}
	dconst_1    = &constants.DCONST_1{}
	fconst_0    = &constants.FCONST_0{}
	fconst_1    = &constants.FCONST_1{}
	fconst_2    = &constants.FCONST_2{}
	iconst_m1   = &constants.ICONST_M1{}
	iconst_0    = &constants.ICONST_0{}
	iconst_1    = &constants.ICONST_1{}
	iconst_2    = &constants.ICONST_2{}
	iconst_3    = &constants.ICONST_3{}
	iconst_4    = &constants.ICONST_4{}
	iconst_5    = &constants.ICONST_5{}
	lconst_0    = &constants.LCONST_0{}
	lconst_1    = &constants.LCONST_1{}
	bipush      = &constants.BIPUSH{}
	sipush      = &constants.SIPUSH{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
