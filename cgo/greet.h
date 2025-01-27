#ifndef _GREET_H
#define _GREET_H

struct Greetee{
    const char *name;
    int year;
};

int greet(const char *name, int year, char *out);
int greet_struct(struct Greetee *g, char *out);

#endif