#ifndef _MATT_H
#define _MATT_H


struct node {
  struct node *next;
  char         value[];
};


int length(struct node *list) {

  if (list == NULL) {
    return 0;
  }

  int length = 1;

  while (list->next != NULL) {
    list = list->next;
    length++;
  }

  return length;
}


// TAIL
struct node* tail(struct node *list) {

  if (list == NULL) {
    return NULL;
  }

  struct node *snaker = list;

  while (snaker->next != NULL) {
    snaker = snaker->next;
  }

  return snaker;
}


// APPEND
void append(struct node **list, char *value) {

  struct node *t = tail(*list);

  struct node *item = malloc(sizeof(struct node));
  item->next = NULL;
  strcpy(item->value, value);

  if (t == NULL) {
    *list = item;
  } else {
    t->next = item;
  }
}

struct node* get_item(struct node *list, char *key) {
  if (list == NULL) {
    return NULL;
  }

  while (list !=  NULL) {
    if (strcmp(list->value, key) == 0) {
      return list;
    }
    list = list->next;
  }

  return NULL;
}


void printwhatever(struct node *list) {
  while (list != NULL) {
    printf("%s ", list->value);
    list = list->next;
  }
}


#endif
