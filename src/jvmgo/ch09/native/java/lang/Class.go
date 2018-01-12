package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/rtda/heap"
)

func init(){
	native.Register("java/lang/Class","getPrimitiveClass","(Ljava/lang/String;)Ljava/lang/Class;",getPrimitiveClass)
	native.Register("java/lang/Class","getName0","()Ljava/lang/String;",getName0)
	native.Register("java/lang/Class","desiredAssertionStatus0","(Ljava/lang/Class;)Z",desiredAssertionStatus)
}


func getPrimitiveClass(frame *rtda.Frame){
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

func getName0(frame *rtda.Frame){
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(),name)
	frame.OperandStack().PushRef(nameObj)
}


func desiredAssertionStatus(frame *rtda.Frame){
	frame.OperandStack().PushBoolean(false)
}