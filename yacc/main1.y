%{
#include <stdio.h>
#include <stdlib.h>

#define YYDEBUG 1
#define INT_TYPE 1
#define CHAR_TYPE 2
#define STR_TYPE 3

double stack[19];
int sp;

void push(double x) {
  stack[sp++] = x;
}

double pop() {
  return stack[--sp];
}

}%

%token char
%token else
%token for
%token func
%token if
%token int
%token main
%token nil
%token switch
%token string
%token var


simple_type: int
           | char
           | str
           ;


%%

yyerror(char *s)
{
  printf("%s\n", s);
}

extern FILE *yyin;

main(int argc, char **argv)
{
  if(argc>1) {
    yyin = fopen(argv[1], "r");
  }
  if((argc>2)&&(!strcmp(argv[2],"-d")))  {
    yydebug = 1;
  }
  if(!yyparse()) {
    fprintf(stderr,"\tO.K.\n");
  }
}
