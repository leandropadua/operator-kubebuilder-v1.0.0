/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	cluster_v1 "github.com/feloy/operator/pkg/apis/cluster/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCdnClusters implements CdnClusterInterface
type FakeCdnClusters struct {
	Fake *FakeClusterV1
	ns   string
}

var cdnclustersResource = schema.GroupVersionResource{Group: "cluster.anevia.com", Version: "v1", Resource: "cdnclusters"}

var cdnclustersKind = schema.GroupVersionKind{Group: "cluster.anevia.com", Version: "v1", Kind: "CdnCluster"}

// Get takes name of the cdnCluster, and returns the corresponding cdnCluster object, and an error if there is any.
func (c *FakeCdnClusters) Get(name string, options v1.GetOptions) (result *cluster_v1.CdnCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cdnclustersResource, c.ns, name), &cluster_v1.CdnCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cluster_v1.CdnCluster), err
}

// List takes label and field selectors, and returns the list of CdnClusters that match those selectors.
func (c *FakeCdnClusters) List(opts v1.ListOptions) (result *cluster_v1.CdnClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cdnclustersResource, cdnclustersKind, c.ns, opts), &cluster_v1.CdnClusterList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cluster_v1.CdnClusterList), err
}

// Watch returns a watch.Interface that watches the requested cdnClusters.
func (c *FakeCdnClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cdnclustersResource, c.ns, opts))

}

// Create takes the representation of a cdnCluster and creates it.  Returns the server's representation of the cdnCluster, and an error, if there is any.
func (c *FakeCdnClusters) Create(cdnCluster *cluster_v1.CdnCluster) (result *cluster_v1.CdnCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cdnclustersResource, c.ns, cdnCluster), &cluster_v1.CdnCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cluster_v1.CdnCluster), err
}

// Update takes the representation of a cdnCluster and updates it. Returns the server's representation of the cdnCluster, and an error, if there is any.
func (c *FakeCdnClusters) Update(cdnCluster *cluster_v1.CdnCluster) (result *cluster_v1.CdnCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cdnclustersResource, c.ns, cdnCluster), &cluster_v1.CdnCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cluster_v1.CdnCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCdnClusters) UpdateStatus(cdnCluster *cluster_v1.CdnCluster) (*cluster_v1.CdnCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cdnclustersResource, "status", c.ns, cdnCluster), &cluster_v1.CdnCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cluster_v1.CdnCluster), err
}

// Delete takes name of the cdnCluster and deletes it. Returns an error if one occurs.
func (c *FakeCdnClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cdnclustersResource, c.ns, name), &cluster_v1.CdnCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCdnClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cdnclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cluster_v1.CdnClusterList{})
	return err
}

// Patch applies the patch and returns the patched cdnCluster.
func (c *FakeCdnClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cluster_v1.CdnCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cdnclustersResource, c.ns, name, data, subresources...), &cluster_v1.CdnCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cluster_v1.CdnCluster), err
}
