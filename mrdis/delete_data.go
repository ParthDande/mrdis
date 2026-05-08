package mrdis 

import("errors")
func (m *mrdis) Delete (key string) error { 
	if key == ""{
		return errors.New("missing key")
	}
    m.mu.Lock() // put the writers lock . we are cahing the state of the data . 
    defer m.mu.Unlock()
	delete(m.Data, key)
	return nil 
}
