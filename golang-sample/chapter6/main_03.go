
package main

import "fmt"

func main(){


	elements := make(map[string]map[string]string)
	elements["H"] = map[string]string{ "OK": "Hydrogen", "haha" : "sdjasf"}
 //	elements["He"] = "Helium"

	name := elements["H"]

	//value1, value2 := name

	fmt.Println(name["OK"], name["haha"])
}
