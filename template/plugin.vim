let s:save_cpo = &cpoptions
set cpoptions&vim

if exists('g:loaded_{{ .Name }}')
  finish
endif
let g:loaded_{{ .Name }} = 1

let &cpoptions = s:save_cpo
unlet s:save_cpo
