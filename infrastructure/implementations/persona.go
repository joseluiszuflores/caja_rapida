package implementations

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"strings"

	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
)

type personaSql struct {
	db *gorm.DB
}

func NewPersonaSql(db *gorm.DB) *personaSql {
	return &personaSql{db: db}
}

func (p2 *personaSql) Registrar(p *entities.Persona) error {
	p.ID = primitive.NewObjectID().Hex()
	p.Correo = strings.ToLower(p.Correo)
	return p2.db.Create(p).Error
}

func (p2 *personaSql) Actualizar(p *entities.Persona) error {
	panic("implement me")
}

func (p2 *personaSql) ConsultarPorId(id string) (*entities.Persona, error) {
	row := new(entities.Persona)
	err := p2.db.Where("id", id).First(row).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return row, err
}

func (p2 *personaSql) ConsultarPorCorreo(correo string) (*entities.Persona, error) {
	row := new(entities.Persona)
	err := p2.db.Where("correo", correo).First(row).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return row, err
}

func (p2 *personaSql) Listar(pagina, limite int) ([]*entities.Persona, int64, error) {

	var rows []*entities.Persona
	var total int64

	if err := p2.db.Model(&entities.Persona{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := p2.db.
		Limit(limite).Offset((pagina - 1) * limite).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}

	return rows, total, nil

}
