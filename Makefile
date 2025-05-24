.PHONY: compile-frontend build-project run-dev

compile-frontend:
	cd frontend && npm install element-plus --save

build-project:
	wails build

run-dev:
	wails dev