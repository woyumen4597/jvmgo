package heap

type Object struct {
	class  *Class
	data  interface{}
	extra interface{}
}

func (o *Object) Data() interface{} {
	return o.data
}

func (o *Object) SetData(data interface{}) {
	o.data = data
}

func (o *Object) Extra() interface{} {
	return o.extra
}

func (o *Object) SetExtra(extra interface{}) {
	o.extra = extra
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) SetRefVar(name,descriptor string,ref *Object){
	field := self.class.getField(name,descriptor,false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId,ref)
}

func (self *Object) GetRefVar(name,descriptor string) *Object{
	field := self.class.getField(name,descriptor,false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}