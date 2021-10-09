original = input()
translation = input()
correct_translation = ""

for letter in original:
    if letter == "t":
        letter = "s"
    correct_translation = letter + correct_translation

if correct_translation == translation:
    print("YES")
else:
    print("NO")

print(correct_translation)