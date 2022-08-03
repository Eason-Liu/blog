package blog

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"time"
)

var validate = validator.New()

func NewCreateBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		CreateAt:          time.Now().Unix(),
		CreateBlogRequest: req,
		Status:            STATUS_DRAF,
	}
}

type Blog struct {
	Id                 int    `json:"id"`               //文章id
	Summary            string `json:"summary" gorm:"-"` //文章摘要信息，通过提前content内容获取
	CreateAt           int64  `json:"create_at"`        //创建时间
	UpdateAt           int64  `json:"update_at"`        //更新时间
	PublishAt          int64  `json:"publish_at"`       //发布时间
	*CreateBlogRequest        //用户提交的数据
	Status             Status `json:"status"` //文章状态  草稿/发布
}

func (b *Blog) Strings() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	Total int64   `json:"total"` //总条数,用于前端分页
	Items []*Blog `json:"items"` //文章列表
}

func (b *BlogSet) Strings() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{}
}

type CreateBlogRequest struct {
	TitleImg  string `json:"title_img"`                     //文章图片
	TitleName string `json:"title_name" valdate:"required"` //文章标题
	SubTitle  string `json:"sub_title"`                     //文章副标题
	Content   string `json:"content" valdate:"required"`    //文章内容
	Author    string `json:"author"`                        //文章作者
}

//校验对象
func (req *CreateBlogRequest) Valdate() error {
	return validate.Struct(req)
}

type UpdateBlogRequest struct {
	BlogId     int
	UpdateMode UpdateMode
	*CreateBlogRequest
}

func NewPutUpdateBlogRequest(id int) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

func NewPatchUpdateBlogRequest(id int) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PATCH,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

func NewDeleteBlogRequest(id int) *DeleteBlogRequest {
	return &DeleteBlogRequest{Id: id}
}

type DeleteBlogRequest struct {
	Id int
}

type QueryBlogRequest struct {
	PageSize   int
	PageNumber int
	Keywords   string
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20, //每页显示多少条
		PageNumber: 1,  //当前页是多少
	}
}

func (req *QueryBlogRequest) Offset() int {
	return (req.PageNumber - 1) * req.PageSize
}

func NewDescribeBlogRequest(id int) *DescribeBlogRequest {
	return &DescribeBlogRequest{Id: id}
}

type DescribeBlogRequest struct {
	Id int
}

type UpdateBlogStatusRequest struct {
	Id     int
	status Status
}
