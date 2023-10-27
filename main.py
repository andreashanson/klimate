import time
devider1 = 8
devider2 = 3
start_time = time.time()
for val in range(1000000):
    if val%devider1 == 0 and val%devider2 == 0:
        print("FIZZBUSS", val)
    elif val%devider1 == 0:
        print("BUSS", val)
    elif val%devider2 == 0:
        print("FIZZ", val)
end_time = time.time()
print("elapsed")
print(end_time - start_time)
