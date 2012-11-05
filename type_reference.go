package otto

type _reference interface {
	GetBase() *_object
	GetValue() Value
	PutValue(Value) bool
	Name() string
	Strict() bool
	Delete() bool
}

// Reference

type _reference_ struct {
    name string
	strict bool
}

func (self _reference_) GetBase() *_object {
	return nil
}

func (self _reference_) Name() string {
	return self.name
}

func (self _reference_) Strict() bool {
	return self.strict
}

func (self _reference_) Delete() {
	panic(hereBeDragons())
}

// PropertyReference

type _propertyReference struct {
	_reference_
    Base *_object
	node _node
}

func newPropertyReference(base *_object, name string, strict bool, node _node) *_propertyReference {
	return &_propertyReference{
		Base: base,
		_reference_: _reference_{
			name: name,
			strict: strict,
		},
		node: node,
	}
}

func (self *_propertyReference) GetBase() *_object {
	return self.Base
}

func (self *_propertyReference) GetValue() Value {
	if self.Base == nil {
		panic(newReferenceError("notDefined", self.name, self.node))
	}
	return self.Base.get(self.name)
}

func (self *_propertyReference) PutValue(value Value) bool {
	if self.Base == nil {
		return false
	}
	self.Base.set(self.name, value, self.Strict())
	return true
}

func (self *_propertyReference) Delete() bool {
	if self.Base == nil {
		// ?
		return false
	}
	return self.Base.delete(self.name, self.Strict())
}

// ArgumentReference

func newArgumentReference(base *_object, name string, strict bool) *_propertyReference {
	if base == nil {
		panic(hereBeDragons())
	}
	return newPropertyReference(base, name, strict, nil)
}

// getIdentifierReference

func getIdentifierReference(environment _environment, name string, strict bool, node _node) _reference {
	if environment == nil {
		return newPropertyReference(nil, name, strict, node)
	}
	if environment.HasBinding(name) {
		return environment.newReference(name, strict)
	}
	return getIdentifierReference(environment.Outer(), name, strict, node)
}
