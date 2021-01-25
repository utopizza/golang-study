package sorting

// HeapSort 堆排序
func HeapSort(arr []int) {
	// 原数组元素复制到一个辅助数组，第一个元素下标为1
	N := len(arr)
	helper := make([]int, N+1)
	for i := range arr {
		helper[i+1] = arr[i]
	}

	// 调整非叶子结点，保证每个节点比左右孩子都大，即构建最大堆
	// 构建完成后，堆顶为数组最大元素
	for k := N / 2; k > 1; k-- {
		sink(helper, k, N)
	}

	// 排序、拆堆
	for N > 1 {
		// 堆顶元素出堆，放置到数组末尾，即与某个叶子节点交换
		swap(helper, 1, N)

		// 既然堆顶元素已经出堆，那么堆大小减1
		N--

		// 由于此时堆顶是一个小节点，因此需要重新调整堆，将该节点下沉到正确位置
		sink(helper, 1, N)
	}

	// 拷贝回原数组
	for i := range arr {
		arr[i] = helper[i+1]
	}
}

func sink(arr []int, k, N int) {
	for k*2 <= N {
		// 找第k元素的左右孩子中，较大的那个
		child := k * 2
		if child+1 <= N && arr[child+1] > arr[child] {
			child++
		}

		// 若第k元素已经比最大孩子大，即堆有序，结束调整
		if arr[k] >= arr[child] {
			break
		}

		// 否则，进行交换，保证每个节点比其左右孩子大
		swap(arr, k, child)

		// 继续往下调整
		k = child
	}
}
