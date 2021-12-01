// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/30/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"github.com/gonyyi/reqtest"
	"testing"
)

func TestNewRollingIndex(t *testing.T) {
	ri := reqtest.NewRollingIndex(3)
	ri = ri.Next()
	gosl.Test(t, 0, ri.Curr())
	gosl.Test(t, "[0]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 1, ri.Curr())
	gosl.Test(t, "[0 1]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 2, ri.Curr())
	gosl.Test(t, "[0 1 2]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 0, ri.Curr())
	gosl.Test(t, "[1 2 0]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 1, ri.Curr())
	gosl.Test(t, "[2 0 1]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 2, ri.Curr())
	gosl.Test(t, "[0 1 2]", fmt.Sprint(ri.List()))
}
