// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var flagExpose = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--cluster-ip", Description: "ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service."},
	prompt.Suggest{Text: "--container-port", Description: "Synonym for --target-port"},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "--external-ip", Description: "Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to expose a service"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to expose a service"},
	prompt.Suggest{Text: "--generator", Description: "The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'."},
	prompt.Suggest{Text: "-l", Description: "Labels to apply to the service created by this call."},
	prompt.Suggest{Text: "--labels", Description: "Labels to apply to the service created by this call."},
	prompt.Suggest{Text: "--load-balancer-ip", Description: "IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used (cloud-provider specific)."},
	prompt.Suggest{Text: "--name", Description: "The name for the newly created object."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--overrides", Description: "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field."},
	prompt.Suggest{Text: "--port", Description: "The port that the service should serve on. Copied from the resource being exposed, if unspecified"},
	prompt.Suggest{Text: "--protocol", Description: "The network protocol for the service to be created. Default is 'TCP'."},
	prompt.Suggest{Text: "--record", Description: "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future."},
	prompt.Suggest{Text: "--selector", Description: "A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.)"},
	prompt.Suggest{Text: "--session-affinity", Description: "If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--target-port", Description: "Name or number for the port on the container that the service should direct traffic to. Optional."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--type", Description: "Type for this service: ClusterIP, NodePort, LoadBalancer, or ExternalName. Default is 'ClusterIP'."},
}
