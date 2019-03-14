package sql_scripts

const (
	SetDatabaseSettings = `SET GLOBAL sql_mode=''`
	SetDatabaseStricts  = `SET sql_mode='STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ZERO_IN_DATE'`

	CreateTableUsers = `CREATE TABLE IF NOT EXISTS users (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  status_account ENUM('active', 'locked', 'blocked', 'deleted') NOT NULL,
  type VARCHAR(64) NOT NULL,
  created DATETIME NOT NULL,
  updated DATETIME NOT NULL,
  mfa_type ENUM ('opt', 'phone') NOT NULL,
  PRIMARY KEY (id))
ENGINE = InnoDB;
`
	CreateTableAuth = `
CREATE TABLE IF NOT EXISTS auth (
   id INT UNSIGNED NOT NULL AUTO_INCREMENT,
   user_id INT UNSIGNED NOT NULL,
   phone_country_code VARCHAR(8) NULL,
   phone_number VARCHAR(16) NULL,
   email VARCHAR(255) NULL,
   password VARCHAR(65) NOT NULL,
   salt VARCHAR(65) NOT NULL,
   created DATETIME NOT NULL,
   updated DATETIME NOT NULL,
   is_email_verified TINYINT(1) NOT NULL DEFAULT 0,
  is_phone_verified TINYINT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  INDEX user_id_idx (user_id ASC),
  UNIQUE INDEX email_unique (email ASC),
  UNIQUE INDEX user_id_unique (user_id ASC),
  UNIQUE INDEX phone (phone_country_code ASC, phone_number ASC),
  CONSTRAINT auth_user_id_fk
     FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;
`
	CreateTableAuthProviders = `
CREATE TABLE IF NOT EXISTS auth_providers (
  provider VARCHAR(64) NOT NULL,
  provider_user_key VARCHAR(128) NOT NULL,
  user_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (provider, provider_user_key),
  INDEX user_id_idx (user_id ASC),
  CONSTRAINT auth_providers_user_id_fk
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;
`

	CreateTableAuthentificationHistory = ` 
CREATE TABLE IF NOT EXISTS authentification_history (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  signin DATETIME NOT NULL,
  meta VARCHAR(255) NOT NULL,
  signout DATETIME NULL,
  INDEX user_id_idx (user_id ASC),
  PRIMARY KEY (id),
  CONSTRAINT authentication_history_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;
`

	CreateTableUsersMfaSecret = `
CREATE TABLE IF NOT EXISTS user_mfa_secret (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  secret VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX user_id_UNIQUE (user_id ASC),
  INDEX user_id_idx (user_id ASC),
  CONSTRAINT user_mfa_secret_user_id_fk
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
`

	CreateTableUsersMfaPhone = `
CREATE TABLE IF NOT EXISTS user_mfa_phone (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  phone VARCHAR(45) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX user_id_UNIQUE (user_id ASC),
  UNIQUE INDEX phone_UNIQUE (phone ASC),
  INDEX user_id_idx (user_id ASC),
  CONSTRAINT user_mfa_phone_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
`
	CreateTableUserMfaCode = `
CREATE TABLE IF NOT EXISTS user_mfa_code (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  code VARCHAR(45) NOT NULL,
  PRIMARY KEY (id),
  INDEX user_id_idx (user_id ASC),
  CONSTRAINT user_mfa_code_user_id_fk
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
`
)
