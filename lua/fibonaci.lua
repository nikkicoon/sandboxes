local q = require("memoization")
local memoize = q.memoize

local function fibo (n)
    if n < 2 then return n end
    return fibo(n - 1) + fibo(n - 2)
end

local function fib(n)
    local f = memoize(fib)
    if n < 2 then return n end
    return f(n - 1) + f(n - 2)
end


print("fibo with cache (fastFib): " .. fib(10))
--local nClock = os.clock()
--print(fib(10))
--print("elapsed time: " .. os.clock() - nClock)
local f = memoize(fibo(12))
print("with memoize cache: ")
--local pClock = os.clock()
--print("elapsed time: " .. os.clock() - pClock)
print(f)
print("without memoize: " .. fibo(10))
--local qClock = os.clock()
--print(fibo(10))
--print("elapsed time: " .. os.clock() - qClock)
