package repository

import (
	"context"
)

type CourseRepository interface {
	GetCourses(ctx context.Context, semester string, category []string, sort string, limit, cursor int) ([]map[string]interface{}, error)
	GetByID(ctx context.Context, courseID string) (map[string]interface{}, error)
	Search(ctx context.Context, keyword string, category []string, limit, cursor int) ([]map[string]interface{}, error)
	UploadResource(ctx context.Context, userID int, courseID string, data map[string]interface{}) (map[string]interface{}, error)
	DownloadTextbook(ctx context.Context, courseID, textbookID string) (string, error)
	AddComment(ctx context.Context, userID int, courseID, content string) (map[string]interface{}, error)
	DeleteComment(ctx context.Context, userID int, courseID string) (map[string]interface{}, error)
	ReplyComment(ctx context.Context, userID int, courseID, commentID, content string) (map[string]interface{}, error)
	DeleteReply(ctx context.Context, userID int, courseID, commentID string) (map[string]interface{}, error)
	AddView(ctx context.Context, courseID string) (int, error)
	CollectCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error)
	UncollectCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error)
	LikeCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error)
	UnlikeCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error)
	GetPending(ctx context.Context, cursor, limit int) ([]map[string]interface{}, error) // 新增方法
}

type courseRepository struct {
	db *Database
}

func NewCourseRepository(db *Database) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) GetCourses(ctx context.Context, semester string, category []string, sort string, limit, cursor int) ([]map[string]interface{}, error) {
	// 实现获取课程列表的逻辑
	return []map[string]interface{}{
		{
			"courseId":     1,
			"resourceType": "course",
			"name":         "软件工程导论",
			"teacher":      []string{"张教授", "李教授"},
			"category":     []string{"专必", "有签到"},
			"semester":     "大二上",
			"credit":       3,
			"cover":        "https://example.com/course1.jpg",
			"views":        1000,
			"loves":        200,
			"collections":  150,
		},
	}, nil
}

func (r *courseRepository) GetByID(ctx context.Context, courseID string) (map[string]interface{}, error) {
	// 实现根据ID获取课程详情的逻辑
	return map[string]interface{}{
		"courseId":     courseID,
		"resourceType": "course",
		"name":         "软件工程导论",
		"catagory":     "专业必修",
		"url_form": []map[string]interface{}{
			{
				"resource_intro": "课程视频",
				"resource_url":   "https://example.com/video1",
				"resource_id":    1,
			},
		},
		"upload_form": []map[string]interface{}{
			{
				"resource_intro":  "课程讲义",
				"resource_upload": "https://example.com/lecture1.pdf",
				"resource_id":     2,
			},
		},
		"contributor":   []string{"user1", "user2"},
		"collections":   150,
		"views":         1000,
		"likes":         200,
		"isliked":       false,
		"iscollected":   false,
		"comment_total": 25,
		"comments":      []map[string]interface{}{},
		"createdAt":     "2023-09-01",
	}, nil
}

func (r *courseRepository) Search(ctx context.Context, keyword string, category []string, limit, cursor int) ([]map[string]interface{}, error) {
	// 实现搜索课程的逻辑
	return []map[string]interface{}{
		{
			"courseId":     2,
			"resourceType": "course",
			"name":         "高级软件工程 - " + keyword,
			"teacher":      []string{"王教授"},
			"category":     []string{"专选", "无签到"},
			"semester":     "大三上",
			"credit":       2,
			"cover":        "https://example.com/course2.jpg",
			"views":        800,
			"loves":        150,
			"collections":  100,
		},
	}, nil
}

func (r *courseRepository) UploadResource(ctx context.Context, userID int, courseID string, data map[string]interface{}) (map[string]interface{}, error) {
	// 实现上传课程资源的逻辑
	return map[string]interface{}{
		"resourceId":   1,
		"resourceType": "teach",
		"resource1": map[string]interface{}{
			"resource_intro": data["description"],
			"resource_url":   data["resource"],
			"resource_id":    1,
		},
		"resource2": map[string]interface{}{
			"resource_intro":  data["description"],
			"resource_upload": data["file"],
			"resource_id":     2,
		},
		"auditStatus":  "pending",
		"submitTime":   "2023-12-01 10:00:00",
		"auditTime":    nil,
		"rejectReason": nil,
	}, nil
}

func (r *courseRepository) DownloadTextbook(ctx context.Context, courseID, textbookID string) (string, error) {
	// 实现下载课本的逻辑
	return "Textbook content for course " + courseID + " textbook " + textbookID, nil
}

func (r *courseRepository) AddComment(ctx context.Context, userID int, courseID, content string) (map[string]interface{}, error) {
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

func (r *courseRepository) DeleteComment(ctx context.Context, userID int, courseID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"commentId":   1,
		"nickname":    "用户",
		"avater":      "https://example.com/avatar.jpg",
		"comment":     "已删除的评论",
		"delete_Date": "2023-12-01 10:00:00",
	}, nil
}

func (r *courseRepository) ReplyComment(ctx context.Context, userID int, courseID, commentID, content string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"commentId":   2,
		"nickname":    "用户",
		"avater":      "https://example.com/avatar.jpg",
		"comment":     content,
		"commentDate": "2023-12-01 10:00:00",
		"isreply":     true,
		"reply_id":    commentID,
		"replies":     []map[string]interface{}{},
	}, nil
}

func (r *courseRepository) DeleteReply(ctx context.Context, userID int, courseID, commentID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"commentId":   2,
		"nickname":    "用户",
		"avater":      "https://example.com/avatar.jpg",
		"comment":     "已删除的回复",
		"delete_Date": "2023-12-01 10:00:00",
	}, nil
}

func (r *courseRepository) AddView(ctx context.Context, courseID string) (int, error) {
	return 1001, nil
}

func (r *courseRepository) CollectCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"iscollected": true,
		"collections": 151,
	}, nil
}

func (r *courseRepository) UncollectCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"iscollected": false,
		"collections": 150,
	}, nil
}

func (r *courseRepository) LikeCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"isliked": true,
		"likes":   201,
	}, nil
}

func (r *courseRepository) UnlikeCourse(ctx context.Context, userID int, courseID string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"isliked": false,
		"likes":   200,
	}, nil
}

func (r *courseRepository) GetPending(ctx context.Context, cursor, limit int) ([]map[string]interface{}, error) {
	return []map[string]interface{}{
		{
			"submitor":     "user1",
			"submitDate":   "2023-12-01 10:00:00",
			"reourceId":    1,
			"resourceType": "course",
			"resourcename": "新课程资源",
			"catagory":     "教学资料",
			"link":         "https://example.com/resource",
			"description":  "课程相关资源",
			"tags":         []string{"讲义", "视频"},
			"file":         "resource.pdf",
		},
	}, nil
}
