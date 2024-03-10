package result

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const defaultErrorKey = "__ERROR"

// P node input/output data transfer object
type P map[string]any

// Success looks for the "success" key.
func (p P) Success() bool {
	v, ok := (p["success"]).(bool)
	if !ok {
		return false
	}

	return v
}

func (p P) getValue(key string) (any, error) {
	v, exists := p[key]
	if !exists {
		return "", fmt.Errorf("key not found: %s", key)
	}

	return v, nil
}

// OutStr get string value from map
func (p P) OutStr(name string, out *string) P {
	v, err := p.getValue(name)
	if err != nil {
		p[defaultErrorKey] = err
		return p
	}

	str, cast := v.(string)
	if !cast {
		p[defaultErrorKey] = fmt.Errorf("key (%s) cannot cast to string", name)
		return p
	}

	*out = str
	return p
}

// OutStructure decodes value of given key to the target pointer specified
func (p P) OutStructure(name string, out any) P {
	v, err := p.getValue(name)
	if err != nil {
		p[defaultErrorKey] = err
		return p
	}

	if err := mapstructure.Decode(v, out); err != nil {
		p[defaultErrorKey] = err
		return p
	}

	return p
}

// OutStr get boolean value from map
func (p P) OutBool(name string, out *bool) P {
	v, err := p.getValue(name)
	if err != nil {
		p[defaultErrorKey] = err
		return p
	}

	b, cast := v.(bool)
	if !cast {
		p[defaultErrorKey] = fmt.Errorf("key (%s) cannot cast to boolean", name)
		return p
	}

	*out = b
	return p
}

func (p P) Error() error {
	err, ok := p[defaultErrorKey]
	if ok && err != nil {
		return err.(error)
	}

	return nil
}

// Out decodes P to the target pointer specified
func (p P) Out(out any) P {
	if err := mapstructure.Decode(p, out); err != nil {
		p[defaultErrorKey] = err
		return p
	}

	return p
}
