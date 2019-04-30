package sql_scripts

const (
	DropTableUsers                   = `drop table users;`
	DropTableAuth                    = `drop table auth;`
	DropTableAuthProviders           = `drop table auth_providers;`
	DropTableAuthentificationHistory = `drop table authentification_history;`
	DropTableUserMfaSecret           = `drop table user_mfa_secret;`
	DropTableUserMfaPhone            = `drop table user_mfa_phone;`
	DropTableUserMfaRecoveryCodes             = `drop table user_mfa_recovery_codes;`
)
