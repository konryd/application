package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: "app.k8s.io", Version: "v1alpha1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Application{},
		&ApplicationList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// CRD Generation
func getFloat(f float64) *float64 {
	return &f
}

var (
	// Define CRDs for resources
	ApplicationCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "applications.app.k8s.io",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "app.k8s.io",
			Version: "v1alpha1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "Application",
				Plural: "applications",
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"componentKinds": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"group": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"kind": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
										},
									},
								},
								"description": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"info": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"name": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"value": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
										},
									},
								},
								"keywords": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "string",
										},
									},
								},
								"links": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"description": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"url": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
										},
									},
								},
								"maintainers": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"email": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"name": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"url": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
										},
									},
								},
								"notes": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"owners": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "string",
										},
									},
								},
								"selector": v1beta1.JSONSchemaProps{
									Type: "object",
									Properties: map[string]v1beta1.JSONSchemaProps{
										"matchExpressions": v1beta1.JSONSchemaProps{
											Type: "array",
											Items: &v1beta1.JSONSchemaPropsOrArray{
												Schema: &v1beta1.JSONSchemaProps{
													Type: "object",
													Properties: map[string]v1beta1.JSONSchemaProps{
														"key": v1beta1.JSONSchemaProps{
															Type: "string",
														},
														"operator": v1beta1.JSONSchemaProps{
															Type: "string",
														},
														"values": v1beta1.JSONSchemaProps{
															Type: "array",
															Items: &v1beta1.JSONSchemaPropsOrArray{
																Schema: &v1beta1.JSONSchemaProps{
																	Type: "string",
																},
															},
														},
													},
												},
											},
										},
										"matchLabels": v1beta1.JSONSchemaProps{
											Type: "object",
											AdditionalProperties: &v1beta1.JSONSchemaPropsOrBool{
												Allows: true,
												//Schema: &,
											},
										},
									},
								},
								"type": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"version": v1beta1.JSONSchemaProps{
									Type: "string",
								},
							},
						},
						"status": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"observedGeneration": v1beta1.JSONSchemaProps{
									Type:   "integer",
									Format: "int64",
								},
							},
						},
					},
				},
			},
		},
	}
)
