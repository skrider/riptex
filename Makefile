RIPTEX_PKG = github.com/skrider/riptex

.PHONY: cp
cp: ; echo -n $(RIPTEX_PKG) | xclip -selection c

.PHONY: test-buffer
test-buffer: ;
	go test $(RIPTEX_PKG)/buffer

.PHONY: debug-buffer
debug-buffer: ;
	dlv test $(RIPTEX_PKG)/buffer

.PHONY: run
run: ;
	go run riptex.go