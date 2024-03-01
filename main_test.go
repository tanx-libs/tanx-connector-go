package main

import (
	"testing"

	"github.com/tanx-libs/tanx-connector-go/crypto-cpp/src/starkware/crypto/ffi/crypto_lib"
)

func BenchmarkSign(b *testing.B) {
	// numGoroutines := 5
	// results := 1000

	// buffer of 5 signatures
	// data := make(chan string, results)
	privKey := "588ad6325783739b3806e27feebeb1c270d4c6875c29517ff1d689e777b13a6"

	m := "312a58fc5637279faed49de9e786b800edc87343e41ccc871bb9193783762ebb"
	for i := 0; i < b.N; i++ {
		crypto_lib.Sign(privKey, m, "")
	}

	// generate := func() {
	// 	for i:= 0; i < results/numGoroutines; i++{
	// 		bytes := make([]byte, 32)
	// 		if _, err := rand.Read(bytes); err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		m = fmt.Sprintf("%x", bytes)
	// 		r, s = client.Sign(privKey, m)
	// 		data <- fmt.Sprintf("0x%s,0x%s,0x%s\n", m, r, s)
	// 	}
	// }

	// for i := 0; i < numGoroutines; i++ {
	// 	go generate()
	// }

	// // file input
	// file, err := os.Create("benchmark.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// writer := bufio.NewWriter(file)

	// fmt.Fprintf(file, "%s,%s,%s\n", "msg_hash", "r", "s")
	// b.ReportAllocs()

	// b.N = results
	// for i := 0; i < b.N; i++ {
	// 	// fmt.Fprintf(writer, <-data)
	// 	<-data
	// 	// if i%500 == 0 {
	// 	// 	writer.Flush()
	// 	// }
	// }

	// writer.Flush()
}
