package schema

type HandlerAnnotation struct {
	CreateGroups []string
	ListGroups   []string
	ReadGroups   []string
	UpdateGroups []string

	ReadEager []string
	ListEager []string
}

func (HandlerAnnotation) Name() string {
	return "HandlerGen"
}
