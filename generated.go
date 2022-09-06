package main

import (
	"context"
	"fmt"

	"ci/hugo/gen/core"

	"github.com/dagger/cloak/engine"
)

func (r *query) hugo(ctx context.Context) (*Hugo, error) {

	return new(Hugo), nil

}

type query struct{}
type hugo struct{}

func main() {
	ctx := context.Background()
	cfg := engine.Config{
		ConfigPath: "./cloak.yml",
		Workdir:    ".",
	}
	err := engine.Start(ctx, &cfg, func(ctx engine.Context) error {
		deb, err := core.Image(ctx, "alpine:latest")
		if err != nil {
			return (err)
		}

		curled, err := core.Exec(ctx, deb.Core.Image.ID, core.ExecInput{
			Args: []string{"sh", "-c", "apk add curl && curl -L https://github.com/gohugoio/hugo/releases/download/v0.102.3/hugo_0.102.3_Linux-64bit.tar.gz | tar -xz"},
			//Args: []string{"sh", "/mnt/run.sh"},
			//Mounts: []core.MountInput{
			//	{Fs: wdID, Path: "/mnt"},
			//},
		})
		if err != nil {
			return fmt.Errorf("in curl: %w", err)
		}

		wd, err := core.Workdir(ctx)
		if err != nil {
			return err
		}

		wdID := wd.Host.Workdir.Read.ID
		_ = wdID

		execResp, err := core.Exec(ctx, curled.Core.Filesystem.ID, core.ExecInput{
			Args: []string{"/hugo", "--help"},
			//Mounts: []core.MountInput{
			//	{
			//		Fs:   wdID,
			//		Path: "/mnt",
			//	},
			//},
			Workdir: "/",
		})
		if err != nil {
			return fmt.Errorf("in final: %w", err)
		}

		outFS := execResp.Core.Filesystem

		println(outFS.Exec.Stdout)
		return nil
	})
	if err != nil {
		panic(err)
	}

}
