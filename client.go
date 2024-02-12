package helmclient

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Client struct {
	Chart    *HelmChart
	DebugLog *action.DebugLog
	Config   *action.Configuration
}

func GetClient(namespace string, chart *HelmChart) (*Client, error) {
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(getKubeConfig(namespace), namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		panic(err)
	}

	return &Client{
		Chart:  chart,
		Config: actionConfig,
	}, nil
}

func getKubeConfig(namespace string) *genericclioptions.ConfigFlags {
	kubeConfigPath := fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	return kube.GetConfig(kubeConfigPath, "", namespace)
}
