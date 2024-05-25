---
--- nikita, 2024, public domain
--- simple example for reusable memoization caching function in lua
---
local M = {}

function M.memoize(fn)
    local cache = {}
    return function (...)
        cache[...] = cache[...] or fn(...)
        return cache[...]
    end
end

return M