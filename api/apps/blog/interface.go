package blog

import "context"

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
}
