package entities

type LoginManager interface {
	ConsultarUsuario(username string) (*Credencial, error)
}

type Credencial struct {
	ID        string
	PersonaID string
	Persona   *Persona
	Usuario   string
	Password  string
}
