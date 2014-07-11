# This bash file can be used to compile gort using gccgo.
# This is usually quite slower than the default go compiler.
go build -gccgoflags "-Ofast" -compiler gccgo src/main.go
