#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#include "hanoi.h"
#include "autosolve.h"
/*
 * Name :  cndoit18
 *
 * Simple Towers of Hanoi game. You can solve it yourself or
 * ask the game to solve it for you.
 */

int InputPin(char *msg, int fm, int to);
bool CheckDone(int tower[NumPins][NumDisks]);

/*
 * The program main entry point
 */
int main()
{
    int tower[NumPins][NumDisks];
    int fm, to;

    /* Reset the tower to all disks on the first pin */
    Reset(tower);
    printf("\n");

    do
    {
        /* Display the tower before the move */
        DisplayTower(tower);

        /* Get a move from the user */
        fm = InputPin("Pin to move from: ", 0, 3);
        if (fm == 0)
        {
            Autosolve(tower);
            break;
        }

        to = InputPin("Pin to move to: ", 1, 3);

        /* Perform the move the user asked */
        MoveDisk(tower, fm, to);

    } while (!CheckDone(tower));

    return 0;
}

/*
 * Get a pin from the user
 */
int InputPin(char *msg, int fm, int to)
{
    int val;

    /* Allow values in the range [fm, to] */
    do
    {
        printf("%s", msg);
        scanf("%d", &val);
        if (val < fm || val > to)
            val = fm - 1;

    } while (val < fm);

    return val;
}

/*
 * Check to see if all disks moved to last pin
 */
bool CheckDone(int tower[NumPins][NumDisks])
{
    return tower[0][0] == 0 && tower[1][0] == 0;
}

/*
 * Display a disk of width p
 */
void NumPinsDisplayDisk(int p)
{
    int i;
    int spaceAround = (MaxDiskSize - p) / 2;

    for (i = 0; i < spaceAround; i++)
        printf(" ");

    if (p == 0)
    {
        printf("|");
    }
    else
    {
        for (i = 0; i < p; i++)
            printf("O");
    }

    for (i = 0; i < spaceAround; i++)
        printf(" ");
}
