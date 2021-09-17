package util

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseHumNum(t *testing.T) {
	tests := []struct {
		name  string
		input string
		res   int
	}{
		{
			"simple number",
			"100",
			100,
		},
		{
			"simple number",
			"-100",
			-100,
		},
		{
			"num with suffix",
			"100K",
			100000,
		},
		{
			"invalid num",
			"100dd",
			MinInt,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, v.res, ParseHumNum(v.input))
		})
	}
}

func TestMonToNum(t *testing.T) {
	tests := []struct {
		name  string
		input string
		res   int
	}{
		{
			"ok",
			"JUN",
			6,
		},
		{
			"invalid mon",
			"BIB",
			0,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, v.res, MonToNum(v.input))
		})
	}
}

var (
	strIn = `adsaf asfas asfsa
asfas rtetrgf dsgfsd
sdgds dgsd sadsa etw
         asfas rtetrgf dsgfsd
afs ewrd
asfas rtetrgf dsgfsd
fasf`
	monIn = `DEC fsfsdf
APR fdsf 123
JAN g214
FEB`
	numIn = `123 463 34
535 53 54
424 656 3
123`
	humIn = `204G
2
2k
3M`
)

func TestSort(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sets  *Settings
		out   string
		ok    bool
	}{
		{
			"strings sort",
			strIn,
			NewSettings(
				NewFlags(-1, false, false, false, false),
				NewIncompatibleFlags(false, false, false),
			),
			`         asfas rtetrgf dsgfsd
adsaf asfas asfsa
afs ewrd
asfas rtetrgf dsgfsd
asfas rtetrgf dsgfsd
fasf
sdgds dgsd sadsa etw`,
			true,
		},
		{
			"strings sort with -b",
			strIn,
			NewSettings(
				NewFlags(-1, false, false, true, false),
				NewIncompatibleFlags(false, false, false),
			),
			`adsaf asfas asfsa
afs ewrd
asfas rtetrgf dsgfsd
         asfas rtetrgf dsgfsd
asfas rtetrgf dsgfsd
fasf
sdgds dgsd sadsa etw`,
			true,
		},
		{
			"strings sort with -u",
			strIn,
			NewSettings(
				NewFlags(-1, false, true, false, false),
				NewIncompatibleFlags(false, false, false),
			),
			`         asfas rtetrgf dsgfsd
adsaf asfas asfsa
afs ewrd
asfas rtetrgf dsgfsd
fasf
sdgds dgsd sadsa etw`,
			true,
		},
		{
			"month sort",
			monIn,
			NewSettings(
				NewFlags(-1, false, false, false, false),
				NewIncompatibleFlags(false, true, false),
			),
			`JAN g214
FEB
APR fdsf 123
DEC fsfsdf`,
			true,
		},
		{
			"numbers sort",
			numIn,
			NewSettings(
				NewFlags(-1, false, false, false, false),
				NewIncompatibleFlags(true, false, false),
			),
			`123 463 34
123
424 656 3
535 53 54`,
			true,
		},
		{
			"numbers sort check",
			numIn,
			NewSettings(
				NewFlags(-1, false, false, false, true),
				NewIncompatibleFlags(true, false, false),
			),
			"строки не отсортированы",
			true,
		},
		{
			"human readable numbers sort",
			humIn,
			NewSettings(
				NewFlags(-1, false, false, false, false),
				NewIncompatibleFlags(false, false, true),
			),
			`2
2k
3M
204G`,
			true,
		},
		{
			"incompatible flags",
			humIn,
			NewSettings(
				NewFlags(-1, false, false, false, false),
				NewIncompatibleFlags(true, true, true),
			),
			"",
			false,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, err := Sort(strings.NewReader(v.input), v.sets)
			if v.ok {
				assert.Nil(t, err)
				assert.Equal(t, v.out, res)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
