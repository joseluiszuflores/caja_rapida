package entities

import "time"

type PersonaManager interface {
	Registrar(p *Persona) error
	Actualizar(p *Persona) error
	ConsultarPorId(id string) (*Persona, error)
	ConsultarPorCorreo(correo string) (*Persona, error)
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
