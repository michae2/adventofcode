#include <algorithm>
#include <iostream>

using namespace std;

int
main(int, char *[])
{
    ios::sync_with_stdio(false);
    unsigned valid = 0, lo, hi;
    char sep, ch;
    string pw;
    while (getline(cin >> lo >> sep >> hi >> ws >> ch >> sep >> ws, pw)) {
        auto n = count(pw.begin(), pw.end(), ch);
        if (lo <= n && n <= hi)
            ++valid;
    }
    cout << valid << endl;
    return 0;
}
