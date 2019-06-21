







requires "krypto.k"
requires "domains.k"

module IELE-DATA
    imports KRYPTO
    imports STRING-BUFFER
    imports ARRAY
    imports BYTES
    imports IELE-CONSTANTS
    imports IELE-COMMON
    imports DOMAINS
    imports COLLECTIONS

    syntax KResult ::= Int





    syntax Int ::= "pow30"  [function]
                 | "pow160" [function]
                 | "pow256" [function]
 // ----------------------------------
    rule pow30  => 2 ^Int 30
    rule pow160 => 2 ^Int 160
    rule pow256 => 2 ^Int 256






    syntax JSONList ::= List{JSON,","}
    syntax JSONKey  ::= String | Int
    syntax JSON     ::= String | Bool
                      | JSONKey ":" JSON
                      | "{" JSONList "}"
                      | "[" JSONList "]"
 // ------------------------------------











    syntax Int ::= chop ( Int ) [function]
 // --------------------------------------
    rule chop ( I:Int ) => bitRangeInt(I, 0, 256) requires I <Int 0  orBool I >=Int pow256
    rule chop ( I:Int ) => I               requires I >=Int 0 andBool I <Int pow256






    syntax Int ::= bool2Word ( Bool ) [function]
 // --------------------------------------------
    rule bool2Word(true)  => 1
    rule bool2Word(false) => 0

    syntax Bool ::= word2Bool ( Int ) [function]
 // --------------------------------------------
    rule word2Bool( 0 ) => false
    rule word2Bool( W ) => true  requires W =/=K 0









    syntax Account ::= ".Account" | Int








    syntax Int ::= #sizeRegs ( Operands ) [function]
                 | #sizeRegs ( Operands , Int ) [function, klabel(#sizeRegsAux)]
 // ----------------------------------------------------------------------------
    rule #sizeRegs(REGS) => #sizeRegs(REGS, 0)
    rule #sizeRegs(REG , REGS, N) => #sizeRegs(REGS, N +Int 1)
    rule #sizeRegs(.Operands, N) => N

    syntax Int ::= #sizeLVals ( LValues ) [function]
                 | #sizeLVals ( LValues , Int ) [function, klabel(#sizeLValuesAux)]
 // -------------------------------------------------------------------------------
    rule #sizeLVals(REGS) => #sizeLVals(REGS, 0)
    rule #sizeLVals(REG , REGS, N) => #sizeLVals(REGS, N +Int 1)
    rule #sizeLVals(.LValues, N) => N

    syntax String ::= IeleName2String ( IeleName ) [function]
                    | IeleNameToken2String ( IeleName ) [function, hook(STRING.token2string)]
    syntax IeleName ::= String2IeleName ( String ) [function, hook(STRING.string2token)]
 // ------------------------------------------------------------------------------------
    rule IeleName2String(I:Int) => Int2String(I)
    rule IeleName2String(N) => IeleNameToken2String(N) [owise]
    syntax String ::= StringIeleName2String ( StringIeleName ) [function, hook(STRING.token2string)]
 // ------------------------------------------------------------------------------------------------

    syntax Int ::= getInt(K) [function]
    syntax IeleName ::= getIeleName(K) [function]
 // ---------------------------------------------
    rule getInt(I:Int) => I
    rule getIeleName(X:IeleName) => X












    syntax Int ::= Int "up/Int" Int [function, klabel(ceilDiv)]
 // -----------------------------------------------------------
    rule I1 up/Int I2 => (I1 +Int (I2 -Int 1)) /Int I2 requires I2 >Int 0








    syntax Int ::= intSize ( Int ) [function]
 // -----------------------------------------
    rule intSize(N) => (log2Int(N) +Int 2) up/Int 64 requires N >Int 0
    rule intSize(0) => 1
    rule intSize(N) => intSize(~Int N) requires N <Int 0

    syntax Int ::= intSizes ( Ints ) [function]
 // -------------------------------------------
    rule intSizes(.Ints) => 0
    rule intSizes(I , INTS) => intSize(I) +Int intSizes(INTS)

    syntax Int ::= intSizes ( Array , Int , Schedule ) [function, klabel(intSizesArr)]
                 | intSizes ( Array , Int , Int , Schedule ) [function, klabel(intSizesAux)]
 // -----------------------------------------------------------------------------
    rule intSizes(ARR::Array, I, SCHED) => intSizes(ARR, I, 0, SCHED)
    rule intSizes(ARR::Array, I, I, _) => 0
    rule intSizes(ARR, I, J, SCHED) => intSize(getInt(ARR [ J ])) +Int intSizes(ARR, I, J +Int 1, SCHED)
      requires SCHED =/=K ALBE [owise]
    rule intSizes(ARR, I, J, ALBE) => getInt(ARR [ J ]) +Int intSizes(ARR, I, J +Int 1, ALBE) [owise]

    syntax Int ::= bitsInWords ( Int , Schedule ) [function]
 // ---------------------------------------------
    rule bitsInWords(I, ALBE) => I up/Int 256
    rule bitsInWords(I, _) => I up/Int 64 [owise]

    syntax Int ::= bytesInWords ( Int ) [function]
 // ----------------------------------------------
    rule bytesInWords(I) => I up/Int 8





    syntax Int ::= powmod(Int, Int, Int) [function]
 // -----------------------------------------------
    rule powmod(W0, W1, W2) => W0 ^%Int W1 W2 requires W2 =/=Int 0





    syntax Int ::= gcdInt(Int, Int)  [function]
                 | #gcdInt(Int, Int) [function]
 // -------------------------------------------
    rule gcdInt(A, B) => #gcdInt(absInt(A), absInt(B)) requires absInt(A) >=Int absInt(B)
    rule gcdInt(A, B) => #gcdInt(absInt(B), absInt(A)) [owise]
    rule #gcdInt(A, 0) => A
    rule #gcdInt(A, B) => #gcdInt(B, A modInt B) [owise]








    syntax Int ::= byte ( Int , Int ) [function]
 // --------------------------------------------
    rule byte(N, W) => bitRangeInt(W, N <<Int 3, 8)





    syntax Int ::= Int "<<Byte" Int [function]
 // ------------------------------------------
    rule N <<Byte M => N <<Int (8 *Int M)







    syntax Int ::= signextend ( Int , Int ) [function]
                 | twos ( Int , Int )       [function]
                 | bswap ( Int , Int )      [function]
 // --------------------------------------------------
    rule signextend(N, W) => signExtendBitRangeInt(W, 0, N <<Int 3)

    rule twos(N, W) => bitRangeInt(W, 0, N <<Int 3)

    rule bswap(N, W) => Bytes2Int(Int2Bytes(N, twos(N, W), BE), LE, Unsigned)





    syntax Int ::= keccak ( Bytes ) [function]
 // ----------------------------------------------
    rule keccak(WS) => #parseHexWord(Keccak256(Bytes2String(WS)))















    syntax WordStack [flatPredicate]
    syntax WordStack ::= ".WordStack" | Int ":" WordStack
 // -----------------------------------------------------









    syntax WordStack ::= WordStack "++" WordStack [function, right]
 // ---------------------------------------------------------------
    rule .WordStack ++ WS' => WS'
    rule (W : WS)   ++ WS' => W : (WS ++ WS')

    syntax WordStack ::= #rev ( WordStack , WordStack ) [function]
 // --------------------------------------------------------------
    rule #rev ( .WordStack , WS ) => WS
    rule #rev ( W : WS1 , WS2 ) => #rev(WS1, W : WS2)

    syntax WordStack ::= #take ( Int , WordStack ) [function]
                       | #take ( Int , WordStack , WordStack ) [function, klabel(#takeAux)]
 // ---------------------------------------------------------------------------------------
    rule #take(N, WS)             => #take(N, WS, .WordStack)
    rule #take(0, _, WS)          => #rev(WS, .WordStack)
    rule #take(N, .WordStack, WS) => #take(N -Int 1, .WordStack, 0 : WS)  requires N >Int 0
    rule #take(N, (W : WS1), WS2) => #take(N -Int 1, WS1,        W : WS2) requires N >Int 0

    syntax WordStack ::= #drop ( Int , WordStack ) [function]
 // ---------------------------------------------------------
    rule #drop(0, WS)         => WS
    rule #drop(N, .WordStack) => .WordStack
    rule #drop(N, (W : WS))   => #drop(N -Int 1, WS) [owise]

    syntax WordStack ::= WordStack "[" Int ".." Int "]" [function]
 // --------------------------------------------------------------
    rule WS::WordStack [ START .. WIDTH ] => #take(chop(WIDTH), #drop(chop(START), WS))





    syntax WordStack ::= WordStack "[" Int ":=" WordStack "]" [function, klabel(assignWordStackRange)]
 // --------------------------------------------------------------------------------------------------
    rule WS1::WordStack [ N := WS2::WordStack ] => #take(N, WS1) ++ WS2 ++ #drop(N +Int #sizeWordStack(WS2), WS1)






    syntax Int ::= #sizeWordStack ( WordStack )       [function, smtlib(sizeWordStack)]
                 | #sizeWordStack ( WordStack , Int ) [function, klabel(sizeWordStackAux), smtlib(sizeWordStackAux)]
 // ----------------------------------------------------------------------------------------------------------------
    rule #sizeWordStack ( WS ) => #sizeWordStack(WS, 0)
    rule #sizeWordStack ( .WordStack, SIZE ) => SIZE
    rule #sizeWordStack ( W : WS, SIZE )     => #sizeWordStack(WS, SIZE +Int 1)





    syntax WordStack ::= #padToWidth ( Int , WordStack ) [function]
 // ---------------------------------------------------------------
    rule #padToWidth(N, WS) => WS                     requires notBool #sizeWordStack(WS) <Int N
    rule #padToWidth(N, WS) => #padToWidth(N, 0 : WS) requires #sizeWordStack(WS) <Int N










    syntax Array ::= ".Array" [function]
 // ---------------------------------------------
    rule .Array => makeArray(pow30, 0)














    syntax Int ::= #asUnsigned ( WordStack ) [function]
 // ---------------------------------------------------
    rule #asUnsigned( .WordStack )    => 0
    rule #asUnsigned( W : .WordStack) => W
    rule #asUnsigned( W0 : W1 : WS )  => #asUnsigned(((W0 <<Int 8) |Int W1) : WS)

    syntax Account ::= #asAccount ( String ) [function]
 // ------------------------------------------------------
    rule #asAccount("") => .Account
    rule #asAccount("0x") => .Account
    rule #asAccount(S::String) => #parseHexWord(S) [owise]

    syntax Bytes ::= Bytes "[" Int ".." Int "]" [function, klabel(bytesRange)]
 // --------------------------------------------------------------------------
    rule B::Bytes [ I .. J ] => padRightBytes(substrBytes(B, I, minInt(lengthBytes(B), I +Int J)), J, 0)
      requires I <Int lengthBytes(B)
    rule B::Bytes [ I .. J ] => padRightBytes(.Bytes, J, 0) [owise]

    syntax Bytes ::= Bytes "[" Int ":=" Bytes "]" [function, klabel(assignBytesRange)]
 // ----------------------------------------------------------------------------------
    rule B::Bytes [ I := B'::Bytes ] => replaceAtBytes(padRightBytes(B, I +Int lengthBytes(B'), 0), I, B')
      requires B' =/=K .Bytes
    rule B::Bytes [ I := B'::Bytes ] => B
      requires B' ==K .Bytes








    syntax Int ::= #addr ( Int ) [function]
 // ---------------------------------------
    rule #addr(W) => W modInt pow256






    syntax Int ::= #newAddr ( Int , Int ) [function]
 // ------------------------------------------------
    rule #newAddr(ACCT, NONCE) => 
      (#parseHexWord(Keccak256(
         Bytes2String(Int2Bytes(32, ACCT, BE)) +String Bytes2String(Int2Bytes(NONCE, BE, Unsigned))
      ))
      &Int #parseHexWord("0xffffffffffffffffffffffffffffffffffffffff0000"))
      |Int
      (ACCT &Int #parseHexWord("0xffff"))
   
    syntax Account ::= #sender ( String , Int , String , String )                                         [function, klabel(#senderAux)]
                     | #sender ( String )                                                                 [function, klabel(#senderAux2)]
 // ---------------------------------------------------------------------------------------------------------------------------------
    rule #sender(HT, TW, TR, TS) => #sender(ECDSARecover(HT, TW, TR, TS))

    rule #sender("")  => .Account
    rule #sender(STR) => #addr(#parseHexWord(Keccak256(STR))) requires STR =/=String ""








    syntax Map ::= #removeZeros ( Map ) [function]
                 | #removeZeros ( List , Map ) [function, klabel(#removeZerosAux)]
 // ------------------------------------------------------------------------------
    rule #removeZeros( M )                                   => #removeZeros(Set2List(keys(M)), M)
    rule #removeZeros( .List, .Map )                         => .Map
    rule #removeZeros( ListItem(KEY) L, KEY |-> 0 REST )     => #removeZeros(L, REST)
    rule #removeZeros( ListItem(KEY) L, KEY |-> VALUE REST ) => KEY |-> VALUE #removeZeros(L, REST) requires VALUE =/=K 0




















    syntax Int ::= #parseHexWord ( String ) [function]
                 | #parseWord    ( String ) [function]
 // --------------------------------------------------
    rule #parseHexWord("")   => 0
    rule #parseHexWord("0x") => 0
    rule #parseHexWord(S)    => String2Base(replaceAll(S, "0x", ""), 16) requires (S =/=String "") andBool (S =/=String "0x")

    rule #parseWord("") => 0
    rule #parseWord(S)  => #parseHexWord(S) requires lengthString(S) >=Int 2 andBool substrString(S, 0, 2) ==String "0x"
    rule #parseWord(S)  => String2Int(S) [owise]

    syntax WordStack ::= #parseByteStack ( String )    [function]
                       | #parseByteStack ( String , WordStack , Int , Int ) [function, klabel(#parseByteStackAux)]
                       | #parseByteStackRaw ( String ) [function]
                       | #parseByteStackRaw ( String , WordStack , Int , Int ) [function, klabel(#parseByteStackRawAux)]
 // --------------------------------------------------------------------------------------------------------------------
    rule #parseByteStack(S) => #fun(STR => #parseByteStack(STR, .WordStack, 0, lengthString(STR)))(replaceAll(S, "0x", ""))
    rule #parseByteStack(_, WS, LEN, LEN) => #rev(WS, .WordStack)
    rule #parseByteStack(S, WS, I, LEN)  => #parseByteStack(S, #parseHexWord(substrString(S, I, I +Int 2)) : WS, I +Int 2, LEN) [owise]
    rule #parseByteStackRaw(S) => #parseByteStackRaw(S, .WordStack, 0, lengthString(S))
    rule #parseByteStackRaw(S, WS, LEN, LEN) => #rev(WS, .WordStack)
    rule #parseByteStackRaw(S, WS, I, LEN) => #parseByteStackRaw(S, ordChar(substrString(S, I, I +Int 1)) : WS, I +Int 1, LEN) [owise]

    syntax Map ::= #parseMap ( JSON ) [function]
 // --------------------------------------------
    rule #parseMap( { .JSONList                   } ) => .Map
    rule #parseMap( { _   : (VALUE:String) , REST } ) => #parseMap({ REST })                                                requires #parseHexWord(VALUE) ==K 0
    rule #parseMap( { KEY : (VALUE:String) , REST } ) => #parseMap({ REST }) [ #parseHexWord(KEY) <- #parseHexWord(VALUE) ] requires #parseHexWord(VALUE) =/=K 0

    syntax Int ::= #parseAddr ( String ) [function]
 // -----------------------------------------------
    rule #parseAddr(S) => #addr(#parseHexWord(S))










    syntax String ::= #unparseByteStack ( WordStack )                [function]
                    | #unparseByteStack ( WordStack , StringBuffer ) [function, klabel(#unparseByteStackAux)]
 // ---------------------------------------------------------------------------------------------------------
    rule #unparseByteStack ( WS ) => #unparseByteStack(WS, .StringBuffer)

    rule #unparseByteStack( .WordStack, BUFFER ) => StringBuffer2String(BUFFER)
    rule #unparseByteStack( W : WS, BUFFER )     => #unparseByteStack(WS, BUFFER +String chrChar(W))

















    syntax String ::= #rlpEncodeWord ( Int )            [function]
                    | #rlpEncodeBytes ( Int , Int )     [function]
                    | #rlpEncodeString ( String )       [function]
                    | #rlpEncodeInts ( Ints ) [function, klabel(rlpEncodeInts), symbol]
                    | #rlpEncodeInts ( StringBuffer, Ints ) [function, klabel(#rlpEncodeIntsAux)]
 // ---------------------------------------------------------------------------------------------
    rule #rlpEncodeWord(0) => "\x80"
    rule #rlpEncodeWord(WORD) => chrChar(WORD) requires WORD >Int 0 andBool WORD <Int 128
    rule #rlpEncodeWord(WORD) => #rlpEncodeLength(Bytes2String(Int2Bytes(WORD, BE, Unsigned)), 128) requires WORD >=Int 128

    rule #rlpEncodeBytes(WORD, LEN) => #rlpEncodeString(Bytes2String(Int2Bytes(LEN, WORD, BE)))

    rule #rlpEncodeString(STR) => STR                        requires lengthString(STR) ==Int 1 andBool ordChar(STR) <Int 128
    rule #rlpEncodeString(STR) => #rlpEncodeLength(STR, 128) [owise]

    rule #rlpEncodeInts(INTS) => #rlpEncodeInts(.StringBuffer, INTS)
    rule #rlpEncodeInts(BUF => BUF +String #rlpEncodeString(Bytes2String(Int2Bytes(I, BE, Signed))), (I , INTS) => INTS) requires I =/=Int 0
    rule #rlpEncodeInts(BUF => BUF +String #rlpEncodeString("\x00"), (0, INTS) => INTS)
    rule #rlpEncodeInts(BUF, .Ints) => #rlpEncodeLength(StringBuffer2String(BUF), 192)

    syntax String ::= #rlpEncodeLength ( String , Int )          [function]
                    | #rlpEncodeLength ( String , Int , String ) [function, klabel(#rlpEncodeLengthAux)]
 // ----------------------------------------------------------------------------------------------------
    rule #rlpEncodeLength(STR, OFFSET) => chrChar(lengthString(STR) +Int OFFSET) +String STR requires lengthString(STR) <Int 56
    rule #rlpEncodeLength(STR, OFFSET) => #rlpEncodeLength(STR, OFFSET, Bytes2String(Int2Bytes(lengthString(STR), BE, Unsigned))) requires lengthString(STR) >=Int 56
    rule #rlpEncodeLength(STR, OFFSET, BL) => chrChar(lengthString(BL) +Int OFFSET +Int 55) +String BL +String STR










    syntax LengthPrefixType ::= "#str" | "#list"
    syntax LengthPrefix ::= LengthPrefixType "(" Int "," Int ")"
 // ------------------------------------------------------------

    syntax Int ::= #loadLen ( WordStack ) [function]
 // ------------------------------------------------
    rule #loadLen ( B0 : WS ) => 1                               requires B0  <Int 128 orBool  B0 >=Int 192
    rule #loadLen ( B0 : WS ) => B0 -Int 128                     requires B0 >=Int 128 andBool B0  <Int 184
    rule #loadLen ( B0 : WS ) => #asUnsigned(#take(B0 -Int 183, WS)) requires B0 >=Int 184 andBool B0  <Int 192

    syntax Int ::= #loadOffset ( WordStack ) [function]
 // ---------------------------------------------------
    rule #loadOffset ( B0 : WS ) => 0           requires B0  <Int 128 orBool  B0 >=Int 192
    rule #loadOffset ( B0 : WS ) => 1           requires B0 >=Int 128 andBool B0  <Int 184
    rule #loadOffset ( B0 : WS ) => B0 -Int 182 requires B0 >=Int 184 andBool B0  <Int 192

    syntax JSON ::= #rlpDecode(String)               [function, klabel(rlpDecode), symbol]
                  | #rlpDecode(String, LengthPrefix) [function, klabel(#rlpDecodeAux)]
 // ----------------------------------------------------------------------------------
    rule #rlpDecode(STR) => #rlpDecode(STR, #decodeLengthPrefix(STR, 0))
    rule #rlpDecode(STR, #str(LEN, POS))  => substrString(STR, POS, POS +Int LEN)
    rule #rlpDecode(STR, #list(LEN, POS)) => [#rlpDecodeList(STR, POS)]

    syntax JSONList ::= #rlpDecodeList(String, Int)               [function]
                      | #rlpDecodeList(String, Int, LengthPrefix) [function, klabel(#rlpDecodeListAux)]
 // ---------------------------------------------------------------------------------------------------
    rule #rlpDecodeList(STR, POS) => #rlpDecodeList(STR, POS, #decodeLengthPrefix(STR, POS)) requires POS <Int lengthString(STR)
    rule #rlpDecodeList(STR, POS) => .JSONList [owise]
    rule #rlpDecodeList(STR, POS, _:LengthPrefixType(L, P)) => #rlpDecode(substrString(STR, POS, L +Int P)) , #rlpDecodeList(STR, L +Int P)

    syntax LengthPrefixType ::= "#str" | "#list"
    syntax LengthPrefix ::= LengthPrefixType "(" Int "," Int ")"
                          | #decodeLengthPrefix ( String , Int )                                [function]
                          | #decodeLengthPrefix ( String , Int , Int )                          [function, klabel(#decodeLengthPrefixAux)]
                          | #decodeLengthPrefixLength ( LengthPrefixType , String , Int , Int ) [function]
                          | #decodeLengthPrefixLength ( LengthPrefixType , Int    , Int , Int ) [function, klabel(#decodeLengthPrefixLengthAux)]
 // --------------------------------------------------------------------------------------------------------------------------------------------
    rule #decodeLengthPrefix(STR, START) => #decodeLengthPrefix(STR, START, ordChar(substrString(STR, START, START +Int 1)))

    rule #decodeLengthPrefix(STR, START, B0) => #str(1, START)                                   requires B0 <Int 128
    rule #decodeLengthPrefix(STR, START, B0) => #str(B0 -Int 128, START +Int 1)                  requires B0 >=Int 128 andBool B0 <Int (128 +Int 56)
    rule #decodeLengthPrefix(STR, START, B0) => #decodeLengthPrefixLength(#str, STR, START, B0)  requires B0 >=Int (128 +Int 56) andBool B0 <Int 192
    rule #decodeLengthPrefix(STR, START, B0) => #list(B0 -Int 192, START +Int 1)                 requires B0 >=Int 192 andBool B0 <Int 192 +Int 56
    rule #decodeLengthPrefix(STR, START, B0) => #decodeLengthPrefixLength(#list, STR, START, B0) [owise]

    rule #decodeLengthPrefixLength(#str,  STR, START, B0) => #decodeLengthPrefixLength(#str,  START, B0 -Int 128 -Int 56 +Int 1, Bytes2Int(String2Bytes(substrString(STR, START +Int 1, START +Int 1 +Int (B0 -Int 128 -Int 56 +Int 1))), BE, Unsigned))
    rule #decodeLengthPrefixLength(#list, STR, START, B0) => #decodeLengthPrefixLength(#list, START, B0 -Int 192 -Int 56 +Int 1, Bytes2Int(String2Bytes(substrString(STR, START +Int 1, START +Int 1 +Int (B0 -Int 192 -Int 56 +Int 1))), BE, Unsigned))
    rule #decodeLengthPrefixLength(TYPE, START, LL, L) => TYPE(L, START +Int 1 +Int LL)

endmodule

