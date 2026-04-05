package mrdis 

func Connect() *mrdis { 

	return &mrdis{
		Data : make (map[string]interface{}),
	}
}