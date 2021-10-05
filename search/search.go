package search

/*
1. 循环退出条件
注意是 low<=high，而不是 low<high.
因为若数据的长度为偶数，查找操作的倒数第二步，low会等于high，故比较运算符选择 <= 。 不然，查找操作将被中断。

2.mid 的取值
实际上，mid=(low+high)/2 这种写法是有问题的。因为如果 low 和 high 比较大的话，两者之和就有可能会溢出。
改进的方法是将 mid 的计算方式写成 low+(high-low)/2。
更进一步，如果要将性能优化到极致的话，我们可以将这里的除以 2 操作转化成位运算 low+((high-low)>>1)。
因为相比除法运算来说，计算机处理位运算要快得多。

3.low 和 high 的更新
low=mid+1，high=mid-1。注意这里的 +1 和 -1，如果直接写成 low=mid 或者 high=mid，就可能会发生死循环。
比如，当 high=3，low=3 时，如果 a[3]不等于 value，就会导致一直循环不退出。
*/

//非递归实现二分查找
func BSearch1(inputArr []int, val int) (index int) {
	low := 0
	high := len(inputArr) - 1

	for low <= high {
		mid := (low + high) / 2
		if inputArr[mid] == val {
			return mid
		} else if inputArr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//递归实现二分查找
func BSearch2(inputArr []int, val int) (index int) {
	var search func(low, high int) int

	search = func(low, high int) (index int) {
		if low > high {
			return -1
		}
		mid := low + ((high - low) >> 1)
		if inputArr[mid] == val {
			return mid
		} else if inputArr[mid] < val {
			return search(mid+1, high)
		} else {
			return search(low, mid-1)
		}
	}
	return search(0, len(inputArr)-1)
}
