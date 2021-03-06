/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/solo-io/qloo/pkg/storage/crd/solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResolverMapLister helps list ResolverMaps.
type ResolverMapLister interface {
	// List lists all ResolverMaps in the indexer.
	List(selector labels.Selector) (ret []*v1.ResolverMap, err error)
	// ResolverMaps returns an object that can list and get ResolverMaps.
	ResolverMaps(namespace string) ResolverMapNamespaceLister
	ResolverMapListerExpansion
}

// resolverMapLister implements the ResolverMapLister interface.
type resolverMapLister struct {
	indexer cache.Indexer
}

// NewResolverMapLister returns a new ResolverMapLister.
func NewResolverMapLister(indexer cache.Indexer) ResolverMapLister {
	return &resolverMapLister{indexer: indexer}
}

// List lists all ResolverMaps in the indexer.
func (s *resolverMapLister) List(selector labels.Selector) (ret []*v1.ResolverMap, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ResolverMap))
	})
	return ret, err
}

// ResolverMaps returns an object that can list and get ResolverMaps.
func (s *resolverMapLister) ResolverMaps(namespace string) ResolverMapNamespaceLister {
	return resolverMapNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResolverMapNamespaceLister helps list and get ResolverMaps.
type ResolverMapNamespaceLister interface {
	// List lists all ResolverMaps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ResolverMap, err error)
	// Get retrieves the ResolverMap from the indexer for a given namespace and name.
	Get(name string) (*v1.ResolverMap, error)
	ResolverMapNamespaceListerExpansion
}

// resolverMapNamespaceLister implements the ResolverMapNamespaceLister
// interface.
type resolverMapNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResolverMaps in the indexer for a given namespace.
func (s resolverMapNamespaceLister) List(selector labels.Selector) (ret []*v1.ResolverMap, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ResolverMap))
	})
	return ret, err
}

// Get retrieves the ResolverMap from the indexer for a given namespace and name.
func (s resolverMapNamespaceLister) Get(name string) (*v1.ResolverMap, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("resolvermap"), name)
	}
	return obj.(*v1.ResolverMap), nil
}
