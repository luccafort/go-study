package adapter

import (
	"testing"
)

func TestAdapterByEmbedded(t *testing.T) {
	var decorator Decorator

	value := "ABC"
	expected := "*ABC*"
	decorator = NewEmbeddedDecorateBanner(value)

	if ret := decorator.Decorate(); ret != expected {
		t.Errorf("Expect decorated value to %s, but %s", expected, ret)
	}
}
