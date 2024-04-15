package helmclient

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

// InstallChart in kubernetes namespace and chose helm release name
func (c *Client) InstallChart(namespace, releaseName string) (*release.Release, error) {
	iCli := action.NewInstall(c.Config)
	iCli.Namespace = namespace
	iCli.ReleaseName = releaseName
	return iCli.Run(c.Chart.Chart, nil)
}

// UpgradeChart in kubernetes namespace to the selected release (release name)
func (c *Client) UpgradeChart(namespace, releaseName string) (*release.Release, error) {
	iCli := action.NewUpgrade(c.Config)
	return iCli.Run(c.Chart.ReleaseName, c.Chart.Chart, nil)
}

// UninstallChart uninstall release from kubernetes
func (c *Client) UninstallChart() (*release.UninstallReleaseResponse, error) {
	iCli := action.NewUninstall(c.Config)
	return iCli.Run(c.Chart.ReleaseName)
}
