queue = [int(n) for n in input().split()]
arrangement = input()

for i in range(queue[1]):
    for j in range(1, len(arrangement)):
        if (arrangement[j-1] + arrangement[j] == "BG"):
            list_arrangement = list(arrangement)
            list_arrangement[j-1] = "G"
            list_arrangement[j] = "B"
            arrangement = "".join(list_arrangement)
            # print(new_arrangement)

            # HARUSNYA J LONCAT DUA KALI KETIKA KONDISI IF TERPENUHI!!!!!!!
            j += 2

    # arrangement = new_arrangement

print(arrangement)