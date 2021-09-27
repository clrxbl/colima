package kubernetes

import (
	_ "embed"
	"github.com/abiosoft/colima/cli"
	"github.com/abiosoft/colima/environment"
	"path/filepath"
	"strconv"
)

func installContainerdDeps(guest environment.GuestActions, r *cli.ActiveCommandChain) {
	// fix cni path
	r.Add(func() error {
		cniDir := "/opt/cni/bin"
		if err := guest.Run("ls", cniDir); err == nil {
			return nil
		}

		if err := guest.Run("sudo", "mkdir", "-p", filepath.Dir(cniDir)); err != nil {
			return err
		}
		return guest.Run("sudo", "ln", "-s", "/var/lib/rancher/k3s/data/current/bin", cniDir)
	})

	// fix cni config
	r.Add(func() error {
		return guest.Run("sudo", "mkdir", "-p", "/etc/cni/net.d")
	})
	r.Add(func() error {
		return guest.Run("sudo", "sh", "-c", "echo "+strconv.Quote(k3sFlannelConflist)+" > /etc/cni/net.d/10-flannel.conflist")
	})
}

//go:embed k3s-flannel.json
var k3sFlannelConflist string
