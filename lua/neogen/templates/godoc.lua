local i = require("neogen.types.template").item

return {
    { nil, " $1", { no_results = true } },
    { "func_name", "// %s", { type = { "func" } } },
    { "func_name", "// @desc $1", { type = { "func" } } },
    { "type_name", " %s $1", { type = { "type" } } },
    { "package_name", " Package %s $1", { type = { "file" } } },
    { i.Parameter,[[// @param %s]] },
    { i.Return, "// @return %s" },
}
