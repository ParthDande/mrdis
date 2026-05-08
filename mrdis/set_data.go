package mrdis 

import (
	"errors"
)
func (m *mrdis) Set (key string , value interface{}) error { 
	if key == "" {
        return errors.New("key cannot be empty")
    }
    m.mu.Lock() // put the lock for writer . set is a write operation . 
    defer m.mu.Unlock()
    m.Data[key] = value
    return nil
}
