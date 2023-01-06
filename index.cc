#include <stdio.h>

void max(int *num1)
{
    printf(" %p ", num1);
    *num1 = 2;
}

int main(int argc, const char **argv)
{
    int a = 1;
    max(&a);
    printf(" %d ", a);
}

