package service

import (
	"context"
	"softeng-platform/internal/model"
	"softeng-platform/internal/repository"
)

type UserService interface {
	GetProfile(ctx context.Context, userID int) (*model.User, error)
	UpdateProfile(ctx context.Context, userID int, req model.UpdateProfileRequest) (*model.User, error)
	GetCollection(ctx context.Context, userID int) (map[string]interface{}, error)
	DeleteCollection(ctx context.Context, userID int, resourceType string, resourceID int) (map[string]interface{}, error)
	GetStatus(ctx context.Context, userID int) (map[string]interface{}, error)
	GetSummit(ctx context.Context, userID int) (map[string]interface{}, error)
	UpdateResourceStatus(ctx context.Context, userID int, resourceType, resourceID, action, state string) (map[string]interface{}, error)
	UpdateEmail(ctx context.Context, userID int, name, password, newEmail, code string) (*model.User, error)
	UpdatePassword(ctx context.Context, userID int, name, email, newPassword, code string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetProfile(ctx context.Context, userID int) (*model.User, error) {
	return s.userRepo.GetByID(ctx, userID)
}

func (s *userService) UpdateProfile(ctx context.Context, userID int, req model.UpdateProfileRequest) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 更新字段
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Description != "" {
		user.Description = req.Description
	}
	if req.FacePhoto != "" {
		user.FacePhoto = req.FacePhoto
	}

	err = s.userRepo.Update(ctx, user)
	return user, err
}

func (s *userService) GetCollection(ctx context.Context, userID int) (map[string]interface{}, error) {
	// 实现获取收藏逻辑
	return map[string]interface{}{
		"message":   "success",
		"resources": []interface{}{},
		"tools":     []interface{}{},
		"teaches":   []interface{}{},
	}, nil
}

func (s *userService) DeleteCollection(ctx context.Context, userID int, resourceType string, resourceID int) (map[string]interface{}, error) {
	// 实现删除收藏逻辑
	return map[string]interface{}{
		"message":   "success",
		"resources": []interface{}{},
		"tools":     []interface{}{},
		"teaches":   []interface{}{},
	}, nil
}

func (s *userService) GetStatus(ctx context.Context, userID int) (map[string]interface{}, error) {
	// 实现获取审核状态逻辑
	return map[string]interface{}{
		"message":   "success",
		"resources": []interface{}{},
		"tools":     []interface{}{},
		"teaches":   []interface{}{},
	}, nil
}

func (s *userService) GetSummit(ctx context.Context, userID int) (map[string]interface{}, error) {
	// 实现获取个人提交逻辑
	return map[string]interface{}{
		"message":   "success",
		"resources": []interface{}{},
		"tools":     []interface{}{},
		"teaches":   []interface{}{},
	}, nil
}

func (s *userService) UpdateResourceStatus(ctx context.Context, userID int, resourceType, resourceID, action, state string) (map[string]interface{}, error) {
	// 实现更新资源状态逻辑
	return map[string]interface{}{
		"message": "success",
		"manipulate": map[string]interface{}{
			"resourceId":   resourceID,
			"resourceType": resourceType,
			"newstatus":    action,
			"oldestatus":   "published", // 假设之前的状态
			"operateTime":  "2023-12-01 10:00:00",
			"operator":     "user",
		},
	}, nil
}

func (s *userService) UpdateEmail(ctx context.Context, userID int, name, password, newEmail, code string) (*model.User, error) {
	// 实现更新邮箱逻辑
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdatePassword(ctx context.Context, userID int, name, email, newPassword, code string) (*model.User, error) {
	// 实现更新密码逻辑
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
