package implementations

import (
	"errors"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (l *loginSql)GuardarCredencial(c *entities.Credencial) error {
	c.ID = primitive.NewObjectID().Hex()
	return l.db.Create(c).Error
}