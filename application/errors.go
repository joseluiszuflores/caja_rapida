package application

import "errors"

var (
	ErrInternalError = errors.New("error interno desconocido")

	// Personas
	ErrCorreoRepetido = errors.New("ya existe una cuenta con ese correo")
)
