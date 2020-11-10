package discovery

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes/scheme"
)

// NewCacheFallbackDiscoveryClient creates a new discovery client which wraps delegate client
// and uses hardcoded discovery information in case the delegate client fails.
func NewCacheFallbackDiscoveryClient(log *logrus.Entry, delegate discovery.DiscoveryInterface) discovery.DiscoveryInterface {
	return &cacheFallbackDiscoveryClient{
		DiscoveryInterface: delegate,
		log:                log,
		asset:              Asset,
	}
}

type cacheFallbackDiscoveryClient struct {
	discovery.DiscoveryInterface
	log   *logrus.Entry
	asset func(name string) ([]byte, error)
}

// ServerResourcesForGroupVersion returns the supported resources for a group and version.
func (d *cacheFallbackDiscoveryClient) ServerResourcesForGroupVersion(groupVersion string) (*metav1.APIResourceList, error) {
	liveResources, err := d.DiscoveryInterface.ServerResourcesForGroupVersion(groupVersion)
	if err != nil {
		filename := filepath.Join(groupVersion, "serverresources.json")
		cachedResources := &metav1.APIResourceList{}
		cacheErr := d.getCached(filename, cachedResources)
		if cacheErr == nil {
			return cachedResources, nil
		}
		d.log.Warnf("discovery cache failed. Fallback to live dynamic client. Error: %s", cacheErr)
	}

	return liveResources, err
}

// ServerResources returns the supported resources for all groups and versions.
// Deprecated: use ServerGroupsAndResources instead.
func (d *cacheFallbackDiscoveryClient) ServerResources() ([]*metav1.APIResourceList, error) {
	_, rs, err := discovery.ServerGroupsAndResources(d)
	return rs, err
}

// ServerGroupsAndResources returns the supported groups and resources for all groups and versions.
func (d *cacheFallbackDiscoveryClient) ServerGroupsAndResources() ([]*metav1.APIGroup, []*metav1.APIResourceList, error) {
	return discovery.ServerGroupsAndResources(d)
}

// ServerGroups returns the supported groups, with information like supported versions and the
// preferred version.
func (d *cacheFallbackDiscoveryClient) ServerGroups() (*metav1.APIGroupList, error) {
	liveGroups, err := d.DiscoveryInterface.ServerGroups()
	if err != nil {
		cachedGroups := &metav1.APIGroupList{}
		cacheErr := d.getCached("servergroups.json", cachedGroups)
		if cacheErr == nil {
			return cachedGroups, nil
		}
		d.log.Info(cacheErr)
	}

	return liveGroups, err
}

// ServerPreferredResources returns the supported resources with the version preferred by the
// server.
func (d *cacheFallbackDiscoveryClient) ServerPreferredResources() ([]*metav1.APIResourceList, error) {
	return discovery.ServerPreferredResources(d)
}

// ServerPreferredNamespacedResources returns the supported namespaced resources with the
// version preferred by the server.
func (d *cacheFallbackDiscoveryClient) ServerPreferredNamespacedResources() ([]*metav1.APIResourceList, error) {
	return discovery.ServerPreferredNamespacedResources(d)
}

func (d *cacheFallbackDiscoveryClient) getCached(filename string, out runtime.Object) error {
	cachedBytes, err := d.asset(filename)
	if err != nil {
		return err
	}

	return runtime.DecodeInto(scheme.Codecs.UniversalDecoder(), cachedBytes, out)
}
