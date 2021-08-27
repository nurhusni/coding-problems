word = input()
upper_count = 0

for letter in word:
    if letter.isupper():
        upper_count += 1

if upper_count > len(word)/2:
    print (word.upper())
else:
    print (word.lower())
