// Copyright (c) 2020 Red Hat, Inc.

package multiclustermonitoring

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestNewVolumeClaimTemplate(t *testing.T) {
	vct := newVolumeClaimTemplate("1Gi")
	if vct.Spec.AccessModes[0] != v1.ReadWriteOnce ||
		vct.Spec.Resources.Requests[v1.ResourceStorage] != resource.MustParse(defaultStorageSize) {
		t.Errorf("Failed to newVolumeClaimTemplate")
	}
}

func TestNewDefaultObservatorium(t *testing.T) {
	obs := newDefaultObservatoriumSpec()
	if obs.Query.Image != defaultThanosImage ||
		obs.Query.Version != defaultThanosVersion {
		t.Errorf("Failed to newDefaultObservatorium")
	}
}