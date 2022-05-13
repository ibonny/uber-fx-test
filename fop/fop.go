package fop

type FopStruct struct {
	optionOne string
	optionTwo string
}

func NewFopStruct(options ...func(*FopStruct)) *FopStruct {
	fs := &FopStruct{"default_one", "default_two"}

	for _, o := range options {
		o(fs)
	}

	return fs
}

func WithOptionOne(a string) func(*FopStruct) {
	return func(fs *FopStruct) {
		fs.optionOne = a
	}
}

func WithOptionTwo(a string) func(*FopStruct) {
	return func(fs *FopStruct) {
		fs.optionTwo = a
	}
}
