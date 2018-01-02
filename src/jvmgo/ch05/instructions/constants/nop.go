package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Excute(frame *rtda.Frame) {
	// nothing to do
}
