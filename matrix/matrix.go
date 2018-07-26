package matrix

import (
	"sync"

	"github.com/jasonrichardsmith/rbac-view/client"
	rbac "k8s.io/api/rbac/v1"
)

type Builder interface {
	Build() (Matrices, error)
}

type MatrixBuilder struct {
	client client.Client
}

func New(c client.Client) Builder {
	return MatrixBuilder{c}
}

func (mb MatrixBuilder) Build() (Matrices, error) {
	ms := Matrices{
		ClusterRoles: NewMatrix(),
		Roles:        NewMatrix(),
	}
	err := ms.ClusterRoles.getClusterLevel(mb.client)
	if err != nil {
		return ms, err
	}
	err = ms.Roles.getNamespaceLevel(mb.client)
	return ms, nil

}

type Matrices struct {
	ClusterRoles Matrix `json:"clusterRoles"`
	Roles        Matrix `json:"roles"`
}

type Matrix struct {
	Objects     []string `json:"objects"`
	Roles       []Role   `json:"roles"`
	objectmutex *sync.Mutex
	rolemutex   *sync.Mutex
	wg          sync.WaitGroup
}

func NewMatrix() Matrix {
	return Matrix{
		Objects:     make([]string, 0),
		Roles:       make([]Role, 0),
		rolemutex:   &sync.Mutex{},
		objectmutex: &sync.Mutex{},
	}
}

type Role struct {
	Name     string              `json:"name"`
	Objects  map[string][]string `json:"objects"`
	Subjects []rbac.Subject      `json:"subjects"`
}

func NewRole() Role {
	return Role{
		Objects:  make(map[string][]string),
		Subjects: make([]rbac.Subject, 0),
	}
}

func (m *Matrix) addObjects(objects []string) {
	m.objectmutex.Lock()
	for _, o := range objects {
		if !m.objectExists(o) {
			m.Objects = append(m.Objects, o)
		}
	}
	m.objectmutex.Unlock()
}

func (m *Matrix) objectExists(object string) bool {
	for _, co := range m.Objects {
		if co == object {
			return true
		}
	}
	return false

}

func (m *Matrix) getClusterLevel(c client.Client) (err error) {
	rbs, err := c.GetClusterRoleBindings()
	if err != nil {
		return err
	}
	for _, rb := range rbs {
		m.wg.Add(1)
		go m.getClusterRole(c, rb)
	}
	m.wg.Wait()
	return
}

func (m *Matrix) getClusterRole(c client.Client, rb rbac.ClusterRoleBinding) (err error) {
	defer m.wg.Done()
	r := NewRole()
	r.Subjects = rb.Subjects
	r.Name = rb.RoleRef.Name
	if rb.RoleRef.Kind == "ClusterRole" {
		cr, err := c.GetClusterRole(rb.RoleRef.Name)
		if err != nil {
			return err
		}
		for _, rule := range cr.Rules {
			go m.addObjects(rule.Resources)
			for _, o := range rule.Resources {
				r.Objects[o] = rule.Verbs
			}
		}
	}
	m.rolemutex.Lock()
	m.Roles = append(m.Roles, r)
	m.rolemutex.Unlock()
	return nil
}

func (m *Matrix) getNamespaceLevel(c client.Client) (err error) {
	rbs, err := c.GetRoleBindings()
	if err != nil {
		return err
	}
	for _, rb := range rbs {
		m.wg.Add(1)
		go m.getRole(c, rb)
	}
	m.wg.Wait()
	return
}

func (m *Matrix) getRole(c client.Client, rb rbac.RoleBinding) (err error) {
	defer m.wg.Done()
	r := NewRole()
	r.Subjects = rb.Subjects
	r.Name = rb.RoleRef.Name
	if rb.RoleRef.Kind == "Role" {
		cr, err := c.GetRole(rb.RoleRef.Name, rb.ObjectMeta.Namespace)
		if err != nil {
			return err
		}
		for _, rule := range cr.Rules {
			go m.addObjects(rule.Resources)
			for _, o := range rule.Resources {
				r.Objects[o] = rule.Verbs
			}
		}
	}
	m.rolemutex.Lock()
	m.Roles = append(m.Roles, r)
	m.rolemutex.Unlock()
	return nil
}
