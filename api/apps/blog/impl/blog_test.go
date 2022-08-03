package impl_test

import (
	"blog/apps/blog"
	"blog/apps/blog/impl"
	"blog/conf"
	"context"
	"github.com/infraboard/mcube/exception"
	"testing"
)

var blogService blog.Service

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.TitleName = "blog5"
	req.Content = "blog2 hello gohper,this is a golang."
	ins, err := blogService.CreateBlog(context.Background(), req)
	if err != nil {
		if v, ok := err.(exception.APIException); ok {
			t.Log(v.ErrorCode())
		}
		t.Fatal(err)
	}

	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	req.Keywords = "blog"
	ins, err := blogService.QueryBlog(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeBlogRequest(2)
	ins, err := blogService.DescribeBlog(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest(2)
	ins, err := blogService.DeleteBlog(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateBlog(t *testing.T) {
	req := blog.NewPutUpdateBlogRequest(3)
	req.Author = "admin"
	ins, err := blogService.UpdateBlog(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdatePatchBlog(t *testing.T) {
	req := blog.NewPutUpdateBlogRequest(4)
	req.TitleImg = "update image"
	ins, err := blogService.UpdateBlog(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestValidate(t *testing.T) {
	ins := blog.NewCreateBlog(blog.NewCreateBlogRequest())
	if err := ins.CreateBlogRequest.Validate(); err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	if err := conf.LoadConfigFromToml("/Users/easonliu/Downloads/blog/api/etc/config.toml"); err != nil {
		panic(err)
	}

	svr := impl.NewImpl()
	if err := svr.Init(); err != nil {
		panic(err)
	}

	blogService = svr
}
