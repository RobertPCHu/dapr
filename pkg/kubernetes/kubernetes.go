package kubernetes

import (
	scheme "github.com/actionscore/actions/pkg/client/clientset/versioned"
	"github.com/actionscore/actions/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func Clients() (kubernetes.Interface, scheme.Interface, error) {
	client := utils.GetKubeClient()
	config := utils.GetConfig()

	eventingClient, err := scheme.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return client, eventingClient, nil
}

func GetDeployment(name, namespace string) (*appsv1.Deployment, error) {
	client := utils.GetKubeClient()

	dep, err := client.AppsV1().Deployments(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return dep, nil
}

func UpdateDeployment(deployment *appsv1.Deployment) error {
	client := utils.GetKubeClient()
	_, err := client.AppsV1().Deployments(deployment.ObjectMeta.Namespace).Update(deployment)
	if err != nil {
		return err
	}

	return nil
}

func CreateService(service *corev1.Service, namespace string) error {
	client := utils.GetKubeClient()

	_, err := client.CoreV1().Services(namespace).Create(service)
	if err != nil {
		return err
	}

	return nil
}

func ServiceExists(name, namespace string) bool {
	client := utils.GetKubeClient()

	_, err := client.CoreV1().Services(namespace).Get(name, meta_v1.GetOptions{})
	return err == nil
}

func GetEndpoints(name, namespace string) (*corev1.Endpoints, error) {
	client := utils.GetKubeClient()

	endpoints, err := client.CoreV1().Endpoints(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

func GetDeploymentsBySelector(selector meta_v1.LabelSelector) ([]appsv1.Deployment, error) {
	client := utils.GetKubeClient()

	s := labels.SelectorFromSet(selector.MatchLabels)

	dep, err := client.AppsV1().Deployments(meta_v1.NamespaceAll).List(meta_v1.ListOptions{
		LabelSelector: s.String(),
	})
	if err != nil {
		return nil, err
	}

	return dep.Items, nil
}