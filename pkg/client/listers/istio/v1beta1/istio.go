/*
Copyright 2019 Banzai Cloud.

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

package v1beta1

import (
	v1beta1 "github.com/banzaicloud/istio-operator/pkg/apis/istio/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IstioLister helps list Istios.
type IstioLister interface {
	// List lists all Istios in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Istio, err error)
	// Istios returns an object that can list and get Istios.
	Istios(namespace string) IstioNamespaceLister
	IstioListerExpansion
}

// istioLister implements the IstioLister interface.
type istioLister struct {
	indexer cache.Indexer
}

// NewIstioLister returns a new IstioLister.
func NewIstioLister(indexer cache.Indexer) IstioLister {
	return &istioLister{indexer: indexer}
}

// List lists all Istios in the indexer.
func (s *istioLister) List(selector labels.Selector) (ret []*v1beta1.Istio, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Istio))
	})
	return ret, err
}

// Istios returns an object that can list and get Istios.
func (s *istioLister) Istios(namespace string) IstioNamespaceLister {
	return istioNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IstioNamespaceLister helps list and get Istios.
type IstioNamespaceLister interface {
	// List lists all Istios in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.Istio, err error)
	// Get retrieves the Istio from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.Istio, error)
	IstioNamespaceListerExpansion
}

// istioNamespaceLister implements the IstioNamespaceLister
// interface.
type istioNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Istios in the indexer for a given namespace.
func (s istioNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Istio, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Istio))
	})
	return ret, err
}

// Get retrieves the Istio from the indexer for a given namespace and name.
func (s istioNamespaceLister) Get(name string) (*v1beta1.Istio, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("istio"), name)
	}
	return obj.(*v1beta1.Istio), nil
}