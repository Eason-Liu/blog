package impl

import (
	"blog/apps/blog"
	"context"
	"errors"
	"github.com/infraboard/mcube/exception"
)

func (i *Impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	//校验
	if err := req.Valdate(); err != nil {
		return nil, exception.NewBadRequest("validate create blog request error, %s", err)
	}

	ins := blog.NewCreateBlog(req)
	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	set := blog.NewBlogSet()
	query := i.DB()

	//关键字查询
	if req.Keywords != "" {
		query = query.Where(
			"title_name LIKE ? OR content LIKE ?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
		)
	}

	//查询总条数
	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}

	//分页查询
	//LIMIT <offset>,<limit>
	query = query.Offset(req.Offset()).Limit(req.PageSize)
	if err := query.WithContext(ctx).Scan(&set.Items).Error; err != nil {
		return nil, err
	}

	return set, nil
}

func (i *Impl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, errors.New("not implment")
}

func (i *Impl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, errors.New("not implment")
}

func (i *Impl) DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error) {
	return nil, errors.New("not implment")
}

func (i *Impl) UpdateBlogStatus(ctx context.Context, req *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, errors.New("not implment")
}
