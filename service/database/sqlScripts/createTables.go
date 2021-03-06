package sqlScripts

const (
	CreateTableUsers = `CREATE TABLE IF NOT EXISTS users (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  kind VARCHAR(255) NOT NULL,
  status_id INT UNSIGNED NOT NULL,
  type VARCHAR(64) NOT NULL,
  created DATETIME NULL,
  updated DATETIME NULL,
  mfa_type VARCHAR(8) NOT NULL,
  PRIMARY KEY (id))
ENGINE = InnoDB;
`
	CreateTableAuthenticationHistory = ` 
CREATE TABLE IF NOT EXISTS authentication_history (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  logged_in DATETIME NOT NULL,
  meta VARCHAR(255) NOT NULL,
  logged_out DATETIME NULL,
  INDEX user_id_idx (user_id ASC),
  PRIMARY KEY (id),
  CONSTRAINT authentication_history_user_id
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;
`
	CreateTableAuth = `
CREATE TABLE IF NOT EXISTS auth (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  phone VARCHAR(45) NULL,
  email VARCHAR(255) NULL,
  password VARCHAR(65) NOT NULL,
  salt VARCHAR(65) NOT NULL,
  created DATETIME NULL,
  updated DATETIME NULL,
  is_email_verified TINYINT(1) NOT NULL DEFAULT 0,
  is_phone_verified TINYINT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  INDEX user_id_idx (user_id ASC),
  UNIQUE INDEX phone_unique (phone ASC),
  UNIQUE INDEX email_unique (email ASC),
  UNIQUE INDEX user_id_unique (user_id ASC),
  CONSTRAINT auth_user_id_fk
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;
`
	CreateTableUserMfaSecret = `
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

	CreateTableUserMfaPhone = `
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
	CreateTableUsersMfaCode = `
CREATE TABLE IF NOT EXISTS users_mfa_code (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  code VARCHAR(45) NOT NULL,
  PRIMARY KEY (id),
  INDEX user_id_idx (user_id ASC),
  CONSTRAINT users_mfa_code_user_id_fk
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
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
)
