package main

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"os"
	"os/exec"
)

func main() {
	echoCmd := exec.Command("ipconfig")

	output, err := echoCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("comb out: %v", err)
		os.Exit(1)
	}

	printHash(output)

}

func printHash(data []byte) {
	single := printSingeHash(data)
	multi := ""
	for i := 0; i <= 5; i++ {
		iter := sumIterOfMultiHash(i, []byte(single))
		fmt.Printf("MultuHash %d %s %d\n", i, single, iter)
		multi = fmt.Sprintf("%s%d", multi, iter)
	}

	fmt.Printf("MultiHash result: %v\n", multi)
}

func printSingeHash(data []byte) string {
	md5Hash := sumMda5(data)
	fmt.Printf("SingleHash mda5(data) %x ", md5Hash)

	crcMd5Hash := sumCrc32(md5Hash)
	fmt.Printf("SingleHash crc32(mda5(data)) %d ", crcMd5Hash)

	crc32Hash := sumCrc32(data)
	fmt.Printf("SingleHash crc32(data) %d ", crc32Hash)
	single := fmt.Sprintf("%d~%d", crc32Hash, crcMd5Hash)

	fmt.Println("Single Hash result", single)

	return single

}

func sumMda5(input []byte) []byte {
	result := md5.Sum(input)

	return result[:]
}

func sumCrc32(input []byte) uint32 {
	return crc32.ChecksumIEEE(input)
}

func sumIterOfMultiHash(th int, input []byte) uint32 {
	return sumCrc32([]byte(fmt.Sprintf("%d%s", th, input)))
}
