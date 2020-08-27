package deploy

import (
	coremodel "sigs.k8s.io/aws-alb-ingress-controller/pkg/model/core"
)

// StackSchema represents the JSON model for stack.
type StackSchema struct {
	// all resources within stack.
	Resources map[string]map[string]interface{} `json:"resources"`
}

// NewStackSchemaBuilder constructs new stackSchemaBuilder.
func NewStackSchemaBuilder() *stackSchemaBuilder {
	return &stackSchemaBuilder{
		resources: make(map[string]map[string]interface{}),
	}
}

var _ coremodel.ResourceVisitor = &stackSchemaBuilder{}

type stackSchemaBuilder struct {
	resources map[string]map[string]interface{}
}

// Visit will visit a resource.
func (b *stackSchemaBuilder) Visit(res coremodel.Resource) error {
	if _, ok := b.resources[res.Type()]; !ok {
		b.resources[res.Type()] = make(map[string]interface{})
	}
	b.resources[res.Type()][res.ID()] = res
	return nil
}

// Build will build StackSchema based on resources visited.
func (b *stackSchemaBuilder) Build() StackSchema {
	return StackSchema{
		Resources: b.resources,
	}
}