/* 
 * trans.c - Matrix transpose B = A^T
 *
 * Each transpose function must have a prototype of the form:
 * void trans(int M, int N, int A[N][M], int B[M][N]);
 *
 * A transpose function is evaluated by counting the number of misses
 * on a 1KB direct mapped cache with a block size of 32 bytes.
 */
#include <stdio.h>
#include "cachelab.h"

int is_transpose(int M, int N, int A[N][M], int B[M][N]);

/* 
 * transpose_submit - This is the solution transpose function that you
 *     will be graded on for Part B of the assignment. Do not change
 *     the description string "Transpose submission", as the driver
 *     searches for that string to identify the transpose function to
 *     be graded. 
 */
char transpose_submit_desc[] = "Transpose submission";
void transpose_submit(int M, int N, int A[N][M], int B[M][N])
{
    if (M == 32)
    {
        int v1, v2, v3, v4, v5, v6, v7, v8;
        for (int n = 0; n < N; n += 8)
        {
            for (int m = 0; m < M; m += 8)
            {
                for (int k = n; k < n + 8; k++)
                {
                    v1 = A[k][m];
                    v2 = A[k][m + 1];
                    v3 = A[k][m + 2];
                    v4 = A[k][m + 3];
                    v5 = A[k][m + 4];
                    v6 = A[k][m + 5];
                    v7 = A[k][m + 6];
                    v8 = A[k][m + 7];
                    B[m][k] = v1;
                    B[m + 1][k] = v2;
                    B[m + 2][k] = v3;
                    B[m + 3][k] = v4;
                    B[m + 4][k] = v5;
                    B[m + 5][k] = v6;
                    B[m + 6][k] = v7;
                    B[m + 7][k] = v8;
                }
            }
        }
    }

    if (M == 64)
    {
        int v1, v2, v3, v4;
        for (int n = 0; n < N; n += 4)
        {
            for (int m = 0; m < M; m += 4)
            {
                for (int k = n; k < n + 4; k++)
                {
                    v1 = A[k][m];
                    v2 = A[k][m + 1];
                    v3 = A[k][m + 2];
                    v4 = A[k][m + 3];
                    B[m][k] = v1;
                    B[m + 1][k] = v2;
                    B[m + 2][k] = v3;
                    B[m + 3][k] = v4;
                }
            }
        }
    }
    else
    {
        int i, j, tmp;

        for (i = 0; i < N; i++)
        {
            for (j = 0; j < M; j++)
            {
                tmp = A[i][j];
                B[j][i] = tmp;
            }
        }
    }
}

/* 
 * You can define additional transpose functions below. We've defined
 * a simple one below to help you get started. 
 */

/* 
 * trans - A simple baseline transpose function, not optimized for the cache.
 */
char trans_desc[] = "Simple row-wise scan transpose";
void trans(int M, int N, int A[N][M], int B[M][N])
{
    int i, j, tmp;

    for (i = 0; i < N; i++)
    {
        for (j = 0; j < M; j++)
        {
            tmp = A[i][j];
            B[j][i] = tmp;
        }
    }
}

/*
 * registerFunctions - This function registers your transpose
 *     functions with the driver.  At runtime, the driver will
 *     evaluate each of the registered functions and summarize their
 *     performance. This is a handy way to experiment with different
 *     transpose strategies.
 */
void registerFunctions()
{
    /* Register your solution function */
    registerTransFunction(transpose_submit, transpose_submit_desc);

    /* Register any additional transpose functions */
    registerTransFunction(trans, trans_desc);
}

/* 
 * is_transpose - This helper function checks if B is the transpose of
 *     A. You can check the correctness of your transpose by calling
 *     it before returning from the transpose function.
 */
int is_transpose(int M, int N, int A[N][M], int B[M][N])
{
    int i, j;

    for (i = 0; i < N; i++)
    {
        for (j = 0; j < M; ++j)
        {
            if (A[i][j] != B[j][i])
            {
                return 0;
            }
        }
    }
    return 1;
}
