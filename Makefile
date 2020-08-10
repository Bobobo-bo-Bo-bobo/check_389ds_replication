GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = check_389ds_replication

depend:
	env GOPATH=$(GOPATH) go get -u gopkg.in/ldap.v3

build:
	env GOPATH=$(GOPATH) go install $(PROGRAMS)

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/lib64/nagios/plugins

strip: build
	strip --strip-all $(BINDIR)/check_389ds_replication

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/check_389ds_replication $(DESTDIR)/usr/lib64/nagios/plugins

clean:
	/bin/rm -f bin/check_389ds_replication

distclean: clean
	rm -rf src/gopkg.in/
	rm -rf src/github.com/
	rm -rf src/golang.org/

uninstall:
	/bin/rm -f $(DESTDIR)/usr/bin

all: depend build strip install

