
current_dir := "$(shell pwd)"
zip_file := "$(current_dir)/build/ansible_modules.zip"
base64_file := "$(current_dir)/build/ansible_modules.txt"

.PHONY: prepare_ansible
prepare_ansible:
	mkdir -p build 
	pushd build && git clone https://github.com/ansible/ansible.git
	cp -r resources/ansible/* build/ansible/lib/ansible/

.PHONY: zip_modules
zip_modules:
	pushd build/ansible/lib && zip -r "$(zip_file)" ansible/modules ansible/module_utils/ ansible/__init__.py && cat "$(zip_file)" | base64 > "$(base64_file)"

.PHONY: gen
gen:
	go run tools/main.go -m  build/ansible/lib/ansible/modules/ -o pkg/module/
	go fmt pkg/module/*.go