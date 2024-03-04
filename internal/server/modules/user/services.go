package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

type userService struct {
	query *sqlc.Queries
}

func UserService() *userService {
	return &userService{
		query: sqlc.New(database.PgPool),
	}
}

func (s *userService) ListUsers() ([]sqlc.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	u, err := s.query.ListUsers(ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return u, nil
}

func (s *userService) CreateUser(input CreateUserInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	hash, e := hash.Generate(input.Password)
	if e != nil {
		return e
	}

	arg := sqlc.CreateUserParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hash,
	}

	_, err := s.query.CreateUser(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return echo.NewHTTPError(http.StatusBadRequest, "User with this email already exists")
			}
			return err
		}
		return err
	}

	return nil
}
