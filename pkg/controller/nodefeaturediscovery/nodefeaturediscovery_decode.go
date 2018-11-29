package nodefeaturediscovery

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"

	securityv1 "github.com/openshift/api/security/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
)

//var nfdNameSpace *corev1.Namespace
var nfdServiceAccount *corev1.ServiceAccount
var nfdClusterRole *rbacv1.ClusterRole
var nfdClusterRoleBinding *rbacv1.ClusterRoleBinding
var nfdSecurityContextConstraint *securityv1.SecurityContextConstraints
var nfdDaemonSet *appsv1.DaemonSet

func decodeManifest(yaml string) interface{} {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode([]byte(yaml), nil, nil)
	if err != nil {
		log.Printf("Error decoding manifest %v\n", err)
		return nil
	}
	return obj
}

func init() {
	nfdServiceAccount = decodeManifest(nfdserviceaccount).(*corev1.ServiceAccount)
	nfdClusterRole = decodeManifest(nfdclusterrole).(*rbacv1.ClusterRole)
	nfdClusterRoleBinding = decodeManifest(nfdclusterrolebinding).(*rbacv1.ClusterRoleBinding)
	nfdSecurityContextConstraint =  decodeManifest(nfdsecuritycontextconstraint).(*securityv1.SecurityContextConstraints)
	nfdDaemonSet = decodeManifest(nfddaemonset).(*appsv1.DaemonSet)
}