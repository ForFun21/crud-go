package service

import (
	"github.com/ForFun21/crud-go/src/configuration/rest_err"
	"github.com/ForFun21/crud-go/src/model"
)

func (*userDomainService) UpdateUser(
	userID string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
