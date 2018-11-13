#include "stdio.h"
#include "time.h"

int arr[1024*1024][8];
int sum = 0;

int main(int argc, const char **argv) {
    clock_t start, finish;
    double Total_time;

    for (int i = 0; i < 1024 * 1024; i += 1) {
        for (int j = 0; j < 8; j++) {
            arr[i][j] = 1;
        }
    }

    start = clock();
    for (int i = 0; i < 1024 * 1024; i += 1) {
        for (int j = 0; j < 8; j++) {
            sum += arr[i][j];
        }
    }
    finish = clock();
    Total_time = (double)(finish - start) / CLOCKS_PER_SEC*1000;
    printf("\n函数运行时间：%0.3f毫秒 ,%d\n", Total_time,sum);

    sum = 0;    
    start = clock();
    for (int i = 0; i < 8; i += 1) {
        for (int j = 0; j < 1024 * 1024; j++) {
            sum += arr[j][i];
        }
    }
    finish = clock();
    Total_time = (double)(finish - start) / CLOCKS_PER_SEC*1000;
    printf("\n函数运行时间：%0.3f毫秒 ,%d\n", Total_time,sum);

    return 0;
}