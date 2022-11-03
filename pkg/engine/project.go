package engine

import (
	"cuelang.org/go/cue/build"
	"cuelang.org/go/cue/cuecontext"
	cueload "cuelang.org/go/cue/load"
	"github.com/go-courier/logr"
	"github.com/octohelm/cuemod/pkg/cuemod"
	"github.com/octohelm/wagon/cuepkg"
	"github.com/octohelm/wagon/pkg/engine/plan"
	"github.com/octohelm/wagon/pkg/engine/plan/task/core"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"os"
	"path"
	"strings"

	_ "github.com/octohelm/wagon/pkg/engine/plan/task"
)

func init() {
	if err := cuepkg.RegistryCueStdlibs(); err != nil {
		panic(err)
	}
}

type Project interface {
	Run(ctx context.Context, action ...string) error
}

type option struct {
	entryFile string
	output    string
}

type OptFunc = func(o *option)

func WithPlan(root string) OptFunc {
	return func(o *option) {
		o.entryFile = root
	}
}

func WithOutput(output string) OptFunc {
	return func(o *option) {
		o.output = output
	}
}

func New(ctx context.Context, opts ...OptFunc) (Project, error) {
	c := &project{}
	for i := range opts {
		opts[i](&c.opt)
	}

	cwd, _ := os.Getwd()
	sourceRoot := path.Join(cwd, c.opt.entryFile)

	if strings.HasSuffix(sourceRoot, ".cue") {
		sourceRoot = path.Dir(sourceRoot)
	}

	c.sourceRoot = sourceRoot

	buildConfig := cuemod.ContextFor(cwd).BuildConfig(ctx)

	instances := cueload.Instances([]string{c.opt.entryFile}, buildConfig)
	if len(instances) != 1 {
		return nil, errors.New("only one package is supported at a time")
	}
	c.instance = instances[0]

	if err := c.instance.Err; err != nil {
		return nil, err
	}

	return c, nil
}

type project struct {
	opt        option
	sourceRoot string
	instance   *build.Instance
}

func (c *project) Run(ctx context.Context, action ...string) error {
	logr.FromContext(ctx).WithValues("name", "Project").Debug("loading...")

	val := cuecontext.New().BuildInstance(c.instance)
	if err := val.Err(); err != nil {
		return err
	}

	cueValue := plan.WrapValue(val)
	workdir := plan.NewWorkdir(c.sourceRoot, "")
	registryAuthStore := plan.NewRegistryAuthStore()

	logr.FromContext(ctx).WithValues("name", "Project").Debug("starting...")

	runner := plan.NewRunner(cueValue, c.opt.output, &core.FS{}, &core.Image{})

	ctx = plan.TaskRunnerFactoryContext.Inject(ctx, core.DefaultFactory)
	ctx = plan.WorkdirContext.Inject(ctx, workdir)
	ctx = plan.RegistryAuthStoreContext.Inject(ctx, registryAuthStore)

	return runner.Run(ctx, action)
}
