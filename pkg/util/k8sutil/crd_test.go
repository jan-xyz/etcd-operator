package k8sutil

import (
	"context"
	"testing"

	"github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestCreateCRD(t *testing.T) {
	cs := fake.NewSimpleClientset()

	err := CreateCRD(context.Background(), cs, v1beta2.EtcdBackupResourcePlural)
	if err != nil {
		t.Errorf("create CRD failed: %v", err)
	}

	actions := cs.Actions()

	actual_num_actions := len(actions)
	expected_num_actions := 1
	if actual_num_actions != expected_num_actions {
		t.Errorf("expect actions=%d, got=%d", expected_num_actions, actual_num_actions)
	}

	actual_resource := actions[0].GetResource()
	expected_resource := schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  "v1",
		Resource: "customresourcedefinitions",
	}
	if actual_resource != expected_resource {
		t.Errorf("expect action=%#v, got=%#v", expected_resource, actual_resource)
	}
	actual_verb := actions[0].GetVerb()
	expected_verb := "create"
	if actual_verb != expected_verb {
		t.Errorf("expect action=%s, got=%s", expected_verb, actual_verb)
	}
}
