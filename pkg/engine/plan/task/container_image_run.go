package task

import (
	"fmt"
	"github.com/go-courier/logr"
	"github.com/octohelm/wagon/pkg/logutil"
	"sort"
	"strings"
	"time"

	"dagger.io/dagger"
	"github.com/octohelm/wagon/pkg/engine/daggerutil"
	"github.com/octohelm/wagon/pkg/engine/plan/task/core"
	"golang.org/x/net/context"
)

func init() {
	core.DefaultFactory.Register(&ImageRun{})
}

type ImageRun struct {
	core.Task

	Input core.Image `json:"input"`

	Mounts     map[string]core.Mount          `json:"mounts"`
	Env        map[string]core.SecretOrString `json:"env"`
	Workdir    string                         `json:"workdir" default:"/"`
	User       string                         `json:"user" default:"root:root"`
	Always     bool                           `json:"always,omitempty"`
	Entrypoint []string                       `json:"entrypoint,omitempty"`
	Command    struct {
		Name  string                       `json:"name"`
		Args  []string                     `json:"args"`
		Flags map[string]core.StringOrBool `json:"flags"`
	} `json:"command,omitempty"`

	Output core.Image `json:"-" wagon:"generated,name=output"`
	Exit   int        `json:"-" wagon:"generated,name=exit"`
}

func (e *ImageRun) Do(ctx context.Context) error {
	return daggerutil.Do(ctx, func(c *dagger.Client) error {
		ct := c.Container().
			WithRootfs(c.Directory(dagger.DirectoryOpts{
				ID: e.Input.Rootfs.DirectoryID(),
			}))

		ct = e.Input.Config.ApplyTo(ct)

		for n := range e.Mounts {
			ct = e.Mounts[n].MountTo(c, ct)
		}

		for k := range e.Env {
			if envVar := e.Env[k]; envVar.Secret != nil {
				ct = ct.WithSecretVariable(k, c.Secret(envVar.Secret.SecretID()))
			} else {
				ct = ct.WithEnvVariable(k, envVar.Value)
			}
		}

		if workdir := e.Workdir; workdir != "" {
			ct = ct.WithWorkdir(workdir)
		}

		if user := e.User; user != "" {
			ct = ct.WithUser(user)
		}

		if entrypoint := e.Entrypoint; len(entrypoint) > 0 {
			ct = ct.WithEntrypoint(entrypoint)
		}

		if e.Always {
			// disable cache
			ct = ct.WithEnvVariable("__WAGON_EXEC_STARTED_AT", time.Now().String())
		}

		args := make([]string, 0)
		if name := e.Command.Name; name != "" {
			args = append(args, name)

			flagNames := make([]string, 0)
			for flag := range e.Command.Flags {
				flagNames = append(flagNames, flag)
			}
			sort.Strings(flagNames)

			for _, flag := range flagNames {
				v := e.Command.Flags[flag]
				if v.Bool != nil {
					args = append(args, flag)
				} else {
					args = append(args, flag, v.String)
				}
			}

			args = append(args, e.Command.Args...)
		}

		_, _ = fmt.Fprint(logutil.Forward(logr.FromContext(ctx).Info), strings.Join(args, " "))

		ct = ct.WithExec(args)

		exitCode, err := ct.ExitCode(ctx)
		if err != nil {
			return err
		}
		e.Exit = exitCode

		id, err := ct.ID(ctx)
		if err != nil {
			return err
		}
		if err := e.Output.Config.Resolve(ctx, c, id); err != nil {
			return err
		}
		e.Output.Platform = e.Input.Platform
		return e.Output.Rootfs.SetDirectoryIDBy(ctx, ct.Rootfs())
	})
}
