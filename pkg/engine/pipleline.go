package engine

import (
	"bytes"
	cueerrors "cuelang.org/go/cue/errors"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type Pipeline struct {
	Action []string `arg:""`
	Plan   string   `flag:",omitempty" alias:"p"`
	Output string   `flag:",omitempty" alias:"o"`
}

func (c *Pipeline) SetDefaults() {
	if c.Plan == "" {
		c.Plan = "./wagon.cue"
	}
}

func (c *Pipeline) Run(ctx context.Context) error {
	p, err := New(
		ctx,
		WithPlan(c.Plan),
		WithOutput(c.Output),
	)
	if err != nil {
		return err
	}

	if err := p.Run(ctx, c.Action...); err != nil {
		// print full cue errors if exists
		if errlist := cueerrors.Errors(err); len(errlist) > 0 {
			buf := bytes.NewBuffer(nil)
			for i := range errlist {
				cueerrors.Print(buf, errlist[i], nil)
			}
			return errors.New(buf.String())
		}
		return err
	}
	return nil
}
