#include <stdio.h>
#include <string.h>

#include "lib/bbb_dht_read.h"

int main(void) {
    int sensor = 22;
    int base = 1, number = 28;
    float humidity = 0, temperature = 0;
    char hum[10], temp[10];

    int result = bbb_dht_read(sensor, base, number, &humidity, &temperature);

    memset(hum, 0, sizeof(hum));
    memset(temp, 0, sizeof(temp));
    sprintf(hum, "%.2f", humidity);
    sprintf(temp, "%.2f", temperature);
    printf("[%s, %s]\n", hum, temp);
    
    return result;
}
