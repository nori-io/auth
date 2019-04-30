package sql_scripts

const (
	DropTableUsers                   = `drop table users;`
	DropTableAuth                    = `drop table auth;`
	DropTableAuthProviders           = `drop table auth_providers;`
	DropTableAuthenticationHistory = `drop table authentication_history;`
	DropTableUserMfaSecret           = `drop table user_mfa_secret;`
	DropTableUserMfaPhone            = `drop table user_mfa_phone;`
	DropTableUserMfaRecoveryCodes    = `drop table user_mfa_recovery_codes;`
)
