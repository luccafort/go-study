package adapter

import "fmt"

type Decorator interface {
	Decorate() string
}

type Banner struct {
	body string
}

func (b *Banner) Body() string {
	return "*" + b.body + "*"
}

type EmbeddedDecoratorBanner struct {
	*Banner
}

var _ = EmbeddedDecoratorBanner{&Banner{""}}

func NewEmbeddedDecorateBanner(body string) *EmbeddedDecoratorBanner {
	return &EmbeddedDecoratorBanner{&Banner{body: body}}
}

func (ebd *EmbeddedDecoratorBanner) Decorate() string {
	return ebd.Body()
}

func main() {
	var decorator Decorator

	value := "ABC"
	expected := "*ABC*"
	decorator = NewEmbeddedDecorateBanner(value)

	if ret := decorator.Decorate(); ret != expected {
		fmt.Errorf("Expect decorated value to %s, but %s", expected, ret)
		return
	}
}
