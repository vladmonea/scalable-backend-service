package users

var users = []User{
	{
		Name: "John",
		Age:  30,
	},
	{
		Name: "Jane",
		Age:  24,
	},
	{
		Name: "Harry",
		Age:  12,
	},
}

func GetUsers() []User {
	return users
}

func AddUser(name string, age int) {
	users = append(users, User{Name: name, Age: age})
}
