package _5_quick

//边切分，边确保大致有序（左边小于中间元素，右边大于中间元素）
//最后从最小的单位逐级返回，每次返回都是有序的
//原地排序版快排，即空间复杂度是O(1)
func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	quickSort(arr, 0, length-1)
	return
}

//p:数组下标起点, r:数组下标终点
func quickSort(arr []int, p, r int) {
	if p >= r {
		return
	}
	index := partition(arr, p, r)
	quickSort(arr, p, index-1)
	quickSort(arr, index+1, r)
}

//空间复杂度O(1)，关键就在该分区函数中。比pivot小，才与索引i的元素交换。循环跳出时，索引比i小的元素，都比pivot小。
func partition(arr []int, p, r int) (index int) {
	pivot := arr[r]
	//注意：i，j不同步。循环结束，i就是arr[r]的位置
	index = p
	for i := p; i < r; i++ {
		if arr[i] < pivot {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index], arr[r] = arr[r], arr[index]
	return index
}

/*O
(n)时间复杂度内，在无序数组中，查找第K大元素.
第一次分区查找，我们需要对大小为 n 的数组执行分区操作，需要遍历 n 个元素。
第二次分区查找，我们只需要对大小为 n/2 的数组执行分区操作，需要遍历 n/2 个元素。
依次类推，分区遍历元素的个数分别为、n/2、n/4、n/8、n/16.……直到区间缩小为 1。

求解等比数列：
该数列长度未知，n表示的是原数组的长度，不是本数列的长度
n+n/2+n/4+n/8+...+1

等比数列求和公式：
[a1 * (1 - q^n) ]/ (1 - q)
其中，n为元素个数

思路1：[推荐]
不关注该数列的长度，因为明确知道最后一个元素是1.
所以，[a1 * (1 - q^m) ]/ (1 - q) = (a1 - am * q) / (1 - q) = (n - 1/2）* 2 = 2n - 1

思路2：
关注该数列的长度，严格按照等比数列求和公式运算。
数列长度k=log2n + 1
所以，n * (1 - (1/2) ^ (log2n + 1)) / (1/2) = n * (1 - 1/2n) / (1/2) = 2n - 1
*/

func QuickSearch(arr []int, k int) (element int) {
	n := len(arr)
	if n < k {
		return -1
	}

	//初始分区为整个数组
	low := 0
	up := n - 1
	for {
		pivot := arr[up]
		//注意：i，j不同步。循环结束，i就是arr[up]的位置
		i := low
		for j := low; j <= up-1; j++ {
			if arr[j] < pivot {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
		arr[i], arr[up] = arr[up], arr[i]

		//第k大，即索引为k-1。
		if i == k-1 { //循环退出的情况
			return arr[i]
		} else if i > k-1 { //k出现在左半边，减小上限，继续搜索
			up = i - 1
		} else { //k出现在右半边，增大下限，继续搜索
			low = i + 1
		}
	}
}
