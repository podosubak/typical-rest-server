package generated

// SourceRecipe is source code recipe for generated.go in typical package
type SourceRecipe struct {
	PackageName  string
	Imports      map[string]string
	Structs      []StructPogo
	Constructors []string
	MockTargets  []string
}

func (r SourceRecipe) String() string {
	var builder Builder
	builder.Printlnf("// Autogenerated by Typical-Go. DO NOT EDIT\n")
	builder.Printlnf("package %s", r.PackageName)

	for packageName := range r.Imports {
		builder.Printlnf(`import %s "%s"`, r.Imports[packageName], packageName)
	}

	for i := range r.Structs {
		builder.Printlnf("%s", r.Structs[i])
	}

	builder.Printlnf("func init() {")
	for i := range r.Constructors {
		builder.Printlnf("Context.AddConstructor(%s)", r.Constructors[i])
	}
	for i := range r.MockTargets {
		builder.Printlnf("Context.AddMockTarget(\"%s\")", r.MockTargets[i])
	}
	builder.Printlnf("}")

	return builder.String()
}

// AddConstructorPogos to add FunctionPogo to constructor
func (r *SourceRecipe) AddConstructorPogos(pogos ...FunctionPogo) {
	for _, pogo := range pogos {
		r.Constructors = append(r.Constructors, pogo.String())
	}
}

// AddConstructors to add constructors
func (r *SourceRecipe) AddConstructors(constructors ...string) {
	r.Constructors = append(r.Constructors, constructors...)
}

// AddMockTargets to add constructors
func (r *SourceRecipe) AddMockTargets(mockTargets ...string) {
	r.MockTargets = append(r.MockTargets, mockTargets...)
}
