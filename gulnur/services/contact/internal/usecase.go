package internal

import (
    "gulnur.net/internal/usecase"
    "gulnur.net/internal/repository"
)

type ContactUsecase struct {
    ContactService *ContactService
}

func NewContactUsecase(contactService *ContactService) *ContactUsecase {
    return &ContactUsecase{
        ContactService: contactService,
    }
}
