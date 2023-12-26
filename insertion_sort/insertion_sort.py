from typing import List


def insertion_sort(nums: List[int]) -> None:
    for i in range(1, len(nums)):
        for j in range(i, 0, -1):
            if nums[j] < nums[j - 1]:
                nums[j], nums[j - 1] = nums[j - 1], nums[j]
            else:
                break


if __name__ == "__main__":
    arr = [1, 5, 23, 4, 7, 2, 10]

    insertion_sort(arr)

    print(arr)
