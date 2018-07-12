default:
	go run main.go

gqlgen:
	${GOPATH}/bin/gqlgen

installdeps:
	go get -u github.com/vektah/gqlgen
	go get
