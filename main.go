package main 

import (
	"fmt"
	"github.com/ParthDande/mrdis/mrdis"
)


func main () { 
	db := mrdis.Connect() 
	// 2. Set some data (Key is string, Value is interface/any)
    db.Set("user_1", "Alice")
    db.Set("login_count", 42)

    // 3. Get the data back
    name := db.Get("user_1")
    count := db.Get("login_count")

    // 4. Print it
    fmt.Println("User Name:", name)
    fmt.Println("Logins:", count)
    
    // 5. Delete a key
    db.Delete("user_1")
}