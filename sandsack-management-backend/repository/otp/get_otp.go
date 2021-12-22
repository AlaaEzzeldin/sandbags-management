package repo_otp

import "gorm.io/gorm"

func GetOTP(db *gorm.DB, otp, reason string) (string, error) {
	query := `select code from otp where code = ? and type = ?`
	var code string
	if err := db.Raw(query, otp, reason).Scan(&code).Error; err != nil {
		return "", err
	}
	return code, nil
}
