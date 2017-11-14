local cjson = require "cjson"
local baseFile = ngx.var.uri:match("[^/]+$")

local function FormatHeaderKey(key)
  local fk = ""
  fk = string.gsub(key, "^%l", string.upper)
  fk = string.gsub(fk, "(-.)", function(a) return string.upper(a) end)
  return fk 
end

local metaFile = io.open(ngx.var.document_root .. "/_" .. baseFile, "r")
if metaFile == nil then 
  ngx.exec("@no_meta", "") 
end

local raw = metaFile:read("*all")
io.close(metaFile)

local meta = cjson.decode(raw)
ngx.var.lua_obj_uri = "/" .. meta["object"]
local mime = meta["mime_type"]

if ngx.var.lua_obj_uri then
  for k, v in pairs(meta) do
    ngx.header['PC-' .. FormatHeaderKey(k)] = v
  end

  ngx.exec("@serve_obj", "")
else
  ngx.exec("@no_meta", "")
end
