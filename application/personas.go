package application

import (
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
	"strings"
)

type personaCasosUso struct {
	personas entities.PersonaManager
	login    entities.LoginManager
	hasher   entities.Hasher
}

func NewPersonaCasosUso(p entities.PersonaManager, l entities.LoginManager, h entities.Hasher) *personaCasosUso {
	return &personaCasosUso{personas: p, login: l, hasher: h}
}

func (cu *personaCasosUso) RegistrarNuevaPersona(p *entities.Persona, password string) error {
	personadb, err := cu.personas.ConsultarPorCorreo(p.Correo)
	if err != nil {
		// ToDo: se grega un log para registrar este error?. En todos los errores desconocidos
		return ErrInternalError
	}

	if personadb != nil {
		return ErrCorreoRepetido
	}
	if err := cu.personas.Registrar(p); err != nil {
		return ErrInternalError
	}

	hashPassword, err := cu.hasher.HashPassword(strings.TrimSpace(password))
	if err != nil {
		return ErrInternalError
	}

	credencial := &entities.Credencial{
		PersonaID: p.ID,
		Usuario:   p.Correo,
		Password:  hashPassword,
	}

	if err := cu.login.GuardarCredencial(credencial); err != nil {
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
