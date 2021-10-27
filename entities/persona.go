package entities

import "time"

type UcPersonas interface {
	RegistrarNuevaPersona(p *Persona, password string) error
	Listar(pagina, limite int) ([]*Persona, int64, error)
}

type PersonaManager interface {
	RegistrarPersonaYCredencial(p *Persona, hashPassword string) error
	ConsultarPorId(id string) (*Persona, error)
	Listar(pagina, limite int) ([]*Persona, int64, error)
}
type Persona struct {
	ID              string
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string
	Correo          string
	Telefono        string
	FechaNacimiento time.Time
}
