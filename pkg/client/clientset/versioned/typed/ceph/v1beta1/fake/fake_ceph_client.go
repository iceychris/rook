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

package fake

import (
	v1beta1 "github.com/rook/rook/pkg/client/clientset/versioned/typed/ceph/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCephV1beta1 struct {
	*testing.Fake
}

func (c *FakeCephV1beta1) Clusters(namespace string) v1beta1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeCephV1beta1) Filesystems(namespace string) v1beta1.FilesystemInterface {
	return &FakeFilesystems{c, namespace}
}

func (c *FakeCephV1beta1) ObjectStores(namespace string) v1beta1.ObjectStoreInterface {
	return &FakeObjectStores{c, namespace}
}

func (c *FakeCephV1beta1) Pools(namespace string) v1beta1.PoolInterface {
	return &FakePools{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCephV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}