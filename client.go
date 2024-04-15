package helmclient

import (
	"fmt"
	"os"
	"time"

	"helm.sh/helm/v3/pkg/action"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Client struct {
	Chart    *HelmChart
	DebugLog *action.DebugLog
	Config   *action.Configuration
}

func GetClient(namespace string, chart *HelmChart) (*Client, error) {
	if os.Getenv("HELMUT") != "" {
		fmt.Println(time.Now(), "INFO:", "setting in cluster kubernetes client config")
		return GetInClusterClient(namespace, chart)
	}

	fmt.Println(time.Now(), "INFO:", "setting kubernetes client from kubeconfig")
	return GetCliClient(namespace, chart)
}

func GetCliClient(namespace string, chart *HelmChart) (*Client, error) {
	actionConfig := new(action.Configuration)
	gconfig := getKubeConfig(namespace)
	fmt.Println("gconfig", gconfig)
	restConfig, err := gconfig.ToRESTConfig()
	if err != nil {
		fmt.Println("fail to create rest config")
	}

	fmt.Printf("to rest config %v", restConfig)

	inClusterConfig, err := inClusterConfig()
	if err != nil {
		fmt.Println("fail to get in cluster config")
	}
	fmt.Printf("in cluster config %v", inClusterConfig)
	if err := actionConfig.Init(getKubeConfig(namespace), namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Printf(format, v)
	}); err != nil {
		panic(err)
	}

	return &Client{
		Chart:  chart,
		Config: actionConfig,
	}, nil
}

func GetInClusterClient(namespace string, chart *HelmChart) (*Client, error) {
	actionConfig := new(action.Configuration)

	inClusterConfig, err := inClusterConfig()
	if err != nil {
		fmt.Println("fail to get in cluster config")
	}
	fmt.Printf("in cluster config %v", inClusterConfig)
	if err := actionConfig.Init(inClusterConfig, namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Printf(format, v)
	}); err != nil {
		panic(err)
	}

	return &Client{
		Chart:  chart,
		Config: actionConfig,
	}, nil
}
