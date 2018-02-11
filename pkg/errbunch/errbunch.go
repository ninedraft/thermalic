package errbunch

import (
	"bytes"
)

type Err string

func (err Err) Error() string {
	return string(err)
}

func (err Err) Wrap(errs ...error) *Wrapper {
	return &Wrapper{
		top:   err,
		chain: errs,
	}
}

type Wrapper struct {
	top          error
	cachedErrMsg string
	chain        []error
}

func (wrapper *Wrapper) Error() string {
	const delim = "\n"
	if wrapper.cachedErrMsg == "" {
		buf := bytes.NewBufferString(wrapper.top.Error())
		if len(wrapper.chain) > 0 {
			buf.WriteString(":" + delim)
			for _, err := range wrapper.chain {
				buf.WriteString(delim + err.Error())
			}
			wrapper.cachedErrMsg = buf.String()
		}
	}
	return wrapper.cachedErrMsg
}

func (wrapper Wrapper) Top() error {
	return wrapper.top
}

func (wrapper Wrapper) ErrChain() []error {
	return wrapper.chain
}
