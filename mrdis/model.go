package mrdis
import ( 
    "sync"
)
type mrdis struct {
    mu sync.RWMutex
	Data map[string]interface{}
}
