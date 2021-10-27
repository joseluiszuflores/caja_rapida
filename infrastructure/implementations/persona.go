package implementations

import (
	"errors"
	"fmt"
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

func (p2 *personaSql) RegistrarPersonaYCredencial(p *entities.Persona, hashPassword string) error {

	p2.db = p2.db.Begin()
	if err := p2.db.Error; err != nil {
		return fmt.Errorf("no se ha podido iniciar transaccion: %w", err)
	}
	defer p2.db.Rollback()

	// Se valida el correo
	personadb, err := p2.consultarPorCorreo(p.Correo)
	if err != nil {
		return err
	}

	if personadb != nil {
		return fmt.Errorf("%w: ya existe una cuenta con ese correo", entities.ErrPolitica)
	}

	// Se inserta registro Persona
	p.ID = primitive.NewObjectID().Hex()
	p.Correo = strings.ToLower(p.Correo)
	if err := p2.db.Create(p).Error; err != nil {
		return fmt.Errorf("no se ha podido crear persona: %w", err)
	}

	// Se crea la credencial
	credencial := entities.Credencial{
		PersonaID: p.ID,
		Usuario:   p.Correo,
		Password:  hashPassword,
	}
	credencial.ID = primitive.NewObjectID().Hex()
	if err := p2.db.Create(&credencial).Error; err != nil {
		return fmt.Errorf("no se ha podido crear credencial: %w", err)
	}

	if err := p2.db.Commit().Error; err != nil {
		return fmt.Errorf("no se ha podido hacer commit: %w", err)
	}
	return nil
}

func (p2 *personaSql) ConsultarPorId(id string) (*entities.Persona, error) {
	row := new(entities.Persona)
	err := p2.db.Where("id", id).First(row).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("no se ha podido consultar por id: %w", err)
	}
	return row, err
}

func (p2 *personaSql) consultarPorCorreo(correo string) (*entities.Persona, error) {
	row := new(entities.Persona)
	err := p2.db.Where("correo", correo).First(row).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("no se ha podido consultar por correo: %w", err)
	}
	return row, nil
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
