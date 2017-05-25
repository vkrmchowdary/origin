// This file was automatically generated by lister-gen

package internalversion

import (
	api "github.com/openshift/origin/pkg/authorization/api"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterPolicyBindingLister helps list ClusterPolicyBindings.
type ClusterPolicyBindingLister interface {
	// List lists all ClusterPolicyBindings in the indexer.
	List(selector labels.Selector) (ret []*api.ClusterPolicyBinding, err error)
	// Get retrieves the ClusterPolicyBinding from the index for a given name.
	Get(name string) (*api.ClusterPolicyBinding, error)
	ClusterPolicyBindingListerExpansion
}

// clusterPolicyBindingLister implements the ClusterPolicyBindingLister interface.
type clusterPolicyBindingLister struct {
	indexer cache.Indexer
}

// NewClusterPolicyBindingLister returns a new ClusterPolicyBindingLister.
func NewClusterPolicyBindingLister(indexer cache.Indexer) ClusterPolicyBindingLister {
	return &clusterPolicyBindingLister{indexer: indexer}
}

// List lists all ClusterPolicyBindings in the indexer.
func (s *clusterPolicyBindingLister) List(selector labels.Selector) (ret []*api.ClusterPolicyBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*api.ClusterPolicyBinding))
	})
	return ret, err
}

// Get retrieves the ClusterPolicyBinding from the index for a given name.
func (s *clusterPolicyBindingLister) Get(name string) (*api.ClusterPolicyBinding, error) {
	key := &api.ClusterPolicyBinding{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("clusterpolicybinding"), name)
	}
	return obj.(*api.ClusterPolicyBinding), nil
}