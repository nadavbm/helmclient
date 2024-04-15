package helmclient

import (
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

// HelmChart
type HelmChart struct {
	ReleaseName string
	Chart       *chart.Chart
}

// GetHelmChart from local os path
func GetHelmChart(releaseName, chartPath string) *HelmChart {
	chart, err := loader.Load(chartPath)
	if err != nil {
		panic(err)
	}

	return &HelmChart{
		ReleaseName: releaseName,
		Chart:       chart,
	}
}
