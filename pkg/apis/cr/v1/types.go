/*
Copyright 2017 The Kubernetes Authors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/jgensler8/openfaas-client-go"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Example is a specification for an Example resource
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   FunctionSpec   `json:"spec"`
	Status FunctionStatus `json:"status,omitempty"`
}

// ExampleSpec is the spec for an Example resource
type FunctionSpec struct {
	swagger.CreateFunctionRequest
}

// ExampleStatus is the status for an Example resource
type FunctionStatus struct {
	State   FunctionState `json:"state,omitempty"`
	Message string       `json:"message,omitempty"`
}

type FunctionState string

const (
	FunctionStateCreated FunctionState = "Created"
	FunctionStateProcessed FunctionState = "Processed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExampleList is a list of Example resources
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Function `json:"items"`
}
