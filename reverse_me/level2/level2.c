#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

#define uint unsigned int

void no(void)
{
	puts("Nope.");
	exit(1);
}

void ok(void)
{
	puts("Good job.");
	return;
}

int main(void)
{
	uint uVar1;
	size_t sVar2;
	int iVar3;
	bool bVar4;
	char buffer[3] = {0};
	char local_3d;
	char local_3c;
	char local_3b;
	int local_3a;
	char password[24];
	char key[9];
	uint passwordOffset;
	// undefined4 local_c;

	// local_c = 0;
	printf("Please enter key: ");
	if (!scanf("%s", password))
	{
		no();
	}
	if (password[1] != '0')
	{
		no();
	}
	if (password[0] != '0')
	{
		no();
	}
	fflush(stdin);
	memset(key, 0, 9);
	key[0] = 'd';
	local_3a = 0;
	passwordOffset = 2;
	int i = 1;
	while (true)
	{
		sVar2 = strlen(key);
		uVar1 = passwordOffset;
		bVar4 = false;
		if (sVar2 < 8)
		{
			sVar2 = strlen(password);
			bVar4 = uVar1 < sVar2;
		}
		if (!bVar4)
			break;
		buffer[0] = password[passwordOffset];
		buffer[1] = password[passwordOffset + 1];
		buffer[2] = password[passwordOffset + 2];
		// printf("key[%d] = %d\n", i, key[i]);
		key[i] = (char)atoi((const char *) &buffer);
		// printf("key: '%s' (atoi'ing: '%c%c%c')\n", key, buffer[0], buffer[1], buffer[2]);
		passwordOffset += 3;
		i += 1;
	}
	// printf("i=%d\n", i);
	key[i] = '\0';
	iVar3 = strcmp(key, "delabere");
	if (iVar3 == 0)
	{
		ok();
	}
	else
	{
		no();
	}
	return 0;
}
