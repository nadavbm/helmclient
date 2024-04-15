package helmclient

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/kube"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

func getKubeConfig(namespace string) *genericclioptions.ConfigFlags {
	kubeConfigPath := fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	return kube.GetConfig(kubeConfigPath, "", namespace)
}

func inClusterConfig() (*genericclioptions.ConfigFlags, error) {
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	return &genericclioptions.ConfigFlags{
		Insecure:    &restConfig.Insecure,
		Username:    &restConfig.Username,
		Password:    &restConfig.Password,
		BearerToken: &restConfig.BearerToken,
		CAFile:      &restConfig.CAFile,
		CertFile:    &restConfig.CertFile,
		KeyFile:     &restConfig.KeyFile,
		APIServer:   &restConfig.Host,
	}, nil
}
