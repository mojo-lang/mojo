package hello_world

const HandlerMethod = `
func (s helloWorld) Echo(ctx context.Context, in *pb.EchoRequest) (*hello_world.Echo, error){
	resp := hello_world.Echo{
		Name: in.Name,
		Message: "Hello, " + in.Name + "!"
	}
	return &resp, nil
}
`
