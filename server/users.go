package server

type User struct {
	name string
	age  int
}

var users = []User{
	{
		name: "John",
		age:  30,
	},
	{
		name: "Jane",
		age:  24,
	},
	{
		name: "Harry",
		age:  12,
	},
}

func getUsers() []User {
	return users
}
