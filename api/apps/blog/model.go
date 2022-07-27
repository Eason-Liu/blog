package blog

type Blog struct {
	Id        string //文章id
	Sumary    string //文章摘要信息，通过提前content内容获取
	CreateAt  string //创建时间
	UpdateAt  string //更新时间
	PubilshAt string //发布时间
	Status    Status //文章状态  草稿/发布
}

type BlogSet struct {
	Items []*Blog //文章列表
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{}
}

type CreateBlogRequest struct {
	TitleImg  string //文章图片
	TitleName string //文章标题
	SubTitle  string //文章副标题
	Content   string //文章内容
	Author    string //文章作者
}

type UpdateBlogRequest struct {
	UpdateMode UpdateMode
	*CreateBlogRequest
}

func NewPutUpdateBlogRequest() *UpdateBlogRequest {
	return &UpdateBlogRequest{
		UpdateMode: UPDATE_MODE_PUT,
	}
}

func NewPatchUpdateBlogRequest() *UpdateBlogRequest {
	return &UpdateBlogRequest{
		UpdateMode: UPDATE_MODE_PATCH,
	}
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

type DescribeBlogRequest struct {
	Id int
}

type UpdateBlogStatusRequest struct {
	Id     int
	status Status
}
