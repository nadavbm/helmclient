package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/nadavbm/helmclient"
)

var releaseName, releaseNamespace, chartPath string

func main() {
	parseCommandArgs()
	chart := helmclient.GetHelmChart(releaseName, chartPath)

	fmt.Println(time.Now(), "INFO:", "getting helm cliemt")
	client, err := helmclient.GetClient(releaseNamespace, chart)
	if err != nil {
		panic(err)
	}

	rel, err := client.InstallChart(releaseNamespace, releaseName)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Now(), "INFO:", "Successfully installed release: ", rel.Name)

	// check in other terminal
	time.Sleep(1 * time.Minute)
	res, err := client.UninstallChart()
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Now(), "INFO:", "Successfully uninstalled release: ", res.Release.Name)
}

func parseCommandArgs() {
	flag.StringVar(&releaseName, "rel", "helmut", "helm release name")
	flag.StringVar(&releaseNamespace, "ns", "helmut", "kubernetes namespace of helm release")
	flag.StringVar(&chartPath, "path", "helmut", "helm chart path")
	flag.Parse()
}
