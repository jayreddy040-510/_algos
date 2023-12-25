from typing import List


def bsearch(sorted_nums: List[int], target: int) -> int:
    i, j = 0, len(sorted_nums) - 1
    while i <= j:
        pivot = i + ((j - i) // 2)
        num = sorted_nums[pivot]
        if num == target:
            return pivot
        elif target > num:
            i = pivot + 1
        else:
            j = pivot - 1
    return -1
