package gohashids

import (
	"math/rand"
	"time"

	"github.com/spf13/cast"
)

var sortMap map[int]string

func init() {

	sortMap = make(map[int]string)
	sortMap[0] = "68553402-5250-41C4-8A11-C4C54967BE21(MIVnj,^s)"
	sortMap[1] = "E1CF511E-0D04-4A23-A973-B8D28D7159D1(AJa>}V2p)"
	sortMap[2] = "DC715F96-023C-4A39-9655-FEA080FED963(hmaT+[zW)"
	sortMap[3] = "B0981BB2-4CE9-4452-9CCC-A86C90413D7B(c*Hu'wh\")"
	sortMap[4] = "C6049B30-F07A-4CA0-B725-F49E5BABBDB6(o;vwYMIB)"
	sortMap[5] = "802BC740-2048-483F-BD6C-058A9F0AC33A(e:%Vyv}))"
	sortMap[6] = "F7BC0C5A-6BF8-4ED7-B067-0C44DDBB3B00(JUVTMC`.)"
	sortMap[7] = "ECA1DA60-EBC4-48B0-98A5-EC5F29E02E69($UN#KA7&)"
	sortMap[8] = "56BC7A41-C53B-4568-BFAE-5BFBF9DB5B03(0{*OnI{2)"
	sortMap[9] = "A6F7BED3-760B-49ED-AF06-4C8A4D40DC6A(vz{_rjKs)"
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getSalt(idx int64) (string, string) {

	newIdx := idx % int64(len(sortMap))

	return cast.ToString(newIdx), sortMap[int(newIdx)]
}

func splitInt(idx int) (int, int) {

	if idx <= 0 {
		return 0, 0
	}

	tmp := rand.Int() % idx

	return tmp, idx - tmp
}

func splitStr(idx int) (string, string) {

	first, end := splitInt(idx)

	return cast.ToString(first), cast.ToString(end)
}
