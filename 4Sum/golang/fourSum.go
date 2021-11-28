package fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	i, j, k, l := 0, 1, 2, len(nums)-1

	result := [][]int{}
	for i < len(nums) {
		for j < len(nums) {
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})

					tmp := nums[k]
					for k < len(nums) && nums[k] == tmp {
						k++
					}

					tmp = nums[l]
					for l >= 0 && nums[l] == tmp {
						l--
					}
				} else if sum < target {
					tmp := nums[k]
					for k < len(nums) && nums[k] == tmp {
						k++
					}
				} else {
					tmp := nums[l]
					for l >= 0 && nums[l] == tmp {
						l--
					}
				}
			}
			tmp := nums[j]
			for j < len(nums) && nums[j] == tmp {
				j++
			}

			k = j + 1
			l = len(nums) - 1
		}

		tmp := nums[i]
		for i < len(nums) && nums[i] == tmp {
			i++
		}

		j = i + 1
		k = j + 1
		l = len(nums) - 1
	}

	return result
}
