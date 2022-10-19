package mergesort

func sort(arr []int64, compare func(a int64, b int64) bool) []int64 {
	if len(arr) <= 1 {
		return arr
	} else {
		arr1 := sort(arr[0:len(arr)/2], compare)
		arr2 := sort(arr[len(arr)/2:], compare)
		var i int = 0
		var j int = 0
		var ans []int64
		for i < len(arr1) || j < len(arr2) {
			if i == len(arr1) {
				ans = append(ans, arr2[j])
				j++
			} else if j == len(arr2) {
				ans = append(ans, arr1[i])
				i++
			} else {
				if compare(arr2[j], arr1[i]) {
					ans = append(ans, arr2[j])
					j++
				} else {
					ans = append(ans, arr1[i])
					i++
				}
			}

		}
		return ans
	}
}

func Mergesort(arr []int64, compare func(a int64, b int64) bool) []int64 {
	if compare == nil {
		compare = func(a int64, b int64) bool {
			return a < b
		}
	}

	ans := sort(arr, compare)

	return ans
}
