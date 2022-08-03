package impl

import (
	"blog/apps/blog"
	"context"
	"errors"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/exception"
)

func (i *Impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	//校验
	if err := req.Validate(); err != nil {
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

func (i *Impl) DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error) {
	ins := blog.NewCreateBlog(blog.NewCreateBlogRequest())

	query := i.DB().Where("id=?", req.Id)
	if err := query.Find(ins).Error; err != nil {
		return nil, err
	}

	if ins.Id == 0 {
		return nil, exception.NewBadRequest("blog %d not found", req.Id)
	}

	return ins, nil
}

func (i *Impl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := i.DB().Delete(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.BlogId))
	if err != nil {
		return nil, err
	}
	switch req.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		ins.CreateBlogRequest = req.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		if err := mergo.MapWithOverwrite(ins.CreateBlogRequest, req.CreateBlogRequest); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("update mode not support %s", req.UpdateMode)
	}

	if err := ins.CreateBlogRequest.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate request error, %s", err)
	}

	if err := i.DB().WithContext(ctx).Updates(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) UpdateBlogStatus(ctx context.Context, req *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, errors.New("not implment")
}
