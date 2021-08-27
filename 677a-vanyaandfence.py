nums = [int(n) for n in input().split()]
heights = [int(height) for height in input().split()]
count_width = 0

for height in heights:
    if height > nums[1]:
        count_width += 2
    else:
        count_width += 1

print(count_width)