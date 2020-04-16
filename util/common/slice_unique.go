package common

func SliceDuplicateInt(data []int) []int {
	if len(data) < 1024 {
		return duplicateIntByLoop(data)
	}
	return duplicateIntByMap(data)
}

func duplicateIntByLoop(data []int) []int {
	result := make([]int, 0)
	var flag bool
	for i := range data {
		flag = true
		for j := range result {
			if data[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, data[i])
		}
	}
	return result
}

func duplicateIntByMap(data []int) []int {
	result := make([]int, 0)
	temp := make(map[int]struct{}, 0) // 存放不重复数据 空struct不占内存
	for _, v := range data {
		if _, ok := temp[v]; !ok {
			result = append(result, v)
			temp[v] = struct{}{}
		}
	}
	return result
}

func SliceDuplicateInt64(data []int64) []int64 {
	if len(data) < 1024 {
		return duplicateInt64ByLoop(data)
	}
	return duplicateInt64ByMap(data)
}

func duplicateInt64ByLoop(data []int64) []int64 {
	result := make([]int64, 0)
	var flag bool
	for i := range data {
		flag = true
		for j := range result {
			if data[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, data[i])
		}
	}
	return result
}

func duplicateInt64ByMap(data []int64) []int64 {
	result := make([]int64, 0)
	temp := make(map[int64]struct{}, 0) // 存放不重复数据 空struct不占内存
	for _, v := range data {
		if _, ok := temp[v]; !ok {
			result = append(result, v)
			temp[v] = struct{}{}
		}
	}
	return result
}

func SliceDuplicateString(data []string) []string {
	if len(data) < 1024 {
		return duplicateStringByLoop(data)
	}
	return duplicateStringByMap(data)
}

func duplicateStringByLoop(data []string) []string {
	result := make([]string, 0)
	var flag bool
	for i := range data {
		flag = true
		for j := range result {
			if data[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, data[i])
		}
	}
	return result
}

func duplicateStringByMap(data []string) []string {
	result := make([]string, 0)
	temp := make(map[string]struct{}, 0) // 存放不重复数据 空struct不占内存
	for _, v := range data {
		if _, ok := temp[v]; !ok {
			result = append(result, v)
			temp[v] = struct{}{}
		}
	}
	return result
}
