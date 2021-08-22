username = input()
distinct = []

for letter in username:
    for char in distinct:
        if letter == char:
            break
        else:
            distinct.append(letter)

print(len(distinct))