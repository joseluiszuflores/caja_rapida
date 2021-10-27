package graph

import (
	"context"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/application"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/configs"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/entities"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/db/postgres"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/hasher"
	"github.com/PerezBautistaAntonioDeJesus/caja_rapida/infrastructure/implementations"
)

func getCasosUsoPersona(ctx context.Context) entities.UcPersonas {
	db := postgres.GetDB().WithContext(ctx)
	p := implementations.NewPersonaSql(db)
	h := hasher.NewHasher256(configs.GetConfig().SecretHashKey)
	cu := application.NewPersonaCasosUso(p, h)
	return cu
}
