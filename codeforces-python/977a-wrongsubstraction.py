nums = [int(num) for num in input().split()]
substracted_num = nums[0]

for i in range(nums[1]):
    if substracted_num % 10 == 0:
        substracted_num /= 10
    else:
        substracted_num -= 1

print(int(substracted_num))
