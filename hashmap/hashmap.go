package hashmap

import (
	"fmt"
	"training-go/structure"
)

func GetMap() {
	v1 := structure.Vertex{
		X: 20,
		Y: 30,
	}

	v2 := structure.Vertex{
		X: 50,
		Y: 60,
	}

	m := map[string]structure.Vertex{
		"01": v1,
		"02": v2,
	}

	delete(m, "01")

	if elem, ok := m["02"]; ok {
		fmt.Println("map:", elem)
	}
}
