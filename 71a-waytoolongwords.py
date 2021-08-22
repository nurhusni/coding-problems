# string solution

n = int(input())
words = []
abbreviation = ""

for i in range (0, n):
    words.append(input())
    if (len(words[i]) > 10):
        abbreviation = words[i][0] + str((len(words[i]) - 2)) + words[i][-1]
        print(abbreviation)
    else:
        print(words[i])