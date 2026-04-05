package mrdis  

import (
	"errors"
)
func (m *mrdis) Get (key string) interface {}{
	if key == ""{
		return errors.New("Missing Key")
	}
	return m.Data[key]
}