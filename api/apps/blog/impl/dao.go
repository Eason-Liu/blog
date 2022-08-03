package impl

import (
	"blog/apps/blog"
	"context"
)

func (i *Impl) save(ctx context.Context, ins *blog.Blog) error {
	return i.DB().WithContext(ctx).Create(ins).Error
}
