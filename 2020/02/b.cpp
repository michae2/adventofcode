#include <iostream>

using namespace std;

int
main(int, char *[])
{
    ios::sync_with_stdio(false);
    unsigned valid = 0, a, b;
    char sep, ch;
    string pw;
    while (getline(cin >> a >> sep >> b >> ws >> ch >> sep >> ws, pw))
        if ((a <= pw.size() && pw[a - 1] == ch) ^
            (b <= pw.size() && pw[b - 1] == ch))
            ++valid;
    cout << valid << endl;
    return 0;
}
