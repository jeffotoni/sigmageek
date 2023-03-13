// @autor: jeffotoni
// @date: 2023-03-13
// gcc -o palindormo main.c -static

#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

bool isPalindrome(const char *str, size_t length)
{
  for (size_t i = 0; i < (length + 1) / 2; i++)
  {
    if (str[i] != str[length - i - 1])
    {
      return false;
    }
  }
  return true;
}

bool isPrime(int n)
{
  if (n <= 1 || n % 2 == 0)
  {
    return false;
  }
  for (int i = 3; i * i <= n; i += 2)
  {
    if (n % i == 0)
    {
      return false;
    }
  }
  return true;
}

int64_t findNextPrimePalindrome(int64_t current)
{
  for (int64_t i = current + 1;; i++)
  {
    char str[100];
    sprintf(str, "%ld", i);
    int length = strlen(str);
    if (isPalindrome(str, length) && isPrime(i))
    {
      return i;
    }
  }
}

int main(int argc, char *argv[])
{
  if (argc != 2)
  {
    printf("use: %s <numero de digitos>\n", argv[0]);
    return 1;
  }
  int digito = atoi(argv[1]);

  FILE *fp = fopen("pi.txt", "r");
  if (fp == NULL)
  {
    perror("Failed to open file");
    return 1;
  }

  bool start = false;
  int maxPrimePalindrome = 0;
  char line[1000000];
  int buffer_size = 1000000;
  char *buffer = (char *)malloc(buffer_size * sizeof(char));

  buffer[digito] = '\0';

  while (fgets(line, buffer_size, fp) != NULL)
  {
    for (int i = 0; i <= (int)strlen(line) - digito; i++)
    {
      strncpy(buffer, line + i, digito);
      buffer[digito] = '\0';
      int n = atoi(buffer);
      if (isPalindrome(line + i, digito) && isPrime(n))
      {
        if (n > maxPrimePalindrome)
        {
          maxPrimePalindrome = n;
        }
        if (!start)
        {
          start = true;
          printf("Primeiro primo palindromo encontrado na posição %d: %d\n", i, n);
        }

        int next = findNextPrimePalindrome(n);
				printf("Proximo primo palindromo encontrado: %d\n", next);
        break;
      }
    }
  }

  if (maxPrimePalindrome > 0) {
		printf("Maior número primo palíndromo encontrado: %d\n", maxPrimePalindrome);
	}

  fclose(fp);
  return 0;
}
