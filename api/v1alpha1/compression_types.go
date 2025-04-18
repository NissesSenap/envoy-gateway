// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

// CompressorType defines the types of compressor library supported by Envoy Gateway.
//
// +kubebuilder:validation:Enum=Gzip;Brotli
type CompressorType string

const (
	GzipCompressorType CompressorType = "Gzip"

	BrotliCompressorType CompressorType = "Brotli"
)

// GzipCompressor defines the config for the Gzip compressor.
// The default values can be found here:
// https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/compression/gzip/compressor/v3/gzip.proto#extension-envoy-compression-gzip-compressor
type GzipCompressor struct{}

// BrotliCompressor defines the config for the Brotli compressor.
// The default values can be found here:
// https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/compression/brotli/compressor/v3/brotli.proto#extension-envoy-compression-brotli-compressor
type BrotliCompressor struct{}

// Compression defines the config of enabling compression.
// This can help reduce the bandwidth at the expense of higher CPU.
type Compression struct {
	// CompressorType defines the compressor type to use for compression.
	//
	// +required
	Type CompressorType `json:"type"`

	// The configuration for Brotli compressor.
	//
	// +optional
	Brotli *BrotliCompressor `json:"brotli,omitempty"`

	// The configuration for GZIP compressor.
	//
	// +optional
	Gzip *GzipCompressor `json:"gzip,omitempty"`
}
