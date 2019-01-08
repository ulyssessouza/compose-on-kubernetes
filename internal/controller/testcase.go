package controller

import (
	latest "github.com/docker/compose-on-kubernetes/api/compose/v1alpha3"
	"github.com/docker/compose-on-kubernetes/internal/stackresources"
)

// TestCase is a serializable type used to combine a stack and its children for a record & replay test scenario
type TestCase struct {
	Stack    *latest.Stack
	Children *stackresources.StackState
}
