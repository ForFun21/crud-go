package service

import (
	"github.com/ForFun21/crud-go/src/configuration/rest_err"
	"github.com/ForFun21/crud-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
