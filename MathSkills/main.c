#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#define MAX_SIZE 10000  

int compare(const void *a, const void *b) { return (*(int *)a - *(int *)b); }

int main(int argc, char *argv[]) {
    if (argc != 2) { printf("Usage: %s <file_path>\n", argv[0]);
        return 1; }

    FILE *file = fopen(argv[1], "r");
    if (!file) { perror("Error opening file");
        return 1; }

    int numbers[MAX_SIZE], count = 0;
    long long sum = 0;

    while (fscanf(file, "%d", &numbers[count]) == 1 && count < MAX_SIZE) { sum += numbers[count++]; }
    
    fclose(file);

    if (count == 0) { printf("Error: No data found in file\n");
        return 1; }

    long long mean = round((double)sum / count);
    
    qsort(numbers, count, sizeof(int), compare);
    long long median = count % 2 ? numbers[count / 2] 
                                 : round((numbers[count / 2 - 1] + numbers[count / 2]) / 2.0);

    long long var_sum = 0;
    for (int i = 0; i < count; i++) 
        var_sum += (long long)(numbers[i] - mean) * (numbers[i] - mean);

    long long variance = round((double)var_sum / count); 
    long long stddev = (long long)round(sqrt(variance));

    printf("Average: %lld\n", mean);
    printf("Median: %lld\n", median);
    printf("Variance: %lld\n", variance);
    printf("Standard Deviation: %lld\n", stddev);

    return 0;
}
