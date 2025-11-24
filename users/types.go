package users

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
