package dao

import (
	"fmt"

	"internet-cafe/pkg/model"
)

type InternetCafeDao struct {
}

var (
	saved *model.InternetCafe
)

func NewInternetCafeDao() *InternetCafeDao {
	return &InternetCafeDao{}
}

func (i *InternetCafeDao) Save(inetCafe *model.InternetCafe) error {
	saved = inetCafe
	return nil
}

func (i *InternetCafeDao) Get() (*model.InternetCafe, error) {
	if saved == nil {
		return nil, fmt.Errorf("internet cafe not found")
	}
	return saved, nil
}
