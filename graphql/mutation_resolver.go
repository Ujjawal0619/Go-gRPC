package main

import "context"

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) createAccount(ctx context.Context, in AccountInput) (*Account, error) {
	panic("unimplemented")
}

func (r *mutationResolver) createProduct(ctx context.Context, in ProductInput) (*Product, error) {
	panic("unimplemented")

}

func (r *mutationResolver) createOrder(ctx context.Context, in OrderInput) (*Order, error) {
	panic("unimplemented")
}
