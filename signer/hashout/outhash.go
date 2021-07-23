package hashout

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"sync"
	"time"
)

type multiHash struct {
	indexOfInteration int
	hash              uint32
}

func printHash(name string, data []byte) string {

	single := printSingleHash(name, data)
	result := printMultiHash(name, single)

	fmt.Printf("%s MultiHash result: %v\n", name, result)

	return result
}

func printSingleHash(name string, data []byte) string {
	md5Hash, crcMd5Hash, crcHash := calculateSingleHash(data)

	single := fmt.Sprintf("%d~%d", crcHash, crcMd5Hash)
	fmt.Printf("%s SingleHash mda5(data) %x SingleHash crc32(mda5(data)) %d SingleHash crc32(data) %d Single Hash result: %s\n",
		name, md5Hash, crcMd5Hash, crcHash, single)

	return single
}

func calculateSingleHash(data []byte) (md5Hash []byte, crcMd5Hash uint32, crcHash uint32) {
	var wg sync.WaitGroup
	md5Hash = sumMda5(data)

	wg.Add(2)
	go func() {
		defer wg.Done()
		crcMd5Hash = sumCrc32(md5Hash)
	}()
	go func() {
		defer wg.Done()
		crcHash = sumCrc32(data)
	}()
	wg.Wait()

	return
}

var ticker = time.NewTicker(time.Millisecond * 10)

func sumMda5(input []byte) []byte {
	for !isFree() {
		time.Sleep(time.Second)
	}

	result := md5.Sum(input)
	return result[:]
}

func isFree() bool {
	select {
	case <-ticker.C:
		return true
	default:
		return false
	}
}

func sumCrc32(input []byte) uint32 {
	time.Sleep(time.Second)
	return crc32.ChecksumIEEE(input)
}

func printMultiHash(name string, singleHash string) string {
	const numberOfIter = 6
	var wg sync.WaitGroup
	wg.Add(numberOfIter)

	multi := make(chan multiHash)
	for i := 0; i < numberOfIter; i++ {
		go func(iterNumber int) {
			defer wg.Done()
			multi <- multiHash{iterNumber, sumIterOfMultiHash(iterNumber, []byte(singleHash))}
		}(i)
	}
	go func() {
		wg.Wait()
		close(multi)
	}()

	result := &outputBuffer{
		single:         singleHash,
		nameOfProgramm: name,
	}

	for hash := range multi {
		result.printHash(hash)
	}

	return result.result
}

func sumIterOfMultiHash(th int, input []byte) uint32 {
	return sumCrc32([]byte(fmt.Sprintf("%d%s", th, input)))
}
