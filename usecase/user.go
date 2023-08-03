package usecase

import (
	"bettersocial/container"
	"bettersocial/helper"
	"bettersocial/model"

	"bettersocial/repository"
	"context"
)

type User struct {
	repository.User
}

func NewUser(c *container.Container) *User {
	return &User{
		User: c.UserRepo(),
	}
}

// func (t *Todo) Add(ctx context.Context, req model.Todo) (*model.Todo, error) {
// 	logger := helper.GetLogger(ctx).WithField("method", "usecase.Add")

// 	if err := req.ValidateAdd(); err != nil {
// 		logger.WithError(err).Warning("Not Valid Request")
// 		return nil, model.NewParameterError(helper.Pointer(err.Error()))
// 	}

// 	res, err := t.Todo.Add(ctx, req)
// 	if err != nil {
// 		logger.WithError(err).Warning("Failed insert todo")
// 		return nil, err
// 	}

// 	return res, nil
// }

func (u *User) CheckUsername(ctx context.Context, username string) error {
	logger := helper.GetLogger(ctx).WithField("method", "usecase.CheckUsername")

	_, err := u.User.GetByUsername(ctx, username)
	if err != nil {
		// return nil if username not exist
		if model.IsNotFoundError(err) {
			return nil
		}
		logger.WithError(err).Warning("Failed getById todo")
		return err
	}

	// return error if username exist
	return model.NewDuplicateError()
}

// func (t *Todo) Fetch(ctx context.Context, req request.FetchTodoRequest) (*model.Pagination[model.Todo], error) {
// 	logger := helper.GetLogger(ctx).WithField("method", "usecase.Fetch")

// 	res, err := t.Todo.Fetch(ctx, &repository.FetchFilter{
// 		SortOrder: "ASC",
// 		SortField: "created_at",
// 		PerPage:   req.PerPage,
// 		Page:      req.Page,
// 	})
// 	if err != nil {
// 		logger.WithError(err).Warning("Failed fetch todo")
// 		return nil, err
// 	}

// 	return res, nil
// }

// func (t *Todo) Update(ctx context.Context, todoId string, req *model.Todo) (*model.Todo, error) {
// 	logger := helper.GetLogger(ctx).WithField("method", "usecase.Update")

// 	if todoId == "" {
// 		logger.Error("missing todoId")
// 		return nil, model.NewParameterError(helper.Pointer("missing todoId"))
// 	}

// 	res, err := t.Todo.Update(ctx, todoId, req)
// 	if err != nil {
// 		logger.WithError(err).Warning("Failed update todo")
// 		return nil, err
// 	}

// 	return res, nil
// }

// func (t *Todo) Delete(ctx context.Context, todoId string) error {
// 	logger := helper.GetLogger(ctx).WithField("method", "usecase.Delete")

// 	if todoId == "" {
// 		logger.Error("missing todoId")
// 		return model.NewParameterError(helper.Pointer("missing todoId"))
// 	}

// 	err := t.Todo.Delete(ctx, todoId)
// 	if err != nil {
// 		logger.WithError(err).Warning("Failed delete todo")
// 		return err
// 	}

// 	return nil
// }
