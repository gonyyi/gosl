// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Bitlag_NewBitflag(t *testing.T) {
	const (
		HDFS gosl.Bitflag = 1 << iota
		LOCALFILE
		SQL
		ORACLE
	)

	f := func(b gosl.Bitflag) string {
		return string(b.Output(nil))
	}

	gosl.Test(t,
		f(gosl.Bitflag(1).Nth(2, 4, 6, 8, 10)),
		"11010101010000000000000000000000")
}


