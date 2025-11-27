package repository

import (
	"context"
)

type ProjectRepository interface {
	GetProjects(ctx context.Context, category string, techStack []string, sort string, limit int, cursor string) ([]map[string]interface{}, error)
	GetByID(ctx context.Context, projectID string) (map[string]interface{}, error)
	Search(ctx context.Context, keyword string, category []string, cursor string, limit int) ([]map[string]interface{}, error)
	Create(ctx context.Context, userID int, data map[string]interface{}) (map[string]interface{}, error)
	Update(ctx context.Context, userID int, projectID string, data map[string]interface{}) (map[string]interface{}, error)
	LikeProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error)
	UnlikeProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error)
	AddComment(ctx context.Context, userID int, projectID, content string) (map[string]interface{}, error)
	DeleteComment(ctx context.Context, userID int, projectID string) (map[string]interface{}, error)
	ReplyComment(ctx context.Context, userID int, projectID, commentID, content string) (map[string]interface{}, error)
	DeleteReply(ctx context.Context, userID int, projectID, commentID string) (map[string]interface{}, error)
	AddView(ctx context.Context, projectID string) (int, error)
	CollectProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error)
	UncollectProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error)
	GetPending(ctx context.Context, cursor, limit int) ([]map[string]interface{}, error) // 新增方法
}

type projectRepository struct {
	db *Database
}

func NewProjectRepository(db *Database) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) GetProjects(ctx context.Context, category string, techStack []string, sort string, limit int, cursor string) ([]map[string]interface{}, error) {
	// 实现获取项目列表的逻辑
	return []map[string]interface{}{
		{
			"projectId":    1,
			"resourceType": "project",
			"name":         "校园社交平台",
			"description":  "基于Go和React的校园社交平台",
			"category":     "实训项目",
			"techStack":    []string{"Go", "React", "PostgreSQL"},
			"likecount":    45,
			"authername":   []string{"张三", "李四"},
			"cover":        "https://example.com/project1.jpg",
			"createdat":    "2023-11-01",
			"loves":        45,
			"collections":  30,
			"views":        200,
		},
	}, nil
}

func (r *projectRepository) GetByID(ctx context.Context, projectID string) (map[string]interface{}, error) {
	// 实现根据ID获取项目详情的逻辑
	return map[string]interface{}{
		"projectId":     projectID,
		"resourceType":  "project",
		"name":          "校园社交平台",
		"description":   "基于Go和React的校园社交平台",
		"githubURL":     "https://github.com/example/social-platform",
		"techStack":     []string{"Go", "React", "PostgreSQL"},
		"catagory":      "实训项目",
		"images":        []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"},
		"likes":         45,
		"views":         200,
		"collections":   30,
		"isliked":       false,
		"iscollected":   false,
		"author":        []string{"张三", "李四"},
		"comment_count": 10,
		"comments":      []map[string]interface{}{},
		"createdAt":     "2023-11-01",
	}, nil
}

func (r *projectRepository) Search(ctx context.Context, keyword string, category []string, cursor string, limit int) ([]map[string]interface{}, error) {
	// 实现搜索项目的逻辑
	return []map[string]interface{}{
		{
			"projectId":    2,
			"resourceType": "project",
			"name":         "电商平台 - " + keyword,
			"description":  "完整的电商平台解决方案",
			"category":     "课程设计",
			"techStack":    []string{"Java", "Spring Boot", "Vue"},
			"likecount":    35,
			"authername":   []string{"王五"},
			"cover":        "https://example.com/project2.jpg",
			"createdat":    "2023-10-15",
			"loves":        35,
			"collections":  25,
			"views":        150,
		},
	}, nil
}

func (r *projectRepository) Create(ctx context.Context, userID int, data map[string]interface{}) (map[string]interface{}, error) {
	// 实现创建项目的逻辑
	return map[string]interface{}{
		"resourceId":   1,
		"resourceType": "resource",
		"resource":     "https://example.com/project",
		"auditStatus":  "pending",
		"submitTime":   "2023-12-01 10:00:00",
		"auditTime":    nil,
		"rejectReason": nil,
	}, nil
}

func (r *projectRepository) Update(ctx context.Context, userID int, projectID string, data map[string]interface{}) (map[string]interface{}, error) {
	// 实现更新项目的逻辑
	return map[string]interface{}{
		"resourceId":   1,
		"resourceType": "resource",
		"resource":     "https://example.com/updated-project",
		"auditStatus":  "pending",
		"submitTime":   "2023-12-01 10:00:00",
		"auditTime":    nil,
		"rejectReason": nil,
	}, nil
}

func (r *projectRepository) LikeProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"likecounts": 46,
		"isliked":    true,
	}, nil
}

func (r *projectRepository) UnlikeProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"likecounts": 45,
		"isliked":    false,
	}, nil
}

func (r *projectRepository) AddComment(ctx context.Context, userID int, projectID, content string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"comment_Id":  1,
		"nickname":    "用户",
		"avater":      "https://example.com/avatar.jpg",
		"comment":     content,
		"commentDate": "2023-12-01 10:00:00",
		"love_count":  0,
		"reply_total": 0,
		"replies":     []map[string]interface{}{},
	}, nil
}

func (r *projectRepository) DeleteComment(ctx context.Context, userID int, projectID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"comment_Id":  1,
		"nickname":    "用户",
		"avater":      "https://example.com/avatar.jpg",
		"comment":     "已删除的评论",
		"delete_Date": "2023-12-01 10:00:00",
	}, nil
}

func (r *projectRepository) ReplyComment(ctx context.Context, userID int, projectID, commentID, content string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"comment_Id":  2,
		"nickname":    "用户",
		"avater":      "https://example.com/avatar.jpg",
		"comment":     content,
		"commentDate": "2023-12-01 10:00:00",
		"love_count":  0,
		"reply_total": 0,
		"replies":     []map[string]interface{}{},
	}, nil
}

func (r *projectRepository) DeleteReply(ctx context.Context, userID int, projectID, commentID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"comment_Id":  2,
		"nickname":    "用户",
		"avater":      "https://example.com/avatar.jpg",
		"comment":     "已删除的回复",
		"delete_Date": "2023-12-01 10:00:00",
	}, nil
}

func (r *projectRepository) AddView(ctx context.Context, projectID string) (int, error) {
	return 201, nil
}

func (r *projectRepository) CollectProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"iscollected": true,
		"collections": 31,
	}, nil
}

func (r *projectRepository) UncollectProject(ctx context.Context, userID int, projectID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"iscollected": false,
		"collections": 30,
	}, nil
}

func (r *projectRepository) GetPending(ctx context.Context, cursor, limit int) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"submitor":     "user2",
			"submitDate":   "2023-12-01 10:00:00",
			"reourceId":    1,
			"resourceType": "project",
			"resourcename": "新项目",
			"catagory":     "实训项目",
			"link":         "https://github.com/example/project",
			"description":  "项目描述",
			"tags":         []string{"Go", "React"},
			"file":         "project.zip",
		},
	}, nil
}
