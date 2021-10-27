package implementations

import (
	"errors"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
	"gorm.io/gorm"
	"strings"
)

type loginSql struct {
	db *gorm.DB
}

func NewLoginSql(db *gorm.DB) *loginSql {
	return &loginSql{db: db}
}

func (l *loginSql) ConsultarUsuario(username string) (*entities.Credencial, error) {
	row := new(entities.Credencial)
	err := l.db.Where("usuario", strings.ToLower(username)).First(row).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return row, err
}
