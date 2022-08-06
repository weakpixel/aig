
current_dir := "$(shell pwd)"
zip_file := "$(current_dir)/build/ansible_modules.zip"
base64_file := "$(current_dir)/build/ansible_modules.zip.base64"
ansible_tag := "v2.13.1"

.PHONY: prepare_ansible
prepare_ansible:
	mkdir -p build 
	pushd build && git clone https://github.com/ansible/ansible.git
	pushd build/ansible && git checkout $(ansible_tag)
	cp -r resources/ansible/* build/ansible/lib/ansible/

.PHONY: zip_modules
zip_modules:
	pushd build/ansible/lib && zip -r "$(zip_file)" ansible/modules ansible/module_utils/ ansible/__init__.py && cat "$(zip_file)" | base64 > "$(base64_file)"

.PHONY: clean
clean:
	rm -fr ./build/

.PHONY: gen
gen:
	rm -fr pkg/module/*
	go run tools/generator.go -m  build/ansible/lib/ansible/modules/ -o pkg/module/ -V $(ansible_tag)
	go fmt pkg/module/*.go

build: gen
	go build -o build/app main.go
