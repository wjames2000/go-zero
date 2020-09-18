package docker

import (
	"errors"

	"github.com/urfave/cli"
	"github.com/wjames2000/go-zero/tools/goctl/gen"
)

func DockerCommand(c *cli.Context) error {
	goFile := c.String("go")
	if len(goFile) == 0 {
		return errors.New("-go can't be empty")
	}

	return gen.GenerateDockerfile(goFile, "-f", "etc/config.yaml")
}
