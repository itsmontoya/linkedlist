package typed

//go:generate go install github.com/joeshaw/gengen

//go:generate gengen -o int github.com/itsmontoya/linkedlist int
//go:generate gengen -o int32 github.com/itsmontoya/linkedlist int32
//go:generate gengen -o int64 github.com/itsmontoya/linkedlist int64
//go:generate gengen -o string github.com/itsmontoya/linkedlist string
//go:generate gengen -o byteslice github.com/itsmontoya/linkedlist []byte
