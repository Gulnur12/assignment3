package internal

import (
    "gulnur.net/internal/delivery"
    "gulnur.net/internal/usecase"
)

type ContactDelivery struct {
    ContactUsecase *usecase.ContactUsecase
}

func NewContactDelivery(contactUsecase *usecase.ContactUsecase) *ContactDelivery {
    return &ContactDelivery{
        ContactUsecase: contactUsecase,
    }
}
