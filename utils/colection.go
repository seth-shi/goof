package utils

func Intersect(nums1 []string, nums2 []string) []string {

	m := make(map[string]int)

	for _, v := range nums2 {
		m[v]++
	}

	var diff []string
	for _, v := range nums1 {

		if _, exists := m[v]; !exists {

			diff = append(diff, v)
		}
	}

	return diff
}
