#!/usr/bin/env lua
print("this is " .. arg[0])

print("length of arg is " .. #arg)

for i = 0, #arg do
    print(arg[i])
end
