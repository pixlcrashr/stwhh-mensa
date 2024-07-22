// Package nullable
// Copyright © 2024 pixlcrashr (Vincent Heins)
package nullable

import (
	"encoding/json"
)

type Nullable[T any] struct {
	hasValue bool
	value    T
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	var v *T
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	n.hasValue = v != nil
	n.value = *v

	return nil
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	var v *T = nil

	if n.hasValue {
		*v = n.value
	}

	return json.Marshal(v)
}

func (n Nullable[T]) HasValue() bool {
	return n.hasValue
}

func (n Nullable[T]) Value() T {
	return n.value
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{
		hasValue: false,
	}
}

func Value[T any](value T) Nullable[T] {
	return Nullable[T]{
		hasValue: true,
		value:    value,
	}
}
