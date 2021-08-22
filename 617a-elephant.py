coordinate = int(input())

if coordinate <= 5:
    print(1)
else:
    minimum_value = int(coordinate/5)
    if coordinate % 5 == 0:
        print(minimum_value)
    else:
        print(minimum_value + 1)