import matplotlib.pyplot as plt
import numpy as np
import csv

day = []
temp = []

# CSV reader
with open('data.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    line_count = 0
    for row in csv_reader:
        if line_count == 0:
            #column_names = .join(row)
            line_count += 1
        else: 
            day.append(int(row[0]))
            temp.append(int(row[1]))
            line_count += 1


print(day)

print(temp)

# Check if Graph valid
if len(day) != len(temp):
    print("Graph not bijective.")

# Graph 
plt.plot(day,temp) # Plot some data on the axes.
plt.xlabel('day')
plt.ylabel('temperature')
plt.show()


