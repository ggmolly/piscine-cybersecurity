#include <string.h>
#include <stdio.h>

int main(void) {
	char password[14] = "__stack_check\0";
	char entered_password[255] = {0};
	printf("Please enter key: ");
	scanf("%s", entered_password);
	if (strcmp(password, entered_password) == 0) {
		printf("Good job.\n");
	} else {
		printf("Nope.\n");
	}
	return (0);
}
