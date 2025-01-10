package variables

import "fmt"

func ShowIntegers() {
	var intCommon int
	intOf32 := int32(10)
	intOf64 := int64(100)

	fmt.Println("intCommon = ", intCommon)
	fmt.Println("intCommon 32 = ", intOf32)
	fmt.Println("intCommon 64 = ", intOf64)
}