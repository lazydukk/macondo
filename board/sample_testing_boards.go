package board

// This file contains some sample filled boards, used solely for testing.

import "github.com/domino14/macondo/alphabet"

// VsWho is an enumeration
type VsWho uint8

const (
	// VsEd was a game I played against Ed, under club games 20150127vEd
	VsEd VsWho = iota
	// VsMatt was a game I played against Matt Graham, 2018 Lake George tourney
	VsMatt
	// VsJeremy was a game I played against Jeremy Hall, 2018-11 Manhattan tourney
	VsJeremy
	// VsOxy is a constructed game that has a gigantic play available.
	VsOxy
	// VsMatt2 at the 2018-11 Manhattan tourney
	VsMatt2
	// VsRoy at the 2011 California Open
	VsRoy
	// VsMacondo1 is poor_endgame_timing.gcg
	VsMacondo1
	// JDvsNB is a test for endgames
	JDvsNB
	// VsAlec at the 2019 Nationals
	VsAlec
	// VsAlec2 same game as above just a couple turns later.
	VsAlec2
	// VsJoey from Lake George 2019
	VsJoey
	// VsCanik from 2019 Nationals
	VsCanik
	// JoeVsPaul, sample endgame given in the Maven paper
	JoeVsPaul
)

// SetToGame sets the board to a specific game in progress. It is used to
// generate test cases.
func (b *GameBoard) SetToGame(alph *alphabet.Alphabet, game VsWho) *TilesInPlay {
	// Set the board to a game
	switch game {
	case VsEd:
		// Quackle generates 219 total unique moves with a rack of AFGIIIS
		return b.SetFromPlaintext(`
cesar: Turn 8
   A B C D E F G H I J K L M N O   -> cesar                    AFGIIIS   182
   ------------------------------     ed                       ADEILNV   226
 1|=     '       =       '     E| --Tracking-----------------------------------
 2|  -       "       "       - N| ?AAAAACCDDDEEIIIIKLNOOOQRRRRSTTTTUVVZ  37
 3|    -       '   '       -   d|
 4|'     -       '       -     U|
 5|        G L O W S   -       R|
 6|  "       "     P E T     " E|
 7|    '       ' F A X I N G   R|
 8|=     '     J A Y   T E E M S|
 9|    B     B O Y '       N    |
10|  " L   D O E     "     U "  |
11|    A N E W         - P I    |
12|'   M O   L E U       O N   '|
13|    E H     '   '     H E    |
14|  -       "       "       -  |
15|=     '       =       '     =|
   ------------------------------
`, alph)
	case VsMatt:
		return b.SetFromPlaintext(`
cesar: Turn 10
   A B C D E F G H I J K L M N O      matt g                   AEEHIIL   341
   ------------------------------  -> cesar                    AABDELT   318
 1|=     '       Z E P   F     =| --Tracking-----------------------------------
 2|  F L U K Y       R   R   -  | AEEEGHIIIILMRUUWWY  18
 3|    -     E X   ' A   U -    |
 4|'   S C A R I E S T   I     '|
 5|        -         T O T      |
 6|  "       " G O   L O     "  |
 7|    '       O R ' E T A '    | ↓
 8|=     '     J A B S   b     =|
 9|    '     Q I   '     A '    | ↓
10|  "       I   N   "   N   "  | ↓
11|      R e S P O N D - D      | ↓
12|' H O E       V       O     '| ↓
13|  E N C O M I A '     N -    | ↓
14|  -       "   T   "       -  |
15|=     V E N G E D     '     =|
   ------------------------------
`, alph)
	case VsJeremy:
		return b.SetFromPlaintext(`
jeremy hall: Turn 13
   A B C D E F G H I J K L M N O   -> jeremy hall              DDESW??   299
   ------------------------------     cesar                    AHIILR    352
 1|=     '       N       '     M| --Tracking-----------------------------------
 2|  -       Z O O N "       A A| AHIILR  6
 3|    -       ' B '       - U N|
 4|'   S -       L       L A D Y|
 5|    T   -     E     Q I   I  |
 6|  " A     P O R N "     N O R|
 7|    B I C E '   A A   D A   E|
 8|=     '     G U V S   O P   F|
 9|    '       '   E T   L A   U|
10|  "       J       R   E   U T|
11|        V O T E   I - R   N E|
12|'     -   G   M I C K I E S '|
13|    -       F E ' T   T H E W|
14|  -       " O R   "   E   X I|
15|=     '     O Y       '     G|
   ------------------------------
`, alph)
	case VsOxy:
		// lol
		return b.SetFromPlaintext(`
cesar: Turn 11
   A B C D E F G H I J K L M N O      rubin                    ADDELOR   345
   ------------------------------  -> cesar                    OXPBAZE   129
 1|= P A C I F Y I N G   '     =| --Tracking-----------------------------------
 2|  I S     "       "       -  | ADDELORRRTVV  12
 3|Y E -       '   '       -    |
 4|' R E Q U A L I F I E D     '|
 5|H   L   -           -        |
 6|E D S     "       "       "  |
 7|N O '     T '   '       '    |
 8|= R A I N W A S H I N G     =|
 9|U M '     O '   '       '    |
10|T "   E   O       "       "  |
11|  W A K E n E R S   -        |
12|' O n E T I M E       -     '|
13|O O T     E ' B '       -    |
14|N -       "   U   "       -  |
15|= J A C U L A T I N G '     =|
   ------------------------------
`, alph)
	case VsMatt2:
		return b.SetFromPlaintext(`
cesar: Turn 8
   A B C D E F G H I J K L M N O   -> cesar                    EEILNT?   237
   ------------------------------     matt graham              EIJPSTW   171
 1|=     '       =       '     R| --Tracking-----------------------------------
 2|  -       "       "     Q - E| AABCDDDEEEEEHIIIIJLLLMNOPRSSSTTTUUVWWY  38
 3|    T I G E R   '     H I   I|
 4|'     -     O F       U     N|
 5|        O C E A N   P R A N K|
 6|  "       "   B A Z A R   "  |
 7|    '       '   '     A '    |
 8|=     '       M O O N Y     =|
 9|    '       D I F       '    |
10|  "       V E G   "       "  |
11|        -     S A n T O O R  |
12|'     -       '     O X     '|
13|    -       ' A G U E   -    |
14|  -       "       "       -  |
15|=     '       =       '     =|
   ------------------------------
`, alph)
	case VsRoy:
		return b.SetFromPlaintext(`
cesar: Turn 10
   A B C D E F G H I J K L M N O      roy                      WZ        427
   ------------------------------  -> cesar                    EFHIKOQ   331
 1|=     '       =     L U R I D| --Tracking-----------------------------------
 2|  - O     "       "       - I| WZ  2
 3|    U       '   P R I C E R S|
 4|O U T R A T E S       O     T|
 5|    V   -           - u     E|
 6|G " I   C O L O N I A L   " N|
 7|A   E S     '   '     T '   D|
 8|N     E       U P B Y E     E|
 9|J   ' R     M   ' O   R '   D|
10|A B   E N " A G A V E S   "  |
11|  L   N O   F   M I X        |
12|' I   A N   I '   D   -     '|
13|  G A T E W A Y s       -    |
14|  H   E   "       "       -  |
15|= T   '       =       '     =|
   ------------------------------
`, alph)
	case VsMacondo1:
		return b.SetFromPlaintext(`
teich: Turn 12
   A B C D E F G H I J K L M N O      cesar                    ENNR      379
   ------------------------------  -> teich                    APRS?     469
 1|J O Y E D     =       '     =| --Tracking-----------------------------------
 2|U - E L   V       "       -  | ENNR  4
 3|G   W O   I '   '       -    |
 4|A I   P   G   '       -     '|
 5|    F E T A         -        |
 6|  Y I R R S       C       "  |
 7|    L   I   O B I A     '    |
 8|U     H A I K A   L   '     =|
 9|N   Z   L   '   ' O F   '    |
10|I T E M I S E D   T O     "  |
11|T   B A N     E   T O        |
12|E   R U G     V   e D -     '|
13|  H A D     ' O '       -    |
14|W E       "   I N Q U E S T  |
15|E X E C       R       M O A N|
   ------------------------------
`, alph)

	case JDvsNB:
		return b.SetFromPlaintext(`
Nathan Benedict: Turn 14
   A B C D E F G H I J K L M N O   -> Nathan Benedict          RR        365
   ------------------------------     JD                       LN        510
 1|G R R L       d     Q '     =| --Tracking-----------------------------------
 2|  E       J   E   G I     -  | LN  2
 3|  M M   L A ' N Y E     -    |
 4|'   O B I     O O N   -     '|
 5|    K I P     T   I F        |
 6|  W E T A "   I   T O     "  |
 7|S O ' C     ' V ' O B   '    |
 8|U T   H E D G E   R   ' V   =|
 9|I   '     U '   '       E    |
10|  "       P       F A A N "  |
11|    C   - E N T A Y L E D    |
12|W   H O A S   A       - I   '|
13|E   I       ' T '       N    |
14|D A Z E D "   O   "     g -  |
15|S     X I     U     U R S A E|
   ------------------------------
`, alph)

	case VsAlec:
		return b.SetFromPlaintext(`
cesar: Turn 11
   A B C D E F G H I J K L M N O      alec                     EGNOQR    420
   ------------------------------  -> cesar                    DGILOPR   369
 1|=     '       L       '     =| --Tracking-----------------------------------
 2|  -       "   A   J       -  | EGNOQR  6
 3|    -       ' T W I N E R    |
 4|'     A       E   V O T E R '|
 5|  B U R A N   E   E -        |
 6|  "   c H U T N E Y S     "  |
 7|    ' A     '   '       '    |
 8|W O O D       S       '     =|
 9|E   ' I F   ' L '       '    |
10|I "   A A "   E   "       "  |
11|R   C   U   P A M   -        |
12|D   A L G U A Z I l   -     C|
13|O   N   H   V E X       F   Y|
14|E - T     K I   T O O N I E S|
15|S     ' M I D I       ' B   T|
   ------------------------------
`, alph)

	case VsAlec2:
		return b.SetFromPlaintext(`
cesar: Turn 12
   A B C D E F G H I J K L M N O      alec                     ENQR      438
   ------------------------------  -> cesar                    DGILOR    383
 1|=     '       L       '     =| --Tracking-----------------------------------
 2|  -       "   A   J O G   -  | ENQR  4
 3|    -       ' T W I N E R    |
 4|'     A       E   V O T E R '|
 5|  B U R A N   E   E -        |
 6|  "   c H U T N E Y S     "  |
 7|    P A     '   '       '    |
 8|W O O D       S       '     =|
 9|E   ' I F   ' L '       '    |
10|I "   A A "   E   "       "  |
11|R   C   U   P A M   -        |
12|D   A L G U A Z I l   -     C|
13|O   N   H   V E X       F   Y|
14|E - T     K I   T O O N I E S|
15|S     ' M I D I       ' B   T|
   ------------------------------
`, alph)

	case VsJoey:
		return b.SetFromPlaintext(`
Joey: Turn 11
   A B C D E F G H I J K L M N O      Cesar                    DIV       412
   ------------------------------  -> Joey                     AEFILMR   371
 1|A I D E R     U       '     =| --Tracking-----------------------------------
 2|b - E   E "   N   Z       -  | DIV  3
 3|A W N   T   ' M ' A T T -    |
 4|L I   C O B L E     O W     '|
 5|O P     U     E     A A      |
 6|N E     C U S T A R D S   Q  |
 7|E R ' O H   '   '   I   ' U  |
 8|S     K     F O B   E R G O T|
 9|    '     H E X Y L S   ' I  |
10|  "     J I N     "       N  |
11|    G O O P     N A I V E s T|
12|' D I R E     '       -     '|
13|    G A Y   '   '       -    |
14|  -       "       "       -  |
15|=     '       =       '     =|
   ------------------------------
`, alph)

	case VsCanik:
		return b.SetFromPlaintext(`
cesar: Turn 12
   A B C D E F G H I J K L M N O      canik                    DEHILOR   389
   ------------------------------  -> cesar                    BGIV      384
 1|=     '       =   A   P I X Y| --Tracking-----------------------------------
 2|  -       "       S   L   -  | DEHILOR  7
 3|    T o W N L E T S   O -    |
 4|'     -       '   U   D A   R|
 5|      G E R A N I A L   U   I|
 6|  "       "       g     T " C|
 7|    '       '   W E     O B I|
 8|=     '     E M U     '   O N|
 9|    '       A I D       G O  |
10|  "       H U N   "     E T  |
11|        Z A   T     -   M E  |
12|' Q   F A K E Y       J O E S|
13|F I V E   E '   '     I T   C|
14|  -       S P O R R A N   - A|
15|=     '     O R E     N     D|
   ------------------------------
`, alph)

	case JoeVsPaul:
		return b.SetFromPlaintext(`
joe: Turn 12
   A B C D E F G H I J K L M N O      joe                      ILMZ      296
   ------------------------------  -> paul                     ?AEINRU   296
 1|=     '   B E R G S   '     =| --Tracking-----------------------------------
 2|  -     P A       U       -  | ILMZ 4
 3|    Q A I D '   ' R     -    |
 4|'     B E E   '   F   T S K '|
 5|  P   E T     V I A T I C    |
 6|M A   T A W       c     H "  |
 7|E S '     I S   ' E     A    |
 8|A T   F O L I A       ' V   =|
 9|L I ' L   E X   E       '    |
10|  N   O   D     N "   Y   "  |
11|  G N U -   C   J E T E      |
12|'   E R     O H O     N     '|
13|    O       G O Y       -    |
14|  I N D O W   U   "       -  |
15|=     ' D O R R       '     =|
   ------------------------------
`, alph)

	}

	return nil
}
