package Solution

// 解法一
/*
 遍历的粗暴解法：
	第一个遍历拿个数，第二个遍历剩下的数
	比较，直到两个数刚好等于给定的这个数就算是成功
	时间复杂度：O（n^2）
*/
func TwoSum1(nums []int, target int) []int {
	
	// 获取数据的长度
	n := len(nums)

	// 遍历数组
	for i, v := range nums {
		// 从剩下的数里，找，是否能够等于目标值的
		for j := i + 1; j < n; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}


// 解法二
/**
	
 */
func TwoSum2(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}


// 解法三
func TwoSum3(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{-1, -1}
	}

	res := []int{-1, -1}

	//	MAP的KEY表示值，MAP的VAL表示nums的下标
	intMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		//	判断MAP中是否纯在一个Key满足 KEY + nums[i] = target
		//	如果满足则返回相关地址,不满足则将数组中的值PUSH到MAP
		if _, ok := intMap[target-nums[i]]; ok {
			res[0] = intMap[target-nums[i]]
			res[1] = i
			break
		}
		intMap[nums[i]] = i
	}
	return res
}
