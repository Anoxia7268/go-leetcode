/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (42.57%)
 * Likes:    7228
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.7M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */
package sf

import "sort"

// @lc code=start
func findMedianSortedArrays4(nums1 []int, nums2 []int) float64 {
	a := len(nums1)
	b := len(nums2)
	m := make([]int, a+b)
	if a != 0 {
		for i := 0; i < a; i++ {
			m[i] = nums1[i]
		}
	}
	if b != 0 {
		for i := 0; i < b; i++ {
			m[i+a] = nums2[i]
		}
	}
	sort.Ints(m)
	if a+b < 2 {
		return float64(m[0])
	}
	if (a+b)%2 != 0 {
		return float64(m[(a+b)/2])
	}
	return float64((m[(a+b)/2-1] + m[(a+b)/2])) / 2
}

// @lc code=end
