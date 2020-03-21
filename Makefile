deps:
	go mod download
	GO111MODULE=off go get -u github.com/myitcv/gobin

generate: deps
	gobin -m -run github.com/go-swagger/go-swagger/cmd/swagger \
	generate client \
	-f swagger.json \
	-A netbox \
	-c plumbing \
	--copyright-file=copyright_header.txt \
	--with-flatten=full \
	--default-scheme=https

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration
