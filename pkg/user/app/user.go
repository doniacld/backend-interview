package app

import (
	"github.com/gustvision/backend-interview/pkg/user"
)

var _ user.App = (*App)(nil)

type App struct {
	user.Store
}
