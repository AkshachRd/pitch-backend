package main

import "context"

type service struct {
	store CardsStore
}

func NewService(store CardsStore) *service {
	return &service{
		store: store,
	}
}

func (s *service) CreateCard(ctx context.Context) error {
	return nil // TODO
}