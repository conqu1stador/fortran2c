#include <stdio.h>
#include <math.h>
 
int main() {
	float X[52], Y[2][50], LITERL[2];
	double S1,S2,S3,S4,S5,T,S,B,D,R,E1,E2,BBAR;
	DATA Y/1,2/;
	int N;

	printf("*   *   *  LINEAR REGRESSION ANALYSIS  *   *   *,//\n");
	printf("HOW MANY PAIRS TO BE ANALYZED?\n");
	scanf("%d", &N);
	if (N > 50) goto l_70;
	printf("Enter one pair at a time\n");
	printf("and separate X from Y with a comma.//\n");
	printf("Enter pair number one : \n");
	scanf("%f%f", &X[1 - 1], &Y[1 - 1][1 - 1]);
	for (int I=2; I <= N; I++) {
		printf("Enter pair number%d, : \n", I);
		scanf("%d%d", &X[I - 1], &Y[1 - 1][I - 1]);
	}
	goto l_90;
	l_70:
	printf("At present this program can only handle 50 data pairs.\n");
	return 0;
	l_90:
	printf("Would you like to review the data?\n");
	scanf("%d", &LITERL[1 - 1]);
	if (LITERL[1 - 1] == 'N') goto l_140;
	for (int I=1; I <= N; I++) {
		printf("DATA PAIR%d,).  %f%f\n", I, X[I - 1], Y[1 - 1][I - 1]);
	}
	l_140:
	printf("Would you like to make any changes?\n");
	scanf("%d", &LITERL[1 - 1]);
	if (LITERL[1 - 1] == 'N') goto l_200;
	printf("To change a data pair, enter the\n");
	printf("Pair Number and the new X, Y pair.\n");
	printf("How many changes shall we make?\n");
	scanf("%d", &ICHANG);
	for (int I=1; I <= ICHANG; I++) {
		printf("CHANGE NUMBER%d, : \n", I);
		scanf("%f%f%f", &C1, &C2, &C3);
		X[C1 - 1]=C2;
		Y[1 - 1][C1 - 1]=C3;
	}
	goto l_90;
	l_200:
	for (int I=1; I <= N; I++) {
		S1=S1+X[I - 1];
		S2=S2+Y[1 - 1][I - 1];
		S3=S3+X[I - 1]*Y[1 - 1][I - 1];
		S4=S4+X[I - 1]*X[I - 1];
		S5=S5+Y[1 - 1][I - 1]*Y[1 - 1][I - 1];
	}
	T=N*S4-S1*S1;
	S=(N*S3-S1*S2)/T;
	B=(S4*S2-S1*S3)/T;
	for (int I=1; I <= N; I++) {
		Y[2 - 1][I - 1]=S*X[I - 1]+B;
		D=D+(Y[2 - 1][I - 1]-Y[2 - 1][I - 1])^2;
	}
	D=D/(N-2);
	E1=sqrt(D*N/T);
	E2=sqrt(D/N*(1+S1*S1/T));
	R=(N*S3-S1*S2)/(sqrt(abs(((N*S4-abs(S1)^2))*(N*S5-abs(S2)^2))));
	printf("/X-VALUEY-OBSY-CALC\n");
	printf("==================\n");
	for (int I=1; I <= N; I++) {
		printf(".10%f%f\n", X[I - 1], Y[1 - 1][I - 1], Y[1 - 1][I - 1]);
	}
	printf("SLOPE = %f\n", S);
	printf("THE ERROR IN THE SLOPE IS +/- %f\n", E1);
	printf("INTERCEPT = %f\n", B);
	printf("THE ERROR IN THE INTERCEPT IS +/- %f\n", E2);
	LITERL[2 - 1]='+';
	if (abs(B) != B) LITERL[2 - 1]='-';
	printf("EQUATION FOR THE BEST LINEAR FIT IS : \n");
	BBAR=abs(B);
	printf("Y[X - 1] =%f, * X ,A1%f,///\n", S, LITERL[2 - 1], BBAR);
	printf("INEAR CORRELATION COEFFICIENT =%f\n", R);
	return 'LINEAR... Execution completed'0;
}
