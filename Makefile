default:
	go run main.go

gqlgen:
	cd graph && ${GOPATH}/bin/gqlgen -typemap typemap.json -schema ../schema.graphql

installdeps:
	go get -u github.com/vektah/gqlgen
	go get
	npm install -g prisma

prisma:
	prisma deploy

playground:
	prisma playground