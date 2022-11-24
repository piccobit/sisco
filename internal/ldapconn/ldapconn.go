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

func (lc *LDAPConn) Authenticate(user string, password string) (string, auth.Permissions, error) {
	var err error
	var group string

	needles := []needle{
		{
			search:      "{user_attribute}",
			replacement: ldap.EscapeFilter(cfg.Config.LdapFilterUserAttribute),
		},
		{
			search:      "{group_attribute}",
			replacement: ldap.EscapeFilter(cfg.Config.LdapFilterGroupAttribute),
		},
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
		return "", auth.Unauthorized, err
	}

	if len(result.Entries) == 0 {
		return "", auth.Unauthorized, err
	}

	dn := result.Entries[0].DN

	err = lc.ldapConn.Bind(dn, password)
	if err != nil {
		return "", auth.Unauthorized, err
	}

	filter = replace(lc.config.LdapFilterGroup, &needles)

	attributes = []string{
		"cn",
		"uid",
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
		return "", auth.Unauthorized, err
	}

	permissions := auth.Unauthorized

	for _, e := range result.Entries {
		tmpGroup := e.GetAttributeValue("uid")

		if strings.EqualFold(tmpGroup, cfg.Config.LdapUsersGroup) {
			group = tmpGroup
			permissions = permissions | auth.User
		}

		if strings.EqualFold(tmpGroup, cfg.Config.LdapServicesGroup) {
			group = tmpGroup
			permissions = permissions | auth.Service
		}

		if strings.EqualFold(tmpGroup, cfg.Config.LdapAdminsGroup) {
			group = tmpGroup
			permissions = permissions | auth.Admin
		}
	}

	return group, permissions, nil
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
