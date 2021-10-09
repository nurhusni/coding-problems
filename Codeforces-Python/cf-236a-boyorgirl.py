word = input()
setWord = set(())

for letter in word:
    setWord.add(letter)

if len(setWord) % 2 == 1:
    print("IGNORE HIM!")
else:
    print("CHAT WITH HER!")

print(setWord)