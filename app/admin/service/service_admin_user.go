package service

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/conf/global"
	"registeruser/util/jwt"
)

const (
	newAdminUserDefaultNickname  = "新用户"
	newAdminUserDefaultAvatarUrl = "https://xupengfei.net/image/logo.jpg"
	newAdminUserDefaultStatus    = 10
)

// 新建一个后台用户服务
func (s *Service) RegisterAdminUser(ctx context.Context, request *request.RequestRegisterAdminUser) (adminUser *dao.AdminUser, err error) {
	adminUser = new(dao.AdminUser)
	if request.Password != request.RepeatPwd {
		err = errors.New("两次密码不一致")
		return
	}
	_, err = s.FindAdminUserByUUID(ctx, request.Username)
	if err == nil {
		err = errors.New("账号已存在")
		return
	}
	_, err = s.FindAdminUserByEmail(ctx, request.Email)
	if err == nil {
		err = errors.New("邮箱已存在")
		return
	}
	// 获取唯一uuid
	for {
		adminUser.UUID = fmt.Sprintf("%s", uuid.NewV4())
		_, err := s.FindAdminUserByUUID(ctx, adminUser.UUID)
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

// 后台用户登录服务
func (s *Service) LoginForAdminUser(ctx context.Context, request *request.RequestAdminUserLogin) (token string, err error) {
	adminUser, err := s.FindAdminUserByUsername(ctx, request.Username)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminUser.PasswordHash), []byte(request.Password))
	if err != nil {
		err = errors.New("密码错误")
		return
	}
	token, err = jwt.NewJWT().CreateToken(&global.JwtClaims{
		UUID: adminUser.UUID,
	})
	if err != nil {
		return
	}
	return
}

func (s *Service) RefreshTokenByAdminUser(ctx context.Context, adminUser *dao.AdminUser) (token string, err error) {
	token, err = jwt.NewJWT().CreateToken(&global.JwtClaims{
		UUID: adminUser.UUID,
	})
	return
}

// 修改后台用户的昵称和头像
func (s *Service) UpdateAdminUserInfo(ctx context.Context, updateData *request.RequestAdminUserUpdateInfo) (adminUser *dao.AdminUser, err error) {
	adminUser = &dao.AdminUser{
		UUID:      updateData.UUID,
		Nickname:  updateData.Nickname,
		AvatarUrl: updateData.AvatarUrl,
	}
	err = s.adminUserModel.UpdateUserInfoByUUID(ctx, adminUser)
	return
}

// 修改后台用户的密码
func (s *Service) ResetPwdForAdminUser(ctx context.Context, updateData *request.RequestAdminUserResetPwd) (adminUser *dao.AdminUser, err error) {
	if updateData.Password != updateData.OldPwd {
		err = errors.New("两次密码不一致")
		return
	}
	// 生成密码
	hash, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	adminUser = &dao.AdminUser{
		UUID:         updateData.UUID,
		PasswordHash: string(hash),
	}
	err = s.adminUserModel.UpdateUserPwdByUUID(ctx, adminUser)
	return
}

// 根据uuid查询后台用户
func (s *Service) FindAdminUserByUUID(ctx context.Context, uuid string) (*dao.AdminUser, error) {
	return s.adminUserModel.FindUserByUUID(ctx, uuid)
}

// 根据username查询后台用户
func (s *Service) FindAdminUserByUsername(ctx context.Context, username string) (*dao.AdminUser, error) {
	return s.adminUserModel.FindUserByUsername(ctx, username)
}

// 根据Email查询后台用户
func (s *Service) FindAdminUserByEmail(ctx context.Context, email string) (*dao.AdminUser, error) {
	return s.adminUserModel.FindUserByEmail(ctx, email)
}