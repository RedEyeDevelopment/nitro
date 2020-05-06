package nitro

import (
	"fmt"

	"github.com/craftcms/nitro/validate"
)

func Mount(name, folder, site string) (*Action, error) {
	if err := validate.MachineName(name); err != nil {
		return nil, err
	}
	if err := validate.Path(folder); err != nil {
		return nil, err
	}
	if err := validate.Hostname(site); err != nil {
		return nil, err
	}

	target := "/nitro/sites/" + site

	return &Action{
		Type:       "mount",
		Output:     fmt.Sprintf("Mounting %s to %s", folder, target),
		UseSyscall: false,
		Args:       []string{"mount", folder, name + ":" + target},
	}, nil
}
