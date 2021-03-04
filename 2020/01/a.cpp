#include <iostream>
#include <unordered_set>

using namespace std;

int
main(int, char *[])
{
    ios::sync_with_stdio(false);
    int e;
    unordered_set<int> es;
    while (cin >> e) {
        int d = 2020 - e;
        if (es.contains(d)) {
            cout << d * e << endl;
            break;
        }
        es.insert(e);
    }
    return 0;
}
