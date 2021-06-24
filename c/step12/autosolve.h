#include "hanoi.h"

#ifndef AUTOSOLVE_H
#define AUTOSOLVE_H

void Autosolve(int tower[NumPins][NumDisks]);
void AutoMove(int tower[NumPins][NumDisks], int num, int fm, int to);
void DisplayTower(int tower[NumPins][NumDisks]);
void MoveDisk(int tower[NumPins][NumDisks], int fm, int to);
void Reset(int tower[NumPins][NumDisks]);

#endif