package impl_test

import (
	"blog/apps/blog"
	"blog/apps/blog/impl"
	"context"
	"testing"
)

var blogService blog.Service

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	ins, err := blogService.CreateBlog(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	blogService = impl.NewImpl()
}
