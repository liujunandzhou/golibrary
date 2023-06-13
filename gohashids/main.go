package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gosuri/uiprogress"
	"github.com/liujunandzhou/golibrary/gohashids/nhashids"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cast"
)

func CheckEncodeDecode() {

	slices := []int64{123456, 123457, 123458}
	for _, data := range slices {

		encodeData, err := nhashids.Encode(data)
		if err != nil {
			log.Printf("Encode failed,err=%v\n", err)
			continue
		}

		fmt.Printf("encode=%s\n", encodeData)

		decodeData, err := nhashids.Decode(encodeData)
		if err != nil {
			log.Printf("Decode failed,err=%v\n", err)
			continue
		}

		fmt.Printf("data=%v encode=%v decode=%v\n", data, encodeData, decodeData)
	}

}

var count int64
var start int64

func init() {
	flag.Int64Var(&count, "count", 10000, "set the count times")
	flag.Int64Var(&start, "start", 100000000, "the start of the number")
}

func main() {

	flag.Parse()

	defer ants.Release()

	var wg sync.WaitGroup

	uiprogress.Start()
	bar := uiprogress.AddBar(int(count)).AppendCompleted().PrependElapsed()

	p, _ := ants.NewPoolWithFunc(10000, func(i interface{}) {

		str, _ := nhashids.Encode(cast.ToInt64(i))

		fmt.Fprintf(os.Stderr, "%s\n", str)

		bar.Incr()

		wg.Done()
	})

	var i int64
	for i = 0; i < count; i++ {

		wg.Add(1)
		_ = p.Invoke(int64(i + start))
	}

	wg.Wait()
}

//对于大文件排序和去重的策略如下：
//sort ip.all -S 80% --parallel=6 -T /data/ipv6/tmp
//-S: 指定使用内存的比率
// --parallel: 使用多核能力
// -T :指定临时目录

//大文件去重思路：
//sort hashids.txt --parallel=6|uniq > test.txt
