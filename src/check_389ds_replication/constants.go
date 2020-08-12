package main

const name = "check_389ds_replication"
const version = "1.0.0-20200812"

const versionText = `%s version %s
Copyright (C) 2020 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

%s is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s

`

const defaultLDAPTimeout = 15

const helpText = `Usage: %s [--ca-cert=<file>] --domain<domain> [--help] [--insecure] --password=<pass>|--password-file=<file> --server=<uri> --user=<user> [--version]

  --ca-cert=<file>          Use <file> as CA certificate for SSL verification.

  --domain=<domain>         LDAP domain. This option is mandatory.

  --help                    This text.

  --insecure                Skip SSL verification of server certificate.

  --password=<pass>         Authenticate using password of <pass>.
  --password-file=<file>    Authenticate using password read from <file>
                            Password is mandatory, either from the command line or read from file.
                            --password and --password-file are mutually exclusive.

  --timeout=<sec>           LDAP connection and search timeout in seconds.
                            Default: %d

  --server=<uri>            URI of LDAP server. This option is mandatory.

  --user=<user>             Authenticate as <user>. This option is mandatory.

  --version                 Show version information.

`

const (
	// OK - Nagios exit code
	OK int = iota
	// WARNING - Nagios exit code
	WARNING
	// CRITICAL - Nagios exit code
	CRITICAL
	// UNKNOWN - Nagios exit code
	UNKNOWN
)

const nsds5TimeFormat = "2006-01-02T15:04:05Z"

var ldapSearchReplicaInfo = []string{
	"cn",
	"nsDS5ReplicaHost",
	"nsds5replicaLastUpdateStatusJSON",
}
