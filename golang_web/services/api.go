package services

import (
	"github.com/gin-gonic/gin"
	"kubeclub/scaffold_go_web/dao"
	"kubeclub/scaffold_go_web/dto"
	"kubeclub/scaffold_go_web/middleware"
)

type ApiService struct {
}


func (as *ApiService) AddUser(c *gin.Context, input *dto.AddUserInput) error {
	user := &dao.User{
		Name:  input.Name,
		Sex:   input.Sex,
		Age:   input.Age,
		Birth: input.Birth,
		Addr:  input.Addr,
	}
	if err := user.Save(c); err != nil {
		middleware.ResponseError(c, 2002, err)
		return err
	}
	return nil
}
