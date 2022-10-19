import time

def custom_range(start=0,stop=0,step=1):
    currentValue=start
    if stop>start:
        while currentValue<stop:
            yield currentValue
            currentValue+=step
    elif stop<start:
        while currentValue>stop:
            yield currentValue
            currentValue+=step
    else:
        yield currentValue
        currentValue+=step
    

if __name__=='__main__':

    start_time = time.time()
    for index in range(2,5):
        print(index)
    print("range function took", time.time() - start_time, "to run")

    start_time = time.time()
    for index in custom_range(2,5):
        print(index)
    print("custom_range function took", time.time() - start_time, "to run")
    

#Testcases
# for index in custom_range(6,0,-2):
#     print(index)
# for index in custom_range(6,-7,-2):
#     print(index)

# Output:
# 2
# 3
# 4
# range function took 2.193450927734375e-05 to run
# 2
# 3
# 4
# custom_range function took 1.3828277587890625e-05 to run
