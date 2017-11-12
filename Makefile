NAME=user

.PHONY: install
install:
	mkdir -p bin
	go build -o bin/$(NAME)
	cf install-plugin -f bin/$(NAME)

.PHONY: uninstall
uninstall:
	cf uninstall-plugin $(NAME)

.PHONY: reinstall
reinstall: uninstall install

.PHONY: clean
clean:
	rm -rf bin
