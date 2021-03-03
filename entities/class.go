package entities

import "fmt"

type Class struct {
	Id    int64
	Name  string
	Roll  int64
	Age   int64
	Email string
}

func (class Class) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nroll: %d\nage: %demail: %s", class.Id, class.Name, class.Roll, class.Age, class.Email)
}
