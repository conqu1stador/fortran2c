#include <stdio.h>
#include <math.h>
 
int main() {

//        The Steve Sharpe (University of Bridgeport);
//        Irvin Levy (Gordon College) Linear Regression Package;
// 
//        *  *  *  *  VARIABLE LIST *  *  *  *;
//                S1 - Sum of X values;
//                S2 - Sum of Y values;
//                S3 - Sum of X * Y;
//                S4 - Sum of X * X;
//                S5 - Sum of Y * Y;
//                D  - Sum of Y(obs) - Y(calc) squared;
//                S  - Slope;
//                B  - Y-intercept;
//                E1 - Error in the slope;
//                E2 - Error in the intercept;
//                T  - Beta function for linear regression;
//                R  - Correlation coefficient;
// 
//        Error function uses (N-2) weighting;
// 
// 
// 
DIMENSION X(52), Y(2,50), LITERL(2);
double S1,S2,S3,S4,S5,T,S,B,D,R,E1,E2,BBAR;
int N;

printf("%d%d", IA, IB);

printf("*   *   *  LINEAR REGRESSION ANALYSIS  *   *   *,//\n");
// 10     FORMAT('0',1X,'*   *   *  LINEAR REGRESSION ANALYSIS  *   *   *',//);
printf("HOW MANY PAIRS TO BE ANALYZED?\n");
// 20     FORMAT(1X,'HOW MANY PAIRS TO BE ANALYZED?'$);
scanf("%d", &N);
if (N > 50) goto l_70;
printf("Enter one pair at a time\n");
// 30     FORMAT(//1X,'Enter one pair at a time');
printf("and separate X from Y with a comma.//\n");
// 40     FORMAT(1X,'and separate X from Y with a comma.'//);
printf("Enter pair number one : \n");
// 50     FORMAT(1X,'Enter pair number one : '$);
scanf("%f%f", &X(1), &Y(1,1));
for (double I=2; I <= N; I++) {
printf("Enter pair number%d, : \n");
// 55             FORMAT(1X,'Enter pair number',I3,' : '$);
scanf("%d%d", &X(I), &Y(1,I));
}
goto l_90;
l_70:
printf("At present this program can only handle 50 data pairs.\n");
// 80     FORMAT(1X,'At present this program can only handle 50 data pairs.');
return;
l_90:
printf("Would you like to review the data?\n");
// 100    FORMAT(1X,'Would you like to review the data?'$);
scanf("%d", &LITERL(1));
//// 110    FORMAT(A1);
if (LITERL(1) == 'N') goto l_140;
for (double I=1; I <= N; I++) {
printf("DATA PAIR%d,).  ,2%f\n", X(I), Y(1,I));
// 120            FORMAT(1X,'DATA PAIR',I3,').  ',2F20.10);
}
l_140:
printf("Would you like to make any changes?\n");
// 150    FORMAT(1X,'Would you like to make any changes?'$);
scanf("%d", &LITERL(1));
if (LITERL(1) == 'N') goto l_200;
printf("To change a data pair, enter the\n");
// 160    FORMAT(1X,'To change a data pair, enter the');
printf("Pair Number and the new X, Y pair.\n");
// 165    FORMAT(1X,'Pair Number and the new X, Y pair.');
printf("How many changes shall we make?\n");
// 180    FORMAT('0',1X,'How many changes shall we make?'$);
scanf("%d", &ICHANG);
for (double I=1; I <= ICHANG; I++) {
printf("CHANGE NUMBER%d, : \n");
// 185            FORMAT(1X,'CHANGE NUMBER',I4,' : '$);
scanf("%f%f%f", &C1, &C2, &C3);
X(C1)=C2;
Y(1,C1)=C3;
}
goto l_90;
l_200:
for (double I=1; I <= N; I++) {
S1=S1+X(I);
S2=S2+Y(1,I);
S3=S3+X(I)*Y(1,I);
S4=S4+X(I)*X(I);
S5=S5+Y(1,I)*Y(1,I);
}
T=N*S4-S1*S1;
S=(N*S3-S1*S2)/T;
B=(S4*S2-S1*S3)/T;
for (double I=1; I <= N; I++) {
Y(2,I)=S*X(I)+B;
D=D+(Y(2,I)-Y(1,I))**2;
}
D=D/(N-2);
E1=sqrt(D*N/T);
E2=sqrt(D/N*(1+S1*S1/T));
R=(N*S3-S1*S2)/;
2    (sqrt(ABS(((N*S4-ABS(S1)**2))*(N*S5-ABS(S2)**2))));
printf("/X-VALUEY-OBSY-CALC\n");
// 230    FORMAT(////,10X,'X-VALUE',20X,'Y-OBS',22X,'Y-CALC');
printf("==================\n");
// 235    FORMAT(10X,'=======',20X,'=====',22X,'======');
for (double I=1; I <= N; I++) {
printf(".10,2%f\n", Y(1,I), Y(2,I));
// 240            FORMAT(F22.10,2F27.10);
}
printf("SLOPE = ,D22.16\n");
// 260    FORMAT(//,1X,'SLOPE = ',D22.16);
printf("THE ERROR IN THE SLOPE IS +/- ,D22.16\n");
// 270    FORMAT(1X,'THE ERROR IN THE SLOPE IS +/- ',D22.16);
printf("INTERCEPT = ,D22.16\n");
// 280    FORMAT(//,1X,'INTERCEPT = ',D22.16);
printf("THE ERROR IN THE INTERCEPT IS +/- ,D22.16\n");
// 285    FORMAT(1X,'THE ERROR IN THE INTERCEPT IS +/- ',D22.16);
LITERL(2)='+';
if (ABS(B) != B) LITERL(2)='-';
printf("EQUATION FOR THE BEST LINEAR FIT IS : \n");
// 290    FORMAT(//,1X,'EQUATION FOR THE BEST LINEAR FIT IS : ');
BBAR=ABS(B);
printf("Y(X) =,D25.16, * X ,A1,D25.16,///\n", (2), ,BBAR);
// 295    FORMAT(1X,'Y(X) =',D25.16,' * X ',A1,D25.16,///);
printf("INEAR CORRELATION COEFFICIENT =,D22.16\n");
// 296    FORMAT(' LINEAR CORRELATION COEFFICIENT =',D22.16);
return 'LINEAR... Execution completed';
}
