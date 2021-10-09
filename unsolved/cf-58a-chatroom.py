word = input()
hello = "h e l l o".split()
print(hello)

for letter in word:
    if len(hello) == 0:
        print("YES")
        break
    elif letter == hello[0]:
        hello.pop(0)
    else:
        print("NO")
        break