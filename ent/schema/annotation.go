package schema

type HandlerAnnotation struct {
	ReadEager []string
}

func (HandlerAnnotation) Name() string {
	return "HandlerGen"
}
