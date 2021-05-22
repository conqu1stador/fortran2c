#include <stdio.h>
#include <math.h>
 
int main() {
int IA,IB,IC;
double S,AREA;
scanf("%d%d%d", &IA, &IB, &IC);
if(IA == 0 || IB == 0 || IC == 0) return 1;
if(IA == 10) goto l_703;
S = (IA + IB + IC) / 2.0;
AREA = sqrt( S * (S - IA) * (S - IB) * (S - IC) );
printf("A= %d  B= %d  C= %d  AREA= %f\n", IA, IB, IC, AREA);
l_703:
return 0;
}

SUBROUTINE PPRNT(P,N,ID);
float P[ID][N];
IGO=1;
NEND=0;
l_1:
NBEG=NEND+1;
NEND=NEND+6;
if != D < N) GO TO 3;
2 NEND=N;
IGO=2;
printf("H06(I3,2X\n", (K), K);
printf("H \n");
for (int I=1; I <= N; I++) {
}
7 FORMAT(1H ,I8,2X,6F10.5);
goto l_1;
8 RETURN;
}
