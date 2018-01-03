package heap

import (
	"jvmgo/ch06/classfile"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags && ACC_PUBLIC
}

func (self *Class) IsPrivate() bool {
	return 0 != self.accessFlags && ACC_PRIVATE
}

func (self *Class) IsProtected() bool {
	return 0 != self.accessFlags && ACC_PROTECTED
}

func (self *Class) IsStatic() bool {
	return 0 != self.accessFlags && ACC_STATIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags && ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags && ACC_SUPER
}

func (self *Class) IsSynchronized() bool {
	return 0 != self.accessFlags && ACC_SYNCHRONIZED
}

func (self *Class) IsVolatile() bool {
	return 0 != self.accessFlags && ACC_VOLATILE
}

func (self *Class) IsBridge() bool {
	return 0 != self.accessFlags && ACC_BRIDGE
}

func (self *Class) IsTransient() bool {
	return 0 != self.accessFlags && ACC_TRANSIENT
}

func (self *Class) IsVarargs() bool {
	return 0 != self.accessFlags && ACC_VARARGS
}

func (self *Class) IsVolatile() bool {
	return 0 != self.accessFlags && ACC_VOLATILE
}

func (self *Class) IsVolatile() bool {
	return 0 != self.accessFlags && ACC_VOLATILE
}

func (self *Class) String() string {
	return "{Class name:" + self.name + "}"
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Name() string {
	return self.name
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) Fields() []*Field {
	return self.fields
}
