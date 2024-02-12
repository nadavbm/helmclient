package helmclient

import (
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

type HelmChart struct {
	ReleaseName string
	Chart       *chart.Chart
}

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
