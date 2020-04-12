package service

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"registeruser/app/admin/entity"
	"registeruser/app/admin/entity/request"
	"registeruser/conf/global"
	"registeruser/util"
)

const (
	newAdminUserDefaultNickname  = "新用户"
	newAdminUserDefaultAvatarUrl = "https://xupengfei.net/image/logo.jpg"
	newAdminUserDefaultStatus    = 10
)

func (s *Service) AdminUserRegister(ctx context.Context, request *request.RequestRegisterAdminUser) (adminUser *entity.AdminUser, err error) {
	adminUser = new(entity.AdminUser)
	if request.Password != request.RepeatPwd {
		err = errors.New("两次密码不一致")
		return
	}
	_, err = s.AdminUserFindByUsername(ctx, request.Username)
	if err == nil {
		err = errors.New("账号已存在")
		return
	}
	_, err = s.AdminUserFindByEmail(ctx, request.Email)
	if err == nil {
		err = errors.New("邮箱已存在")
		return
	}
	// 获取唯一uuid
	for {
		adminUser.UUID = fmt.Sprintf("%s", uuid.NewV4())
		_, err := s.AdminUserFindByUUID(ctx, adminUser.UUID)
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

func (s *Service) AdminUserLogin(ctx context.Context, request *request.RequestAdminUserLogin) (token string, err error) {
	adminUser, err := s.AdminUserFindByUsername(ctx, request.Username)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminUser.PasswordHash), []byte(request.Password))
	if err != nil {
		return
	}
	token, err = util.NewJWT().CreateToken(&global.JwtClaims{
		UUID: adminUser.UUID,
	})
	if err != nil {
		return
	}
	return
}

func (s *Service) AdminUserRefreshToken(ctx context.Context, adminUser *entity.AdminUser) (token string, err error) {
	token, err = util.NewJWT().CreateToken(&global.JwtClaims{
		UUID: adminUser.UUID,
	})
	return
}

func (s *Service) AdminUserUpdateInfo(ctx context.Context, updateData *request.RequestAdminUserUpdateInfo) (adminUser *entity.AdminUser, err error) {
	adminUser = &entity.AdminUser{
		UUID:      updateData.UUID,
		Nickname:  updateData.Nickname,
		AvatarUrl: updateData.AvatarUrl,
	}
	err = s.adminUserModel.UpdateUserInfoByUUID(ctx, adminUser)
	return
}

func (s *Service) AdminUserResetPwd(ctx context.Context, updateData *request.RequestAdminUserResetPwd) (adminUser *entity.AdminUser, err error) {
	if updateData.Password != updateData.OldPwd {
		err = errors.New("两次密码不一致")
		return
	}
	// 生成密码
	hash, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	adminUser = &entity.AdminUser{
		UUID:         updateData.UUID,
		PasswordHash: string(hash),
	}
	err = s.adminUserModel.UpdateUserPwdByUUID(ctx, adminUser)
	return
}

func (s *Service) AdminUserFindByUUID(ctx context.Context, uuid string) (*entity.AdminUser, error) {
	return s.adminUserModel.FindUserByUUID(ctx, uuid)
}

func (s *Service) AdminUserFindByUsername(ctx context.Context, username string) (*entity.AdminUser, error) {
	return s.adminUserModel.FindUserByUsername(ctx, username)
}

func (s *Service) AdminUserFindByEmail(ctx context.Context, email string) (*entity.AdminUser, error) {
	return s.adminUserModel.FindUserByEmail(ctx, email)
}
