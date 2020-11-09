package schema

type (
	// Used on a schema to pass options to the handler generator.
	HandlerAnnotation struct {
		SkipGeneration bool
		CreateGroups   []string
		ListGroups     []string
		ReadGroups     []string
		UpdateGroups   []string
	}
	// Used on fields pass options to the handler generator.
	FieldAnnotation struct {
		Create              bool
		CreateValidationTag string
		Patch               bool
		PatchValidationTag  string
	}
)

func (HandlerAnnotation) Name() string {
	return "HandlerGen"
}

func (FieldAnnotation) Name() string {
	return "FieldGen"
}
