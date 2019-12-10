#include <string.h>
#include <netdb.h>
#include <unistd.h>
#include <stdio.h>

int main(int argc, char* argv[])
{
  char hn[254];
  char *dn;
  struct hostent *hp;
  char dname[254];

  gethostname(hn, 254);
  printf("%s\n", hn);
  getdomainname(dname, 254);
  printf("domainname: %s\n", dname);
  hp = gethostbyname(hn);
  dn = strchr(hp->h_name, '.');
  if ( dn != NULL ) {
      printf("%s\n", ++dn);
    }
   else {
       printf("No domain name available through gethostbyname().\n");
     }
  
    return 0;
} 
