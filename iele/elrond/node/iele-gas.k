







module IELE-GAS
    imports IELE-DATA
    imports IELE-CONFIGURATION
    imports IELE-COMMON
    imports IELE-INFRASTRUCTURE
    imports IELE-PRECOMPILED












    syntax InternalOp ::= "#deductGas"
 // ----------------------------------
    rule <k> #gas [ OP ] => #memory [ OP ] ~> #compute [ OP , SCHED ] ~> #deductGas ... </k> <schedule> SCHED </schedule>

    rule <k> G:Int ~> #deductGas => #exception OUT_OF_GAS ... </k> <gas> GAVAIL                  </gas> requires GAVAIL <Int G
    rule <k> G:Int ~> #deductGas => .                     ... </k> <gas> GAVAIL => GAVAIL -Int G </gas> <previousGas> _ => GAVAIL </previousGas> requires GAVAIL >=Int G





























    syntax InternalOp ::= "#memory" "[" Instruction "]"
 // ---------------------------------------------------





















    rule #memory [ REG = not   W       ] => #registerDelta(REG, intSize(W))
    rule #memory [ REG = and   W0 , W1 ] => #registerDelta(REG, minInt(intSize(W0), intSize(W1)))
    rule #memory [ REG = or    W0 , W1 ] => #registerDelta(REG, maxInt(intSize(W0), intSize(W1)))
    rule #memory [ REG = xor   W0 , W1 ] => #registerDelta(REG, maxInt(intSize(W0), intSize(W1)))
    rule <k> #memory [ REG = shift W0 , W1 ] => #registerDelta(REG, maxInt(1, intSize(W0) +Int bitsInWords(W1, SCHED))) ... </k> <schedule> SCHED </schedule>
    rule #memory [ REG = log2  W       ] => #registerDelta(REG, 2)







    rule #memory [ REG = iszero _     ] => #registerDelta(REG, 1)
    rule #memory [ REG = cmp _  _ , _ ] => #registerDelta(REG, 1)

















    rule #memory [ REG = add W0 , W1 ] => #registerDelta(REG, maxInt(intSize(W0), intSize(W1)) +Int 1)
    rule #memory [ REG = sub W0 , W1 ] => #registerDelta(REG, maxInt(intSize(W0), intSize(W1)) +Int 1)
    rule #memory [ REG = mul W0 , W1 ] => #registerDelta(REG, intSize(W0) +Int intSize(W1))
    rule #memory [ REG = div W0 , W1 ] => #registerDelta(REG, maxInt(1, intSize(W0) -Int intSize(W1) +Int 1))
    rule #memory [ REG = mod W0 , W1 ] => #registerDelta(REG, minInt(intSize(W0), intSize(W1)))
    rule #memory [ REG = exp W0 , W1 ] => #registerDelta(REG, #adjustedBitLength(intSize(W0), W0) *Int W1 /Int 64)








    rule #memory [ REG = addmod _  , _  , W2 ] => #registerDelta(REG, intSize(W2))
    rule #memory [ REG = mulmod W0 , W1 , W2 ] => #registerDelta(REG, intSize(W2))
    rule #memory [ REG = expmod _  , _  , W2 ] => #registerDelta(REG, intSize(W2))







    rule <k> #memory [ REG = sha3 _ ] => #registerDelta(REG, bitsInWords(256, SCHED)) ... </k> <schedule> SCHED </schedule>









    rule #memory [ REG = byte  INDEX , _ ] => #registerDelta(REG, bytesInWords(1))
    rule #memory [ REG = sext  WIDTH , _ ] => #registerDelta(REG, bytesInWords(chop(WIDTH)))
    rule #memory [ REG = twos  WIDTH , _ ] => #registerDelta(REG, bytesInWords(chop(WIDTH)))
    rule #memory [ REG = bswap WIDTH , _ ] => #registerDelta(REG, bytesInWords(chop(WIDTH)))







    rule #memory [ REG = call @iele.gas         ( .Ints ) ] => #registerDelta(REG, 1)
    rule #memory [ REG = call @iele.gasprice    ( .Ints ) ] => #registerDelta(REG, 1)
    rule #memory [ REG = call @iele.gaslimit    ( .Ints ) ] => #registerDelta(REG, 1)
    rule #memory [ REG = call @iele.number      ( .Ints ) ] => #registerDelta(REG, 1)
    rule #memory [ REG = call @iele.msize       ( .Ints ) ] => #registerDelta(REG, 1)
    rule #memory [ REG = call @iele.codesize    ( .Ints ) ] => #registerDelta(REG, 1)
    rule #memory [ REG = call @iele.extcodesize ( _     ) ] => #registerDelta(REG, 1)
    rule #memory [ REG = calladdress _ at _               ] => #registerDelta(REG, 1)





    rule #memory [ REG = call @iele.beneficiary ( .Ints ) ] => #registerDelta(REG, bytesInWords(20))
    rule #memory [ REG = call @iele.address     ( .Ints ) ] => #registerDelta(REG, bytesInWords(20))
    rule #memory [ REG = call @iele.origin      ( .Ints ) ] => #registerDelta(REG, bytesInWords(20))
    rule #memory [ REG = call @iele.caller      ( .Ints ) ] => #registerDelta(REG, bytesInWords(20))





    rule <k> #memory [ REG = call @iele.timestamp   ( .Ints ) ] => #registerDelta(REG, bitsInWords(256, SCHED)) ... </k> <schedule> SCHED </schedule>
    rule <k> #memory [ REG = call @iele.difficulty  ( .Ints ) ] => #registerDelta(REG, bitsInWords(256, SCHED)) ... </k> <schedule> SCHED </schedule>
    rule <k> #memory [ REG = call @iele.callvalue   ( .Ints ) ] => #registerDelta(REG, bitsInWords(256, SCHED)) ... </k> <schedule> SCHED </schedule>
    rule <k> #memory [ REG = call @iele.blockhash   ( _     ) ] => #registerDelta(REG, bitsInWords(256, SCHED)) ... </k> <schedule> SCHED </schedule>
    rule <k> #memory [ REG = call @iele.balance     ( _     ) ] => #registerDelta(REG, bitsInWords(256, SCHED)) ... </k> <schedule> SCHED </schedule>








    rule <k> #memory [ DEST = % SRC:Int ] => #registerDelta(DEST, intSize(getInt(REGS [ SRC ]))) ... </k>
         <regs> REGS </regs>
    rule <k> #memory [ DEST = SRC:Int ] => #registerDelta(DEST, intSize(SRC)) ... </k>










    rule <k> #memory [ REGS = call @ NAME ( ARGS ) ] => #memoryDelta(REGISTERS -Int #sizeRegs(ARGS) +Int intSizes(ARGS) +Int Gcallmemory < SCHED >) ... </k>
         <schedule> SCHED </schedule>
         <funcId> NAME </funcId>
         <nregs> REGISTERS </nregs>

    rule <k> #memory [ _ = call (IDX:Int => @ FUNC) ( _ ) ] ... </k>
         <funcLabels> ... IDX |-> FUNC ... </funcLabels>
    // this will throw an exception, so the gas cost doesn't really matter
    rule <k> #memory [ _ = call IDX:Int ( _ ) ] => . ... </k>
         <funcLabels> LABELS </funcLabels>
      requires notBool IDX in_keys(LABELS)










    rule <k> #memory [ ret ARGS ] => #memoryDelta(0 -Int intSizes(REGS, NREGS, SCHED) -Int Gcallmemory < SCHED >) ... </k>
         <schedule> SCHED </schedule>
         <fid> NAME </fid>
         <regs> REGS </regs>
         <funcId> NAME </funcId>
         <nregs> NREGS </nregs>
















    rule #memory [ REG = load INDEX1 , INDEX2 , WIDTH ] => #registerDelta(REG, bytesInWords(chop(WIDTH)))
    rule #memory [ store _ ,  INDEX1 , INDEX2 , WIDTH ] => #memoryExpand(INDEX1, bytesInWords(chop(INDEX2) +Int chop(WIDTH))) requires chop(WIDTH) >Int 0
    rule #memory [ store _ ,  INDEX1 , INDEX2 , WIDTH ] => .K requires chop(WIDTH) ==Int 0

    rule <k> #memory [ REG = load INDEX ] => #registerDelta(REG, bytesInWords(lengthBytes(LM))) ... </k>
         <localMem>... INDEX |-> LM ...</localMem>
    rule <k> #memory [ REG = load INDEX ] => #registerDelta(REG, 0)... </k>
         <localMem> LM </localMem>
      requires notBool INDEX in_keys(LM)
    rule #memory [ store VALUE ,  INDEX ] => #memoryDelta(INDEX, intSize(VALUE))








    rule <k> #memory [ REG = sload INDEX ] => #registerDelta(REG, intSize(VALUE)) ... </k>
         <id> ACCT </id>
         <account>
           <acctID> ACCT </acctID>
           <storage> ... INDEX |-> VALUE ... </storage>
           ...
         </account>

    rule <k> (.K => #lookupStorage(ACCT, INDEX)) ~> #memory [ _ = sload INDEX ] ... </k>
         <id> ACCT </id>
         <account>
           <acctID> ACCT </acctID>
           <storage> STORAGE </storage>
           ...
         </account>
      requires notBool INDEX in_keys(STORAGE)





    rule #memory [ sstore _ , _   ] => .














    rule #memory [ br _           ] => .
    rule #memory [ br _ , _       ] => .
    rule #memory [ revert _       ] => .
    rule #memory [ log _          ] => .
    rule #memory [ log _ , _      ] => .
    rule #memory [ selfdestruct _ ] => .

    rule <k> #memory [ ret _          ] => . ... </k> <localCalls> .List </localCalls>

    rule #memory [ _ = call _ at _ ( _ ) send _ , gaslimit _ ] => .
    rule #memory [ _ = staticcall _ at _ ( _ ) gaslimit _ ] => .
    rule #memory [ _ , _ = create _ ( _ ) send _ ] => .
    rule #memory [ _ , _ = copycreate _ ( _ ) send _ ] => .

    rule #memory [ ECREC ] => .
    rule #memory [ SHA256 ] => .
    rule #memory [ RIP160 ] => .
    rule #memory [ ID ] => .
    rule #memory [ ECADD ] => .
    rule #memory [ ECMUL ] => .
    rule #memory [ ECPAIRING ] => .












    syntax InternalOp ::= #registerDelta ( LValue , Int )
 // -----------------------------------------------------
    rule <k> #registerDelta(% REG, NEWSIZE) => #deductMemory(PEAK) ... </k>
         <currentMemory> CURR </currentMemory>
         <regs> REGS </regs>
         <peakMemory> PEAK => maxInt(PEAK, CURR +Int NEWSIZE -Int intSize(getInt(REGS [ REG ]))) </peakMemory>






    syntax InternalOp ::= #registerDeltas ( LValues , Ints )
 // --------------------------------------------------------
    rule #registerDeltas(REG, REGS, INT, INTS) => #registerDelta(REG, intSize(INT)) ~> #registerDeltas(REGS, INTS)
    rule #registerDeltas(.LValues, _) => .K
    rule #registerDeltas(_, .Ints) => .K











    syntax InternalOp ::= #memoryExpand  ( Int , Int )
                        | #memoryDelta   ( Int , Int )
                        | #memoryDelta   ( Int ) [klabel(memoryDirectDelta)]
 // ------------------------------------------------------------------------
    rule <k> #memoryExpand(INDEX, NEWSIZE) => #deductMemory(PEAK) ... </k>
         <localMem>... INDEX |-> LM ...</localMem>
         <currentMemory> CURR => CURR +Int maxInt(0, NEWSIZE -Int bytesInWords((lengthBytes(LM)))) </currentMemory>
         <peakMemory> PEAK => maxInt(PEAK, CURR +Int maxInt(0, NEWSIZE -Int bytesInWords((lengthBytes(LM))))) </peakMemory>

    rule <k> #memoryExpand(INDEX, NEWSIZE) => #deductMemory(PEAK) ... </k>
         <localMem> LM </localMem>
         <currentMemory> CURR => CURR +Int NEWSIZE </currentMemory>
         <peakMemory> PEAK => maxInt(PEAK, CURR +Int NEWSIZE) </peakMemory>
      requires notBool INDEX in_keys(LM)

    rule <k> #memoryDelta(INDEX, NEWSIZE) => #deductMemory(PEAK) ... </k>
         <localMem>... INDEX |-> LM ...</localMem>
         <currentMemory> CURR => CURR +Int NEWSIZE -Int bytesInWords((lengthBytes(LM))) </currentMemory>
         <peakMemory> PEAK => maxInt(PEAK, CURR +Int NEWSIZE -Int bytesInWords((lengthBytes(LM)))) </peakMemory>

    rule <k> #memoryDelta(INDEX, NEWSIZE) => #deductMemory(PEAK) ... </k>
         <localMem> LM </localMem>
         <currentMemory> CURR => CURR +Int NEWSIZE </currentMemory>
         <peakMemory> PEAK => maxInt(PEAK, CURR +Int NEWSIZE) </peakMemory>
      requires notBool INDEX in_keys(LM)

    rule <k> #memoryDelta(DELTA) => #deductMemory(PEAK) ... </k>
         <currentMemory> CURR => CURR +Int DELTA </currentMemory>
         <peakMemory> PEAK => maxInt(PEAK, CURR +Int DELTA) </peakMemory>

    syntax InternalOp ::= #deductMemory ( Int )
 // -------------------------------------------
    rule <k> #deductMemory(OLDPEAK) => Cmem(SCHED, NEWPEAK) -Int Cmem(SCHED, OLDPEAK) ~> #deductGas ... </k>
         <schedule> SCHED </schedule>
         <peakMemory> NEWPEAK </peakMemory>








    syntax Int ::= Cmem ( Schedule , Int )       [function, memo]
                 | Cpricedmem ( Schedule, Int )  [function]
 // -------------------------------------------------------
    rule Cmem(SCHED, N)  => (Cpricedmem(SCHED, N) *Int Gmemory < SCHED >) +Int ((Cpricedmem(SCHED, N) *Int Cpricedmem(SCHED, N)) /Int Gquadcoeff < SCHED >)
    rule Cpricedmem(SCHED, N) => maxInt(0, N -Int Smemallowance < SCHED > )












    syntax InternalOp ::= "#compute" "[" Instruction "," Schedule "]"
 // -----------------------------------------------------------------









    rule #compute [ _ = not   W,       SCHED ] => Gnot < SCHED > +Int intSize(W) *Int Gnotword < SCHED >
    rule #compute [ _ = and   W0 , W1, SCHED ] => Gbitwise < SCHED > +Int minInt(intSize(W0), intSize(W1)) *Int Gbitwiseword < SCHED >
    rule #compute [ _ = or    W0 , W1, SCHED ] => Gbitwise < SCHED > +Int maxInt(intSize(W0), intSize(W1)) *Int Gbitwiseword < SCHED >
    rule #compute [ _ = xor   W0 , W1, SCHED ] => Gnot < SCHED > +Int maxInt(intSize(W0), intSize(W1)) *Int Gnotword < SCHED >
    rule #compute [ _ = shift W0 , W1, SCHED ] => Gbitwise < SCHED > +Int maxInt(1, intSize(W0) +Int bitsInWords(W1, SCHED)) *Int Gbitwiseword < SCHED >
    rule #compute [ _ = log2 W,        SCHED ] => Glogarithm < SCHED > +Int intSize(W) *Int Glogarithmword < SCHED >







    rule #compute [ _ = iszero W,      SCHED ] => Giszero < SCHED >
    rule #compute [ _ = cmp _ W0 , W1, SCHED ] => Gcmp < SCHED > +Int minInt(intSize(W0), intSize(W1)) *Int Gcmpword < SCHED >









    rule #compute [ _ = add W0 , W1, SCHED ] => Gadd < SCHED > +Int maxInt(intSize(W0), intSize(W1)) *Int Gaddword < SCHED >
    rule #compute [ _ = sub W0 , W1, SCHED ] => Gadd < SCHED > +Int maxInt(intSize(W0), intSize(W1)) *Int Gaddword < SCHED >
    rule #compute [ _ = mul W0 , W1, SCHED ] => Cmul(SCHED, intSize(W0), intSize(W1))
    rule #compute [ _ = div W0 , W1, SCHED ] => Cdiv(SCHED, intSize(W0), intSize(W1))
    rule #compute [ _ = mod W0 , W1, SCHED ] => Cdiv(SCHED, intSize(W0), intSize(W1))
    rule #compute [ _ = exp W0 , W1, SCHED ] => Cexp(SCHED, intSize(W0), W0, W1)










    rule #compute [ _ = addmod W0 , W1 , W2, SCHED ] => Gadd < SCHED > +Int maxInt(intSize(W0), intSize(W1)) *Int Gaddword < SCHED > +Int Cdiv(SCHED, maxInt(intSize(W0), intSize(W1)) +Int 1, intSize(W2))
    rule #compute [ _ = mulmod W0 , W1 , W2, SCHED ] => Cmul(SCHED, intSize(W0), intSize(W1)) +Int Cdiv(SCHED, intSize(W0) +Int intSize(W1), intSize(W2))
    rule #compute [ _ = expmod W0 , W1 , W2, SCHED ] => Cexpmod(SCHED, intSize(W0), intSize(W1), intSize(W2), W2)







    rule <k> #compute [ _ = sha3 W0, SCHED ] => Gsha3 < SCHED > +Int bytesInWords(lengthBytes(LM)) *Int Gsha3word < SCHED > ... </k>
         <localMem>... W0 |-> LM ...</localMem>

    rule <k> #compute [ _ = sha3 W0, SCHED ] => Gsha3 < SCHED > ... </k>
         <localMem> LM </localMem>
      requires notBool W0 in_keys(LM)








    rule #compute [ _ = byte _ , _, SCHED ] => Gbyte < SCHED >
    rule #compute [ _ = twos  WIDTH, _, SCHED ] => Gtwos  < SCHED > +Int maxInt(1, bytesInWords(chop(WIDTH))) *Int Gtwosword  < SCHED >
    rule #compute [ _ = sext  WIDTH, _, SCHED ] => Gsext  < SCHED > +Int maxInt(1, bytesInWords(chop(WIDTH))) *Int Gsextword  < SCHED >
    rule #compute [ _ = bswap WIDTH, _, SCHED ] => Gbswap < SCHED > +Int maxInt(1, bytesInWords(chop(WIDTH))) *Int Gbswapword < SCHED >







    rule #compute [ _ = call @iele.gas         ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.gasprice    ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.gaslimit    ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.number      ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.msize       ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.codesize    ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.beneficiary ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.address     ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.origin      ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.caller      ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.timestamp   ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.difficulty  ( _ ), SCHED ] => Greadstate < SCHED >
    rule #compute [ _ = call @iele.callvalue   ( _ ), SCHED ] => Greadstate < SCHED >





    rule #compute [ _ = call @iele.blockhash   ( _ ), SCHED ] => Gblockhash < SCHED >







    rule #compute [ _ = call @iele.balance     ( _ ), SCHED ] => Gbalance     < SCHED >
    rule #compute [ _ = call @iele.extcodesize ( _ ), SCHED ] => Gextcodesize < SCHED >
    rule #compute [ _ = calladdress _ at _,           SCHED ] => Gcalladdress < SCHED >







    rule <k> #compute [ DEST = % SRC:Int, SCHED ] => Gmove < SCHED > *Int intSize(getInt(REGS [ SRC ])) ... </k>
         <regs> REGS </regs>
      requires notBool Gnewmove << SCHED >>

    rule <k> #compute [ DEST = % SRC:Int, SCHED ] => Gmove < SCHED > ... </k>
         <regs> REGS </regs>
      requires Gnewmove << SCHED >>

    rule <k> #compute [ DEST = SRC:Int, SCHED ] => Gmove < SCHED > *Int intSize(SRC) ... </k>
      requires notBool Gnewmove << SCHED >>

    rule <k> #compute [ DEST = SRC:Int, SCHED ] => Gmove < SCHED > ... </k>
      requires Gnewmove << SCHED >>








    rule #compute [ br _, SCHED ] => Gbr < SCHED >
    rule #compute [ br _ , _, SCHED ] => Gbrcond < SCHED >








    rule <k> #compute [ _ = call @ NAME ( ARGS ), SCHED ] => Gcallreg < SCHED > *Int REGISTERS +Int intSizes(ARGS) *Int Gmove < SCHED > +Int Glocalcall < SCHED > ... </k>
         <funcId> NAME </funcId>
         <nregs> REGISTERS </nregs>
      requires notBool Gnewmove << SCHED >>
    rule <k> #compute [ _ = call @ NAME ( ARGS ), SCHED ] => Gcallreg < SCHED > *Int REGISTERS +Int Glocalcall < SCHED > ... </k>
         <funcId> NAME </funcId>
         <nregs> REGISTERS </nregs>
      requires Gnewmove << SCHED >>

    rule <k> #compute [ _ = call (IDX:Int => @ FUNC) ( _ ), _ ] ... </k>
         <funcLabels> ... IDX |-> FUNC ... </funcLabels>
    // this will throw an exception, so the gas cost doesn't really matter
    rule <k> #compute [ _ = call IDX:Int ( _ ), _ ] => 0 ... </k>
         <funcLabels> LABELS </funcLabels>
      requires notBool IDX in_keys(LABELS)





    rule <k> #compute [ ret ARGS::NonEmptyInts, SCHED ] => Gmove < SCHED > *Int #sizeRegs(ARGS) +Int 8 ... </k>
         <localCalls> ListItem(_) ... </localCalls>
      requires notBool Gnewmove << SCHED >>
    rule <k> #compute [ ret _::NonEmptyInts, SCHED ] => 0 ... </k>
         <localCalls> .List </localCalls>
      requires notBool Gnewmove << SCHED >>

    rule <k> #compute [ ret _::NonEmptyInts, SCHED ] => 0 ... </k>
      requires Gnewmove << SCHED >>
    rule #compute [ revert _, SCHED ] => 0











    rule <k> #compute [ _, RETS::LValues = call _ at ACCTTO ( ARGS ) send VALUE , gaslimit GCAP, SCHED ] => Ccall(SCHED, #accountEmpty(ACCTTO), GCAP *Int Sgasdivisor < SCHED >, GAVAIL, VALUE, #sizeLVals(RETS), Ccallarg(SCHED, ARGS)) ... </k>
         <gas> GAVAIL </gas>

    rule <k> #compute [ _, RETS::LValues = staticcall _ at ACCTTO ( ARGS ) gaslimit GCAP, SCHED ] => Ccall(SCHED, #accountEmpty(ACCTTO), GCAP *Int Sgasdivisor < SCHED >, GAVAIL, 0, #sizeLVals(RETS), Ccallarg(SCHED, ARGS)) ... </k>
         <gas> GAVAIL </gas>







    rule <k> #compute [ log IDX, SCHED ]                                 => (Glog < SCHED > +Int (Glogdata < SCHED > *Int bytesInWords(lengthBytes(LM))) +Int (0 *Int Glogtopic < SCHED >)) ... </k> <localMem>... IDX |-> LM ...</localMem>
    rule <k> #compute [ log IDX , _:Int, SCHED ]                         => (Glog < SCHED > +Int (Glogdata < SCHED > *Int bytesInWords(lengthBytes(LM))) +Int (1 *Int Glogtopic < SCHED >)) ... </k> <localMem>... IDX |-> LM ...</localMem>
    rule <k> #compute [ log IDX , _:Int , _:Int, SCHED ]                 => (Glog < SCHED > +Int (Glogdata < SCHED > *Int bytesInWords(lengthBytes(LM))) +Int (2 *Int Glogtopic < SCHED >)) ... </k> <localMem>... IDX |-> LM ...</localMem>
    rule <k> #compute [ log IDX , _:Int , _:Int , _:Int, SCHED ]         => (Glog < SCHED > +Int (Glogdata < SCHED > *Int bytesInWords(lengthBytes(LM))) +Int (3 *Int Glogtopic < SCHED >)) ... </k> <localMem>... IDX |-> LM ...</localMem>
    rule <k> #compute [ log IDX , _:Int , _:Int , _:Int,  _:Int, SCHED ] => (Glog < SCHED > +Int (Glogdata < SCHED > *Int bytesInWords(lengthBytes(LM))) +Int (4 *Int Glogtopic < SCHED >)) ... </k> <localMem>... IDX |-> LM ...</localMem>

    rule <k> #compute [ log IDX, SCHED ] ... </k>
         <localMem> LM (.Map => IDX |-> .Bytes) </localMem>
      requires notBool IDX in_keys(LM)
    rule <k> #compute [ log IDX, _, SCHED ] ... </k>
         <localMem> LM (.Map => IDX |-> .Bytes) </localMem>
      requires notBool IDX in_keys(LM)








    rule <k> #compute [ _ = load INDEX, SCHED ] => Gloadcell < SCHED > +Int bytesInWords(lengthBytes(LM)) *Int Gloadword < SCHED > ... </k>
         <localMem>... INDEX |-> LM ...</localMem>
    rule <k> #compute [ _ = load INDEX, SCHED ] => Gloadcell < SCHED > ... </k>
         <localMem> LM </localMem>
      requires notBool INDEX in_keys(LM)

    rule #compute [ _ = load INDEX , OFFSET , WIDTH, SCHED ] => Gload < SCHED > +Int bytesInWords(WIDTH) *Int Gloadword < SCHED >

    rule #compute [ store VALUE , INDEX, SCHED ] => Gstorecell < SCHED > +Int intSize(VALUE) *Int Gstoreword < SCHED >
    rule #compute [ store VALUE , INDEX , OFFSET , WIDTH, SCHED ] => Gstore < SCHED > +Int bytesInWords(WIDTH) *Int Gstoreword < SCHED >







        rule <k> #compute [ _ = sload INDEX, SCHED ] => Gsload < SCHED > +Int Gsloadkey < SCHED > *Int intSize(INDEX) +Int Gsloadword < SCHED > *Int intSize(VALUE) ... </k>
         <id> ACCT </id>
         <account>
           <acctID> ACCT </acctID>
           <storage> ... INDEX |-> VALUE ... </storage>
           ...
         </account>





    rule <k> #compute [ sstore VALUE , INDEX, SCHED ] => Csstore(SCHED, INDEX, VALUE, OLDVALUE) ... </k>
         <id> ACCT </id>
         <account>
           <acctID> ACCT </acctID>
           <storage> ... INDEX |-> OLDVALUE ... </storage>
           ...
         </account>

    rule <k> (.K => #lookupStorage(ACCT, INDEX)) ~> #compute [ sstore VALUE , INDEX, _ ] ... </k>
         <id> ACCT </id>
         <account>
           <acctID> ACCT </acctID>
           <storage> STORAGE </storage>
           ...
         </account>
      requires notBool INDEX in_keys(STORAGE)








    rule #compute [ _ , _ = create _ ( ARGS ) send _, SCHED ] => Gcreate < SCHED > +Int Gmove < SCHED > *Int Ccallarg(SCHED, ARGS)
    rule #compute [ _ , _ = copycreate _ ( ARGS ) send _, SCHED ] => Gcopycreate < SCHED > +Int Gmove < SCHED > *Int Ccallarg(SCHED, ARGS)





    rule <k> #compute [ selfdestruct ACCTTO, SCHED ] => Cselfdestruct(SCHED, #accountEmpty(ACCTTO), BAL) ... </k>
         <id> ACCTFROM </id>
         <account>
           <acctID> ACCTFROM </acctID>
           <balance> BAL </balance>
           ...
         </account>







    rule <k> #compute [ ECREC, SCHED ]  => Gecrec < SCHED > ... </k>
    rule <k> #compute [ SHA256, SCHED ] => Gsha256 < SCHED > +Int Gsha256word < SCHED > *Int bytesInWords(maxInt(LEN, intSize(DATA))) ... </k> <callData> LEN , DATA , .Ints </callData>
    rule <k> #compute [ RIP160, SCHED ] => Grip160 < SCHED > +Int Grip160word < SCHED > *Int bytesInWords(maxInt(LEN, intSize(DATA))) ... </k> <callData> LEN , DATA , .Ints </callData>
    rule <k> #compute [ ID, SCHED ]     => 0 ... </k>

    rule #compute [ ECADD, SCHED ] => Gecadd < SCHED >
    rule #compute [ ECMUL, SCHED ] => Gecmul < SCHED >
    rule <k> #compute [ ECPAIRING, SCHED ] => Gecpairing < SCHED > +Int LEN *Int Gecpairingpair < SCHED > ... </k> <callData> LEN , _ </callData>







    syntax Int ::= Csstore ( Schedule , Int , Int , Int ) [function]
 // ----------------------------------------------------------
    rule Csstore(SCHED, INDEX, VALUE, OLDVALUE) => Gsstore < SCHED > +Int Gsstorekey < SCHED > *Int intSize(INDEX) +Int Gsstoreword < SCHED > *Int intSize(VALUE) +Int #if VALUE =/=Int 0 andBool OLDVALUE ==Int 0 #then Gsstoresetkey < SCHED > *Int intSize(INDEX) +Int Gsstoreset < SCHED > *Int intSize(VALUE) #else maxInt(0, Gsstoreset < SCHED > *Int (intSize(VALUE) -Int intSize(OLDVALUE))) #fi

    syntax Operand ::= Ccall    ( Schedule , BExp , Int , Int , Int , Int , Int ) [strict(2)]
                     | Ccallgas ( Schedule , BExp , Int , Int , Int , Int , Int ) [strict(2)]
    syntax Int ::= Cgascap  ( Schedule , Int , Int , Int )                        [function]
                 | Cextra   ( Schedule , Bool , Int , Int , Int )                 [function]
                 | Cxfer    ( Schedule , Int )                                    [function]
                 | Cnew     ( Schedule , Bool , Int )                             [function]
                 | Ccallarg ( Schedule , Operands )                               [function]
 // ----------------------------------------------------------------------------------------
    rule Ccall(SCHED, ISEMPTY:Bool, GCAP, GAVAIL, VALUE, RETS, ARGS) => Cextra(SCHED, ISEMPTY, VALUE, RETS, ARGS) +Int Cgascap(SCHED, GCAP, GAVAIL, Cextra(SCHED, ISEMPTY, VALUE, RETS, ARGS))

    rule Ccallgas(SCHED, ISEMPTY:Bool, GCAP, GAVAIL, 0, RETS, ARGS)     => Cgascap(SCHED, GCAP, GAVAIL, Cextra(SCHED, ISEMPTY,     0, RETS, ARGS))
    rule Ccallgas(SCHED, ISEMPTY:Bool, GCAP, GAVAIL, VALUE, RETS, ARGS) => Cgascap(SCHED, GCAP, GAVAIL, Cextra(SCHED, ISEMPTY, VALUE, RETS, ARGS)) +Int Gcallstipend < SCHED > requires VALUE =/=K 0

    rule Cgascap(SCHED, GCAP, GAVAIL, GEXTRA) => minInt(#allBut64th(GAVAIL -Int GEXTRA), GCAP) requires GAVAIL >=Int GEXTRA andBool notBool Gstaticcalldepth << SCHED >>
    rule Cgascap(SCHED, GCAP, GAVAIL, GEXTRA) => GCAP                                          requires GAVAIL <Int  GEXTRA orBool Gstaticcalldepth << SCHED >>

    rule Cextra(SCHED, ISEMPTY, VALUE, RETS, ARGS) => Gcall < SCHED > +Int Cnew(SCHED, ISEMPTY, VALUE) +Int Cxfer(SCHED, VALUE) +Int Gcallreg < SCHED > *Int (RETS +Int ARGS)

    rule Cxfer(SCHED, 0) => 0
    rule Cxfer(SCHED, N) => Gcallvalue < SCHED > requires N =/=K 0

    rule Cnew(SCHED, ISEMPTY:Bool, VALUE) => Gnewaccount < SCHED >
      requires         ISEMPTY andBool VALUE =/=Int 0
    rule Cnew(SCHED, ISEMPTY:Bool, VALUE) => 0
      requires notBool ISEMPTY orBool  VALUE  ==Int 0

    rule Ccallarg(SCHED, ARGS) => intSizes(ARGS)
      requires notBool Gnewmove << SCHED >>
    rule Ccallarg(SCHED, ARGS) => #sizeRegs(ARGS)
      requires Gnewmove << SCHED >>

    syntax Operand ::= Cselfdestruct ( Schedule , BExp , Int ) [strict(2)]
 // ----------------------------------------------------------------------
    rule Cselfdestruct(SCHED, ISEMPTY:Bool, BAL) => Gselfdestruct < SCHED > +Int Gnewaccount < SCHED >
      requires ISEMPTY andBool (        Gselfdestructnewaccount << SCHED >>) andBool BAL =/=Int 0
    rule Cselfdestruct(SCHED, ISEMPTY:Bool, BAL) => Gselfdestruct < SCHED >
      requires ISEMPTY andBool (notBool Gselfdestructnewaccount << SCHED >>  orBool  BAL  ==Int 0)
    rule Cselfdestruct(SCHED, ISEMPTY:Bool, BAL) => Gselfdestruct < SCHED >
      requires notBool ISEMPTY

    syntax KResult ::= Bool
    syntax BExp ::= Bool
                  | #accountEmpty(Int)
 // ------------------------------------------------------------------------------------------
    rule <k> #accountEmpty(ACCT) => notBool EXISTS ... </k>
         <account>
           <acctID> ACCT </acctID>
           <exists> EXISTS </exists>
           ...
         </account>
    rule <k> (.K => #loadAccount ACCT) ~> #accountEmpty(ACCT) ... </k>
         <activeAccounts> ACCTS </activeAccounts>
      requires notBool ACCT in ACCTS

    syntax Int ::= #allBut64th ( Int ) [function]
 // ---------------------------------------------
    rule #allBut64th(N) => N -Int (N /Int 64)

    syntax Int ::= G0 ( Schedule , WordStack , Ints ) [function, klabel(G0create)]
                 | G0 ( Schedule , String , Ints )    [function, klabel(G0call)]
                 | G0 ( Schedule , WordStack , Bool ) [function, klabel(G0aux)]
 // ---------------------------------------------------------------------------
    rule G0(SCHED, .WordStack, true)  => Gtxcreate    < SCHED >
    rule G0(SCHED, .WordStack, false) => Gtransaction < SCHED >

    rule G0(SCHED, WS, ARGS) => G0(SCHED, #parseByteStackRaw(#rlpEncodeLength(#rlpEncodeString(#unparseByteStack(WS)) +String #rlpEncodeInts(ARGS), 192)), true)
    rule G0(SCHED, FUNC, ARGS) => G0(SCHED, #parseByteStackRaw(#rlpEncodeLength(#rlpEncodeString(FUNC) +String #rlpEncodeInts(ARGS), 192)), false)

    rule G0(SCHED, 0 : REST, ISCREATE::Bool) => Gtxdatazero    < SCHED > +Int G0(SCHED, REST, ISCREATE)
    rule G0(SCHED, N : REST, ISCREATE::Bool) => Gtxdatanonzero < SCHED > +Int G0(SCHED, REST, ISCREATE) requires N =/=Int 0

    syntax Int ::= "G*" "(" Int "," Int "," Int ")" [function]
 // ----------------------------------------------------------
    rule G*(GAVAIL, GLIMIT, REFUND) => GAVAIL +Int minInt((GLIMIT -Int GAVAIL)/Int 2, REFUND)

    syntax Int ::= Cmul     ( Schedule , Int , Int )             [function]
                 | Ckara    ( Int , Int )                        [function]
                 | Cdiv     ( Schedule , Int , Int )             [function]
                 | Cexp     ( Schedule , Int , Int , Int )       [function]
                 | Cexpmod  ( Schedule , Int , Int , Int , Int ) [function]
 // -----------------------------------------------------------------------
    rule Cmul(SCHED, L1, L2) => Cmul(SCHED, L2, L1)
      requires L2 >Int L1

    rule Cmul(SCHED, L1, L2) =>
        Gmulkara < SCHED > *Int Ckara(L1, L2) +Int
        Gmulword < SCHED > *Int (L1 +Int L2) +Int
        Gmul < SCHED >
      [owise]
      // Note that if L2 is low enough (< 32) then #overApproxKara(L2) = L2 * L2

    rule Ckara(L1, L2) => L1 *Int #overApproxKara(L2) /Int L2
      requires L1 >=Int L2

    rule Ckara(L1, L2) => L2 *Int #overApproxKara(L1) /Int L1 [owise]

    rule Cdiv(SCHED, L1, L2) =>
        Gdivkara < SCHED > *Int Ckara(L1 -Int L2 +Int 1, L2) +Int
        Gdivword < SCHED > *Int L1 +Int
        Gdiv < SCHED >
      requires L1 >=Int L2
    
    rule Cdiv(SCHED, L1, L2) =>
        Gdivword < SCHED > *Int L1 +Int
        Gdiv < SCHED >
      requires notBool Gnewarith << SCHED >>
      [owise]

    rule Cdiv(SCHED, L1, L2) =>
        Gdiv < SCHED >
      requires Gnewarith << SCHED >>
      [owise]

    rule Cexp(SCHED, L1, W1, W2) =>
        Gexpkara < SCHED > *Int #overApproxKara(#adjustedBitLength(L1, W1) *Int W2 /Int 64) +Int
        Gexpword < SCHED > *Int L1 +Int
        Gexp < SCHED >
      requires notBool Gnewarith << SCHED >>

    rule Cexp(SCHED, L1, W1, W2) =>
        Gexpkara < SCHED > *Int #overApproxKara(#adjustedBitLength(L1, W1) *Int W2 /Int 64) +Int
        Gexpword < SCHED > *Int #adjustedBitLength(L1, W1) +Int
        Gexp < SCHED >
      requires Gnewarith << SCHED >>

    rule Cexpmod(SCHED, LB, LEX, LM, EX) =>
        ((Gexpmodkara < SCHED > *Int #overApproxKara(LM) *Int #adjustedBitLength(LEX, EX)) up/Int 10) +Int
        Gexpmodmod < SCHED > *Int LM +Int
        Gexpmodexp < SCHED > *Int #adjustedBitLength(LEX, EX) +Int
        Gexpmod < SCHED >
      requires LB <=Int LM

    rule Cexpmod(SCHED, LB, LEX, LM, EX) =>
        ((Gexpmodkara < SCHED > *Int #overApproxKara(LM) *Int #adjustedBitLength(LEX, EX)) up/Int 10) +Int
        Gexpmodmod < SCHED > *Int LM +Int
        Gexpmodexp < SCHED > *Int #adjustedBitLength(LEX, EX) +Int
        Cdiv(SCHED, LB, LM) +Int
        Gexpmod < SCHED >
      [owise]

















    syntax Int ::= #overApproxKara ( Int )                 [function]
 // -----------------------------------------------------------------
    rule #overApproxKara(N) => #if N <=Int   32 #then N *Int N
                         #else #if N <=Int 1024 #then N *Int N /Int  4 +Int  48 *Int N -Int    768
                         #else                        N *Int N /Int 16 +Int 432 *Int N -Int 197376
                         #fi #fi














    syntax Int ::= #adjustedBitLength(Int, Int) [function]
                 | #adjustedBitLength(Int) [function, klabel(#adjustedBitLengthAux)]
 // --------------------------------------------------------------------------------
    rule #adjustedBitLength(LEX, EX) => maxInt(1, #if LEX <=Int 1 #then 0 #else 64 *Int (LEX -Int 1) #fi +Int #adjustedBitLength(twos(8, EX)))

    rule #adjustedBitLength(0) => 0
    rule #adjustedBitLength(N) => log2Int(N) [owise]















    syntax Bool ::= ScheduleFlag "<<" Schedule ">>" [function]
 // ----------------------------------------------------------

    syntax ScheduleFlag ::= "Gselfdestructnewaccount" | "Gstaticcalldepth"
                          | "Gnewmove"                | "Gnewarith"
 // ---------------------------------------------------------------





    syntax Int ::= ScheduleConst "<" Schedule ">" [function]
 // --------------------------------------------------------
 
    syntax ScheduleConst ::= "Gmove"       | "Greadstate"  | "Gadd"        | "Gaddword"       | "Gmul"        | "Gmulword"     | "Gmulkara"
                           | "Gdiv"        | "Gdivword"    | "Gdivkara"    | "Gexpkara"       | "Gexpword"    | "Gexp"         | "Gexpmodkara"    | "Gexpmodmod"
                           | "Gexpmodexp"  | "Gexpmod"     | "Gnot"        | "Gnotword"       | "Gbitwise"    | "Gbitwiseword" | "Glogarithm"     | "Glogarithmword"
                           | "Gbyte"       | "Gtwos"       | "Gtwosword"   | "Gsext"          | "Gsextword"   | "Gbswap"       | "Gbswapword"     | "Giszero"
                           | "Gcmp"        | "Gcmpword"    | "Gbr"         | "Gbrcond"        | "Gblockhash"  | "Gsha3"        | "Gsha3word"      | "Gloadcell"
                           | "Gload"       | "Gloadword"   | "Gstorecell"  | "Gstore"         | "Gstoreword"  | "Gbalance"     | "Gextcodesize"   | "Gcalladdress"
                           | "Glog"        | "Glogdata"    | "Glogtopic"   | "Gsstore"        | "Gsstoreword" | "Gsstorekey"   | "Gsstoreset"     | "Gsstoresetkey"
                           | "Gsload"      | "Gsloadkey"   | "Gsloadword"  | "Gselfdestruct"  | "Gcallmemory" | "Gcallreg"     | "Glocalcall"     | "Gcallstipend" 
                           | "Gcall"       | "Gcallvalue"  | "Gnewaccount" | "Gcreate"        | "Gcopycreate" | "Gcodedeposit" | "Gecrec"         | "Gsha256word"
                           | "Gsha256"     | "Grip160word" | "Grip160"     | "Gecadd"         | "Gecmul"      | "Gecpairing"   | "Gecpairingpair" | "Gtransaction"
                           | "Gtxcreate"   | "Gmemory"     | "Gquadcoeff"  | "Gtxdatanonzero" | "Gtxdatazero" | "Rsstoreset"   | "Rselfdestruct"  | "Rb"
                           | "Sgasdivisor" | "Smemallowance"
 // ---------------------------------------------------------------------------------------------------------------------------------------------------------------







    syntax Schedule ::= "DEFAULT"
 // -----------------------------
    rule Gmove          < DEFAULT > => 3
    rule Greadstate     < DEFAULT > => 2
    rule Gadd           < DEFAULT > => 0
    rule Gaddword       < DEFAULT > => 3
    rule Gmul           < DEFAULT > => 0
    rule Gmulword       < DEFAULT > => 2
    rule Gmulkara       < DEFAULT > => 3
    rule Gdiv           < DEFAULT > => 0
    rule Gdivword       < DEFAULT > => 2
    rule Gdivkara       < DEFAULT > => 3
    rule Gexpkara       < DEFAULT > => 50
    rule Gexpword       < DEFAULT > => 10
    rule Gexp           < DEFAULT > => 0
    rule Gexpmodkara    < DEFAULT > => 40
    rule Gexpmodmod     < DEFAULT > => 0
    rule Gexpmodexp     < DEFAULT > => 0
    rule Gexpmod        < DEFAULT > => 0
    rule Gnot           < DEFAULT > => 0
    rule Gnotword       < DEFAULT > => 3
    rule Gbitwise       < DEFAULT > => 0
    rule Gbitwiseword   < DEFAULT > => 3
    rule Glogarithm     < DEFAULT > => 0
    rule Glogarithmword < DEFAULT > => 3
    rule Gbyte          < DEFAULT > => 3
    rule Gtwos          < DEFAULT > => 0
    rule Gtwosword      < DEFAULT > => 5
    rule Gsext          < DEFAULT > => 0
    rule Gsextword      < DEFAULT > => 5
    rule Gbswap         < DEFAULT > => 0
    rule Gbswapword     < DEFAULT > => 5
    rule Giszero        < DEFAULT > => 3
    rule Gcmp           < DEFAULT > => 0
    rule Gcmpword       < DEFAULT > => 3
    rule Gbr            < DEFAULT > => 8
    rule Gbrcond        < DEFAULT > => 10
    rule Gblockhash     < DEFAULT > => 20
    rule Gsha3          < DEFAULT > => 30
    rule Gsha3word      < DEFAULT > => 6
    rule Gloadcell      < DEFAULT > => 3
    rule Gload          < DEFAULT > => 0
    rule Gloadword      < DEFAULT > => 3
    rule Gstorecell     < DEFAULT > => 3
    rule Gstore         < DEFAULT > => 0
    rule Gstoreword     < DEFAULT > => 3
    rule Gbalance       < DEFAULT > => 400
    rule Gextcodesize   < DEFAULT > => 700
    rule Gcalladdress   < DEFAULT > => 700
    rule Glog           < DEFAULT > => 375
    rule Glogdata       < DEFAULT > => 8
    rule Glogtopic      < DEFAULT > => 375
    rule Gsstore        < DEFAULT > => 1000
    rule Gsstoreword    < DEFAULT > => 500
    rule Gsstorekey     < DEFAULT > => 500
    rule Gsstoreset     < DEFAULT > => 1875
    rule Rsstoreset     < DEFAULT > => 1875
    rule Gsstoresetkey  < DEFAULT > => 1875
    rule Gsload         < DEFAULT > => 50
    rule Gsloadkey      < DEFAULT > => 100
    rule Gsloadword     < DEFAULT > => 50
    rule Gselfdestruct  < DEFAULT > => 0
    rule Gcallmemory    < DEFAULT > => 2
    rule Gcallreg       < DEFAULT > => 3
    rule Glocalcall     < DEFAULT > => 11
    rule Gcall          < DEFAULT > => 40
    rule Gcallstipend   < DEFAULT > => 2300
    rule Gcallvalue     < DEFAULT > => 9000
    rule Gnewaccount    < DEFAULT > => 25000
    rule Gcreate        < DEFAULT > => 32000
    rule Gcopycreate    < DEFAULT > => 33000
    rule Gcodedeposit   < DEFAULT > => 200
    rule Gecrec         < DEFAULT > => 3000
    rule Gsha256        < DEFAULT > => 60
    rule Gsha256word    < DEFAULT > => 3
    rule Grip160        < DEFAULT > => 600
    rule Grip160word    < DEFAULT > => 30
    rule Gecadd         < DEFAULT > => 500
    rule Gecmul         < DEFAULT > => 40000
    rule Gecpairing     < DEFAULT > => 100000
    rule Gecpairingpair < DEFAULT > => 80000
    rule Gmemory        < DEFAULT > => 1
    rule Gquadcoeff     < DEFAULT > => 8192
    rule Gtransaction   < DEFAULT > => 21000
    rule Gtxcreate      < DEFAULT > => 53000
    rule Gtxdatazero    < DEFAULT > => 4
    rule Gtxdatanonzero < DEFAULT > => 68
    rule Rselfdestruct  < DEFAULT > => 24000
    rule Rb             < DEFAULT > => 3 *Int (10 ^Int 18)

    rule Gselfdestructnewaccount << DEFAULT >> => false
    rule Gstaticcalldepth        << DEFAULT >> => true
    rule Gnewmove                << DEFAULT >> => false
    rule Gnewarith               << DEFAULT >> => false

    rule Smemallowance  < DEFAULT > => 4096
    rule Sgasdivisor    < DEFAULT > => 1







 // Albe
 // --------------------------
    rule Gcall         < ALBE > => 700
    rule Gselfdestruct < ALBE > => 5000
    rule SCHEDCONST    < ALBE > => SCHEDCONST < DEFAULT > [owise]

    rule Gselfdestructnewaccount << ALBE >> => true
    rule Gstaticcalldepth        << ALBE >> => false
    rule SCHEDCONST              << ALBE >> => SCHEDCONST << DEFAULT >> [owise]







 // Danse
 // ---------------------------
    rule Gmove          < DANSE > => 2000
    rule Greadstate     < DANSE > => 2000
    rule Gadd           < DANSE > => 2800
    rule Gaddword       < DANSE > => 1
    rule Gmul           < DANSE > => 4900
    rule Gmulword       < DANSE > => 4
    rule Gmulkara       < DANSE > => 4
    rule Gdiv           < DANSE > => 4900
    rule Gdivword       < DANSE > => 5
    rule Gdivkara       < DANSE > => 8
    rule Gexpkara       < DANSE > => 2
    rule Gexpword       < DANSE > => 80
    rule Gexp           < DANSE > => 5300
    rule Gexpmodkara    < DANSE > => 15
    rule Gexpmodmod     < DANSE > => 180
    rule Gexpmodexp     < DANSE > => 8
    rule Gexpmod        < DANSE > => 6000
    rule Gnot           < DANSE > => 2700
    rule Gnotword       < DANSE > => 3
    rule Gbitwise       < DANSE > => 2900
    rule Gbitwiseword   < DANSE > => 1
    rule Glogarithm     < DANSE > => 2300
    rule Glogarithmword < DANSE > => 1
    rule Gbyte          < DANSE > => 2500
    rule Gtwos          < DANSE > => 3100
    rule Gtwosword      < DANSE > => 1
    rule Gsext          < DANSE > => 3300
    rule Gsextword      < DANSE > => 5
    rule Gbswap         < DANSE > => 3300
    rule Gbswapword     < DANSE > => 10
    rule Giszero        < DANSE > => 1800
    rule Gcmp           < DANSE > => 2500
    rule Gcmpword       < DANSE > => 1
    rule Gbr            < DANSE > => 5000
    rule Gbrcond        < DANSE > => 5000
    rule Gblockhash     < DANSE > => 20000
    rule Gsha3          < DANSE > => 8300
    rule Gsha3word      < DANSE > => 20
    rule Gloadcell      < DANSE > => 2900
    rule Gload          < DANSE > => 3300
    rule Gloadword      < DANSE > => 3
    rule Gstorecell     < DANSE > => 2800
    rule Gstore         < DANSE > => 3900
    rule Gstoreword     < DANSE > => 4
    rule Gbalance       < DANSE > => 400000
    rule Gextcodesize   < DANSE > => 700000
    rule Gcalladdress   < DANSE > => 700000
    rule Glog           < DANSE > => 375000
    rule Glogdata       < DANSE > => 8000
    rule Glogtopic      < DANSE > => 375000
    rule Gsstore        < DANSE > => 4950000
    rule Gsstoreword    < DANSE > => 300
    rule Gsstorekey     < DANSE > => 400
    rule Gsstoreset     < DANSE > => 1875000
    rule Gsstoresetkey  < DANSE > => 1875000
    rule Gsload         < DANSE > => 190000
    rule Gsloadkey      < DANSE > => 8000
    rule Gsloadword     < DANSE > => 2000
    rule Gselfdestruct  < DANSE > => 0
    rule Gcallmemory    < DANSE > => 2
    rule Gcallreg       < DANSE > => 1000
    rule Glocalcall     < DANSE > => 6800
    rule Gcall          < DANSE > => 40000
    rule Gcallstipend   < DANSE > => 2300000
    rule Gcallvalue     < DANSE > => 9000000
    rule Gnewaccount    < DANSE > => 25000000
    rule Gcreate        < DANSE > => 32000000
    rule Gcopycreate    < DANSE > => 33000000
    rule Gcodedeposit   < DANSE > => 200000
    rule Gecrec         < DANSE > => 3000000
    rule Gsha256        < DANSE > => 25000
    rule Gsha256word    < DANSE > => 30
    rule Grip160        < DANSE > => 25000
    rule Grip160word    < DANSE > => 30
    rule Gecadd         < DANSE > => 35000
    rule Gecmul         < DANSE > => 1700000
    rule Gecpairing     < DANSE > => 100000000
    rule Gecpairingpair < DANSE > => 26000000
    rule Gmemory        < DANSE > => 750
 
    rule Sgasdivisor < DANSE > => 1000
    rule SCHEDCONST  < DANSE > => SCHEDCONST < ALBE > [owise]

    rule Gnewmove      << DANSE >> => true
    rule Gnewarith     << DANSE >> => true
    rule SCHEDCONST    << DANSE >> => SCHEDCONST << ALBE >> [owise]
    
endmodule

