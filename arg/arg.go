package arg

type Args struct {
	Even  bool
	Limit int
}

var Arg = Args{
	Even:  true,
	Limit: 20,
}
