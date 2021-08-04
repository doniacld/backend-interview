package user

import "context"

type User struct {
	ID   string
	Name string
}

type Filter struct {
	ID string
}

type Store interface {
	Fetch(context.Context, Filter) (User, error)
	FetchMany()
}

type App interface {
	Store
}
