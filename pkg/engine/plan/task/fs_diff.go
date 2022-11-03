package task

import (
	"context"

	"dagger.io/dagger"
	"github.com/octohelm/wagon/pkg/engine/daggerutil"
	"github.com/octohelm/wagon/pkg/engine/plan/task/core"
)

func init() {
	core.DefaultFactory.Register(&Diff{})
}

type Diff struct {
	core.Task

	Upper core.FS `json:"upper"`
	Lower core.FS `json:"lower"`

	Output core.FS `json:"-" wagon:"generated,name=output"`
}

func (e *Diff) Do(ctx context.Context) error {
	return daggerutil.Do(ctx, func(c *dagger.Client) error {
		upper := c.Directory(dagger.DirectoryOpts{
			ID: e.Upper.DirectoryID(),
		})
		lower := c.Directory(dagger.DirectoryOpts{
			ID: e.Lower.DirectoryID(),
		})
		return e.Output.SetDirectoryIDBy(ctx, lower.Diff(upper))
	})
}
