# matrix = []
column = 0
row = 0

for i in range(5):
    # matrix.append([int(submatrix) for submatrix in input().split()])

    dummy_matrix = input()

    if sum(dummy_matrix) == 0:
        continue

    elif sum(dummy_matrix == 1):
        for index, cell in enumerate(dummy_matrix):
            if cell == 1:
                row = i + 1
                column = index + 1
                break

print(row, column)
