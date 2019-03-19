#include "test.h"

#include "stdio.h"
#include "time.h"

void show3Times(char *strA)
{
	for (int i = 0; i < 3; i ++) {
		puts(strA);
	}

	time_t timeT = time(NULL);

	struct tm *timeInfoT = localtime(&timeT);

	printf("%s", asctime(timeInfoT));

}
