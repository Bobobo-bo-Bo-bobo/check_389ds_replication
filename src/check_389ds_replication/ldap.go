package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"

	ldap "github.com/go-ldap/ldap/v3"
)

func connectAndBind(u string, i bool, ca string, t time.Duration, us string, p string) (*ldap.Conn, error) {
	var c *ldap.Conn
	var err error

	dc := &net.Dialer{Timeout: t}
	tlscfg := &tls.Config{}

	if i {
		tlscfg.InsecureSkipVerify = true
	}

	if ca != "" {
		tlscfg.RootCAs = x509.NewCertPool()
		cadata, err := ioutil.ReadFile(ca)
		if err != nil {
			return c, err
		}
		tlsok := tlscfg.RootCAs.AppendCertsFromPEM(cadata)
		if !tlsok {
			return c, fmt.Errorf("Internal error while adding CA data to CA pool")
		}
	}

	c, err = ldap.DialURL(u, ldap.DialWithDialer(dc), ldap.DialWithTLSConfig(tlscfg))
	if err != nil {
		return c, err
	}

	err = c.Bind(us, p)
	return c, err
}

func getReplicationStatus(c *ldap.Conn, d string, t int) ([]replicationStatus, error) {
	var search *ldap.SearchRequest
	var result []replicationStatus

	rplc := strings.NewReplacer("=", "\\=", ",", "\\,")
	// convert LDAP domain into the path used to access replication status object
	base := "dc\\=" + rplc.Replace(strings.Replace(d, ".", ",dc=", -1))
	base = fmt.Sprintf("cn=replica,cn=%s,cn=mapping tree,cn=config", base)

	search = ldap.NewSearchRequest(base, ldap.ScopeSingleLevel, ldap.NeverDerefAliases, 0, t, false, "(objectClass=nsds5replicationagreement)", ldapSearchReplicaInfo, nil)

	sres, err := c.Search(search)
	if err != nil {
		return nil, err
	}

	for _, r := range sres.Entries {
		result = append(result, replicationStatus{
			ReplicationAgreement:  r.GetAttributeValue("cn"),
			ReplicationHost:       r.GetAttributeValue("nsDS5ReplicaHost"),
			ReplicationJSONStatus: []byte(r.GetAttributeValue("nsds5replicaLastUpdateStatusJSON")),
		})
	}

	return result, nil
}
