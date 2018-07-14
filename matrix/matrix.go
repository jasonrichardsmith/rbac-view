package matrix

import (
	"errors"
	"log"
	"sort"
	"sync"

	"github.com/jasonrichardsmith/rbac-view/client"
	"github.com/jasonrichardsmith/rbac-view/colors"
	rbac "k8s.io/api/rbac/v1"
)

var verbmutex = &sync.Mutex{}
var objectmutex = &sync.Mutex{}

type Matrix struct {
	Type      string
	Verbs     map[string]string
	Objects   []string
	Roles     []Role
	rolemutex *sync.Mutex
	wg        sync.WaitGroup
}

type Role struct {
	Name     string
	Objects  map[string][]string
	Subjects []rbac.Subject
}

func NewRole() Role {
	return Role{
		Objects:  make(map[string][]string),
		Subjects: make([]rbac.Subject, 0),
	}
}

func New(roletype string) Matrix {
	return Matrix{
		Type:      roletype,
		Verbs:     make(map[string]string),
		Objects:   make([]string, 0),
		Roles:     make([]Role, 0),
		rolemutex: &sync.Mutex{},
	}
}

func (m *Matrix) addVerbs(verbs []string) {
	verbmutex.Lock()
	for _, v := range verbs {
		if _, ok := m.Verbs[v]; !ok {
			m.Verbs[v] = colors.GetUnique()
		}
	}
	verbmutex.Unlock()
}

func (m *Matrix) addObjects(objects []string) {
	objectmutex.Lock()
	for _, o := range objects {
		if !m.objectExists(o) {
			m.Objects = append(m.Objects, o)
		}
	}
	objectmutex.Unlock()

}

func (m *Matrix) objectExists(object string) bool {
	for _, co := range m.Objects {
		if co == object {
			return true
		}
	}
	return false

}

func (m *Matrix) Build(c client.Client) (err error) {
	if m.Type == "ClusterRoles" {
		err = m.getClusterLevel(c)

	} else if m.Type == "Roles" {
		return errors.New("Not built yet")
		// err := c.NameSpaceLevel()
	} else {
		return errors.New("Invalid type")
	}
	log.Println("sorting string")
	sort.Strings(m.Objects)
	sort.Slice(m.Roles, func(i, j int) bool {
		return m.Roles[i].Name < m.Roles[j].Name
	})
	return

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
			go m.addVerbs(rule.Verbs)
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
