package safe

func (b Block) Do() error {
	err := make(chan error)
	go func() {
		if b.Finally != nil {
			defer b.Finally()
		}
		if b.Catch != nil {
			defer func() {
				if r := recover(); r != nil {
					err <- b.Catch(r)

				}
			}()
		}

		if b.Try != nil {
			if e := b.Try(); e != nil {
				err <- e
			}
		}
		err <- nil
	}()

	return <-err
}

func Throw(e Exception) {
	panic(e)
}

type Block struct {
	Try     func() error
	Catch   func(Exception) error
	Finally func() error
}

type Exception interface{}
