package usecase

import (
	"go-mission/pkg/model"
	"go-mission/pkg/sql"

	"github.com/google/uuid"
)

func CreateUser(u string) {
	data := new(model.User)
	data.Name = u
	uuidV4 := uuid.New()
	data.Token = uuidV4.String()
	sql.CreateUser(data)
}
