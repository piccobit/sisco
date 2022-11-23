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

	needles := []needle{
		{
			search:      "{user}",
			replacement: ldap.EscapeFilter(user),
		},
	}

	filter := replace(lc.config.LdapFilterUser, &needles)

	attributes := []string{
		"dn",
	}

	searchReq := ldap.NewSearchRequest(
		lc.config.LdapBaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err := lc.ldapConn.Search(searchReq)
	if err != nil {
		return auth.Unknown, err
	}

	if len(result.Entries) == 0 {
		return auth.Unknown, err
	}

	dn := result.Entries[0].DN

	err = lc.ldapConn.Bind(dn, password)
	if err != nil {
		return auth.Unknown, err
	}

	filter = replace(lc.config.LdapFilterGroup, &needles)

	attributes = []string{
		"cn",
	}

	searchReq = ldap.NewSearchRequest(
		lc.config.LdapBaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		[]ldap.Control{},
	)

	result, err = lc.ldapConn.Search(searchReq)
	if err != nil {
		return auth.Unknown, err
	}

	permissions := auth.Unknown

	for _, e := range result.Entries {
		if strings.EqualFold(e.DN, cfg.Config.LdapUsersGroupDN) {
			permissions = permissions | auth.User
		}

		if strings.EqualFold(e.DN, cfg.Config.LdapServicesGroupDN) {
			permissions = permissions | auth.Service
		}

		if strings.EqualFold(e.DN, cfg.Config.LdapAdminsGroupDN) {
			permissions = permissions | auth.Admin
		}
	}

	return permissions, nil
}

type needle struct {
	search      string
	replacement string
}

func replace(haystack string, needles *[]needle) string {
	for _, e := range *needles {
		haystack = strings.Replace(
			haystack,
			e.search,
			e.replacement,
			-1,
		)
	}

	return haystack
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
