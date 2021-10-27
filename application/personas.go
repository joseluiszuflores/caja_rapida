package application

import (
	"errors"
	"log"
	"strings"

	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
)

type personaCasosUso struct {
	personas entities.PersonaManager
	hasher   entities.Hasher
}

func NewPersonaCasosUso(p entities.PersonaManager, h entities.Hasher) *personaCasosUso {
	return &personaCasosUso{personas: p, hasher: h}
}

func (cu *personaCasosUso) RegistrarNuevaPersona(p *entities.Persona, password string) error {
	hashPassword, err := cu.hasher.HashPassword(strings.TrimSpace(password))
	if err != nil {
		return ErrInternalError
	}

	if err := cu.personas.RegistrarPersonaYCredencial(p, hashPassword); err != nil {
		if errors.Is(err, entities.ErrPolitica) {
			return err
		}
		log.Printf("%s", err)
		return ErrInternalError
	}
	return nil
}

func (cu *personaCasosUso) Listar(pagina, limite int) ([]*entities.Persona, int64, error) {
	personas, total, err := cu.personas.Listar(pagina, limite)
	if err != nil {
		return nil, 0, ErrInternalError
	}

	return personas, total, nil
}
