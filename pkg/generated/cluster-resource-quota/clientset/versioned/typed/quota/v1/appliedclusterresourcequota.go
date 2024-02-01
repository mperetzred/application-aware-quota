/*
Copyright 2023 The AAQ Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	quotav1 "github.com/openshift/api/quota/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	scheme "kubevirt.io/application-aware-quota/pkg/generated/cluster-resource-quota/clientset/versioned/scheme"
)

// AppliedClusterResourceQuotasGetter has a method to return a AppliedClusterResourceQuotaInterface.
// A group's client should implement this interface.
type AppliedClusterResourceQuotasGetter interface {
	AppliedClusterResourceQuotas(namespace string) AppliedClusterResourceQuotaInterface
}

// AppliedClusterResourceQuotaInterface has methods to work with AppliedClusterResourceQuota resources.
type AppliedClusterResourceQuotaInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*quotav1.AppliedClusterResourceQuota, error)
	List(ctx context.Context, opts v1.ListOptions) (*quotav1.AppliedClusterResourceQuotaList, error)
	AppliedClusterResourceQuotaExpansion
}

// appliedClusterResourceQuotas implements AppliedClusterResourceQuotaInterface
type appliedClusterResourceQuotas struct {
	client rest.Interface
	ns     string
}

// newAppliedClusterResourceQuotas returns a AppliedClusterResourceQuotas
func newAppliedClusterResourceQuotas(c *QuotaV1Client, namespace string) *appliedClusterResourceQuotas {
	return &appliedClusterResourceQuotas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appliedClusterResourceQuota, and returns the corresponding appliedClusterResourceQuota object, and an error if there is any.
func (c *appliedClusterResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *quotav1.AppliedClusterResourceQuota, err error) {
	result = &quotav1.AppliedClusterResourceQuota{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appliedclusterresourcequotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppliedClusterResourceQuotas that match those selectors.
func (c *appliedClusterResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *quotav1.AppliedClusterResourceQuotaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &quotav1.AppliedClusterResourceQuotaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appliedclusterresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}