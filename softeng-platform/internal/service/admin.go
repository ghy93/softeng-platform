package service

import (
	"context"
	"softeng-platform/internal/repository"
)

type AdminService interface {
	GetPending(ctx context.Context, itemType string, cursor, limit int, sort string) (map[string]interface{}, error)
	ReviewItem(ctx context.Context, itemID, action, rejectReason string) error
}

type adminService struct {
	toolRepo    repository.ToolRepository
	courseRepo  repository.CourseRepository
	projectRepo repository.ProjectRepository
}

func NewAdminService(toolRepo repository.ToolRepository, courseRepo repository.CourseRepository, projectRepo repository.ProjectRepository) AdminService {
	return &adminService{
		toolRepo:    toolRepo,
		courseRepo:  courseRepo,
		projectRepo: projectRepo,
	}
}

func (s *adminService) GetPending(ctx context.Context, itemType string, cursor, limit int, sort string) (map[string]interface{}, error) {
	var data []map[string]interface{}
	var err error

	switch itemType {
	case "工具":
		data, err = s.toolRepo.GetPending(ctx, cursor, limit)
	case "课程":
		data, err = s.courseRepo.GetPending(ctx, cursor, limit)
	case "项目":
		data, err = s.projectRepo.GetPending(ctx, cursor, limit)
	case "评论":
		// 获取待审核评论
		data = []map[string]interface{}{}
	default:
		data = []map[string]interface{}{}
	}

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total":  len(data),
		"cursor": cursor,
		"data":   data,
	}, nil
}

func (s *adminService) ReviewItem(ctx context.Context, itemID, action, rejectReason string) error {
	// 根据itemID的类型执行相应的审核操作
	// 这里需要实现具体的审核逻辑
	return nil
}
