package gomod

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
)

func run(workdir string, cmdline ...string) (string, error) {
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Dir = workdir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func LocalRevInfo(workdir string) (*RevInfo, error) {
	if _, err := os.Stat(".git/shallow"); err == nil {
		if _, err := run(workdir, "git", "fetch", "--unshallow"); err != nil {
			return nil, errors.Wrap(err, "git fetch --unshallow failed")
		}
	}

	out, err := run(workdir, "git", "log", "--no-decorate", "-n1", `--format=format:%H %ct`, "--")
	if err != nil {
		return nil, errors.Wrap(err, "git log failed")
	}

	d, err := run(workdir, "git", "describe", "--always", "--tags", "--match", `v*`, "--dirty")
	if err != nil {
		return nil, errors.Wrap(err, "git describe run failed")
	}

	desc := d
	if !strings.HasPrefix(desc, "v") {
		desc = "v0.0.0"
	}

	out = fmt.Sprintf("%s %s", out, strings.TrimSpace(desc))

	vTag := strings.Split(out, " ")
	t, err := strconv.ParseInt(vTag[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid time from git log: %q", out)
	}

	name := "main"
	offset := 0

	if len(vTag) == 3 {
		name = strings.Split(vTag[2], "-g")[0]

		// git describe will return like v0.1.3-8
		if i := strings.LastIndex(name, "-"); i > -1 {
			o, err := strconv.ParseInt(name[i+1:], 10, 64)
			if err == nil {
				name = name[0:i]
				offset = int(o)
			}
		}
	}

	if strings.Contains(d, "-dirty") {
		name += "-dirty"
	}

	return &RevInfo{
		Name:   name,
		Offset: offset,
		Short:  vTag[0][0:12],
		Time:   time.Unix(t, 0).UTC(),
	}, nil
}

type RevInfo struct {
	Name   string
	Offset int
	Short  string
	Time   time.Time
}

func (v RevInfo) Version() string {
	if v.Offset == 0 {
		f := strings.Split(module.PseudoVersion(semver.Major(v.Name), "", v.Time, v.Short), "-")
		f[0] = v.Name
		return strings.Join(f, "-")
	}
	return module.PseudoVersion(semver.Major(v.Name), v.Name, v.Time, v.Short)
}
