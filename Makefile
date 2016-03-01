vendor:
	glide install

install: install-pigeon install-pigeon-app
install-pigeon:
	go install github.com/kaneshin/pigeon/tools/cmd/pigeon
install-pigeon-app:
	go install github.com/kaneshin/pigeon/tools/cmd/pigeon-app

run-pigeon: install-pigeon
	lime -bin=/tmp/pigeon-bin -ignore-pattern=\(\\.git\|vendor\) -build-pattern=.* ./tools/cmd/pigeon
run-pigeon-app: install-pigeon-app
	lime -i -bin=/tmp/pigeon-bin -ignore-pattern=\(\\.git\|vendor\) ./tools/cmd/pigeon-app -port=8080 -- -label
