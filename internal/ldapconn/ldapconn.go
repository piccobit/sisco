package ldapconn

import (
	"errors"
	"fmt"
	"log"
	"sisco/internal/auth"
	"strings"

	"github.com/go-ldap/ldap"
	"sisco/internal/cfg"
)

type LDAPConn struct {
	config            *cfg.Configuration
	ldapConn          *ldap.Conn
	reconnectAttempts int
}

func New(cfg *cfg.Configuration) (*LDAPConn, error) {
	var err error

	lc := LDAPConn{
		config: cfg,
	}

	if lc.ldapConn == nil {
		lc.ldapConn, err = ldap.DialURL(cfg.LdapURL)
		if err != nil {
			return nil, err
		}
	}

	err = lc.ldapConn.Bind(cfg.LdapBindDN, cfg.LdapBindPassword)
	if err != nil && ldap.IsErrorWithCode(err, 200) {
		log.Println("Connection is closed, trying to reconnect...")
		if err := lc.reconnect(); err != nil {
			log.Println("error while trying to reconnect")
			return nil, err
		}
	}

	return &lc, nil
}

func (lc *LDAPConn) Authenticate(user string, password string) (auth.Permissions, error) {
	var err error

	// We check first if this is an 'admin' token.

	permissions := auth.Unknown

	filter := replace(lc.config.LdapFilterAdminsDN, "{user}", ldap.EscapeFilter(user))

	searchReq := ldap.NewSearchRequest(lc.config.LdapBaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, []string{"dn"}, []ldap.Control{})

	result, err := lc.ldapConn.Search(searchReq)
	if err != nil {
		return auth.Unknown, err
	}

	if len(result.Entries) != 0 {
		permissions = permissions | auth.Admin
	} else {
		filter = replace(lc.config.LdapFilterUsersDN, "{user}", ldap.EscapeFilter(user))

		searchReq = ldap.NewSearchRequest(lc.config.LdapBaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, []string{"dn"}, []ldap.Control{})

		result, err = lc.ldapConn.Search(searchReq)
		if err != nil {
			return auth.Admin, err
		}

		if len(result.Entries) == 0 {
			return auth.Admin, errors.New("user not found")
		}
	}

	dn := result.Entries[0].DN

	err = lc.ldapConn.Bind(dn, password)
	if err != nil {
		return auth.Admin, err
	}

	return permissions, nil
}

func replace(haystack string, needle string, replacement string) string {
	res := strings.Replace(
		haystack,
		needle,
		replacement,
		-1,
	)

	return res
}

func (lc *LDAPConn) reconnect() error {
	var err error

	lc.ldapConn.Close()

	lc.ldapConn, err = ldap.DialURL(lc.config.LdapURL)
	if err != nil {
		return err
	}

	err = lc.ldapConn.Bind(lc.config.LdapBindDN, lc.config.LdapBindPassword)
	if err != nil {
		if err != nil {
			return errors.New(fmt.Sprintf("error while trying to reconnect: %v", err))
		}
	}

	return nil
}
