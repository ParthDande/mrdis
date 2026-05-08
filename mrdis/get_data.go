package mrdis  

import (
	"errors"
)
func (m *mrdis) Get (key string) interface {}{
	if key == ""{
		return errors.New("Missing Key")
	}
    m.mu.RLock() // put lock for Reader 
    defer m.mu.RUnlock()
	return m.Data[key]
}
