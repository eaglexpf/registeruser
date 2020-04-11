package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"registeruser/app/admin/entity"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/model"
	"registeruser/app/admin/model/mysql"
	"registeruser/conf/log"
)

const (
	newAdminUserDefaultNickname  = "新用户"
	newAdminUserDefaultAvatarUrl = "https://xupengfei.net/image/logo.jpg"
	newAdminUserDefaultStatus    = 10
)

func NewService() *Service {
	return &Service{
		adminUserModel: mysql.NewAdminUserModel(),
	}
}

type Service struct {
	adminUserModel model.AdminUserModel
}

func (s *Service) FindUserByUUID(ctx context.Context, uuid string) (*entity.AdminUser, error) {
	return s.adminUserModel.FindUserByUUID(ctx, uuid)
}

func (s *Service) FindUserByUsername(ctx context.Context, username string) (*entity.AdminUser, error) {
	return s.adminUserModel.FindUserByUsername(ctx, username)
}

func (s *Service) FindUserByEmail(ctx context.Context, email string) (*entity.AdminUser, error) {
	return s.adminUserModel.FindUserByEmail(ctx, email)
}

func (s *Service) FindUserByMobile(ctx context.Context, mobile string) (*entity.AdminUser, error) {
	return s.adminUserModel.FindUserByMobile(ctx, mobile)
}

func (s *Service) RegisterUser(ctx context.Context, request *request.RequestRegisterAdminUser) (adminUser *entity.AdminUser, err error) {
	adminUser = new(entity.AdminUser)
	if request.Password != request.RepeatPwd {
		err = errors.New("两次密码不一致")
		return
	}
	//if strings.Compare(request.Password, request.RepeatPwd) == 0 {
	//	err = errors.New("两次密码不一致")
	//	return
	//}
	old_user, err := s.FindUserByUsername(ctx, request.Username)
	log.Error("debug:", old_user, err)
	if err == nil {
		err = errors.New("账号已存在")
		return
	}
	_, err = s.FindUserByEmail(ctx, request.Email)
	if err == nil {
		err = errors.New("邮箱已存在")
		return
	}
	// 获取唯一uuid
	for {
		adminUser.UUID = fmt.Sprintf("%s", uuid.NewV4())
		_, err := s.FindUserByUUID(ctx, adminUser.UUID)
		if err != nil {
			break
		}
	}
	// 生成密码
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	adminUser.UserName = request.Username
	adminUser.PasswordHash = string(hash)
	adminUser.Email = request.Email
	adminUser.Nickname = request.Nickname
	if adminUser.Nickname == "" {
		adminUser.Nickname = newAdminUserDefaultNickname
	}
	adminUser.AvatarUrl = request.AvatarUrl
	if adminUser.AvatarUrl == "" {
		adminUser.AvatarUrl = newAdminUserDefaultAvatarUrl
	}
	adminUser.Status = newAdminUserDefaultStatus
	err = s.adminUserModel.InsertUser(ctx, adminUser)
	if err != nil {
		return
	}
	return
}
