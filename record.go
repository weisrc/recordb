package recordb

type Record struct {
	Type  uint16
	Name  string
	Value string
	Wild  bool
	Ttl   uint
	Iat   uint
	Next  *Record
	Prev  *Record
}
