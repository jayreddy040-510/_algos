from typing import List
# any x of i <= x of i + 1


def bubble_sort(nums: List[int]) -> None:
    # sort in place
    for i in range(len(nums)):
        for j in range(len(nums) - 1 - i):
            if nums[j] > nums[j + 1]:
                nums[j], nums[j+1] = nums[j+1], nums[j]


if __name__ == "__main__":
    arr = [1, 4, 2, 4, 11, 5, 8, 7]
    bubble_sort(arr)
    print(arr)
