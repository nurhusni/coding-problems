n = int(input())
count = 0
maximum = 0

for i in range(n):
    passengers = [int(passenger) for passenger in input().split()]
    count = count + passengers[1] - passengers[0]
    if count > maximum:
        maximum = count

print(maximum)