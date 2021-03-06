package search

/*
十一-1
*
递增数组的旋转数组，找到最小值。

二分查找的变种，时间复杂度O(logn)。
*/
func searchMinInRotateSortArr1(arr []int) (index int) {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}

	start := 0
	end := len(arr) - 1
	mid := start
	for arr[start] >= arr[end] { //没有旋转，arr[0]就是min
		if end-start == 1 {
			mid = end
			break
		}

		mid = (end-start)>>1 + start
		if arr[start] == arr[end] && arr[start] == arr[mid] { //首，尾，中全都相等，无法缩小区间，只能顺序遍历
			index = start
			for i := start + 1; i <= end; i++ {
				if arr[i] < arr[index] {
					index = i
				}
			}
			return index
		}

		if arr[mid] >= arr[start] {
			start = mid
		} else {
			end = mid
		}
	}
	return mid
}

//通过索引缩小查找边界，可读性更好
//start，end的三种大小关系，区分出三种大的情况。
//小于：全局递增
//大于：二分讨论，mid与start比较
//等于：二分讨论，mid与start比较；mid与start相等，会出现遍历查找
func searchMinInRotateSortArr2(arr []int) (index int) {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}

	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (end-start)>>1 + start

		if arr[start] < arr[end] { //全局递增数据，相当于没有任何旋转
			return start
		} else if arr[start] > arr[end] {
			if arr[mid] > arr[start] {
				start = mid + 1
			} else if arr[mid] == arr[start] {
				start = mid + 1
			} else {
				if arr[mid-1] > arr[mid] {
					return mid
				} else {
					end = mid - 1
				}
			}
		} else {
			if arr[mid] > arr[start] {
				start = mid + 1
			} else if arr[mid] == arr[start] { //首，中，尾三个相等，无法缩小边界，遍历查找
				index = start
				for i := start + 1; i <= end; i++ {
					if arr[i] < arr[index] {
						index = i
					}
				}
				return index
			} else {
				if arr[mid-1] > arr[mid] {
					return mid
				} else {
					end = mid - 1
				}
			}
		}
	}

	return -1
}
