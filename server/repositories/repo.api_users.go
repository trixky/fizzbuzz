package repositories

import (
	"errors"
	"strings"

	"fizzbuzz.com/v1/database"
	"fizzbuzz.com/v1/models"
	"fizzbuzz.com/v1/tools"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Login_api_users(pseudo string, password string) (api_user *models.Api_users, session_token string, err error) {
	api_user = new(models.Api_users)
	hashed_password_str := tools.Hash_password(password) // hash the password

	if err = database.Postgres.Table("api_users").Where(map[string]interface{}{"pseudo": pseudo, "password": hashed_password_str}).First(api_user).Error; err != nil {
		// If postgres could not find the user.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If postgres could not find the user with this pseudo or/and password.
			return nil, session_token, nil
		}
		// postgres failed.
		return nil, session_token, err
	}

	if api_user.Blocked {
		// If user is blocked.
		return api_user, session_token, nil
	}

	session_token = uuid.NewV4().String() // generates a new session token (uuid)

	if err := database.Redis.Append("tokens>"+pseudo, "|"+session_token).Err(); err != nil {
		// If redis failed to add the token key to token keys.
		return api_user, session_token, err
	}
	if err = database.Redis.Set("token>"+session_token, pseudo, 0).Err(); err != nil {
		return api_user, session_token, err
	}
	return api_user, session_token, nil
}

func Register_api_users(pseudo string, password string) (pseudo_already_taken bool, err error) {
	var pseudo_count int64

	if err := database.Postgres.Table("api_users").Where("pseudo = ?", pseudo).Count(&pseudo_count).Error; err != nil {
		// If postgres failed to search a user with this pseudo.
		return false, err
	}

	if pseudo_count > 0 {
		// If a user with this pseudo already exists.
		return true, nil
	}

	hashed_password_str := tools.Hash_password(password) // hash the password

	if database.Postgres.Create(&models.Api_users{Pseudo: pseudo, Password: hashed_password_str}).Error != nil {
		// If postgres failed to create the new user.
		return false, err
	}

	return false, nil
}

func Is_admin_api_users(pseudo string) (is_admin bool, err error) {
	api_user := models.Api_users{}

	if err = database.Postgres.Table("api_users").Where(map[string]interface{}{"pseudo": pseudo, "admin": true}).First(&api_user).Error; err != nil {
		// If postgres could not find the user.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If user don't have the required privileges.
			return false, nil
		}
		// postgres failed.
		return false, err
	}
	return true, nil
}

func Block_api_users(pseudo string, block_status string) (user_not_found bool, err error) {
	if result := database.Postgres.Table("api_users").Where("pseudo = ?", pseudo).Update("blocked", block_status); result.Error != nil {
		// If postgres failed to update the 'blocked' status.
		return true, result.Error
	} else if result.RowsAffected > 0 {
		if block_status == "true" {
			// If the new 'blocked' status is 'true'.
			tokens_key := "tokens>" + pseudo
			if result := database.Redis.Keys(tokens_key); result.Err() != nil {
				// If redis failed to get token keys.
				return false, result.Err()
			} else if len(result.Val()) > 0 {
				result := database.Redis.Get(tokens_key)
				if result.Err() != nil {
					// If redis failed to get token values.
					return false, result.Err()
				}
				tokens := strings.Split(result.Val(), "|")[1:]
				for _, token := range tokens {
					database.Redis.Del("token>" + token) // delete each token key
				} // ignore err
				database.Redis.Del(tokens_key) // delete token keys // ignore err
			}
		}
		return false, nil
	}
	// user finded with this pseudo.
	return true, nil
}
