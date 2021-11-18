package bisect

// 合っているかわからない

// 探査対象リスト array の中にすでに key が存在する場合、既存の key の左側の位置を挿入箇所として返す
func BisectLeft(array []int, key int) int {
	var left, right, middle int
	left = 0
	right = len(array)
	for left < right {
		middle = ((right - left) / 2) + left
		if key < array[middle] {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
