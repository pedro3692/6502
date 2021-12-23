package register

type StrackPointer struct {
	data byte
}

func (sp *StrackPointer) Reset() {
	sp.data = 0xff
}

func (sp *StrackPointer) Inc() {
	if sp.data == 0x00 {
		sp.data = 0xff

		return
	}

	sp.data--
}

func (sp *StrackPointer) Dec() {
	if sp.data == 0xff {
		sp.data = 0x00

		return
	}

	sp.data++
}

func (r StrackPointer) Read() byte {
	return r.data
}
