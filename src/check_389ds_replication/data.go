package main

type replicationStatus struct {
	ReplicationAgreement  string
	ReplicationHost       string
	ReplicationJSONStatus []byte
}

type status struct {
	Critical []string
	Warning  []string
	Ok       []string
	Unknown  []string
}

type updateStatus struct {
	State      string `ini:"state"`
	LDAPRC     string `ini:"ldap_rc"`
	LDAPRCText string `ini:"ldap_rc_text"`
	ReplRC     string `ini:"repl_rc"`
	ReplRCText string `ini:"repl_rc_text"`
	Date       string `ini:"date"`
	Message    string `ini:"message"`
}
