package container

import (
	"bettersocial/config"
	"bettersocial/repository"

	"gorm.io/gorm"
)

type Container struct {
	db     *gorm.DB
	config config.Config

	// repo
	userRepo repository.User
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Db() *gorm.DB {
	return c.db
}

func (c *Container) SetDb(db *gorm.DB) {
	c.db = db
}

func (c *Container) Config() config.Config {
	return c.config
}

func (c *Container) SetConfig(config config.Config) {
	c.config = config
}

func (c *Container) UserRepo() repository.User {
	return c.userRepo
}

func (c *Container) SetTodoRepo(userRepo repository.User) {
	c.userRepo = userRepo
}
