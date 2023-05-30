DESTDIR =
PREFIX  =/usr/local


all:
clean:
install:
update:
## -- AUTO-GO --
GO_PROGRAMS += bin/go-dialog$(EXE) 
.PHONY all-go: $(GO_PROGRAMS)
all:     all-go
install: install-go
clean:   clean-go
deps:
bin/go-dialog$(EXE): deps 
	go build -o $@ $(GO_DIALOG_FLAGS) $(GO_CONF) ./cmd/go-dialog
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp bin/go-dialog$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(GO_PROGRAMS)
## -- AUTO-GO --
