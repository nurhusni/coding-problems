size = [int(a) for a in input().split()]
row = ""

for n in range(1, size[0] + 1):
    row = ""

    for m in range(1, size[1] + 1):
        if (n % 2 == 1):
            row += "#"
        elif (n % 4 == 0):
            if (m == 1):
                row += "#"
            else:
                row += "."
        elif (n % 2 == 0):
            if (m == size[1]):
                row += "#"
            else:
                row += "."
        else:
            row += "."

    print(row)