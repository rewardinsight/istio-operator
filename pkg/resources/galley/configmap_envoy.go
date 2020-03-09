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

package galley

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/banzaicloud/istio-operator/pkg/resources/templates"
)

func (r *Reconciler) configMapEnvoy() runtime.Object {
	return &apiv1.ConfigMap{
		ObjectMeta: templates.ObjectMeta(configMapNameEnvoy, nil, r.Config),
		Data: map[string]string{
			"envoy.yaml.tmpl": r.envoyConfig(),
		},
	}
}

func (r *Reconciler) envoyConfig() string {
	return `
admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 15000
static_resources:
  clusters:
  - name: in.9901
    http2_protocol_options: {}
    connect_timeout: 1.000s
    hosts:
    - socket_address:
        address: 127.0.0.1
        port_value: 9901
    circuit_breakers:
      thresholds:
      - max_connections: 100000
        max_pending_requests: 100000
        max_requests: 100000
        max_retries: 3
  listeners:
  - name: "15019"
    address:
      socket_address:
        address: 0.0.0.0
        port_value: 15019
    filter_chains:
    - filters:
      - name: envoy.http_connection_manager
        config:
          codec_type: HTTP2
          stat_prefix: "15010"
          stream_idle_timeout: 0s
          http2_protocol_options:
            max_concurrent_streams: 1073741824
          access_log:
          - name: envoy.file_access_log
            config:
              path: /dev/stdout
          http_filters:
          - name: envoy.router
          route_config:
            name: "15019"
            virtual_hosts:
            - name: istio-galley
              domains:
              - '*'
              routes:
              - match:
                  prefix: /
                route:
                  cluster: in.9901
                  timeout: 0.000s
      tls_context:
        common_tls_context:
          alpn_protocols:
          - h2
          tls_certificates:
          - certificate_chain:
              filename: /etc/certs/cert-chain.pem
            private_key:
              filename: /etc/certs/key.pem
          validation_context:
            trusted_ca:
              filename: /etc/certs/root-cert.pem
        require_client_certificate: true
`
}
