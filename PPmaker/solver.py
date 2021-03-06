#!/usr/bin/env python3

from ortools.sat.python import cp_model
import json
import sys

with open('PPmaker/PP.json') as a:
    arrays = json.load(a)

# The order of the projective plane to be arranged
N = len(arrays[0])-1

class CellsSolutionPrinter(cp_model.CpSolverSolutionCallback):
    def __init__(self, cells):
        cp_model.CpSolverSolutionCallback.__init__(self)
        self.__cells = cells
        self.__solution_count = 0

    # Called when a solution is found 
    def on_solution_callback(self):
        self.__solution_count += 1
        solutionset = []
        for row in self.__cells:
            # TODO: set printing here as an option, already done in main
            print([self.Value(cell) for cell in row])
            solutionset.append([self.Value(cell) for cell in row])
        # export the solved set to json
        with open('PPmaker/arrangedPP.json', 'w') as f:
            json.dump(solutionset, f)
        print() # TODO: set printing as option

    def solution_count(self):
        return self.__solution_count


def main():
    # Creates the solver
    model = cp_model.CpModel()

    # Create the variables
    cells = []
    for i in range(N * N + N + 1):
        row = []
        for j in range(N + 1):
            row.append(model.NewIntVarFromDomain(cp_model.Domain.FromValues(arrays[i]), 'cell: %i' % (i * (N + 1) + j)))
        cells.append(row)

    # Add the constraints
    # All rows must be different
    for row in cells:
        model.AddAllDifferent(row)

    # All columns must be different
    for i in range(N + 1):
        column = []
        for j in range(N * N + N + 1):
            column.append(cells[j][i])
        model.AddAllDifferent(column)

    # Arrange first row in increasing order
    for i, row in enumerate(cells):
        if i == 0:
            for j in range(N+1):
                model.Add(row[j] == 2 + j*N)
        else:
            model.AddAllDifferent(row)

    # Solve the model
    solver = cp_model.CpSolver()
    solution_printer = CellsSolutionPrinter(cells)

    # Enumerate all solutions if asked to
    if len(sys.argv) > 1:
        if sys.argv[1] == 'enum':
            solver.parameters.enumerate_all_solutions = True
        else:
            solver.parameters.enumerate_all_solutions = False

    print(solver.Solve(model, solution_printer))

    # Statistics
    print('\nStatistics')
    print(f'  conflicts      : {solver.NumConflicts()}')
    print(f'  branches       : {solver.NumBranches()}')
    print(f'  wall time      : {solver.WallTime()} s')

    if solver.parameters.enumerate_all_solutions:
        print(f'  solution count : {solution_printer.solution_count()}')


if __name__ == '__main__':
    main()
