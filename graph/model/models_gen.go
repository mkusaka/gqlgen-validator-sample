// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text           string          `json:"text" validate:"max=1"`
	UserID         string          `json:"userId" validate:"max=1"`
	RepeatEveryDay *RepeatEveryDay `json:"repeatEveryDay"`
}

type RepeatEveryDay struct {
	Days     int    `json:"days" validate:"max=1"`
	Time     string `json:"time" validate:"max=1"`
	Timezone string `json:"timezone" validate:"max=1"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
