package heap

func lookupMethodInInterfaces(ifaces []*Class, name string, descriptor string) *Method {
	for _,iface := range ifaces{
		for _,method := range iface.methods{
			if method.name == name && method.descriptor == descriptor{
				return method
			}
		}
	}
	return nil
}


func LookupMethodInClass(class *Class, name string, descriptor string) *Method {
	for c := class;c!=nil;c = c.superClass{
		for _,method := range c.methods{
			if method.name == name && method.descriptor == descriptor{
				return method
			}
		}
	}
	return nil
}

