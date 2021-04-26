package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {

	m := map[string][]Vertex{
		"Bell Labs": []Vertex{
			{40, -74},
			{2, 3},
		},
		"Google": []Vertex{
			{37, -122},
			{1, 2},
		},
	}

	m["Bell"] = append(m["Bell"], Vertex{1, 2})
	fmt.Println(m["Bell"])
}
