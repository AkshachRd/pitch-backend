package main

import ("context")

type CardsService interface {
	CreateCard(ctx context.Context) error
}

type CardsStore interface {
	Create(ctx context.Context) error
}