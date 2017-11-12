local cjson = require "cjson"
local baseFile = ngx.var.uri:match("[^/]+$")
local metaFile = io.open(ngx.var.document_root .. "/_" .. baseFile, "r")
local raw = metaFile:read("*all")
io.close(metaFile)

local meta = cjson.decode(raw)
ngx.var.lua_obj_uri = meta["object"]
local mime = meta["mime_type"]

if ngx.var.lua_obj_uri then
  ngx.exec(@serve_obj)
else
  ngx.log(ngx.ERR, "obj invalid.")
  ngx.exec("@no_meta", "")
end
