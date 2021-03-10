package v1beta1

import (
	"github.com/example/memcached-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this Memcached to the Hub version (v1).
func (src *Memcached) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha1.Memcached)
	size := src.Spec.ReplicaSize
	dst.Spec.Size = size
	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Memcached) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha1.Memcached)
	size := src.Spec.Size
	dst.Spec.ReplicaSize = size
	return nil
}
