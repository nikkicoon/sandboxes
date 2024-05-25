local function gcd (p, q)
    if q == 0 then return p end
    local r = p % q
    -- print("remainder:" .. r)
    gcd(q, r)
end

print(gcd(5,25))
