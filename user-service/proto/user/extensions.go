package shipper_user

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("UUID creation something went wrong: %s", err)
		return err
	}
	return scope.SetColumn("Id", uuid.String())
}
