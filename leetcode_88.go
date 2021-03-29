func merge(nums1 []int, m int, nums2 []int, n int) {

	//排完nums2後剩下的nums1會是正確的排序
	for n > 0 {
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}

//Time Complexity: O(N)
//Space Complexity: O(1)