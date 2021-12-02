// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/01/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)

func TestNewRollingIndex(t *testing.T) {
	ri := gosl.NewRollingIndex(3)
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
