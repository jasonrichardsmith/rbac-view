package client

import (
	"flag"
	"os"
	"path/filepath"

	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig *string

func init() {
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

type Client struct {
	k8sclient *kubernetes.Clientset
}

func New() (client Client, err error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	var config *rest.Config
	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return
	}

	// create the clientset
	client.k8sclient, err = kubernetes.NewForConfig(config)
	return
}

func (c *Client) GetClusterRoleBindings() ([]rbac.ClusterRoleBinding, error) {
	rolebindings, err := c.k8sclient.Rbac().ClusterRoleBindings().List(metav1.ListOptions{})
	return rolebindings.Items, err
}

func (c *Client) GetClusterRole(name string) (*rbac.ClusterRole, error) {
	return c.k8sclient.Rbac().ClusterRoles().Get(name, metav1.GetOptions{})
}

func (c *Client) GetRoleBindings() ([]rbac.RoleBinding, error) {
	rb := make([]rbac.RoleBinding, 0)
	nss, err := c.k8sclient.Core().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return rb, err
	}
	for _, ns := range nss.Items {
		rolebindings, err := c.k8sclient.Rbac().RoleBindings(ns.ObjectMeta.Name).List(metav1.ListOptions{})
		if err != nil {
			return rb, err
		}
		rb = append(rb, rolebindings.Items...)

	}
	return rb, nil
}

func (c *Client) GetRole(name string, namespace string) (*rbac.Role, error) {
	return c.k8sclient.Rbac().Roles(namespace).Get(name, metav1.GetOptions{})
}
