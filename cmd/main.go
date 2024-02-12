package main

import (
	"fmt"
	"time"

	"github.com/nadavbm/helmclient"
)

const (
	releaseName = "reloader"
	chartPath   = "/your/path/to/chart/reloader/"
)

func main() {
	chart := helmclient.GetHelmChart(releaseName, chartPath)

	releaseNamespace := "default"
	client, err := helmclient.GetClient(releaseNamespace, chart)
	if err != nil {
		panic(err)
	}

	rel, err := client.InstallChart(releaseNamespace, releaseName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully installed release: ", rel.Name)

	// check in other terminal
	time.Sleep(1 * time.Minute)
	res, err := client.UninstallChart()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully uninstalled release: ", res.Release.Name)
}
