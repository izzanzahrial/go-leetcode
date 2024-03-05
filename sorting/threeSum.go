package sorting

import (
	"math/rand"
	"sort"
	"time"
)

// https://leetcode.com/problems/3sum/submissions/

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			threeSum := num + nums[left] + nums[right]
			if threeSum > 0 {
				right--
			} else if threeSum < 0 {
				left++
			} else {
				result = append(result, []int{num, nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}

		}
	}

	return result
}

type RandomizedSet struct {
    set map[int]bool
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        set: make(map[int]bool),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if ok := this.set[val]; !ok {
        this.set[val] = true
        return true
    }

    return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if ok := this.set[val]; ok {
        delete(this.set, val)
		return true
    }

	return false
}


func (this *RandomizedSet) GetRandom() int {
    rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(this.set))
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */