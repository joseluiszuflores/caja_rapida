package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/configs"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/hasher"

	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/application"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/generated"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/model"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/graph/translators"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/db/postgres"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/implementations"
)

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return "token", nil
}

func (r *mutationResolver) RegistrarNuevaPersona(ctx context.Context, input model.NewUsuarioSistema) (*model.Persona, error) {
	db := postgres.GetDB().WithContext(ctx).Begin()
	if db.Error != nil {
		return nil, application.ErrInternalError
	}
	defer db.Rollback()

	p := implementations.NewPersonaSql(db)
	l := implementations.NewLoginSql(db)
	h := hasher.NewHasher256(configs.GetConfig().SecretHashKey)
	cu := application.NewPersonaCasosUso(p, l, h)

	persona := translators.NuevaPersonaAPIToEntity(&input)
	if err := cu.RegistrarNuevaPersona(persona, input.Password); err != nil {
		return nil, err
	}
	if err := db.Commit().Error; err != nil {
		return nil, application.ErrInternalError
	}
	result := translators.PersonaEntityToAPI(persona)

	return result, nil
}

func (r *queryResolver) Resultados(ctx context.Context) ([]*model.Persona, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }