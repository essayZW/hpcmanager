.PHONY: clean
clean:
	rm -rf ./build
.PHONY: fss
fss:
	cp ../fss/build/main ./build/go/
.PHONY: docker
docker:
	sudo docker build . -t hpcmanager/proxy:1.0
