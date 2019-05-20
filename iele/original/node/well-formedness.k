





require "iele-syntax.k"
require "data.k"











module IELE-CONSTANTS
    syntax Mode ::= "NORMAL" [klabel(NORMAL), symbol] | "VMTESTS"
    syntax Schedule ::= "ALBE"
                      | "DANSE"
endmodule


module IELE-WELL-FORMEDNESS
    imports IELE-COMMON
    imports IELE-DATA
    imports IELE-CONSTANTS
    imports DOMAINS
    imports DEFAULT-CONFIGURATION








    syntax IeleName ::= "Main" [token]
    syntax Schedule

    configuration <well-formedness>
                    <typeChecking> false </typeChecking>
                    <well-formedness-schedule> $SCHEDULE:Schedule </well-formedness-schedule>
                    <contracts> .Set </contracts>
                    <currentContract>
                      <types> intrinsicTypes </types>
                      <contractName> Main </contractName>
                      <declaredContracts> .Set </declaredContracts>
                      <functionBodies> .K </functionBodies>
                      <currentFunction>
                        <functionName> deposit </functionName>
                        <labels> .Set </labels>
                        <currentInstructions> .K </currentInstructions>
                      </currentFunction>
                    </currentContract>
                  </well-formedness>








    syntax Type ::= "int" | Types "->" ReturnType [klabel(funType)]
    syntax Types ::= List{Type,","} [klabel(typeList)]
    syntax ReturnType ::= Types | "unknown"
    syntax priorities typeList > funType

    syntax Types ::= ints(Int) [function]
 // -------------------------------------
    rule ints(0) => .Types
    rule ints(N) => int , ints(N -Int 1) [owise]






    rule CONTRACT1::ContractDefinition CONTRACT2::ContractDefinition CONTRACTS => CONTRACT1 ~> CONTRACT2 CONTRACTS
    rule CONTRACT::ContractDefinition .Contract => CONTRACT

    rule <k> contract NAME { DEFINITIONS } => checkName(NAME) ~> DEFINITIONS ... </k>
         <contracts> CONTRACTS => CONTRACTS SetItem(NAME) </contracts>
         (_:CurrentContractCell => <currentContract>
           <contractName> NAME </contractName>
           ...
         </currentContract>)
      requires notBool NAME in CONTRACTS

    rule DEF::TopLevelDefinition DEFS => DEF ~> DEFS
    rule <k> .TopLevelDefinitions => BODIES ... </k>
         <functionBodies> BODIES </functionBodies>
         <types> ... init |-> _ -> .Types </types>






    rule <k> external contract NAME => . ... </k>
         <contracts> CONTRACTS </contracts>
         <declaredContracts> DECLARED => DECLARED SetItem(NAME) </declaredContracts>
      requires NAME in CONTRACTS andBool notBool NAME in DECLARED

    rule <k> (@ NAME = _)::GlobalDefinition => checkName(NAME) ... </k>
         <types> TYPES => TYPES NAME |-> int </types>
      requires notBool NAME in_keys(TYPES)

    rule <k> define @ init ( ARGS::FunctionParameters ) { BLOCKS } => . ... </k>
         <types> TYPES => TYPES init |-> ints(#sizeNames(ARGS)) -> .Types </types>
         <functionBodies> .K => processFunction(init) ~> BLOCKS ... </functionBodies>
      requires notBool init in_keys(TYPES)

    rule <k> define @ NAME ( ARGS::FunctionParameters ) { BLOCKS } => checkName(NAME) ~> checkArgs(ARGS) ... </k>
         <types> TYPES => TYPES NAME |-> (ints(#sizeNames(ARGS)) -> unknown) </types>
         <functionBodies> .K => processFunction(NAME) ~> BLOCKS ... </functionBodies>
      requires notBool NAME in_keys(TYPES) andBool NAME =/=K init

    rule <k> define public @ NAME ( ARGS::FunctionParameters ) { BLOCKS } => checkName(NAME) ~> checkArgs(ARGS) ... </k>
         <types> TYPES => TYPES NAME |-> ints(#sizeNames(ARGS)) -> unknown </types>
         <functionBodies> .K => processFunction(NAME) ~> BLOCKS ... </functionBodies>
      requires notBool NAME in_keys(TYPES) andBool NAME =/=K init

    syntax KItem ::= processFunction(IeleName)
 // ------------------------------------------
    rule <k> processFunction(NAME) => . ... </k>
         (_:CurrentFunctionCell => <currentFunction>
           <functionName> NAME </functionName>
           ...
         </currentFunction>)

    syntax FunctionParameters ::= Int
    syntax Int ::= #sizeNames(FunctionParameters) [function]
 // --------------------------------------------------------
    rule #sizeNames(I:Int) => I
    rule #sizeNames(.LocalNames) => 0
    rule #sizeNames(N , NAMES) => 1 +Int #sizeNames(NAMES)

    syntax KItem ::= checkArgs(FunctionParameters)
                   | checkNameArgs(LocalNames)
                   | checkIntArgs(LocalNames, Int)
 // ----------------------------------------------
    rule checkArgs(.LocalNames) => .
    rule checkArgs(_:Int) => . 
    rule checkArgs(% N:NumericIeleName , ARGS) => checkIntArgs(% N, ARGS, 0)
    rule checkArgs(% N, ARGS) => checkNameArgs(ARGS) requires notBool isNumericIeleName(N)

    rule checkNameArgs(% N, ARGS) => checkName(N) ~> checkNameArgs(ARGS) requires notBool isNumericIeleName(N)
    rule checkNameArgs(.LocalNames) => .

    rule checkIntArgs( % N , ARGS , I) => checkIntArgs(ARGS, I +Int 1)
      requires String2Int(IeleName2String(N)) ==Int I
    rule checkIntArgs(.LocalNames, _) => .






    rule <k> BLOCK:UnlabeledBlock BLOCKS => BLOCKS ... </k>
         <typeChecking> true </typeChecking>
         <currentInstructions> .K => BLOCK ... </currentInstructions>
    rule <k> BLOCK::LabeledBlock BLOCKS => BLOCK ~> BLOCKS ... </k>
         <typeChecking> true </typeChecking>
    rule <k> .LabeledBlocks => INSTRS ... </k>
         <typeChecking> true </typeChecking>
         <currentInstructions> INSTRS </currentInstructions>

    rule <k> NAME : BLOCK::Instructions => . ... </k>
         <typeChecking> true </typeChecking>
         <labels> LABELS => LABELS SetItem(NAME) </labels>
         <currentInstructions> .K => BLOCK ... </currentInstructions>
      requires notBool NAME in LABELS

    rule <k> INSTR::Instruction INSTRS::Instructions => check ~> INSTR ~> INSTRS ... </k>
         <typeChecking> true </typeChecking>
    rule .Instructions => .










    syntax KResult ::= Type
                     | Types
    syntax KItem ::= "check"
 // ------------------------

    rule check ~> LVAL = OP1 => checkLVal(LVAL) ~> checkOperand(OP1)
    rule check ~> LVAL = load OP1 => checkLVal(LVAL) ~> checkOperand(OP1)
    rule check ~> LVAL = load OP1, OP2, OP3 => checkLVal(LVAL) ~> checkOperands(OP1, OP2, OP3)
    rule check ~> store OP1, OP2 => checkOperands(OP1, OP2)
    rule check ~> store OP1, OP2, OP3, OP4 => checkOperands(OP1, OP2, OP3, OP4)

    rule check ~> LVAL = sload OP1 => checkLVal(LVAL) ~> checkOperand(OP1)
    rule check ~> sstore OP1, OP2 => checkOperands(OP1, OP2)

    rule check ~> LVAL = iszero OP1 => checkLVal(LVAL) ~> checkOperand(OP1)
    rule check ~> LVAL = not    OP1 => checkLVal(LVAL) ~> checkOperand(OP1)

    rule check ~> LVAL = add OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = mul OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = sub OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = div OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = exp OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = mod OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)

    rule check ~> LVAL = addmod OP1, OP2, OP3 => checkLVal(LVAL) ~> checkOperands(OP1, OP2, OP3)
    rule check ~> LVAL = mulmod OP1, OP2, OP3 => checkLVal(LVAL) ~> checkOperands(OP1, OP2, OP3)
    rule check ~> LVAL = expmod OP1, OP2, OP3 => checkLVal(LVAL) ~> checkOperands(OP1, OP2, OP3)

    rule check ~> LVAL = byte  OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = sext  OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = twos  OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = bswap OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)

    rule check ~> LVAL = log2 OP1 => checkLVal(LVAL) ~> checkOperands(OP1)

    rule check ~> LVAL = and   OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = or    OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = xor   OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)
    rule check ~> LVAL = shift OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)

    rule check ~> LVAL = cmp _ OP1, OP2 => checkLVal(LVAL) ~> checkOperands(OP1, OP2)

    rule check ~> LVAL = sha3 OP1 => checkLVal(LVAL) ~> checkOperand(OP1)
    rule check ~> log OP1                     => checkOperand(OP1)
    rule check ~> log OP1, OP2                => checkOperands(OP1, OP2)
    rule check ~> log OP1, OP2, OP3           => checkOperands(OP1, OP2, OP3)
    rule check ~> log OP1, OP2, OP3, OP4      => checkOperands(OP1, OP2, OP3, OP4)
    rule check ~> log OP1, OP2, OP3, OP4, OP5 => checkOperands(OP1, OP2, OP3, OP4, OP5)

    rule check ~> revert OP1 => checkOperand(OP1)
    rule check ~> selfdestruct OP1 => checkOperand(OP1)







    rule <k> check ~> br NAME => . ... </k>
         <labels> ... SetItem(NAME) </labels>

    rule <k> check ~> br OP1, NAME => checkOperand(OP1) ... </k>
         <labels> ... SetItem(NAME) </labels>







    rule <k> check ~> RETS = call @ NAME ( ARGS ) => checkLVals(RETS) ~> checkOperands(ARGS) ... </k>
         <types> ... NAME |-> ARGTYPES -> RETTYPES:Types </types>
         <well-formedness-schedule> SCHED </well-formedness-schedule>
      requires ints(#sizeRegs(ARGS)) ==K ARGTYPES andBool ints(#sizeLVals(RETS)) ==K RETTYPES andBool checkInit(NAME, SCHED)

    rule <k> check ~> RETS = call @ NAME ( ARGS ) => checkLVals(RETS) ~> checkOperands(ARGS) ... </k>
         <types> ... NAME |-> ARGTYPES -> (unknown => ints(#sizeLVals(RETS))) </types>
         <well-formedness-schedule> SCHED </well-formedness-schedule>
      requires ints(#sizeRegs(ARGS)) ==K ARGTYPES andBool checkInit(NAME, SCHED)

    rule <k> check ~> RETS = call % NAME ( ARGS ) => checkLVals(RETS) ~> checkOperands(ARGS) ... </k>

    rule check ~> STATUS, RETS = call NAME at OP1 ( ARGS ) send OP2 , gaslimit OP3 => checkLVals(STATUS, RETS) ~> checkOperands(OP1 , OP2 , OP3 , ARGS)
    rule check ~> STATUS, RETS = staticcall NAME at OP1 ( ARGS ) gaslimit OP2 => checkLVals(STATUS, RETS) ~> checkOperands(OP1 , OP2 , ARGS)

    rule check ~> RET = calladdress NAME at OP => checkLVal(RET) ~> checkOperand(OP)

    rule <k> check ~> ret OPS => checkOperands(OPS) ... </k>
         <functionName> NAME </functionName>
         <types> ... NAME |-> _ -> RETTYPES:Types </types>
      requires ints(#sizeRegs(OPS)) ==K RETTYPES

    rule <k> check ~> ret OPS => checkOperands(OPS) ... </k>
         <functionName> NAME </functionName>
         <types> ... NAME |-> _ -> (unknown => ints(#sizeRegs(OPS))) </types>







    rule <k> check ~> STATUS , RET = create NAME ( ARGS ) send OP1 => checkLVals(STATUS, RET) ~> checkOperands(OP1 , ARGS) ... </k>
         <declaredContracts> ... SetItem(NAME) </declaredContracts>

    rule check ~> STATUS , RET = copycreate OP1 ( ARGS ) send OP2 => checkLVals(STATUS, RET) ~> checkOperands(OP1 , OP2 , ARGS)










    syntax Map ::= "intrinsicTypes" [function]
 // ------------------------------------------
    rule intrinsicTypes =>
    (iele.invalid |-> .Types -> .Types
    (iele.gas |-> .Types -> int, .Types
    (iele.gasprice |-> .Types -> int, .Types
    (iele.gaslimit |-> .Types -> int, .Types
    (iele.beneficiary |-> .Types -> int, .Types
    (iele.timestamp |-> .Types -> int, .Types
    (iele.number |-> .Types -> int, .Types
    (iele.difficulty |-> .Types -> int, .Types
    (iele.address |-> .Types -> int, .Types
    (iele.origin |-> .Types -> int, .Types
    (iele.caller |-> .Types -> int, .Types
    (iele.callvalue |-> .Types -> int, .Types
    (iele.msize |-> .Types -> int, .Types
    (iele.codesize |-> .Types -> int, .Types
    (iele.blockhash |-> int -> int, .Types
    (iele.balance |-> int -> int, .Types
    (iele.extcodesize |-> int -> int, .Types
    )))))))))))))))))
    








    syntax KItem ::= checkName(IeleName)
 // ------------------------------------
    rule checkName(NAME) => .
      requires lengthString(IeleName2String(NAME)) <Int 5 orBool substrString(IeleName2String(NAME), 0, 5) =/=String "iele."

    syntax Bool ::= checkInit(IeleName, Schedule) [function]
 // --------------------------------------------------------
    rule checkInit(init, SCHED) => SCHED =/=K ALBE
    rule checkInit(...) => true [owise]





    syntax String ::= unescape(String) [function]
                    | unescape(String, Int, StringBuffer) [function, klabel(unescapeAux)]
 // -------------------------------------------------------------------------------------
    rule unescape(S) => unescape(S, 1, .StringBuffer)
    rule unescape(S, IDX, SB) => unescape(S, IDX +Int 1, SB +String substrString(S, IDX, IDX +Int 1))
      requires IDX <Int lengthString(S) -Int 1 andBool substrString(S, IDX, IDX +Int 1) =/=K "\\"
    rule unescape(S, IDX, SB) => unescape(S, IDX +Int 3, SB +String chrChar(String2Base(substrString(S, IDX +Int 1, IDX +Int 3), 16)))
      requires IDX <Int lengthString(S) -Int 1 andBool substrString(S, IDX, IDX +Int 1) ==K "\\"
    rule unescape(S, IDX, SB) => StringBuffer2String(SB)
      requires IDX ==Int lengthString(S) -Int 1
    rule `StringIeleName`(NAME:StringIeleName) => String2IeleName(unescape(StringIeleName2String(NAME)))






    syntax KItem ::= checkOperand(Operand)
                   | checkOperands(Operands)
 // ----------------------------------------
    rule checkOperands(OP , OPS) => checkOperand(OP) ~> checkOperands(OPS)
    rule checkOperands(.Operands) => .

    rule checkOperand(% NAME) => .
    rule checkOperand(_:IntConstant) => .
    rule checkOperand(@ NAME) => .






    syntax KItem ::= checkLVal(LValue)
                   | checkLVals(LValues)
 // --------------------------------------
    rule checkLVals(LVAL , LVALS) => checkLVal(LVAL) ~> checkLVals(LVALS)
    rule checkLVals(.LValues) => .

    rule checkLVal(% NAME) => checkName(NAME)
endmodule

module IELE-WELL-FORMEDNESS-STANDALONE
    imports DEFAULT-CONFIGURATION
    imports IELE-WELL-FORMEDNESS

    configuration <k/> <well-formedness/> <exit-code exit=""> 1 </exit-code>

    rule <typeChecking> false => true </typeChecking>
    rule <k> . </k> <exit-code> 1 => 0 </exit-code>

endmodule

