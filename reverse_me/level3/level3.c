#include <string.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define ulong unsigned long

void no() {
	puts("Nope.");
	exit(1);
}

void ok() {
	puts("Good job.");
	return ;
}

int main(void)
{
	ulong uVar1;
	int iVar2;
	size_t sVar3;
	bool bVar4;
	char buffer[4];
	char input[31];
	char key[9];
	ulong passwordOffset;
	int ascii_diff;

	printf("Please enter key: ");
	if (!scanf("%s", input))
	{
		no();
	}
	if (input[0] != '4')
	{
		no();
	}
	if (input[1] != '2')
	{
		no();
	}
	fflush(stdin);
	memset(key, 0, 9);
	key[0] = '*';
	passwordOffset = 2;
	int i = 1;
	while (true)
	{
		sVar3 = strlen(key);
		uVar1 = passwordOffset;
		bVar4 = false;
		if (sVar3 < 8)
		{
			sVar3 = strlen(input);
			bVar4 = uVar1 < sVar3;
		}
		if (!bVar4)
			break;
		buffer[0] = input[passwordOffset];
		buffer[1] = input[passwordOffset + 1];
		buffer[2] = input[passwordOffset + 2];
		key[i] = (char)atoi((const char *)&buffer);
		passwordOffset += 3;
		i++;
	}
	key[i] = '\0';
	// printf("key='%s'\n", key);
	ascii_diff = strcmp(key, "********");
	// printf("ascii diff = %d\n", ascii_diff);
	if (ascii_diff == -2)
		no();
	else if (ascii_diff == -1)
		no();
	else if (ascii_diff == 0)
		ok();
	else if (ascii_diff == 1)
		no();
	else if (ascii_diff == 2)
		no();
	else if (ascii_diff == 3)
		no();
	else if (ascii_diff == 4)
		no();
	else if (ascii_diff == 5)
		no();
	else if (ascii_diff == 115)
		no();
	else
	{
		no();
	}
	return 0;
}
