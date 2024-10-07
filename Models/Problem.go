package Models

type Problem struct {
	ID           uint       `gorm:"primaryKey"`
	Category     string     `json:"Category"`
	Type         string     `json:"type"`
	StartingCode string     `json:"startingCode"`
	Description  string     `json:"description"`
	TestCases    []TestCase `gorm:"foreignKey:ProblemID"`
}

//func init() {
//	CreateProblem()
//}
//func CreateProblem() {
//	//problems := []Problem{
//	//	{
//	//		Type:         "Addition",
//	//		StartingCode: "function value(a, b) { return a + b; }",
//	//		Description:  "Write a function that takes two numbers and returns their sum.",
//	//		TestCases: []TestCase{
//	//			{Input: `[5, 10]`, Expected: `15`},
//	//			{Input: `[2, 3]`, Expected: `5`},
//	//			{Input: `[-1, -1]`, Expected: `-2`},
//	//			{Input: `[0, 0]`, Expected: `0`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Two Sum",
//	//		StartingCode: "function twoSum(nums, target) { /* implementation */ }",
//	//		Description:  "Given an array of integers, return indices of the two numbers such that they add up to a specific target.",
//	//		TestCases: []TestCase{
//	//			{Input: `[2, 7, 11, 15], target = 9`, Expected: `[0, 1]`},
//	//			{Input: `[3, 2, 4], target = 6`, Expected: `[1, 2]`},
//	//			{Input: `[3, 3], target = 6`, Expected: `[0, 1]`},
//	//			{Input: `[1, 5, 3, 6], target = 8`, Expected: `[1, 2]`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Remove Duplicates from Sorted Array",
//	//		StartingCode: "function removeDuplicates(nums) { /* implementation */ }",
//	//		Description:  "Given a sorted array, remove the duplicates in-place such that each element appears only once and return the new length.",
//	//		TestCases: []TestCase{
//	//			{Input: `[1, 1, 2]`, Expected: `2`},
//	//			{Input: `[0,0,1,1,2]`, Expected: `3`},
//	//			{Input: `[1,2,2,3,4]`, Expected: `4`},
//	//			{Input: `[5,5,5,5]`, Expected: `1`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Maximum Subarray",
//	//		StartingCode: "function maxSubArray(nums) { /* implementation */ }",
//	//		Description:  "Find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.",
//	//		TestCases: []TestCase{
//	//			{Input: `[-2,1,-3,4,-1,2,1,-5,4]`, Expected: `6`},
//	//			{Input: `[1]`, Expected: `1`},
//	//			{Input: `[5,4,-1,7,8]`, Expected: `23`},
//	//			{Input: `[-1,-2,-3,-4]`, Expected: `-1`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Best Time to Buy and Sell Stock",
//	//		StartingCode: "function maxProfit(prices) { /* implementation */ }",
//	//		Description:  "Given an array where the i-th element is the price of a given stock on the i-th day, find the maximum profit you can achieve.",
//	//		TestCases: []TestCase{
//	//			{Input: `[7,1,5,3,6,4]`, Expected: `5`},
//	//			{Input: `[7,6,4,3,1]`, Expected: `0`},
//	//			{Input: `[1,2,3,4,5]`, Expected: `4`},
//	//			{Input: `[3,2,6,5,0,0]`, Expected: `4`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Product of Array Except Self",
//	//		StartingCode: "function productExceptSelf(nums) { /* implementation */ }",
//	//		Description:  "Given an array nums of length n, return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].",
//	//		TestCases: []TestCase{
//	//			{Input: `[1,2,3,4]`, Expected: `[24,12,8,6]`},
//	//			{Input: `[0, 1]`, Expected: `[1, 0]`},
//	//			{Input: `[4, 5, 1, 8]`, Expected: `[40, 32, 160, 20]`},
//	//			{Input: `[1, 1, 1, 1]`, Expected: `[1, 1, 1, 1]`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Rotate Array",
//	//		StartingCode: "function rotate(nums, k) { /* implementation */ }",
//	//		Description:  "Given an array, rotate the array to the right by k steps, where k is non-negative.",
//	//		TestCases: []TestCase{
//	//			{Input: `[1,2,3,4,5,6,7], k = 3`, Expected: `[5,6,7,1,2,3,4]`},
//	//			{Input: `[1,2,3,4,5], k = 2`, Expected: `[4,5,1,2,3]`},
//	//			{Input: `[0,1,2], k = 4`, Expected: `[2,0,1]`},
//	//			{Input: `[1], k = 1`, Expected: `[1]`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Find Minimum in Rotated Sorted Array",
//	//		StartingCode: "function findMin(nums) { /* implementation */ }",
//	//		Description:  "Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand. Find the minimum element.",
//	//		TestCases: []TestCase{
//	//			{Input: `[3,4,5,1,2]`, Expected: `1`},
//	//			{Input: `[4,5,6,7,0,1,2]`, Expected: `0`},
//	//			{Input: `[1]`, Expected: `1`},
//	//			{Input: `[2,1]`, Expected: `1`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Merge Intervals",
//	//		StartingCode: "function merge(intervals) { /* implementation */ }",
//	//		Description:  "Given a collection of intervals, merge all overlapping intervals.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[1,3],[2,6],[8,10],[15,18]]`, Expected: `[[1,6],[8,10],[15,18]]`},
//	//			{Input: `[[1,4],[4,5]]`, Expected: `[[1,5]]`},
//	//			{Input: `[[1,2],[3,4],[5,6]]`, Expected: `[[1,2],[3,4],[5,6]]`},
//	//			{Input: `[[1,4],[2,3],[6,8],[7,9]]`, Expected: `[[1,4],[6,9]]`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Contains Duplicate",
//	//		StartingCode: "function containsDuplicate(nums) { /* implementation */ }",
//	//		Description:  "Given an integer array, return true if any value appears at least twice in the array, and return false if every element is distinct.",
//	//		TestCases: []TestCase{
//	//			{Input: `[1,2,3,1]`, Expected: `true`},
//	//			{Input: `[1,2,3,4]`, Expected: `false`},
//	//			{Input: `[1,1,1,3,3,4,3,2,4,2]`, Expected: `true`},
//	//			{Input: `[]`, Expected: `false`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Single Number",
//	//		StartingCode: "function singleNumber(nums) { /* implementation */ }",
//	//		Description:  "Given a non-empty array of integers, every element appears twice except for one. Find that single one.",
//	//		TestCases: []TestCase{
//	//			{Input: `[2,2,1]`, Expected: `1`},
//	//			{Input: `[4,1,2,1,2]`, Expected: `4`},
//	//			{Input: `[1]`, Expected: `1`},
//	//			{Input: `[1,2,3,2,1]`, Expected: `3`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//	{
//	//		Type:         "Intersection of Two Arrays II",
//	//		StartingCode: "function intersect(nums1, nums2) { /* implementation */ }",
//	//		Description:  "Given two arrays, write a function to compute their intersection.",
//	//		TestCases: []TestCase{
//	//			{Input: `nums1 = [1,2,2,1], nums2 = [2,2]`, Expected: `[2,2]`},
//	//			{Input: `nums1 = [4,9,5], nums2 = [9,4,9,8,4]`, Expected: `[4,9]`},
//	//			{Input: `nums1 = [1,1,2], nums2 = [1,1]`, Expected: `[1,1]`},
//	//			{Input: `nums1 = [1,2,3], nums2 = [4,5,6]`, Expected: `[]`},
//	//		},
//	//		Category: "Array",
//	//	},
//	//}
//	//
//	//problems := []Problem{
//	//	{
//	//		Type:         "Valid Anagram",
//	//		StartingCode: "function isAnagram(s, t) { /* implementation */ }",
//	//		Description:  "Given two strings s and t, return true if t is an anagram of s, and false otherwise.",
//	//		TestCases: []TestCase{
//	//			{Input: `("anagram", "nagaram")`, Expected: `true`},
//	//			{Input: `("rat", "car")`, Expected: `false`},
//	//			{Input: `("listen", "silent")`, Expected: `true`},
//	//			{Input: `("hello", "world")`, Expected: `false`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Longest Substring Without Repeating Characters",
//	//		StartingCode: "function lengthOfLongestSubstring(s) { /* implementation */ }",
//	//		Description:  "Given a string s, find the length of the longest substring without repeating characters.",
//	//		TestCases: []TestCase{
//	//			{Input: `("abcabcbb")`, Expected: `3`},
//	//			{Input: `("bbbbb")`, Expected: `1`},
//	//			{Input: `("pwwkew")`, Expected: `3`},
//	//			{Input: `("")`, Expected: `0`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Reverse String",
//	//		StartingCode: "function reverseString(s) { /* implementation */ }",
//	//		Description:  "Write a function that reverses a string. The input string is given as an array of characters.",
//	//		TestCases: []TestCase{
//	//			{Input: `(["h","e","l","l","o"])`, Expected: `["o","l","l","e","h"]`},
//	//			{Input: `(["H","a","n","n","a","h"])`, Expected: `["h","a","n","n","a","H"]`},
//	//			{Input: `(["a","b","c"])`, Expected: `["c","b","a"]`},
//	//			{Input: `(["r","e","v","e","r","s","e"])`, Expected: `["e","s","r","e","v","e","r"]`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "String to Integer (atoi)",
//	//		StartingCode: "function myAtoi(s) { /* implementation */ }",
//	//		Description:  "Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.",
//	//		TestCases: []TestCase{
//	//			{Input: `("42")`, Expected: `42`},
//	//			{Input: `("   -42")`, Expected: `-42`},
//	//			{Input: `("4193 with words")`, Expected: `4193`},
//	//			{Input: `("words and 987")`, Expected: `0`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Palindrome String",
//	//		StartingCode: "function isPalindrome(s) { /* implementation */ }",
//	//		Description:  "Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.",
//	//		TestCases: []TestCase{
//	//			{Input: `("A man, a plan, a canal: Panama")`, Expected: `true`},
//	//			{Input: `("race a car")`, Expected: `false`},
//	//			{Input: `(" ")`, Expected: `true`},
//	//			{Input: `("No 'x' in Nixon")`, Expected: `true`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Group Anagrams",
//	//		StartingCode: "function groupAnagrams(strs) { /* implementation */ }",
//	//		Description:  "Given an array of strings strs, group the anagrams together. You can return the answer in any order.",
//	//		TestCases: []TestCase{
//	//			{Input: `(["eat","tea","tan","ate","nat","bat"])`, Expected: `[["bat"],["nat","tan"],["ate","eat","tea"]]`},
//	//			{Input: `([""])`, Expected: `[[""]]`},
//	//			{Input: `(["a"])`, Expected: `[["a"]]`},
//	//			{Input: `(["abcd","bcda","adcb","dcba"])`, Expected: `[["abcd","bcda","dcba","adcb"]]`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Implement strStr()",
//	//		StartingCode: "function strStr(haystack, needle) { /* implementation */ }",
//	//		Description:  "Implement strStr(). Given two strings haystack and needle, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.",
//	//		TestCases: []TestCase{
//	//			{Input: `("hello", "ll")`, Expected: `2`},
//	//			{Input: `("aaaaa", "bba")`, Expected: `-1`},
//	//			{Input: `("mississippi", "issip")`, Expected: `4`},
//	//			{Input: `("abc", "")`, Expected: `0`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Count and Say",
//	//		StartingCode: "function countAndSay(n) { /* implementation */ }",
//	//		Description:  "The count-and-say sequence is a sequence of digit strings defined by the recursive formula. Given an integer n, generate the nth term of the sequence.",
//	//		TestCases: []TestCase{
//	//			{Input: `(1)`, Expected: `"1"`},
//	//			{Input: `(4)`, Expected: `"1211"`},
//	//			{Input: `(5)`, Expected: `"111221"`},
//	//			{Input: `(6)`, Expected: `"312211"`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Longest Palindromic Substring",
//	//		StartingCode: "function longestPalindrome(s) { /* implementation */ }",
//	//		Description:  "Given a string s, return the longest palindromic substring in s.",
//	//		TestCases: []TestCase{
//	//			{Input: `("babad")`, Expected: `"bab"`},
//	//			{Input: `("cbbd")`, Expected: `"bb"`},
//	//			{Input: `("a")`, Expected: `"a"`},
//	//			{Input: `("ac")`, Expected: `"a"`},
//	//		},
//	//		Category: "String",
//	//	},
//	//	{
//	//		Type:         "Replace All Spaces",
//	//		StartingCode: "function replaceSpaces(s) { /* implementation */ }",
//	//		Description:  "Given a string, replace all spaces with '%20'.",
//	//		TestCases: []TestCase{
//	//			{Input: `("Mr John Smith    ", 13)`, Expected: `"Mr%20John%20Smith"`},
//	//			{Input: `("Hello World  ", 11)`, Expected: `"Hello%20World"`},
//	//			{Input: `("   Leading and trailing spaces    ", 38)`, Expected: `"Leading%20and%20trailing%20spaces"`},
//	//			{Input: `("SingleSpace", 12)`, Expected: `"SingleSpace"`},
//	//		},
//	//		Category: "String",
//	//	},
//	//}
//	//problems := []Problem{
//	//	{
//	//		Type:         "Maximum Depth of Binary Tree",
//	//		StartingCode: "function maxDepth(root) { /* implementation */ }",
//	//		Description:  "Given a binary tree, return its maximum depth. The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.",
//	//		TestCases: []TestCase{
//	//			{Input: `([3,9,20,null,null,15,7])`, Expected: `3`},
//	//			{Input: `([1,null,2])`, Expected: `2`},
//	//			{Input: `([])`, Expected: `0`},
//	//			{Input: `([1])`, Expected: `1`},
//	//		},
//	//		Category: "Tree",
//	//	},
//	//	{
//	//		Type:         "Binary Tree Inorder Traversal",
//	//		StartingCode: "function inorderTraversal(root) { /* implementation */ }",
//	//		Description:  "Given the root of a binary tree, return its inorder traversal as an array of node values.",
//	//		TestCases: []TestCase{
//	//			{Input: `([1,null,2,3])`, Expected: `[1,3,2]`},
//	//			{Input: `([1])`, Expected: `[1]`},
//	//			{Input: `([])`, Expected: `[]`},
//	//			{Input: `([4,2,5,1,3])`, Expected: `[1,2,3,4,5]`},
//	//		},
//	//		Category: "Tree",
//	//	},
//	//	{
//	//		Type:         "Lowest Common Ancestor of a Binary Tree",
//	//		StartingCode: "function lowestCommonAncestor(root, p, q) { /* implementation */ }",
//	//		Description:  "Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.",
//	//		TestCases: []TestCase{
//	//			{Input: `([3,5,1,6,2,0,8], 5, 1)`, Expected: `3`},
//	//			{Input: `([3,5,1,6,2,0,8], 5, 4)`, Expected: `null`},
//	//			{Input: `([1,2,3], 2, 3)`, Expected: `1`},
//	//			{Input: `([1,2], 1, 2)`, Expected: `1`},
//	//		},
//	//		Category: "Tree",
//	//	},
//	//	{
//	//		Type:         "Binary Tree Level Order Traversal",
//	//		StartingCode: "function levelOrder(root) { /* implementation */ }",
//	//		Description:  "Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).",
//	//		TestCases: []TestCase{
//	//			{Input: `([3,9,20,null,null,15,7])`, Expected: `[[3],[9,20],[15,7]]`},
//	//			{Input: `([1])`, Expected: `[[1]]`},
//	//			{Input: `([])`, Expected: `[]`},
//	//			{Input: `([1,2,3,4])`, Expected: `[[1],[2,3],[4]]`},
//	//		},
//	//		Category: "Tree",
//	//	},
//	//	{
//	//		Type:         "Symmetric Tree",
//	//		StartingCode: "function isSymmetric(root) { /* implementation */ }",
//	//		Description:  "Given a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).",
//	//		TestCases: []TestCase{
//	//			{Input: `([1,2,2,3,4,4,3])`, Expected: `true`},
//	//			{Input: `([1,2,2,null,3,null,3])`, Expected: `false`},
//	//			{Input: `([])`, Expected: `true`},
//	//			{Input: `([1])`, Expected: `true`},
//	//		},
//	//		Category: "Tree",
//	//	},
//	//	{
//	//		Type:         "Diameter of Binary Tree",
//	//		StartingCode: "function diameterOfBinaryTree(root) { /* implementation */ }",
//	//		Description:  "Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree.",
//	//		TestCases: []TestCase{
//	//			{Input: `([1,2,3,4,5])`, Expected: `3`},
//	//			{Input: `([1])`, Expected: `0`},
//	//			{Input: `([1,2])`, Expected: `1`},
//	//			{Input: `([1,2,3,4,5,6])`, Expected: `4`},
//	//		},
//	//		Category: "Tree",
//	//	},
//	//	{
//	//		Type:         "Construct Binary Tree from Preorder and Inorder Traversal",
//	//		StartingCode: "function buildTree(preorder, inorder) { /* implementation */ }",
//	//		Description:  "Given two integer arrays, preorder and inorder, where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.",
//	//		TestCases: []TestCase{
//	//			{Input: `([3,9,20,15,7], [9,3,15,20,7])`, Expected: `[3,9,20,null,null,15,7]`},
//	//			{Input: `([1,2], [2,1])`, Expected: `[1,2]`},
//	//			{Input: `([1], [1])`, Expected: `[1]`},
//	//			{Input: `([1,2,3], [2,1,3])`, Expected: `[1,2,3]`},
//	//		},
//	//		Category: "Tree",
//	//	},
//	//}
//	//problems := []Problem{
//	//	{
//	//		Type:         "Valid Parentheses",
//	//		StartingCode: "function isValid(s) { /* implementation */ }",
//	//		Description:  "Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. An input string is valid if the brackets are closed in the correct order.",
//	//		TestCases: []TestCase{
//	//			{Input: `"()"`, Expected: `true`},
//	//			{Input: `"(]"`, Expected: `false`},
//	//			{Input: `"([)]"`, Expected: `false`},
//	//			{Input: `"{[]}"`, Expected: `true`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Min Stack",
//	//		StartingCode: "function MinStack() { /* implementation */ }",
//	//		Description:  "Design a stack that supports push, pop, top, and retrieving the minimum element in constant time. Implement the MinStack class.",
//	//		TestCases: []TestCase{
//	//			{Input: `["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]`, Expected: `[-3,0,-2]`},
//	//			{Input: `["MinStack","push","push","getMin","pop","top","getMin"]\n[[],[1],[2],[],[],[],[]]`, Expected: `[1]`},
//	//			{Input: `["MinStack","push","getMin"]\n[[],[3],[]]`, Expected: `[3]`},
//	//			{Input: `["MinStack","push","push","getMin","pop","top","getMin"]\n[[],[1],[0],[],[],[],[]]`, Expected: `[0]`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Evaluate Reverse Polish Notation",
//	//		StartingCode: "function evalRPN(tokens) { /* implementation */ }",
//	//		Description:  "Evaluate the value of an arithmetic expression in Reverse Polish Notation. Valid operators are '+', '-', '*', and '/'.",
//	//		TestCases: []TestCase{
//	//			{Input: `["2","1","+","3","*"]`, Expected: `9`},
//	//			{Input: `["4","13","5","/","+"]`, Expected: `6`},
//	//			{Input: `["10","6","9","3","-11","*","/","*","17","+","5","+"]`, Expected: `22`},
//	//			{Input: `["18","5","4","*","+"]`, Expected: `38`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Daily Temperatures",
//	//		StartingCode: "function dailyTemperatures(temperatures) { /* implementation */ }",
//	//		Description:  "Given an array of integers temperatures representing daily temperatures, return an array such that answer[i] is the number of days you have to wait until a warmer temperature. If there is no future day for which this is possible, keep answer[i] = 0 instead.",
//	//		TestCases: []TestCase{
//	//			{Input: `[73,74,75,71,69,72,76,73]`, Expected: `[1,1,4,2,1,1,0,0]`},
//	//			{Input: `[30,40,50,60]`, Expected: `[1,1,1,0]`},
//	//			{Input: `[30,60,90]`, Expected: `[1,1,0]`},
//	//			{Input: `[50,40,30,20]`, Expected: `[0,0,0,0]`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Largest Rectangle in Histogram",
//	//		StartingCode: "function largestRectangleArea(heights) { /* implementation */ }",
//	//		Description:  "Given an array of integers heights representing the histogram's bar heights, return the area of the largest rectangle that can be formed in the histogram.",
//	//		TestCases: []TestCase{
//	//			{Input: `[2,1,5,6,2,3]`, Expected: `10`},
//	//			{Input: `[2,4]`, Expected: `4`},
//	//			{Input: `[1,1,1,1,1]`, Expected: `5`},
//	//			{Input: `[6,2,5,4,5,1,6]`, Expected: `12`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Next Greater Element I",
//	//		StartingCode: "function nextGreaterElement(nums1, nums2) { /* implementation */ }",
//	//		Description:  "You are given two arrays, nums1 and nums2. For each element in nums1, find the next greater element in nums2. If there is no greater element, return -1.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[4],[1,2,1,4,3,2]]`, Expected: `[-1]`},
//	//			{Input: `[[2,4],[1,2,3,4]]`, Expected: `[3,-1]`},
//	//			{Input: `[[1,3,2],[2,1,3,1,0]]`, Expected: `[3,-1]`},
//	//			{Input: `[[5],[5,4,3,2,1]]`, Expected: `[-1]`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Remove K Digits",
//	//		StartingCode: "function removeKdigits(num, k) { /* implementation */ }",
//	//		Description:  "Given a string num representing a non-negative integer, remove k digits from the number to minimize its value.",
//	//		TestCases: []TestCase{
//	//			{Input: `["1432219", "3"]`, Expected: `"1219"`},
//	//			{Input: `["10200", "1"]`, Expected: `"200"`},
//	//			{Input: `["10", "2"]`, Expected: `"0"`},
//	//			{Input: `["123456789", "9"]`, Expected: `"1"`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Basic Calculator II",
//	//		StartingCode: "function calculate(s) { /* implementation */ }",
//	//		Description:  "Implement a basic calculator to evaluate a simple expression string containing non-negative integers, '+', '-', '*', and '/'.",
//	//		TestCases: []TestCase{
//	//			{Input: `"3+2*2"`, Expected: `7`},
//	//			{Input: `" 3/2 "`, Expected: `1`},
//	//			{Input: `" 3+5 "`, Expected: `8`},
//	//			{Input: `" 14-3*2 "`, Expected: `8`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Trapping Rain Water",
//	//		StartingCode: "function trap(height) { /* implementation */ }",
//	//		Description:  "Given n non-negative integers representing the height of bars in a histogram, compute how much water it can trap after raining.",
//	//		TestCases: []TestCase{
//	//			{Input: `[0,1,0,2,1,0,1,3,2,1,2,1]`, Expected: `6`},
//	//			{Input: `[4,2,0,3,2,5]`, Expected: `9`},
//	//			{Input: `[1,0,2]`, Expected: `1`},
//	//			{Input: `[2,0,2]`, Expected: `2`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Score of Parentheses",
//	//		StartingCode: "function scoreOfParentheses(s) { /* implementation */ }",
//	//		Description:  "Given a balanced parentheses string s, return the score of the string based on the following rules: '()' has a score of 1, and 'AB' has a score of A + B, while '(A)' has a score of 2 * A.",
//	//		TestCases: []TestCase{
//	//			{Input: `"()"`, Expected: `1`},
//	//			{Input: `"(())"`, Expected: `2`},
//	//			{Input: `"(()(()))"`, Expected: `6`},
//	//			{Input: `"()()"`, Expected: `2`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Binary Tree Postorder Traversal",
//	//		StartingCode: "function postorderTraversal(root) { /* implementation */ }",
//	//		Description:  "Given the root of a binary tree, return the postorder traversal of its nodes' values.",
//	//		TestCases: []TestCase{
//	//			{Input: `[1,null,2,3]`, Expected: `[3,2,1]`},
//	//			{Input: `[1]`, Expected: `[1]`},
//	//			{Input: `[1,2]`, Expected: `[2,1]`},
//	//			{Input: `[1,null,2]`, Expected: `[2,1]`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//	{
//	//		Type:         "Binary Tree Inorder Traversal",
//	//		StartingCode: "function inorderTraversal(root) { /* implementation */ }",
//	//		Description:  "Given the root of a binary tree, return the inorder traversal of its nodes' values.",
//	//		TestCases: []TestCase{
//	//			{Input: `[1,null,2,3]`, Expected: `[1,3,2]`},
//	//			{Input: `[1]`, Expected: `[1]`},
//	//			{Input: `[1,2]`, Expected: `[2,1]`},
//	//			{Input: `[1,null,2]`, Expected: `[1,2]`},
//	//		},
//	//		Category: "Stack",
//	//	},
//	//}
//	//problems := []Problem{
//	//	{
//	//		Type:         "Implement Queue using Stacks",
//	//		StartingCode: "function MyQueue() { /* implementation */ }",
//	//		Description:  "Implement a first-in-first-out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, pop, peek, and empty).",
//	//		TestCases: []TestCase{
//	//			{Input: `["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]`, Expected: `[1,1,false]`},
//	//			{Input: `["MyQueue","push","pop","empty"]\n[[],[1],[],[]]`, Expected: `[1,false]`},
//	//			{Input: `["MyQueue","push","push","peek","pop","empty"]\n[[],[3],[4],[],[],[]]`, Expected: `[3,3,false]`},
//	//			{Input: `["MyQueue","push","pop","push","pop","empty"]\n[[],[5],[],[6],[],[]]`, Expected: `[5,false]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Implement Stack using Queues",
//	//		StartingCode: "function MyStack() { /* implementation */ }",
//	//		Description:  "Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, pop, top, and empty).",
//	//		TestCases: []TestCase{
//	//			{Input: `["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]`, Expected: `[2,2,false]`},
//	//			{Input: `["MyStack","push","pop","empty"]\n[[],[1],[],[]]`, Expected: `[1,false]`},
//	//			{Input: `["MyStack","push","push","top","pop","empty"]\n[[],[3],[4],[],[],[]]`, Expected: `[4,4,false]`},
//	//			{Input: `["MyStack","push","pop","push","pop","empty"]\n[[],[5],[],[6],[],[]]`, Expected: `[6,false]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Sliding Window Maximum",
//	//		StartingCode: "function maxSlidingWindow(nums, k) { /* implementation */ }",
//	//		Description:  "You are given an array of integers nums, and an integer k. Find the maximum value in each sliding window of size k.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[1,3,-1,-3,5,3,6,7],3]`, Expected: `[3,3,5,5,6,7]`},
//	//			{Input: `[[1],1]`, Expected: `[1]`},
//	//			{Input: `[[1,-1],1]`, Expected: `[1,-1]`},
//	//			{Input: `[[9,10,3,2,5],3]`, Expected: `[10,10,5]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Number of Islands",
//	//		StartingCode: "function numIslands(grid) { /* implementation */ }",
//	//		Description:  "Given a 2D grid of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.",
//	//		TestCases: []TestCase{
//	//			{Input: `[['1','1','1','1','0'],['1','1','0','0','0'],['1','0','1','0','0'],['0','0','0','0','0']]`, Expected: `2`},
//	//			{Input: `[['1','0','1','1','0'],['0','0','0','1','0'],['1','1','1','0','0']]`, Expected: `1`},
//	//			{Input: `[['0','0','0','0'],['0','0','0','0']]`, Expected: `0`},
//	//			{Input: `[['1']]`, Expected: `1`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Course Schedule",
//	//		StartingCode: "function canFinish(numCourses, prerequisites) { /* implementation */ }",
//	//		Description:  "There are numCourses courses you have to take, labeled from 0 to numCourses-1. You are given an array of prerequisites where prerequisites[i] = [a, b] indicates that you must take course b before course a. Return true if you can finish all courses.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[2,[[1,0]]]]`, Expected: `true`},
//	//			{Input: `[[2,[[1,0],[0,1]]]]`, Expected: `false`},
//	//			{Input: `[[3,[[0,1],[0,2],[1,2]]]]`, Expected: `true`},
//	//			{Input: `[[1,[[0,0]]]]`, Expected: `false`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Binary Tree Level Order Traversal",
//	//		StartingCode: "function levelOrder(root) { /* implementation */ }",
//	//		Description:  "Given the root of a binary tree, return the level order traversal of its nodes' values (from left to right, level by level).",
//	//		TestCases: []TestCase{
//	//			{Input: `[3,9,20,null,null,15,7]`, Expected: `[[3],[9,20],[15,7]]`},
//	//			{Input: `[1]`, Expected: `[[1]]`},
//	//			{Input: `[1,null,2]`, Expected: `[[1],[2]]`},
//	//			{Input: `[1,2,3,4,5]`, Expected: `[[1],[2,3],[4,5]]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Top K Frequent Elements",
//	//		StartingCode: "function topKFrequent(nums, k) { /* implementation */ }",
//	//		Description:  "Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[1,1,1,2,2,3],2]`, Expected: `[1,2]`},
//	//			{Input: `[[1],1]`, Expected: `[1]`},
//	//			{Input: `[[1,2],2]`, Expected: `[1,2]`},
//	//			{Input: `[[1,2,3,4,4,4,5,5,5,5,6],3]`, Expected: `[4,5,6]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Valid Anagram",
//	//		StartingCode: "function isAnagram(s, t) { /* implementation */ }",
//	//		Description:  "Given two strings s and t, return true if t is an anagram of s, and false otherwise.",
//	//		TestCases: []TestCase{
//	//			{Input: `["anagram","nagaram"]`, Expected: `true`},
//	//			{Input: `["rat","car"]`, Expected: `false`},
//	//			{Input: `["hello","llohe"]`, Expected: `true`},
//	//			{Input: `["a","b"]`, Expected: `false`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Kth Largest Element in an Array",
//	//		StartingCode: "function findKthLargest(nums, k) { /* implementation */ }",
//	//		Description:  "Given an integer array nums and an integer k, return the kth largest element in the array.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[3,2,1,5,6,4],2]`, Expected: `5`},
//	//			{Input: `[[3,2,3,1,2,4,5,5,6],4]`, Expected: `4`},
//	//			{Input: `[[1],1]`, Expected: `1`},
//	//			{Input: `[[5,3,2,4,1],3]`, Expected: `3`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Merge k Sorted Lists",
//	//		StartingCode: "function mergeKLists(lists) { /* implementation */ }",
//	//		Description:  "You are given an array of k linked-lists lists, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[1,4,5],[1,3,4],[2,6]]`, Expected: `[1,1,2,3,4,4,5,6]`},
//	//			{Input: `[[1,2,3],[4,5,6]]`, Expected: `[1,2,3,4,5,6]`},
//	//			{Input: `[[1]]`, Expected: `[1]`},
//	//			{Input: `[[]]`, Expected: `[]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Implement Circular Queue",
//	//		StartingCode: "function MyCircularQueue(k) { /* implementation */ }",
//	//		Description:  "Design your implementation of the circular queue. The implementation should support the following operations: `MyCircularQueue(k)`, `enQueue(value)`, `deQueue()`, `Front()`, `Rear()`, and `isEmpty()`, `isFull()`.",
//	//		TestCases: []TestCase{
//	//			{Input: `["MyCircularQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]\n[[3],[1],[2],[],[],[],[3],[]]`, Expected: `[2,false,3]`},
//	//			{Input: `["MyCircularQueue","enQueue","enQueue","deQueue","Front","Rear"]\n[[3],[1],[2],[],[],[]]`, Expected: `[1,2]`},
//	//			{Input: `["MyCircularQueue","enQueue","deQueue","Front","Rear"]\n[[1],[1],[],[],[]]`, Expected: `[1,-1]`},
//	//			{Input: `["MyCircularQueue","enQueue","enQueue","enQueue","isFull","deQueue","enQueue","Rear"]\n[[3],[1],[2],[3],[],[],[4],[]]`, Expected: `[true,4]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//	{
//	//		Type:         "Binary Tree Right Side View",
//	//		StartingCode: "function rightSideView(root) { /* implementation */ }",
//	//		Description:  "Given the root of a binary tree, return the values of the nodes you can see ordered from top to bottom when looking from the right side.",
//	//		TestCases: []TestCase{
//	//			{Input: `[1,2,3,null,5,null,4]`, Expected: `[1,3,4]`},
//	//			{Input: `[1,null,2]`, Expected: `[1,2]`},
//	//			{Input: `[1,2]`, Expected: `[1,2]`},
//	//			{Input: `[1,2,3,4,5,6,7]`, Expected: `[1,3,7]`},
//	//		},
//	//		Category: "Queue",
//	//	},
//	//}
//	//problems := []Problem{
//	//	{
//	//		Type:         "Number of Connected Components in an Undirected Graph",
//	//		StartingCode: "function countComponents(n, edges) { /* implementation */ }",
//	//		Description:  "You have a graph that consists of n nodes labeled from 0 to n - 1. You are given an integer n and an array of edges where edges[i] = [a, b] indicates that there is an edge between nodes a and b. Return the number of connected components in the graph.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[5,[[0,1],[1,2],[3,4]]]]`, Expected: `2`},
//	//			{Input: `[[4,[[0,1],[0,2],[1,2]]]]`, Expected: `1`},
//	//			{Input: `[[5,[[0,1],[0,2],[0,3],[1,4]]]]`, Expected: `1`},
//	//			{Input: `[[6,[[0,1],[0,2],[3,4],[5]]]]`, Expected: `3`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Clone Graph",
//	//		StartingCode: "function cloneGraph(node) { /* implementation */ }",
//	//		Description:  "Given a reference to a node in a connected undirected graph, return a deep copy (clone) of the graph.",
//	//		TestCases: []TestCase{
//	//			{Input: `[1,[2,4],[1,3]]`, Expected: `[[1],[2,4],[1,3],[2],[1]]`},
//	//			{Input: `[2,[1,3],[2]]`, Expected: `[[2],[1,3],[2]]`},
//	//			{Input: `[3,[3]]`, Expected: `[[3],[3]]`},
//	//			{Input: `[]`, Expected: `[]`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Course Schedule II",
//	//		StartingCode: "function findOrder(numCourses, prerequisites) { /* implementation */ }",
//	//		Description:  "There are numCourses courses you have to take, labeled from 0 to numCourses-1. You are given an array of prerequisites where prerequisites[i] = [a, b] indicates that you must take course b before course a. Return the order of courses you should take to finish all courses, or return an empty array if it is impossible.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[4,[[1,0],[2,0],[3,1],[3,2]]]]`, Expected: `[0,1,2,3]`},
//	//			{Input: `[[2,[[1,0],[0,1]]]]`, Expected: `[]`},
//	//			{Input: `[[3,[[0,1],[0,2],[1,2]]]]`, Expected: `[0,1,2]`},
//	//			{Input: `[[2,[[0,1]]]]`, Expected: `[0,1]`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Graph Valid Tree",
//	//		StartingCode: "function validTree(n, edges) { /* implementation */ }",
//	//		Description:  "Given n nodes labeled from 0 to n - 1 and a list of undirected edges, write a function to check whether these edges make up a valid tree.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[5,[[0,1],[0,2],[0,3],[1,4]]]]`, Expected: `true`},
//	//			{Input: `[[5,[[0,1],[1,2],[2,3],[1,4],[4,0]]]]`, Expected: `false`},
//	//			{Input: `[[4,[[0,1],[0,2],[0,3]]]]`, Expected: `true`},
//	//			{Input: `[[1,[]]]`, Expected: `true`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Shortest Path in a Grid with Obstacles Elimination",
//	//		StartingCode: "function shortestPath(grid, k) { /* implementation */ }",
//	//		Description:  "You are given a m x n grid where each cell can be either 0 (empty) or 1 (obstacle). You can eliminate at most k obstacles. Find the shortest path from the top-left corner to the bottom-right corner.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[0,0,0],[1,1,0],[0,0,0]],1`, Expected: `6`},
//	//			{Input: `[[0,1,1],[1,1,1],[0,0,0]],1`, Expected: `4`},
//	//			{Input: `[[0,0,0],[1,0,0],[0,0,0]],2`, Expected: `4`},
//	//			{Input: `[[0,0,1],[1,0,0],[0,0,0]],1`, Expected: `-1`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Max Area of Island",
//	//		StartingCode: "function maxAreaOfIsland(grid) { /* implementation */ }",
//	//		Description:  "Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical). You may assume all four edges of the grid are surrounded by water. Find the maximum area of an island in the given 2D array.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[0,0,0,0],[0,1,1,0],[0,1,1,0],[0,0,0,0]]`, Expected: `4`},
//	//			{Input: `[[1,1,0,0],[0,1,1,0],[0,0,0,0]]`, Expected: `3`},
//	//			{Input: `[[0,0],[0,0]]`, Expected: `0`},
//	//			{Input: `[[1]]`, Expected: `1`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Minimum Spanning Tree",
//	//		StartingCode: "function minSpanningTree(n, edges) { /* implementation */ }",
//	//		Description:  "Given a number of nodes n and a list of edges with weights, find the minimum spanning tree of the graph.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[3,[[0,1,1],[0,2,2],[1,2,3]]]]`, Expected: `2`},
//	//			{Input: `[[4,[[0,1,5],[1,2,6],[0,2,1],[2,3,1]]]]`, Expected: `7`},
//	//			{Input: `[[5,[[0,1,10],[1,2,10],[2,3,10],[3,4,10]]]]`, Expected: `40`},
//	//			{Input: `[[2,[[0,1,1]]]]`, Expected: `1`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Is Graph Bipartite?",
//	//		StartingCode: "function isBipartite(graph) { /* implementation */ }",
//	//		Description:  "A bipartite graph is a graph that can be colored using two colors such that no two adjacent nodes have the same color. Given a graph represented as an adjacency list, determine if it is bipartite.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[1,3],[0,2],[1,3],[0,2]]`, Expected: `true`},
//	//			{Input: `[[1,2,3],[0,2],[0,3],[0,1]]`, Expected: `false`},
//	//			{Input: `[[1,2],[2,3],[3,0]]`, Expected: `true`},
//	//			{Input: `[[1],[0]]`, Expected: `true`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Pacific Atlantic Water Flow",
//	//		StartingCode: "function pacificAtlantic(heights) { /* implementation */ }",
//	//		Description:  "There is an m x n matrix heights representing the height of each cell in the grid. Water can flow from a cell to another cell in the four cardinal directions (up, down, left, right) if the neighboring cell has a height equal to or lower than the current cell's height. Find the list of grid coordinates where water can flow to both the Pacific and Atlantic oceans.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,5,4,4,5],[5,1,1,2,4]]`, Expected: `[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]`},
//	//			{Input: `[[1]]`, Expected: `[[0,0]]`},
//	//			{Input: `[[2,1],[1,2]]`, Expected: `[[0,0],[0,1],[1,0],[1,1]]`},
//	//			{Input: `[[5,4],[4,5]]`, Expected: `[[0,0],[0,1],[1,0],[1,1]]`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Longest Increasing Path in a Matrix",
//	//		StartingCode: "function longestIncreasingPath(matrix) { /* implementation */ }",
//	//		Description:  "Given an integer matrix, find the length of the longest increasing path in the matrix. From each cell, you can move to four directions (up, down, left, or right). You may not move diagonally or move outside of the boundary.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[9,9,4],[6,6,8],[2,1,1]]`, Expected: `4`},
//	//			{Input: `[[3,4,5],[3,2,6],[2,2,1]]`, Expected: `4`},
//	//			{Input: `[[1]]`, Expected: `1`},
//	//			{Input: `[[7,8,9],[5,6,10],[3,4,5]]`, Expected: `4`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//	{
//	//		Type:         "Frog Jump",
//	//		StartingCode: "function canCross(stones) { /* implementation */ }",
//	//		Description:  "A frog is initially positioned at the first stone, and it can jump to the next stone if the difference in their positions is equal to k (the last jump distance). Implement a function to determine if the frog can cross to the last stone.",
//	//		TestCases: []TestCase{
//	//			{Input: `[[0,1,3,5,6,8,12,17]]`, Expected: `true`},
//	//			{Input: `[[0,1,2,3,4,8,9,11]]`, Expected: `false`},
//	//			{Input: `[[0,1,2,3,4,5,6,7,8,9,10]]`, Expected: `true`},
//	//			{Input: `[[0,1,2,3,4,5,6,8,9]]`, Expected: `false`},
//	//		},
//	//		Category: "Graph",
//	//	},
//	//}
//
//	if err := DB.DB().Create(&problems).Error; err != nil {
//		fmt.Println("Error creating problem:", err)
//	}
//
//}
