number = input()
do_print = True 

if int(number) % 4 == 0 or int(number) % 7 == 0:
    print("YES")
else:
    for digit in number:
        if digit not in {"4", "7"}:
            print("NO")
            do_print = False
            break
    
    if do_print:
        print("YES")

# THE TEST CASE NUMBER 26 SHOULD BE "YES", BUT I DONT SEE WHY
# input: 799 should be "YES", but why?