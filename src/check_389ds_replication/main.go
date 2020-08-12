package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var _showVersion = flag.Bool("version", false, "Show version")
	var _showHelp = flag.Bool("help", false, "Show help text")
	var server = flag.String("server", "", "LDAP server")
	var domain = flag.String("domain", "", "LDAP domain")
	var user = flag.String("user", "", "LDAP user")
	var password = flag.String("password", "", "LDAP password")
	var passwordfile = flag.String("password-file", "", "LDAP password file")
	var insecure = flag.Bool("insecure", false, "Skip SSL verification")
	var cacert = flag.String("ca-cert", "", "CA certificate for SSL")
	var timeout = flag.Uint("timeout", defaultLDAPTimeout, "Connect and LDAP timeout")
	var err error

	flag.Usage = showUsage
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "Error: Too many arguments. Use --help to see a list of all parameters.")
		os.Exit(UNKNOWN)
	}

	if *_showHelp {
		showUsage()
		os.Exit(OK)
	}

	if *_showVersion {
		showVersion()
		os.Exit(OK)
	}

	if *domain == "" {
		fmt.Fprintln(os.Stderr, "Error: Missing mandatory parameter for LDAP domain. Use --help to see a list of all parameters.")
		os.Exit(UNKNOWN)
	}
	if *server == "" {
		fmt.Fprintln(os.Stderr, "Error: Missing mandatory parameter for LDAP server URI. Use --help to see a list of all parameters.")
		os.Exit(UNKNOWN)
	}
	if *user == "" {
		fmt.Fprintln(os.Stderr, "Error: Missing mandatory parameter for LDAP user. Use --help to see a list of all parameters.")
		os.Exit(UNKNOWN)
	}

	if *password != "" && *passwordfile != "" {
		fmt.Fprintln(os.Stderr, "Error: --password and --password-file are mutually.")
		os.Exit(UNKNOWN)
	}
	if *password == "" && *passwordfile == "" {
		fmt.Fprintln(os.Stderr, "Error: A password must be provided for LDAP authentication. Use --help to see a list of all parameters.")
		os.Exit(UNKNOWN)
	}

	if *passwordfile != "" {
		*password, err = readSingleLine(*passwordfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't read password from file: %s\n", err.Error())
			os.Exit(UNKNOWN)
		}
	}
	// sanity checks
	if *timeout == 0 {
		fmt.Fprintf(os.Stderr, "Error: Timout limit must be greater than zero")
		os.Exit(UNKNOWN)
	}

	con, err := connectAndBind(*server, *insecure, *cacert, time.Duration(*timeout)*time.Second, *user, *password)
	if err != nil {
		fmt.Println("CRITICAL - " + err.Error())
		os.Exit(CRITICAL)
	}
	defer con.Close()

	rs, err := getReplicationStatus(con, *domain, int(*timeout))
	if err != nil {
		fmt.Println("CRITICAL - " + err.Error())
		os.Exit(CRITICAL)
	}

	status := parseStatus(rs)

	if len(status.Unknown) > 0 {
		fmt.Printf("UNKNOWN - %s; %s; %s; %s\n", strings.Join(status.Unknown, ", "), strings.Join(status.Critical, ", "), strings.Join(status.Warning, ", "), strings.Join(status.Ok, ", "))
		os.Exit(UNKNOWN)
	}
	if len(status.Critical) > 0 {
		fmt.Printf("CRITICAL - %s; %s; %s\n", strings.Join(status.Critical, ", "), strings.Join(status.Warning, ", "), strings.Join(status.Ok, ", "))
		os.Exit(CRITICAL)
	}
	if len(status.Warning) > 0 {
		fmt.Printf("WARNING - %s; %s\n", strings.Join(status.Warning, ", "), strings.Join(status.Ok, ", "))
		os.Exit(WARNING)
	}
	if len(status.Ok) > 0 {
		fmt.Printf("OK - %s\n", strings.Join(status.Ok, ", "))
		os.Exit(OK)
	}

	fmt.Printf("UNKNOWN - No data for LDAP replication agreements found on this server\n")
	os.Exit(UNKNOWN)
}
