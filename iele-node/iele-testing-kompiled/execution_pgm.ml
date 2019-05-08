
open Prelude
open Constants
open Constants.K
open Run
let () = Sys.catch_break true
let () = Gc.set { (Gc.get()) with Gc.minor_heap_size = 33554432 }let () = CONFIG.set_sys_argv ()
external load_kore_term : unit -> string = "load_kore_term"
external load_marshal_term : unit -> string = "load_marshal_term"
let unserializedKMap = match Lexer.parse_k_binary_string (load_kore_term ()) with
  [Map(SortMap, Lbl_Map_, m)] -> m
| _ -> failwith "kore_term is not of sort Map"
let serialized = (Marshal.from_string (load_marshal_term ()) 0 : Prelude.k)
let serializedKMap = match serialized with
  [Map(SortMap, Lbl_Map_, m)] -> m
| _ -> failwith "marshal_term is not of sort Map"
let completeMap = let conflict key val1 _ = Some val1 in
    [(Map (SortMap, Lbl_Map_, KMap.union conflict unserializedKMap serializedKMap))]
let input = (let module Def = (val Plugin.get ()) in Def.eval (KApply(LblinitGeneratedTopCell, [completeMap])) interned_bottom)
let try_match (c: k) (config: k) (guard: int) : k Subst.t = match c with 
(*{| rule `<generatedTop>`(_1,_2,`<exit-code>`(_0),_3,_4,_5,_6,_7,_8) requires isInt(_0) ensures #token("true","Bool") []|}*)
| (KApply9(Lbl'_LT_'generatedTop'_GT_',(var_1_16690),(var_2_16691),(KApply1(Lbl'_LT_'exit'Hyph'code'_GT_',((Int _ as var_0_16692) :: [])) :: []),(var_3_16693),(var_4_16694),(var_5_16695),(var_6_16696),(var_7_16697),(var_8_16698)) :: []) when true && (true) -> (Subst.add "_2" var_2_16691 (Subst.add "_3" var_3_16693 (Subst.add "_8" var_8_16698 (Subst.add "_1" var_1_16690 (Subst.add "_6" var_6_16696 (Subst.add "_7" var_7_16697 (Subst.add "_5" var_5_16695 (Subst.add "_0" [var_0_16692] (Subst.add "_4" var_4_16694 Subst.empty)))))))))
| _ -> raise(Stuck c)
let _ = try let res, _ = run_no_thread_opt(input) (!CONFIG.depth) in let subst = try_match res res (-1) in
let code = get_exit_code subst in
exit code
with Stuck(res) -> (prerr_endline "Execution failed (configuration dumped)";
let out = open_out !CONFIG.output_file in output_string out (print_k res);
exit 139)
