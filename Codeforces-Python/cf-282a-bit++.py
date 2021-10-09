n = int(input())
result = 0

for i in range(n):
    operation = input()
    if operation[1] == "+":
        result += 1
    elif operation[1] == "-":
        result -= 1

print(result)
