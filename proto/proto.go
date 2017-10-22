package proto

type Args struct {
	PWD string
}

type Reply struct {
	T1,T2 int64
}

type Clock interface {
	Sync(args *Args, reply *Reply) error
}
