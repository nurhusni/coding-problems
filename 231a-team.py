# brute force solution

n = int(input())
solutions = []
total = 0

for i in range(0, n):
    solutions = [int(solution) for solution in input().split()]
    if (sum(solutions) >= 2):
        total += 1

print(total)