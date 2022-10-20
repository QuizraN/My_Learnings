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
    [print(i) for i in range(2,5)]
    print("range function took", time.time() - start_time, "to run")

    start_time = time.time()
    [print(i) for i in custom_range(2,5)]
    print("custom_range function took", time.time() - start_time, "to run")

