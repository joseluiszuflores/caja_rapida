package translators

import (
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/model"
)

func NuevaPersonaAPIToEntity(p *model.NewUsuarioSistema) *entities.Persona {
	persona := &entities.Persona{
		Nombre:   p.Nombre,
		Correo:   p.Correo,
		Telefono: p.Telefono,
	}

	if p.ApellidoPaterno != nil {
		persona.ApellidoPaterno = *p.ApellidoPaterno
	}
	if p.ApellidoMaterno != nil {
		persona.ApellidoMaterno = *p.ApellidoMaterno
	}
	return persona
}

func PersonaEntityToAPI(p *entities.Persona) *model.Persona {
	return &model.Persona{
		ID: p.ID,
		ApellidoMaterno: &p.ApellidoMaterno,
		ApellidoPaterno: &p.ApellidoPaterno,
		Nombre:   p.Nombre,
		Correo:   p.Correo,
		Telefono: p.Telefono,
	}
}