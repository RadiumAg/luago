int main(int argc, const char** argv) {
    struct a {
        int c;
        int b;
    };
    a c = {
        c:1,
        b:2
    };

    int g = 1;
    int *k = &g;
     *k = 1;
     a *d = &c;
}