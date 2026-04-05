package mrdis 

import("errors")
func (m *mrdis) Delete (key string) error { 
	if key == ""{
		return errors.New("missing key")
	}
	delete(m.Data, key)
	return nil 
}