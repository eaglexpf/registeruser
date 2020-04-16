package service

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"registeruser/app/admin/entity/dao"
	"registeruser/app/admin/entity/request"
	"registeruser/app/admin/entity/response"
	"registeruser/conf/global"
	"registeruser/conf/log"
	"registeruser/util/jwt"
)

const (
	newAdminUserDefaultNickname  = "新用户"
	newAdminUserDefaultAvatarUrl = "https://xupengfei.net/image/logo.jpg"
	newAdminUserDefaultStatus    = 10
)

// 新建一个后台用户服务
func (s *Service) RegisterAdminUser(ctx context.Context, request *request.RequestRegisterAdminUser) (adminUser *dao.AdminUser, err error) {
	err = s.AuthCheck(ctx, "admin.user.register")
	if err != nil {
		return
	}
	adminUser = new(dao.AdminUser)
	if request.Password != request.RepeatPwd {
		// 重复密码不一致
		err = response.ERROR_USER_PASSWORD_EQUAL
		return
	}
	_, err = s.FindAdminUserByUUID(ctx, request.Username)
	if err == nil {
		// 账号已存在
		err = response.ERROR_USER_UNIQUE_NAME
		return
	}
	_, err = s.FindAdminUserByEmail(ctx, request.Email)
	if err == nil {
		// 邮箱已存在
		err = response.ERROR_USER_UNIQUE_EMAIL
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
		// 密码加密失败
		err = response.ERROR_USER_PASSWORD_CRYPT
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
		// 用户注册失败
		log.Errorf("用户注册失败，失败原因：%s", err.Error())
		err = response.ERROR_USER_REGISTER_DB
	}
	return
}

// 后台用户登录服务
func (s *Service) LoginForAdminUser(ctx context.Context, request *request.RequestAdminUserLogin) (token string, err error) {
	err = s.AuthCheck(ctx, "admin.user.login")
	if err != nil {
		return
	}
	adminUser, err := s.FindAdminUserByUsername(ctx, request.Username)
	if err != nil {
		// 无效的账号
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminUser.PasswordHash), []byte(request.Password))
	if err != nil {
		// 密码错误
		err = response.ERROR_USER_PASSWORD_FALSE
		return
	}
	token, err = jwt.NewJWT().CreateToken(&global.JwtClaims{
		UUID: adminUser.UUID,
	})
	if err != nil {
		// jwt生成失败
		err = response.ERROR_JWT_CREATE
	}
	return
}

func (s *Service) RefreshTokenByAdminUser(ctx context.Context, adminUser *dao.AdminUser) (token string, err error) {
	err = s.AuthCheck(ctx, "admin.user.refresh_token")
	if err != nil {
		return
	}
	token, err = jwt.NewJWT().CreateToken(&global.JwtClaims{
		UUID: adminUser.UUID,
	})
	if err != nil {
		// jwt生成失败
		err = response.ERROR_JWT_CREATE
	}
	return
}

// 修改后台用户的昵称和头像
func (s *Service) UpdateAdminUserInfo(ctx context.Context, updateData *request.RequestAdminUserUpdateInfo) (adminUser *dao.AdminUser, err error) {
	err = s.AuthCheck(ctx, "admin.user.update_info")
	if err != nil {
		return
	}
	adminUser, err = s.FindAdminUserByUUID(ctx, updateData.UUID)
	if err != nil {
		// 无效的UUID
		return
	}
	adminUser = &dao.AdminUser{
		UUID:      updateData.UUID,
		Nickname:  updateData.Nickname,
		AvatarUrl: updateData.AvatarUrl,
	}
	err = s.adminUserModel.UpdateUserInfoByUUID(ctx, adminUser)
	if err != nil {
		// 用户修改失败
		log.Errorf("修改用户基本信息失败，失败原因：%s", err.Error())
		err = response.ERROR_USER_UPDATE_INFO_DB
	}
	return
}

// 修改后台用户的密码
func (s *Service) ResetPwdForAdminUser(ctx context.Context, updateData *request.RequestAdminUserResetPwd) (adminUser *dao.AdminUser, err error) {
	err = s.AuthCheck(ctx, "admin.user.update_pwd")
	if err != nil {
		return
	}
	if updateData.Password != updateData.OldPwd {
		// 两次新密码不一致
		err = response.ERROR_USER_PASSWORD_EQUAL
		return
	}
	// 生成密码
	hash, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
	if err != nil {
		// 密码生成失败
		err = response.ERROR_USER_PASSWORD_CRYPT
		return
	}
	// 查询用户是否存在
	adminUser, err = s.FindAdminUserByUUID(ctx, updateData.UUID)
	if err != nil {
		// 无效的UUID
		return
	}
	adminUser = &dao.AdminUser{
		UUID:         updateData.UUID,
		PasswordHash: string(hash),
	}
	err = s.adminUserModel.UpdateUserPwdByUUID(ctx, adminUser)
	if err != nil {
		// 用户修改失败
		log.Errorf("修改用户密码失败，失败原因：%s", err.Error())
		err = response.ERROR_USER_UPDATE_PASSWORD_DB
	}
	return
}

// 根据uuid查询后台用户
func (s *Service) FindAdminUserByUUID(ctx context.Context, uuid string) (*dao.AdminUser, error) {
	data, err := s.adminUserModel.FindUserByUUID(ctx, uuid)
	if err != nil {
		err = response.ERROR_USER_FIND_BY_UUID
	}
	return data, err
}

// 根据username查询后台用户
func (s *Service) FindAdminUserByUsername(ctx context.Context, username string) (*dao.AdminUser, error) {
	data, err := s.adminUserModel.FindUserByUsername(ctx, username)
	if err != nil {
		err = response.ERROR_USER_FIND_BY_NAME
	}
	return data, err
}

// 根据Email查询后台用户
func (s *Service) FindAdminUserByEmail(ctx context.Context, email string) (*dao.AdminUser, error) {
	data, err := s.adminUserModel.FindUserByEmail(ctx, email)
	if err != nil {
		err = response.ERROR_USER_FIND_BY_EMAIL
	}
	return data, err
}
