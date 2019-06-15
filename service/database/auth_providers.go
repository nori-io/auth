package database

import (
	"database/sql"

	"github.com/nori-io/nori-common/logger"
)

type authProviders struct {
	db  *sql.DB
	log logger.Writer
}

/*func (a *authProviders) Update_Email(model *AuthProvidersModel) error {
	if model.UserId == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth_providers SET email=? , updated=? WHERE id = ? ",
		model.Email, time.Now(), model.Id)
	return err
}

func (a *authProviders) Update_PhoneNumber_CountryCode(model *AuthProvidersModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth_providers SET  phone_country_code =? , phone_number = ?, updated=? WHERE id = ?",
		model.PhoneCountryCode, model.PhoneNumber, time.Now(), model.Id)

	return err
}

func (a *authProviders) Update_Password_Salt(model **AuthProvidersModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}

	salt, err := CreateSalt()
	if err != nil {
		return err
	}

	password, err := Hash([]byte(model.Password), salt)
	if err != nil {
		return err
	}
	_, err = a.db.Exec("UPDATE auth_providers SET password=? , salt=? , updated=? WHERE id = ? ",
		password, salt, time.Now(), model.Id)
	return err
}

func (a *authProviders) Update_IsEmailVerified(model **AuthProvidersModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth_providers SET is_email_verified=? , updated=? WHERE id = ? ",
		model.IsEmailVerified, time.Now(), model.Id)

	return err
}

func (a *authProviders) Update_IsPhoneVerified(model **AuthProvidersModel) error {
	if model.Id == 0 {
		return errors.New("Empty model")
	}
	_, err := a.db.Exec("UPDATE auth_providers SET is_phone_verified=? , updated=? WHERE id = ? ",
		model.IsPhoneVerified, time.Now(), model.Id)

	return err
}

func (a *authProviders) FindByEmail(email string) (model **AuthProvidersModel, err error) {

	rows, err := a.db.Query("SELECT id, email,password,salt FROM auth_providers WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}

	if rows == nil {
		return model, err
	}

	model = &auth_providersModel{}

	defer rows.Close()
	for rows.Next() {
		var m auth_providersModel

		rows.Scan(&m.Id, &m.Email, &m.Password, &m.Salt)
		model.Id = m.Id
		model.Email = m.Email
		model.Password = m.Password
		model.Salt = m.Salt

	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return model, nil
}


*/
