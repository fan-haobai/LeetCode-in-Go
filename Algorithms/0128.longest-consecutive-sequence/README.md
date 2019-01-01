# [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/)

## 题目
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

For example,
```
Given [100, 4, 200, 1, 3, 2],
The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.
```

Your algorithm should run in O(n) complexity.

## 解题思路

见程序注释，本道题采用了排序的方法，其实更优的方法是用 hashMap 解决，思路如下：

```Java
public int longestConsecutive(int[] nums)
{
    Set<Integer> set = new HashSet<>();
    for(int n: nums) set.add(n);

    int max = 0;
    for(int n: nums) {
        int count = 0;

        if(set.isEmpty()) break;

        // 对于数组中的数，将其左右相连的数都remove掉，这样set会越来越小
        int val = n;
        while(set.remove(val--))
            count ++;

        val = n;
        while(set.remove(++val))
            count ++;
            
        // 判断完每个连续的数块，就更新一次最大值
        max = Math.max(count,max);
    }

    return max;
}
```
