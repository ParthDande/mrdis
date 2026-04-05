package mrdis 

import (
	"errors"
)
func (m *mrdis) Set (key string , value interface{}) error { 
	if key == "" {
        return errors.New("key cannot be empty")
    }

    m.Data[key] = value
    return nil
}