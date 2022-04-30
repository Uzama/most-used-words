package processor

// sort the array (max heap sort) and return max 10 items
func HeapSort(arr []WordCount) []WordCount {

	// heapify the given array (max heap)
	heapify(arr)

	// remove first 10 items
	length := len(arr) - 10

	// if the item counts less than or equal to 10, then remove all item
	if len(arr) <= 10 {
		length = 0
	}

	sortedArray := make([]WordCount, 0)

	for i := len(arr) - 1; i >= length; i-- {

		// swap root with last item
		arr[0], arr[i] = arr[i], arr[0]

		// heapify the swaped item to get correct position
		heapifyDown(arr, 0, i)

		// push removed (max) item into sorted array
		sortedArray = append(sortedArray, arr[i])
	}

	return sortedArray
}

// heapify each element (first half is enough, since second half is children of first half)
func heapify(arr []WordCount) {
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		heapifyDown(arr, i, len(arr))
	}
}

// heapify down the element til the element get correct position
func heapifyDown(heap []WordCount, low, high int) {

	root := low

	for {
		// compare with left child
		child := root*2 + 1

		if child >= high {
			break
		}

		// check whether right child is greater than left child
		if child+1 < high && heap[child].Count < heap[child+1].Count {
			// if right child is greater than left, then compare root with right child
			child++
		}

		// compare with greatest child (left or right)
		if heap[root].Count < heap[child].Count {
			// swap the root with the child (left or right)
			heap[root], heap[child] = heap[child], heap[root]
			root = child

			continue
		}

		break
	}
}
