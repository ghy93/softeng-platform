package repository

import (
	"context"
)

type ToolRepository interface {
	GetTools(ctx context.Context, category, tags []string, sort, cursor string, pageSize int) ([]map[string]interface{}, error)
	GetByID(ctx context.Context, resourceID string) (map[string]interface{}, error)
	Search(ctx context.Context, keyword, cursor string, pageSize int) ([]map[string]interface{}, error)
	Create(ctx context.Context, userID int, data map[string]interface{}) (map[string]interface{}, error)
	AddLike(ctx context.Context, userID int, resourceID string) error
	GetLikes(ctx context.Context, resourceID string) (int, error)
	GetPending(ctx context.Context, cursor, limit int) ([]map[string]interface{}, error) // 新增方法
}

type toolRepository struct {
	db *Database
}

func NewToolRepository(db *Database) ToolRepository {
	return &toolRepository{db: db}
}

func (r *toolRepository) GetTools(ctx context.Context, category, tags []string, sort, cursor string, pageSize int) ([]map[string]interface{}, error) {
	// 实现获取工具列表的逻辑
	return []map[string]interface{}{
		{
			"resourceId":   1,
			"resourceType": "tool",
			"resourceName": "Sample Tool",
			"description":  "This is a sample tool",
			"image":        "https://example.com/image.jpg",
			"catagory":     "软件开发",
			"tags":         []string{"免费", "AI工具"},
			"views":        100,
			"collections":  50,
			"loves":        25,
			"contributors": []string{"user1", "user2"},
			"createdat":    "2023-01-01",
		},
	}, nil
}

func (r *toolRepository) GetByID(ctx context.Context, resourceID string) (map[string]interface{}, error) {
	// 实现根据ID获取工具详情的逻辑
	return map[string]interface{}{
		"resourceId":         resourceID,
		"resourceType":       "tool",
		"resourceName":       "Sample Tool",
		"resourceLink":       "https://example.com",
		"description_detail": "Detailed description of the tool",
		"catagory":           "软件开发",
		"image":              []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"},
		"tags":               []string{"免费", "AI工具", "高效"},
		"views":              150,
		"collections":        75,
		"loves":              30,
		"iscollected":        false,
		"isliked":            false,
		"comment_count":      5,
		"comments":           []map[string]interface{}{},
		"createdDate":        "2023-01-01",
	}, nil
}

func (r *toolRepository) Search(ctx context.Context, keyword, cursor string, pageSize int) ([]map[string]interface{}, error) {
	// 实现搜索工具的逻辑
	return []map[string]interface{}{
		{
			"resourceId":   2,
			"resourceType": "tool",
			"resourceName": "Search Result Tool",
			"description":  "This tool matches the search: " + keyword,
			"image":        "https://example.com/search.jpg",
			"catagory":     "论文阅读",
			"tags":         []string{"搜索", "工具"},
			"views":        80,
			"collections":  40,
			"loves":        20,
			"contributors": []string{"user3"},
			"createdat":    "2023-02-01",
		},
	}, nil
}

func (r *toolRepository) Create(ctx context.Context, userID int, data map[string]interface{}) (map[string]interface{}, error) {
	// 实现创建工具的逻辑
	return map[string]interface{}{
		"resourceId":   3,
		"resourceType": "tool",
		"resource":     "https://example.com/new-tool",
		"auditStatus":  "pending",
		"submitTime":   "2023-12-01 10:00:00",
		"auditTime":    nil,
		"rejectReason": nil,
	}, nil
}

func (r *toolRepository) AddLike(ctx context.Context, userID int, resourceID string) error {
	// 实现添加点赞的逻辑
	return nil
}

func (r *toolRepository) GetLikes(ctx context.Context, resourceID string) (int, error) {
	// 实现获取点赞数的逻辑
	return 1, nil
}

func (r *toolRepository) GetPending(ctx context.Context, cursor, limit int) ([]map[string]interface{}, error) {
	// 实现获取待审核工具的逻辑
	return []map[string]interface{}{
		{
			"submitor":     "user1",
			"submitDate":   "2023-12-01 10:00:00",
			"reourceId":    1,
			"resourceType": "tool",
			"resourcename": "新工具",
			"catagory":     "软件开发",
			"link":         "https://example.com/tool",
			"description":  "这是一个新工具",
			"tags":         []string{"AI", "免费"},
			"file":         "tool.zip",
		},
	}, nil
}
