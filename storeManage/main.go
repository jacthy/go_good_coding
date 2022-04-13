package main

type People struct {
	Name string
	Value string
}

func getMap() [2]string {
	pp := []*People{
		{Name: "A",Value: "a"},
		{Name: "B",Value: "B"},
	}
	m := [2]string{pp[0].Value,pp[1].Value}
	for i, people := range pp {
		m[i]=people.Value
	}
	return m
}

func main() {
	peopleMap := getMap()
	for key, val := range peopleMap {
		println(key,val)
	}
}

func test(string2 string)  {

}