package utils

/**
 * 获取两个slice的交集
 */
func SliceIntersect(sliceA []string, sliceB []string) []string {
	sliceRet := make([]string, 0, 1000)
	for key, _ := range sliceA {
		for _key, _ := range sliceB {
			if sliceA[key] == sliceB[_key] {
				sliceRet = append(sliceRet, sliceA[key])
			}
		}
	}
	return sliceRet
}
