GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = check_389ds_replication

depend:
	# go mod will handle dependencies

build:
	cd $(CURDIR)/src/check_389ds_replication && go get check_389ds_replication/src/check_389ds_replication && go build -o $(CURDIR)/bin/check_389ds_replication

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/lib64/nagios/plugins

strip: build
	strip --strip-all $(BINDIR)/check_389ds_replication

ifneq (, $(shell which upx 2>/dev/null))
	upx -9 $(BINDIR)/check_389ds_replication
endif

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
	/bin/rm -f $(DESTDIR)/usr/lib64/nagios/plugins/check_389ds_replication

all: depend build strip install

