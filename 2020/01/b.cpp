#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

int
main(int, char *[])
{
    int e;
    vector<int> es;
    unordered_map<int, int> tab;
    while (cin >> e) {
        int d = 2020 - e;
        auto i = tab.find(d);
        if (i != tab.end()) {
            cout << i->second * e << endl;
            break;
        }
        for (int e0 : es)
            tab[e0 + e] = e0 * e;
        es.push_back(e);
    }
    return 0;
}
