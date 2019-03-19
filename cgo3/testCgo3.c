#include "cgo3.h"

#include <stdio.h>
#include <string.h>

int main()
{
    char strT[] = "随机数";

    int lenT = strlen(strT);

    GoString goStrT = {strT, lenT};

    printInGo(goStrT);

    GoInt randomT = getRandomInt(100);

    printf("%s: %lld\n", strT, randomT);

    return 0;
}