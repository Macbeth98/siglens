// Copyright (c) 2021-2024 SigScalr, Inc.
//
// This file is part of SigLens Observability Solution
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/valyala/fastrand"
)

func Benchmark_randomBody(b *testing.B) {
	m := make(map[string]interface{})
	f := gofakeit.NewUnlocked(int64(fastrand.Uint32n(1_000)))
	for i := 0; i < b.N; i++ {
		randomizeBody(f, m, true, gofakeit.NewUnlocked(10))
	}
}
