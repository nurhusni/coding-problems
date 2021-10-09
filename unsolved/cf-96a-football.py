players = input()
count = 0

for i in range(1, len(players) + 1):
    if count == 0 and players[i] == players[i-1]:
        count += 2
    elif players[i] == players[i-1]:
        count += 1
        if count == 7:
            break
    elif players[i] != players[i-1]:
        count = 0

if (count == 7):
    print("YES")
elif (count < 7):
    print("NO")