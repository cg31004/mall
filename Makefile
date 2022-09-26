.PHONY: default
default:
	echo 'please input command'

.PHONY: setup
setup:
	sh ./script/setup_local.sh

.PHONY: clean
clean:
	sh ./script/clean_local.sh

.PHONY: mock
mock:
	sh ./script/setup_mock.sh