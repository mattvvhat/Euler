#include <string.h>

struct Item {
  char *key;
  int  val;
};

struct Map {
  int         Size;
  int         Index;
  struct Item *List;
};

// create
struct Map* create(int  n) {
  struct Map *whatever;
  whatever = malloc(sizeof(struct Map));
  whatever->Index = 0;
  whatever->Size = n;
  whatever->List = malloc(whatever->Size * sizeof(struct Item));
  return whatever;
}

void set(struct Map *m, char *key, int val) {
  m->List[m->Index].key = key;
  m->List[m->Index].val = val;
  ++m->Index;
}

int get(struct Map *m, char *key) {

  for (int i=0; i < m->Size; i++) {

    if (m->List[i].key == NULL) {
      continue;
    }

    if (strcmp(m->List[i].key, key) == 0) {
      return m->List[i].val;
    }
  }

  return 2;
}
