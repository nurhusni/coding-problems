// https://codeforces.com/problemset/problem/734/A

// Wrong at test 10
// input: 534204
// output: 106840
// expected: 106841

#include <iostream>

using namespace std;

int main() 
{
    int coordinate;
    cin >> coordinate;

    if (coordinate <= 5)
    {
        cout << 1;
    }
    else
    {
        int minimumValue = (int)(coordinate/5 + 0.5);

        if (minimumValue % 5 == 0)
        {
            cout << minimumValue;
        }
        else
        {
            cout << minimumValue + 1;
        }
    }
}