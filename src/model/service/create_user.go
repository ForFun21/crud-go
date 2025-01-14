package service

import (
	"fmt"

	logger "github.com/ForFun21/crud-go/src/configuration/Logger"
	"github.com/ForFun21/crud-go/src/configuration/rest_err"
	"github.com/ForFun21/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreatUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init creatUser model.", zap.String("journey", "creatUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
