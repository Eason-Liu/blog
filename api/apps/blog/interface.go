package blog

import "context"

type Service interface {
	CreatebBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	UpdateBlog()
	Deleteblog()
	QueryBlog()
	DescribeBlog()
	UpdateBlogStatus()
}
