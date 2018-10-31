package client

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"sync"

	log "github.com/Sirupsen/logrus"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig      string
	ErrNoKubeConfig = errors.New("Kubeconfig has not been set, use the kubeconfig flag to point at a Kubeconfig")
)

func init() {
	kc := defaultKubeConfig()
	flag.StringVar(&kubeconfig, "kubeconfig", kc, "absolute path to the kubeconfig file")
}

func defaultKubeConfig() string {
	if kc := os.Getenv("KUBECONFIG"); kc != "" {
		return kc
	}
	if h := os.Getenv("HOME"); h != "" {
		return filepath.Join(h, ".kube", "config")
	}
	if h := os.Getenv("USERPROFILE"); h != "" {
		return filepath.Join(h, ".kube", "config")
	}
	return ""
}

type Client struct {
	k8sclient *kubernetes.Clientset
}

func New() (client Client, err error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if kubeconfig == "" {
		return client, ErrNoKubeConfig
	}
	var config *rest.Config
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return
	}

	// create the clientset
	client.k8sclient, err = kubernetes.NewForConfig(config)
	return
}

func (c *Client) GetClusterRoleBindings() ([]rbac.ClusterRoleBinding, error) {
	log.Info("Retrieving ClusterRoleBindings")
	rolebindings, err := c.k8sclient.Rbac().ClusterRoleBindings().List(metav1.ListOptions{})
	log.Infof("Retrieved %v ClusterRoleBindings", len(rolebindings.Items))
	return rolebindings.Items, err
}

func (c *Client) GetClusterRole(name string) (*rbac.ClusterRole, error) {
	log.Infof("Retrieving ClusterRole %v", name)
	return c.k8sclient.Rbac().ClusterRoles().Get(name, metav1.GetOptions{})
}

func (c *Client) GetRoleBindings() ([]rbac.RoleBinding, error) {
	log.Infof("Retrieving RoleBindings")
	rb := make([]rbac.RoleBinding, 0)
	nss, err := c.k8sclient.Core().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return rb, err
	}
	rbchan := make(chan []rbac.RoleBinding)
	errs := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(len(nss.Items))
	for _, ns := range nss.Items {
		go c.getNSRoleBindings(ns.ObjectMeta.Name, rbchan, errs)
	}
	go func() {
		for rbs := range rbchan {
			rb = append(rb, rbs...)
			wg.Done()
		}
	}()

	wg.Wait()

	return rb, err
}

func (c *Client) getNSRoleBindings(namespace string, rb chan<- []rbac.RoleBinding, errs chan<- error) {
	log.Infof("Retrieving RoleBindings for Namespace %v", namespace)
	rolebindings, err := c.k8sclient.Rbac().RoleBindings(namespace).List(metav1.ListOptions{})
	if err != nil {
		errs <- err
		return
	}
	log.Infof("Retrieved %v RoleBindings for Namespace %v", len(rolebindings.Items), namespace)
	rb <- rolebindings.Items
}

func (c *Client) GetRole(name string, namespace string) (*rbac.Role, error) {
	log.Infof("Retrieving Role %v in namespace %v", name, namespace)
	return c.k8sclient.Rbac().Roles(namespace).Get(name, metav1.GetOptions{})
}
