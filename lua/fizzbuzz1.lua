local function f(n)
    for i = 1, n, 1 do
        if i % 15 == 0 then
            print("FizzBuzz")
        elseif i % 3 == 0 then
            print("Fizz")
        elseif i % 5 == 0 then
            print("Buzz")
        else
            print(i)
        end
    end
end

f(30)
