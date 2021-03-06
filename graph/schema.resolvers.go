package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/generated"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/model"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/translators"
)

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return "token", nil
}

func (r *mutationResolver) RegistrarNuevaPersona(ctx context.Context, input model.NewUsuarioSistema) (*model.Persona, error) {
	cu := getCasosUsoPersona(ctx)

	persona := translators.NuevaPersonaAPIToEntity(&input)
	if err := cu.RegistrarNuevaPersona(persona, input.Password); err != nil {
		return nil, err
	}

	result := translators.PersonaEntityToAPI(persona)

	return result, nil
}

func (r *queryResolver) Personas(ctx context.Context, page *int, limit *int) ([]*model.Persona, error) {
	pagina := 1
	limite := 10
	if page != nil {
		pagina = *page
		if *page <= 0 {
			pagina = 1
		}
	}
	if limit != nil {
		limite = *limit
		if limite <= 1 || limite >= 50{
			limite = 10
		}
	}
	cu := getCasosUsoPersona(ctx)

	ps, _, err := cu.Listar(pagina, limite)
	if err != nil {
		return nil, err
	}
	personas := translators.PersonasEntityToAPI(ps)

	return personas, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
