package images

var imagelessGVKs = map[string]struct{}{
	// Core v1 resources
	"v1.Namespace":             {},
	"v1.ServiceAccount":        {},
	"v1.Service":               {},
	"v1.ConfigMap":             {},
	"v1.Secret":                {},
	"v1.PersistentVolumeClaim": {},
	"v1.PersistentVolume":      {},
	"v1.Endpoints":             {},
	"v1.LimitRange":            {},
	"v1.ResourceQuota":         {},
	"v1.Node":                  {},
	"v1.Binding":               {},
	"v1.ComponentStatus":       {},

	// Storage resources
	"storage.k8s.io/v1.StorageClass":          {},
	"storage.k8s.io/v1.CSIDriver":             {},
	"storage.k8s.io/v1.CSINode":               {},
	"storage.k8s.io/v1.CSIStorageCapacity":    {},
	"storage.k8s.io/v1.VolumeAttachment":      {},
	"storage.k8s.io/v1.VolumeAttributesClass": {},

	// API extensions
	"apiextensions.k8s.io/v1.CustomResourceDefinition": {},

	// Networking resources
	"networking.k8s.io/v1.Ingress":       {},
	"networking.k8s.io/v1.IngressClass":  {},
	"networking.k8s.io/v1.NetworkPolicy": {},
	"networking.k8s.io/v1.IPAddress":     {},
	"networking.k8s.io/v1.ServiceCIDR":   {},
	"discovery.k8s.io/v1.EndpointSlice":  {},

	// Autoscaling resources
	"autoscaling/v1.HorizontalPodAutoscaler": {},
	"autoscaling/v2.HorizontalPodAutoscaler": {},

	// RBAC resources
	"rbac.authorization.k8s.io/v1.ClusterRole":        {},
	"rbac.authorization.k8s.io/v1.ClusterRoleBinding": {},
	"rbac.authorization.k8s.io/v1.Role":               {},
	"rbac.authorization.k8s.io/v1.RoleBinding":        {},

	// Policy resources
	"policy/v1.PodDisruptionBudget": {},

	// Admission registration resources
	"admissionregistration.k8s.io/v1.MutatingWebhookConfiguration":        {},
	"admissionregistration.k8s.io/v1.ValidatingWebhookConfiguration":      {},
	"admissionregistration.k8s.io/v1.ValidatingAdmissionPolicy":           {},
	"admissionregistration.k8s.io/v1.ValidatingAdmissionPolicyBinding":    {},
	"admissionregistration.k8s.io/v1beta1.MutatingAdmissionPolicy":        {},
	"admissionregistration.k8s.io/v1beta1.MutatingAdmissionPolicyBinding": {},

	// Certificates resources
	"certificates.k8s.io/v1.CertificateSigningRequest":   {},
	"certificates.k8s.io/v1beta1.ClusterTrustBundle":     {},
	"certificates.k8s.io/v1alpha1.PodCertificateRequest": {},

	// Coordination resources
	"coordination.k8s.io/v1.Lease":               {},
	"coordination.k8s.io/v1beta1.LeaseCandidate": {},

	// Authorization resources
	"authorization.k8s.io/v1.LocalSubjectAccessReview": {},
	"authorization.k8s.io/v1.SelfSubjectAccessReview":  {},
	"authorization.k8s.io/v1.SelfSubjectRulesReview":   {},
	"authorization.k8s.io/v1.SubjectAccessReview":      {},

	// Authentication resources
	"authentication.k8s.io/v1.TokenRequest":      {},
	"authentication.k8s.io/v1.TokenReview":       {},
	"authentication.k8s.io/v1.SelfSubjectReview": {},

	// Scheduling resources
	"scheduling.k8s.io/v1.PriorityClass": {},

	// Resource management
	"resource.k8s.io/v1.DeviceClass":           {},
	"resource.k8s.io/v1.ResourceClaim":         {},
	"resource.k8s.io/v1.ResourceClaimTemplate": {},
	"resource.k8s.io/v1.ResourceSlice":         {},
	"resource.k8s.io/v1alpha3.DeviceTaintRule": {},

	// Flow control resources
	"flowcontrol.apiserver.k8s.io/v1.FlowSchema":                 {},
	"flowcontrol.apiserver.k8s.io/v1.PriorityLevelConfiguration": {},

	// API registration resources
	"apiregistration.k8s.io/v1.APIService": {},

	// Events
	"events.k8s.io/v1.Event": {},

	// Apps resources (metadata only)
	"apps/v1.ControllerRevision": {},

	// Node resources
	"node.k8s.io/v1.RuntimeClass": {},

	// Storage migration resources
	"storagemigration.k8s.io/v1alpha1.StorageVersionMigration": {},
	"internal.apiserver.k8s.io/v1alpha1.StorageVersion":        {},

	// Prometheus/monitoring.coreos.com resources
	"monitoring.coreos.com/v1.PodMonitor":               {},
	"monitoring.coreos.com/v1.PrometheusRule":           {},
	"monitoring.coreos.com/v1.ServiceMonitor":           {},
	"monitoring.coreos.com/v1alpha1.AlertmanagerConfig": {},

	// Strimzi Kafka resources
	"kafka.strimzi.io/v1beta2.KafkaNodePool": {},

	// Tekton resources
	"tekton.dev/v1beta1.TaskRun":                        {},
	"triggers.tekton.dev/v1beta1.TriggerBinding":        {},
	"triggers.tekton.dev/v1beta1.ClusterTriggerBinding": {},
	"triggers.tekton.dev/v1alpha1.ClusterInterceptor":   {},

	// Kyverno resources
	"kyverno.io/v1.ClusterPolicy": {},
	"kyverno.io/v1.Policy":        {},

	// cert-manager resources
	"cert-manager.io/v1.Certificate":   {},
	"cert-manager.io/v1.ClusterIssuer": {},
	"cert-manager.io/v1.Issuer":        {},

	// KServe resources (imageless ones)
	"serving.kserve.io/v1alpha1.InferenceGraph":      {},
	"serving.kserve.io/v1alpha1.TrainedModel":        {},
	"serving.kserve.io/v1alpha1.LocalModelCache":     {},
	"serving.kserve.io/v1alpha1.LocalModelNode":      {},
	"serving.kserve.io/v1alpha1.LocalModelNodeGroup": {},
}
