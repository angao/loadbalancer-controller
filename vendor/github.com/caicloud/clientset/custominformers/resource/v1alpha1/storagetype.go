/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	customclient "github.com/caicloud/clientset/customclient"
	internalinterfaces "github.com/caicloud/clientset/custominformers/internalinterfaces"
	v1alpha1 "github.com/caicloud/clientset/listers/resource/v1alpha1"
	resourcev1alpha1 "github.com/caicloud/clientset/pkg/apis/resource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageTypeInformer provides access to a shared informer and lister for
// StorageTypes.
type StorageTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorageTypeLister
}

type storageTypeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageTypeInformer constructs a new informer for StorageType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageTypeInformer(client customclient.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageTypeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageTypeInformer constructs a new informer for StorageType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageTypeInformer(client customclient.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().StorageTypes().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().StorageTypes().Watch(options)
			},
		},
		&resourcev1alpha1.StorageType{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageTypeInformer) defaultInformer(client customclient.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageTypeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageTypeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha1.StorageType{}, f.defaultInformer)
}

func (f *storageTypeInformer) Lister() v1alpha1.StorageTypeLister {
	return v1alpha1.NewStorageTypeLister(f.Informer().GetIndexer())
}