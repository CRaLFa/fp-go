// Copyright (c) 2023 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package either

import (
	L "github.com/IBM/fp-go/lazy"
	M "github.com/IBM/fp-go/monoid"
)

// AlternativeMonoid is the alternative [Monoid] for an [Either]
func AlternativeMonoid[E, A any](m M.Monoid[A]) M.Monoid[Either[E, A]] {
	return M.AlternativeMonoid(
		Of[E, A],
		MonadMap[E, A, func(A) A],
		MonadAp[A, E, A],
		MonadAlt[E, A],
		m,
	)
}

// AltMonoid is the alternative [Monoid] for an [Either]
func AltMonoid[E, A any](zero L.Lazy[Either[E, A]]) M.Monoid[Either[E, A]] {
	return M.AltMonoid(
		zero,
		MonadAlt[E, A],
	)
}
