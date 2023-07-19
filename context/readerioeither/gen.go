package readerioeither

// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-07-19 16:18:34.1521763 +0200 CEST m=+0.011558001

import (
	"context"

	G "github.com/IBM/fp-go/context/readerioeither/generic"
)

// Eitherize0 converts a function with 0 parameters returning a tuple into a function with 0 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize0]
func Eitherize0[F ~func(context.Context) (R, error), R any](f F) func() ReaderIOEither[R] {
	return G.Eitherize0[ReaderIOEither[R]](f)
}

// Eitherize1 converts a function with 1 parameters returning a tuple into a function with 1 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize1]
func Eitherize1[F ~func(context.Context, T0) (R, error), T0, R any](f F) func(T0) ReaderIOEither[R] {
	return G.Eitherize1[ReaderIOEither[R]](f)
}

// Eitherize2 converts a function with 2 parameters returning a tuple into a function with 2 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize2]
func Eitherize2[F ~func(context.Context, T0, T1) (R, error), T0, T1, R any](f F) func(T0, T1) ReaderIOEither[R] {
	return G.Eitherize2[ReaderIOEither[R]](f)
}

// Eitherize3 converts a function with 3 parameters returning a tuple into a function with 3 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize3]
func Eitherize3[F ~func(context.Context, T0, T1, T2) (R, error), T0, T1, T2, R any](f F) func(T0, T1, T2) ReaderIOEither[R] {
	return G.Eitherize3[ReaderIOEither[R]](f)
}

// Eitherize4 converts a function with 4 parameters returning a tuple into a function with 4 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize4]
func Eitherize4[F ~func(context.Context, T0, T1, T2, T3) (R, error), T0, T1, T2, T3, R any](f F) func(T0, T1, T2, T3) ReaderIOEither[R] {
	return G.Eitherize4[ReaderIOEither[R]](f)
}

// Eitherize5 converts a function with 5 parameters returning a tuple into a function with 5 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize5]
func Eitherize5[F ~func(context.Context, T0, T1, T2, T3, T4) (R, error), T0, T1, T2, T3, T4, R any](f F) func(T0, T1, T2, T3, T4) ReaderIOEither[R] {
	return G.Eitherize5[ReaderIOEither[R]](f)
}

// Eitherize6 converts a function with 6 parameters returning a tuple into a function with 6 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize6]
func Eitherize6[F ~func(context.Context, T0, T1, T2, T3, T4, T5) (R, error), T0, T1, T2, T3, T4, T5, R any](f F) func(T0, T1, T2, T3, T4, T5) ReaderIOEither[R] {
	return G.Eitherize6[ReaderIOEither[R]](f)
}

// Eitherize7 converts a function with 7 parameters returning a tuple into a function with 7 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize7]
func Eitherize7[F ~func(context.Context, T0, T1, T2, T3, T4, T5, T6) (R, error), T0, T1, T2, T3, T4, T5, T6, R any](f F) func(T0, T1, T2, T3, T4, T5, T6) ReaderIOEither[R] {
	return G.Eitherize7[ReaderIOEither[R]](f)
}

// Eitherize8 converts a function with 8 parameters returning a tuple into a function with 8 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize8]
func Eitherize8[F ~func(context.Context, T0, T1, T2, T3, T4, T5, T6, T7) (R, error), T0, T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7) ReaderIOEither[R] {
	return G.Eitherize8[ReaderIOEither[R]](f)
}

// Eitherize9 converts a function with 9 parameters returning a tuple into a function with 9 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize9]
func Eitherize9[F ~func(context.Context, T0, T1, T2, T3, T4, T5, T6, T7, T8) (R, error), T0, T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8) ReaderIOEither[R] {
	return G.Eitherize9[ReaderIOEither[R]](f)
}

// Eitherize10 converts a function with 10 parameters returning a tuple into a function with 10 parameters returning a [ReaderIOEither[R]]
// The inverse function is [Uneitherize10]
func Eitherize10[F ~func(context.Context, T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, error), T0, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) ReaderIOEither[R] {
	return G.Eitherize10[ReaderIOEither[R]](f)
}
