/**
 * A simple reminder about how computers work...
 */


#define _GNU_SOURCE

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "table.h"
#include "what.h"


int main(void) {

  struct node *head = NULL;

  append(&head, "10");
  printf("LENGTH => %d\n", length(head));

  append(&head, "79");
  printf("LENGTH => %d\n", length(head));

  append(&head, "79");
  append(&head, "79");
  append(&head, "665");

  printwhatever(head);
  printf("\n");


  struct node *found = get_item(head, "665");

  printf("===\n\n%s\n", found->value);


  /*
  append(head, 10);
  append(head, 13);
  append(head, 68);
  */

  exit(EXIT_SUCCESS);


  struct Map *c = create(100);

  set(c, "red", 10);
  set(c, "dead", 68);

  printf("[red] => %d\n", get(c, "red"));
  printf("[blue] => %d\n", get(c, "blue"));
  printf("[dead] => %d\n", get(c, "dead"));
  printf("[revolver] => %d\n", get(c, "revolver"));

  exit(EXIT_SUCCESS);


  // New Program
  // New Program


  FILE *fp;
  char *line = NULL;
  size_t len = 0;
  ssize_t read;

  fp = fopen("p079_keylog.txt", "r");

  if (fp == NULL) {
    exit(EXIT_FAILURE);
  }

  while ( (read = getline(&line, &len, fp)) != -1) {
    printf("LENGTH => %zd\n", read);
  }

  fclose(fp);

  if (line) {
    free(line);
  }

  exit(EXIT_SUCCESS);
}
