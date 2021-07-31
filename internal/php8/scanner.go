// line internal/php8/scanner.rl:1
package php8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VKCOM/php-parser/pkg/token"
)

// line internal/php8/scanner.go:15
const lexer_start int = 121
const lexer_first_final int = 121
const lexer_error int = 0

const lexer_en_main int = 121
const lexer_en_html int = 124
const lexer_en_php int = 131
const lexer_en_property int = 490
const lexer_en_nowdoc int = 497
const lexer_en_heredoc int = 501
const lexer_en_backqote int = 508
const lexer_en_template_string int = 514
const lexer_en_heredoc_end int = 520
const lexer_en_string_var int = 522
const lexer_en_string_var_index int = 528
const lexer_en_string_var_name int = 538
const lexer_en_halt_compiller_open_parenthesis int = 540
const lexer_en_halt_compiller_close_parenthesis int = 544
const lexer_en_halt_compiller_close_semicolon int = 548
const lexer_en_halt_compiller_end int = 552

// line internal/php8/scanner.rl:17

func initLexer(lex *Lexer) {

	// line internal/php8/scanner.go:43
	{
		lex.cs = lexer_start
		lex.top = 0
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

	// line internal/php8/scanner.rl:21
}

func (lex *Lexer) Lex() *token.Token {
	eof := lex.pe
	var tok token.ID

	tkn := lex.tokenPool.Get()

	lblStart := 0
	lblEnd := 0

	_, _ = lblStart, lblEnd

	// line internal/php8/scanner.go:67
	{
		var _widec int16
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 121:
			goto st121
		case 122:
			goto st122
		case 1:
			goto st1
		case 123:
			goto st123
		case 124:
			goto st124
		case 125:
			goto st125
		case 126:
			goto st126
		case 127:
			goto st127
		case 128:
			goto st128
		case 129:
			goto st129
		case 2:
			goto st2
		case 3:
			goto st3
		case 4:
			goto st4
		case 130:
			goto st130
		case 5:
			goto st5
		case 131:
			goto st131
		case 132:
			goto st132
		case 133:
			goto st133
		case 6:
			goto st6
		case 134:
			goto st134
		case 135:
			goto st135
		case 136:
			goto st136
		case 137:
			goto st137
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 138:
			goto st138
		case 139:
			goto st139
		case 140:
			goto st140
		case 141:
			goto st141
		case 142:
			goto st142
		case 143:
			goto st143
		case 144:
			goto st144
		case 145:
			goto st145
		case 11:
			goto st11
		case 12:
			goto st12
		case 146:
			goto st146
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto st40
		case 41:
			goto st41
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 45:
			goto st45
		case 46:
			goto st46
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 66:
			goto st66
		case 67:
			goto st67
		case 147:
			goto st147
		case 148:
			goto st148
		case 149:
			goto st149
		case 150:
			goto st150
		case 151:
			goto st151
		case 68:
			goto st68
		case 152:
			goto st152
		case 69:
			goto st69
		case 70:
			goto st70
		case 153:
			goto st153
		case 71:
			goto st71
		case 154:
			goto st154
		case 72:
			goto st72
		case 73:
			goto st73
		case 74:
			goto st74
		case 155:
			goto st155
		case 156:
			goto st156
		case 157:
			goto st157
		case 75:
			goto st75
		case 76:
			goto st76
		case 158:
			goto st158
		case 77:
			goto st77
		case 159:
			goto st159
		case 160:
			goto st160
		case 161:
			goto st161
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 162:
			goto st162
		case 163:
			goto st163
		case 82:
			goto st82
		case 164:
			goto st164
		case 165:
			goto st165
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 86:
			goto st86
		case 166:
			goto st166
		case 87:
			goto st87
		case 88:
			goto st88
		case 89:
			goto st89
		case 90:
			goto st90
		case 167:
			goto st167
		case 168:
			goto st168
		case 169:
			goto st169
		case 170:
			goto st170
		case 171:
			goto st171
		case 172:
			goto st172
		case 91:
			goto st91
		case 173:
			goto st173
		case 174:
			goto st174
		case 92:
			goto st92
		case 175:
			goto st175
		case 176:
			goto st176
		case 177:
			goto st177
		case 93:
			goto st93
		case 178:
			goto st178
		case 179:
			goto st179
		case 180:
			goto st180
		case 181:
			goto st181
		case 182:
			goto st182
		case 183:
			goto st183
		case 184:
			goto st184
		case 185:
			goto st185
		case 186:
			goto st186
		case 187:
			goto st187
		case 188:
			goto st188
		case 189:
			goto st189
		case 94:
			goto st94
		case 95:
			goto st95
		case 190:
			goto st190
		case 191:
			goto st191
		case 192:
			goto st192
		case 193:
			goto st193
		case 194:
			goto st194
		case 195:
			goto st195
		case 196:
			goto st196
		case 197:
			goto st197
		case 198:
			goto st198
		case 199:
			goto st199
		case 200:
			goto st200
		case 201:
			goto st201
		case 202:
			goto st202
		case 203:
			goto st203
		case 204:
			goto st204
		case 205:
			goto st205
		case 206:
			goto st206
		case 207:
			goto st207
		case 208:
			goto st208
		case 209:
			goto st209
		case 210:
			goto st210
		case 211:
			goto st211
		case 212:
			goto st212
		case 213:
			goto st213
		case 214:
			goto st214
		case 215:
			goto st215
		case 216:
			goto st216
		case 217:
			goto st217
		case 218:
			goto st218
		case 219:
			goto st219
		case 220:
			goto st220
		case 221:
			goto st221
		case 222:
			goto st222
		case 223:
			goto st223
		case 224:
			goto st224
		case 225:
			goto st225
		case 226:
			goto st226
		case 227:
			goto st227
		case 228:
			goto st228
		case 229:
			goto st229
		case 230:
			goto st230
		case 231:
			goto st231
		case 232:
			goto st232
		case 233:
			goto st233
		case 234:
			goto st234
		case 235:
			goto st235
		case 236:
			goto st236
		case 237:
			goto st237
		case 238:
			goto st238
		case 239:
			goto st239
		case 240:
			goto st240
		case 241:
			goto st241
		case 242:
			goto st242
		case 243:
			goto st243
		case 244:
			goto st244
		case 245:
			goto st245
		case 246:
			goto st246
		case 247:
			goto st247
		case 248:
			goto st248
		case 249:
			goto st249
		case 250:
			goto st250
		case 251:
			goto st251
		case 252:
			goto st252
		case 253:
			goto st253
		case 254:
			goto st254
		case 255:
			goto st255
		case 256:
			goto st256
		case 257:
			goto st257
		case 258:
			goto st258
		case 259:
			goto st259
		case 260:
			goto st260
		case 261:
			goto st261
		case 262:
			goto st262
		case 263:
			goto st263
		case 264:
			goto st264
		case 265:
			goto st265
		case 266:
			goto st266
		case 267:
			goto st267
		case 268:
			goto st268
		case 269:
			goto st269
		case 270:
			goto st270
		case 271:
			goto st271
		case 272:
			goto st272
		case 273:
			goto st273
		case 274:
			goto st274
		case 275:
			goto st275
		case 276:
			goto st276
		case 277:
			goto st277
		case 278:
			goto st278
		case 279:
			goto st279
		case 280:
			goto st280
		case 281:
			goto st281
		case 282:
			goto st282
		case 283:
			goto st283
		case 284:
			goto st284
		case 285:
			goto st285
		case 286:
			goto st286
		case 287:
			goto st287
		case 288:
			goto st288
		case 289:
			goto st289
		case 290:
			goto st290
		case 291:
			goto st291
		case 292:
			goto st292
		case 293:
			goto st293
		case 294:
			goto st294
		case 295:
			goto st295
		case 296:
			goto st296
		case 297:
			goto st297
		case 298:
			goto st298
		case 299:
			goto st299
		case 300:
			goto st300
		case 301:
			goto st301
		case 302:
			goto st302
		case 303:
			goto st303
		case 304:
			goto st304
		case 305:
			goto st305
		case 306:
			goto st306
		case 307:
			goto st307
		case 308:
			goto st308
		case 309:
			goto st309
		case 310:
			goto st310
		case 311:
			goto st311
		case 312:
			goto st312
		case 313:
			goto st313
		case 314:
			goto st314
		case 315:
			goto st315
		case 316:
			goto st316
		case 317:
			goto st317
		case 318:
			goto st318
		case 319:
			goto st319
		case 320:
			goto st320
		case 321:
			goto st321
		case 322:
			goto st322
		case 323:
			goto st323
		case 324:
			goto st324
		case 325:
			goto st325
		case 326:
			goto st326
		case 327:
			goto st327
		case 328:
			goto st328
		case 329:
			goto st329
		case 330:
			goto st330
		case 331:
			goto st331
		case 332:
			goto st332
		case 333:
			goto st333
		case 334:
			goto st334
		case 335:
			goto st335
		case 336:
			goto st336
		case 337:
			goto st337
		case 338:
			goto st338
		case 339:
			goto st339
		case 340:
			goto st340
		case 341:
			goto st341
		case 342:
			goto st342
		case 343:
			goto st343
		case 344:
			goto st344
		case 345:
			goto st345
		case 346:
			goto st346
		case 347:
			goto st347
		case 96:
			goto st96
		case 348:
			goto st348
		case 349:
			goto st349
		case 350:
			goto st350
		case 351:
			goto st351
		case 352:
			goto st352
		case 353:
			goto st353
		case 354:
			goto st354
		case 355:
			goto st355
		case 356:
			goto st356
		case 357:
			goto st357
		case 358:
			goto st358
		case 359:
			goto st359
		case 360:
			goto st360
		case 361:
			goto st361
		case 362:
			goto st362
		case 363:
			goto st363
		case 364:
			goto st364
		case 365:
			goto st365
		case 366:
			goto st366
		case 367:
			goto st367
		case 368:
			goto st368
		case 369:
			goto st369
		case 370:
			goto st370
		case 371:
			goto st371
		case 372:
			goto st372
		case 373:
			goto st373
		case 374:
			goto st374
		case 375:
			goto st375
		case 376:
			goto st376
		case 377:
			goto st377
		case 378:
			goto st378
		case 379:
			goto st379
		case 380:
			goto st380
		case 381:
			goto st381
		case 382:
			goto st382
		case 383:
			goto st383
		case 384:
			goto st384
		case 385:
			goto st385
		case 386:
			goto st386
		case 387:
			goto st387
		case 388:
			goto st388
		case 389:
			goto st389
		case 390:
			goto st390
		case 391:
			goto st391
		case 392:
			goto st392
		case 393:
			goto st393
		case 394:
			goto st394
		case 395:
			goto st395
		case 396:
			goto st396
		case 397:
			goto st397
		case 398:
			goto st398
		case 399:
			goto st399
		case 400:
			goto st400
		case 401:
			goto st401
		case 402:
			goto st402
		case 403:
			goto st403
		case 404:
			goto st404
		case 405:
			goto st405
		case 406:
			goto st406
		case 407:
			goto st407
		case 408:
			goto st408
		case 409:
			goto st409
		case 410:
			goto st410
		case 411:
			goto st411
		case 412:
			goto st412
		case 413:
			goto st413
		case 414:
			goto st414
		case 415:
			goto st415
		case 416:
			goto st416
		case 417:
			goto st417
		case 418:
			goto st418
		case 419:
			goto st419
		case 420:
			goto st420
		case 97:
			goto st97
		case 98:
			goto st98
		case 99:
			goto st99
		case 100:
			goto st100
		case 101:
			goto st101
		case 102:
			goto st102
		case 421:
			goto st421
		case 422:
			goto st422
		case 103:
			goto st103
		case 423:
			goto st423
		case 424:
			goto st424
		case 425:
			goto st425
		case 426:
			goto st426
		case 427:
			goto st427
		case 428:
			goto st428
		case 429:
			goto st429
		case 430:
			goto st430
		case 431:
			goto st431
		case 432:
			goto st432
		case 433:
			goto st433
		case 434:
			goto st434
		case 435:
			goto st435
		case 436:
			goto st436
		case 437:
			goto st437
		case 438:
			goto st438
		case 439:
			goto st439
		case 440:
			goto st440
		case 441:
			goto st441
		case 442:
			goto st442
		case 443:
			goto st443
		case 444:
			goto st444
		case 445:
			goto st445
		case 446:
			goto st446
		case 447:
			goto st447
		case 448:
			goto st448
		case 449:
			goto st449
		case 450:
			goto st450
		case 451:
			goto st451
		case 452:
			goto st452
		case 453:
			goto st453
		case 454:
			goto st454
		case 455:
			goto st455
		case 456:
			goto st456
		case 457:
			goto st457
		case 458:
			goto st458
		case 459:
			goto st459
		case 460:
			goto st460
		case 461:
			goto st461
		case 462:
			goto st462
		case 463:
			goto st463
		case 464:
			goto st464
		case 465:
			goto st465
		case 466:
			goto st466
		case 467:
			goto st467
		case 468:
			goto st468
		case 469:
			goto st469
		case 470:
			goto st470
		case 471:
			goto st471
		case 472:
			goto st472
		case 473:
			goto st473
		case 474:
			goto st474
		case 475:
			goto st475
		case 476:
			goto st476
		case 477:
			goto st477
		case 478:
			goto st478
		case 479:
			goto st479
		case 480:
			goto st480
		case 481:
			goto st481
		case 482:
			goto st482
		case 483:
			goto st483
		case 484:
			goto st484
		case 485:
			goto st485
		case 486:
			goto st486
		case 487:
			goto st487
		case 488:
			goto st488
		case 489:
			goto st489
		case 490:
			goto st490
		case 491:
			goto st491
		case 492:
			goto st492
		case 104:
			goto st104
		case 493:
			goto st493
		case 494:
			goto st494
		case 495:
			goto st495
		case 105:
			goto st105
		case 496:
			goto st496
		case 497:
			goto st497
		case 0:
			goto st0
		case 498:
			goto st498
		case 499:
			goto st499
		case 500:
			goto st500
		case 501:
			goto st501
		case 502:
			goto st502
		case 106:
			goto st106
		case 503:
			goto st503
		case 504:
			goto st504
		case 505:
			goto st505
		case 506:
			goto st506
		case 507:
			goto st507
		case 508:
			goto st508
		case 107:
			goto st107
		case 108:
			goto st108
		case 509:
			goto st509
		case 510:
			goto st510
		case 511:
			goto st511
		case 512:
			goto st512
		case 513:
			goto st513
		case 514:
			goto st514
		case 109:
			goto st109
		case 110:
			goto st110
		case 515:
			goto st515
		case 516:
			goto st516
		case 517:
			goto st517
		case 518:
			goto st518
		case 519:
			goto st519
		case 520:
			goto st520
		case 521:
			goto st521
		case 522:
			goto st522
		case 523:
			goto st523
		case 524:
			goto st524
		case 525:
			goto st525
		case 111:
			goto st111
		case 526:
			goto st526
		case 112:
			goto st112
		case 113:
			goto st113
		case 527:
			goto st527
		case 528:
			goto st528
		case 529:
			goto st529
		case 530:
			goto st530
		case 531:
			goto st531
		case 532:
			goto st532
		case 533:
			goto st533
		case 534:
			goto st534
		case 114:
			goto st114
		case 115:
			goto st115
		case 535:
			goto st535
		case 116:
			goto st116
		case 536:
			goto st536
		case 537:
			goto st537
		case 538:
			goto st538
		case 539:
			goto st539
		case 117:
			goto st117
		case 540:
			goto st540
		case 541:
			goto st541
		case 542:
			goto st542
		case 118:
			goto st118
		case 543:
			goto st543
		case 544:
			goto st544
		case 545:
			goto st545
		case 546:
			goto st546
		case 119:
			goto st119
		case 547:
			goto st547
		case 548:
			goto st548
		case 549:
			goto st549
		case 550:
			goto st550
		case 120:
			goto st120
		case 551:
			goto st551
		case 552:
			goto st552
		case 553:
			goto st553
		case 554:
			goto st554
		case 555:
			goto st555
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 1:
			goto st_case_1
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 130:
			goto st_case_130
		case 5:
			goto st_case_5
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 6:
			goto st_case_6
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 146:
			goto st_case_146
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 68:
			goto st_case_68
		case 152:
			goto st_case_152
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 153:
			goto st_case_153
		case 71:
			goto st_case_71
		case 154:
			goto st_case_154
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 158:
			goto st_case_158
		case 77:
			goto st_case_77
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 82:
			goto st_case_82
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 166:
			goto st_case_166
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 167:
			goto st_case_167
		case 168:
			goto st_case_168
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 91:
			goto st_case_91
		case 173:
			goto st_case_173
		case 174:
			goto st_case_174
		case 92:
			goto st_case_92
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 93:
			goto st_case_93
		case 178:
			goto st_case_178
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 181:
			goto st_case_181
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
		case 192:
			goto st_case_192
		case 193:
			goto st_case_193
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 201:
			goto st_case_201
		case 202:
			goto st_case_202
		case 203:
			goto st_case_203
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 206:
			goto st_case_206
		case 207:
			goto st_case_207
		case 208:
			goto st_case_208
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 213:
			goto st_case_213
		case 214:
			goto st_case_214
		case 215:
			goto st_case_215
		case 216:
			goto st_case_216
		case 217:
			goto st_case_217
		case 218:
			goto st_case_218
		case 219:
			goto st_case_219
		case 220:
			goto st_case_220
		case 221:
			goto st_case_221
		case 222:
			goto st_case_222
		case 223:
			goto st_case_223
		case 224:
			goto st_case_224
		case 225:
			goto st_case_225
		case 226:
			goto st_case_226
		case 227:
			goto st_case_227
		case 228:
			goto st_case_228
		case 229:
			goto st_case_229
		case 230:
			goto st_case_230
		case 231:
			goto st_case_231
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 236:
			goto st_case_236
		case 237:
			goto st_case_237
		case 238:
			goto st_case_238
		case 239:
			goto st_case_239
		case 240:
			goto st_case_240
		case 241:
			goto st_case_241
		case 242:
			goto st_case_242
		case 243:
			goto st_case_243
		case 244:
			goto st_case_244
		case 245:
			goto st_case_245
		case 246:
			goto st_case_246
		case 247:
			goto st_case_247
		case 248:
			goto st_case_248
		case 249:
			goto st_case_249
		case 250:
			goto st_case_250
		case 251:
			goto st_case_251
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 256:
			goto st_case_256
		case 257:
			goto st_case_257
		case 258:
			goto st_case_258
		case 259:
			goto st_case_259
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 262:
			goto st_case_262
		case 263:
			goto st_case_263
		case 264:
			goto st_case_264
		case 265:
			goto st_case_265
		case 266:
			goto st_case_266
		case 267:
			goto st_case_267
		case 268:
			goto st_case_268
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
		case 271:
			goto st_case_271
		case 272:
			goto st_case_272
		case 273:
			goto st_case_273
		case 274:
			goto st_case_274
		case 275:
			goto st_case_275
		case 276:
			goto st_case_276
		case 277:
			goto st_case_277
		case 278:
			goto st_case_278
		case 279:
			goto st_case_279
		case 280:
			goto st_case_280
		case 281:
			goto st_case_281
		case 282:
			goto st_case_282
		case 283:
			goto st_case_283
		case 284:
			goto st_case_284
		case 285:
			goto st_case_285
		case 286:
			goto st_case_286
		case 287:
			goto st_case_287
		case 288:
			goto st_case_288
		case 289:
			goto st_case_289
		case 290:
			goto st_case_290
		case 291:
			goto st_case_291
		case 292:
			goto st_case_292
		case 293:
			goto st_case_293
		case 294:
			goto st_case_294
		case 295:
			goto st_case_295
		case 296:
			goto st_case_296
		case 297:
			goto st_case_297
		case 298:
			goto st_case_298
		case 299:
			goto st_case_299
		case 300:
			goto st_case_300
		case 301:
			goto st_case_301
		case 302:
			goto st_case_302
		case 303:
			goto st_case_303
		case 304:
			goto st_case_304
		case 305:
			goto st_case_305
		case 306:
			goto st_case_306
		case 307:
			goto st_case_307
		case 308:
			goto st_case_308
		case 309:
			goto st_case_309
		case 310:
			goto st_case_310
		case 311:
			goto st_case_311
		case 312:
			goto st_case_312
		case 313:
			goto st_case_313
		case 314:
			goto st_case_314
		case 315:
			goto st_case_315
		case 316:
			goto st_case_316
		case 317:
			goto st_case_317
		case 318:
			goto st_case_318
		case 319:
			goto st_case_319
		case 320:
			goto st_case_320
		case 321:
			goto st_case_321
		case 322:
			goto st_case_322
		case 323:
			goto st_case_323
		case 324:
			goto st_case_324
		case 325:
			goto st_case_325
		case 326:
			goto st_case_326
		case 327:
			goto st_case_327
		case 328:
			goto st_case_328
		case 329:
			goto st_case_329
		case 330:
			goto st_case_330
		case 331:
			goto st_case_331
		case 332:
			goto st_case_332
		case 333:
			goto st_case_333
		case 334:
			goto st_case_334
		case 335:
			goto st_case_335
		case 336:
			goto st_case_336
		case 337:
			goto st_case_337
		case 338:
			goto st_case_338
		case 339:
			goto st_case_339
		case 340:
			goto st_case_340
		case 341:
			goto st_case_341
		case 342:
			goto st_case_342
		case 343:
			goto st_case_343
		case 344:
			goto st_case_344
		case 345:
			goto st_case_345
		case 346:
			goto st_case_346
		case 347:
			goto st_case_347
		case 96:
			goto st_case_96
		case 348:
			goto st_case_348
		case 349:
			goto st_case_349
		case 350:
			goto st_case_350
		case 351:
			goto st_case_351
		case 352:
			goto st_case_352
		case 353:
			goto st_case_353
		case 354:
			goto st_case_354
		case 355:
			goto st_case_355
		case 356:
			goto st_case_356
		case 357:
			goto st_case_357
		case 358:
			goto st_case_358
		case 359:
			goto st_case_359
		case 360:
			goto st_case_360
		case 361:
			goto st_case_361
		case 362:
			goto st_case_362
		case 363:
			goto st_case_363
		case 364:
			goto st_case_364
		case 365:
			goto st_case_365
		case 366:
			goto st_case_366
		case 367:
			goto st_case_367
		case 368:
			goto st_case_368
		case 369:
			goto st_case_369
		case 370:
			goto st_case_370
		case 371:
			goto st_case_371
		case 372:
			goto st_case_372
		case 373:
			goto st_case_373
		case 374:
			goto st_case_374
		case 375:
			goto st_case_375
		case 376:
			goto st_case_376
		case 377:
			goto st_case_377
		case 378:
			goto st_case_378
		case 379:
			goto st_case_379
		case 380:
			goto st_case_380
		case 381:
			goto st_case_381
		case 382:
			goto st_case_382
		case 383:
			goto st_case_383
		case 384:
			goto st_case_384
		case 385:
			goto st_case_385
		case 386:
			goto st_case_386
		case 387:
			goto st_case_387
		case 388:
			goto st_case_388
		case 389:
			goto st_case_389
		case 390:
			goto st_case_390
		case 391:
			goto st_case_391
		case 392:
			goto st_case_392
		case 393:
			goto st_case_393
		case 394:
			goto st_case_394
		case 395:
			goto st_case_395
		case 396:
			goto st_case_396
		case 397:
			goto st_case_397
		case 398:
			goto st_case_398
		case 399:
			goto st_case_399
		case 400:
			goto st_case_400
		case 401:
			goto st_case_401
		case 402:
			goto st_case_402
		case 403:
			goto st_case_403
		case 404:
			goto st_case_404
		case 405:
			goto st_case_405
		case 406:
			goto st_case_406
		case 407:
			goto st_case_407
		case 408:
			goto st_case_408
		case 409:
			goto st_case_409
		case 410:
			goto st_case_410
		case 411:
			goto st_case_411
		case 412:
			goto st_case_412
		case 413:
			goto st_case_413
		case 414:
			goto st_case_414
		case 415:
			goto st_case_415
		case 416:
			goto st_case_416
		case 417:
			goto st_case_417
		case 418:
			goto st_case_418
		case 419:
			goto st_case_419
		case 420:
			goto st_case_420
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 421:
			goto st_case_421
		case 422:
			goto st_case_422
		case 103:
			goto st_case_103
		case 423:
			goto st_case_423
		case 424:
			goto st_case_424
		case 425:
			goto st_case_425
		case 426:
			goto st_case_426
		case 427:
			goto st_case_427
		case 428:
			goto st_case_428
		case 429:
			goto st_case_429
		case 430:
			goto st_case_430
		case 431:
			goto st_case_431
		case 432:
			goto st_case_432
		case 433:
			goto st_case_433
		case 434:
			goto st_case_434
		case 435:
			goto st_case_435
		case 436:
			goto st_case_436
		case 437:
			goto st_case_437
		case 438:
			goto st_case_438
		case 439:
			goto st_case_439
		case 440:
			goto st_case_440
		case 441:
			goto st_case_441
		case 442:
			goto st_case_442
		case 443:
			goto st_case_443
		case 444:
			goto st_case_444
		case 445:
			goto st_case_445
		case 446:
			goto st_case_446
		case 447:
			goto st_case_447
		case 448:
			goto st_case_448
		case 449:
			goto st_case_449
		case 450:
			goto st_case_450
		case 451:
			goto st_case_451
		case 452:
			goto st_case_452
		case 453:
			goto st_case_453
		case 454:
			goto st_case_454
		case 455:
			goto st_case_455
		case 456:
			goto st_case_456
		case 457:
			goto st_case_457
		case 458:
			goto st_case_458
		case 459:
			goto st_case_459
		case 460:
			goto st_case_460
		case 461:
			goto st_case_461
		case 462:
			goto st_case_462
		case 463:
			goto st_case_463
		case 464:
			goto st_case_464
		case 465:
			goto st_case_465
		case 466:
			goto st_case_466
		case 467:
			goto st_case_467
		case 468:
			goto st_case_468
		case 469:
			goto st_case_469
		case 470:
			goto st_case_470
		case 471:
			goto st_case_471
		case 472:
			goto st_case_472
		case 473:
			goto st_case_473
		case 474:
			goto st_case_474
		case 475:
			goto st_case_475
		case 476:
			goto st_case_476
		case 477:
			goto st_case_477
		case 478:
			goto st_case_478
		case 479:
			goto st_case_479
		case 480:
			goto st_case_480
		case 481:
			goto st_case_481
		case 482:
			goto st_case_482
		case 483:
			goto st_case_483
		case 484:
			goto st_case_484
		case 485:
			goto st_case_485
		case 486:
			goto st_case_486
		case 487:
			goto st_case_487
		case 488:
			goto st_case_488
		case 489:
			goto st_case_489
		case 490:
			goto st_case_490
		case 491:
			goto st_case_491
		case 492:
			goto st_case_492
		case 104:
			goto st_case_104
		case 493:
			goto st_case_493
		case 494:
			goto st_case_494
		case 495:
			goto st_case_495
		case 105:
			goto st_case_105
		case 496:
			goto st_case_496
		case 497:
			goto st_case_497
		case 0:
			goto st_case_0
		case 498:
			goto st_case_498
		case 499:
			goto st_case_499
		case 500:
			goto st_case_500
		case 501:
			goto st_case_501
		case 502:
			goto st_case_502
		case 106:
			goto st_case_106
		case 503:
			goto st_case_503
		case 504:
			goto st_case_504
		case 505:
			goto st_case_505
		case 506:
			goto st_case_506
		case 507:
			goto st_case_507
		case 508:
			goto st_case_508
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 509:
			goto st_case_509
		case 510:
			goto st_case_510
		case 511:
			goto st_case_511
		case 512:
			goto st_case_512
		case 513:
			goto st_case_513
		case 514:
			goto st_case_514
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 515:
			goto st_case_515
		case 516:
			goto st_case_516
		case 517:
			goto st_case_517
		case 518:
			goto st_case_518
		case 519:
			goto st_case_519
		case 520:
			goto st_case_520
		case 521:
			goto st_case_521
		case 522:
			goto st_case_522
		case 523:
			goto st_case_523
		case 524:
			goto st_case_524
		case 525:
			goto st_case_525
		case 111:
			goto st_case_111
		case 526:
			goto st_case_526
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 527:
			goto st_case_527
		case 528:
			goto st_case_528
		case 529:
			goto st_case_529
		case 530:
			goto st_case_530
		case 531:
			goto st_case_531
		case 532:
			goto st_case_532
		case 533:
			goto st_case_533
		case 534:
			goto st_case_534
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 535:
			goto st_case_535
		case 116:
			goto st_case_116
		case 536:
			goto st_case_536
		case 537:
			goto st_case_537
		case 538:
			goto st_case_538
		case 539:
			goto st_case_539
		case 117:
			goto st_case_117
		case 540:
			goto st_case_540
		case 541:
			goto st_case_541
		case 542:
			goto st_case_542
		case 118:
			goto st_case_118
		case 543:
			goto st_case_543
		case 544:
			goto st_case_544
		case 545:
			goto st_case_545
		case 546:
			goto st_case_546
		case 119:
			goto st_case_119
		case 547:
			goto st_case_547
		case 548:
			goto st_case_548
		case 549:
			goto st_case_549
		case 550:
			goto st_case_550
		case 120:
			goto st_case_120
		case 551:
			goto st_case_551
		case 552:
			goto st_case_552
		case 553:
			goto st_case_553
		case 554:
			goto st_case_554
		case 555:
			goto st_case_555
		}
		goto st_out
	tr0:
		lex.cs = 121
		// line internal/php8/scanner.rl:130
		(lex.p) = (lex.te) - 1
		{
			lex.cs = 124
			lex.ungetCnt(1)
		}
		goto _again
	tr180:
		lex.cs = 121
		// line internal/php8/scanner.rl:130
		lex.te = (lex.p) + 1
		{
			lex.cs = 124
			lex.ungetCnt(1)
		}
		goto _again
	tr182:
		lex.cs = 121
		// line internal/php8/scanner.rl:130
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.cs = 124
			lex.ungetCnt(1)
		}
		goto _again
	tr183:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:127
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st121
	st121:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof121
		}
	st_case_121:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:2359
		if lex.data[(lex.p)] == 35 {
			goto tr181
		}
		goto tr180
	tr181:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st122
	st122:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof122
		}
	st_case_122:
		// line internal/php8/scanner.go:2374
		if lex.data[(lex.p)] == 33 {
			goto st1
		}
		goto tr182
	tr3:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st1
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		// line internal/php8/scanner.go:2396
		switch lex.data[(lex.p)] {
		case 10:
			goto tr2
		case 13:
			goto tr3
		}
		goto st1
	tr2:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st123
	st123:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof123
		}
	st_case_123:
		// line internal/php8/scanner.go:2421
		goto tr183
	tr4:
		lex.cs = 124
		// line internal/php8/scanner.rl:143
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.te)
			lex.cs = 131
		}
		goto _again
	tr7:
		lex.cs = 124
		// line internal/php8/scanner.rl:147
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
			lex.cs = 131
		}
		goto _again
	tr188:
		// line internal/php8/scanner.rl:137
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(tkn)
			tok = token.T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 124
				goto _out
			}
		}
		goto st124
	tr190:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:137
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("<")
			lex.setTokenPosition(tkn)
			tok = token.T_INLINE_HTML
			{
				(lex.p)++
				lex.cs = 124
				goto _out
			}
		}
		goto st124
	tr196:
		lex.cs = 124
		// line internal/php8/scanner.rl:143
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.te)
			lex.cs = 131
		}
		goto _again
	tr197:
		lex.cs = 124
		// line internal/php8/scanner.rl:152
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ECHO
			lex.cs = 131
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr199:
		lex.cs = 124
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:147
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(lex.te - lex.ts - 5)
			lex.addFreeFloatingToken(tkn, token.T_OPEN_TAG, lex.ts, lex.ts+5)
			lex.cs = 131
		}
		goto _again
	st124:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:2511
		switch lex.data[(lex.p)] {
		case 10:
			goto tr185
		case 13:
			goto tr186
		case 60:
			goto st128
		}
		goto st125
	tr186:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st125
	tr191:
		// line internal/php8/scanner.rl:54

		goto st125
	tr193:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st125
	st125:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof125
		}
	st_case_125:
		// line internal/php8/scanner.go:2556
		switch lex.data[(lex.p)] {
		case 10:
			goto tr185
		case 13:
			goto tr186
		case 60:
			goto st127
		}
		goto st125
	tr185:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st126
	tr192:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st126
	st126:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof126
		}
	st_case_126:
		// line internal/php8/scanner.go:2597
		switch lex.data[(lex.p)] {
		case 10:
			goto tr192
		case 13:
			goto tr193
		case 60:
			goto tr194
		}
		goto tr191
	tr194:
		// line internal/php8/scanner.rl:54

		goto st127
	st127:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof127
		}
	st_case_127:
		// line internal/php8/scanner.go:2616
		switch lex.data[(lex.p)] {
		case 10:
			goto tr185
		case 13:
			goto tr186
		case 60:
			goto st127
		case 63:
			goto tr188
		}
		goto st125
	st128:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof128
		}
	st_case_128:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr185
		case 13:
			goto tr186
		case 60:
			goto st127
		case 63:
			goto tr195
		}
		goto st125
	tr195:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st129
	st129:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof129
		}
	st_case_129:
		// line internal/php8/scanner.go:2654
		switch lex.data[(lex.p)] {
		case 61:
			goto tr197
		case 80:
			goto st2
		case 112:
			goto st2
		}
		goto tr196
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 72:
			goto st3
		case 104:
			goto st3
		}
		goto tr4
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 80:
			goto st4
		case 112:
			goto st4
		}
		goto tr4
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 9:
			goto tr7
		case 10:
			goto tr8
		case 13:
			goto tr9
		case 32:
			goto tr7
		}
		goto tr4
	tr8:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st130
	st130:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof130
		}
	st_case_130:
		// line internal/php8/scanner.go:2721
		goto tr199
	tr9:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st5
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		// line internal/php8/scanner.go:2740
		if lex.data[(lex.p)] == 10 {
			goto tr8
		}
		goto tr4
	tr10:
		// line internal/php8/scanner.rl:161
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st131
	tr12:
		lex.cs = 131
		// line NONE:1
		switch lex.act {
		case 10:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 11:
			{
				(lex.p) = (lex.te) - 1

				s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, 2, 0)

				if err == nil {
					lex.setTokenPosition(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}

				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 12:
			{
				(lex.p) = (lex.te) - 1

				base := 10
				if lex.data[lex.ts] == '0' {
					base = 8
				}

				s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, base, 0)

				if err == nil {
					lex.setTokenPosition(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}

				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 13:
			{
				(lex.p) = (lex.te) - 1

				s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
				_, err := strconv.ParseInt(s, 16, 0)

				if err == nil {
					lex.setTokenPosition(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						goto _out
					}
				}

				lex.setTokenPosition(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					goto _out
				}
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAME_RELATIVE
				{
					(lex.p)++
					goto _out
				}
			}
		case 15:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAME_QUALIFIED
				{
					(lex.p)++
					goto _out
				}
			}
		case 18:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ABSTRACT
				{
					(lex.p)++
					goto _out
				}
			}
		case 19:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ARRAY
				{
					(lex.p)++
					goto _out
				}
			}
		case 20:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_AS
				{
					(lex.p)++
					goto _out
				}
			}
		case 21:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_BREAK
				{
					(lex.p)++
					goto _out
				}
			}
		case 22:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CALLABLE
				{
					(lex.p)++
					goto _out
				}
			}
		case 23:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CASE
				{
					(lex.p)++
					goto _out
				}
			}
		case 24:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CATCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 25:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CLASS
				{
					(lex.p)++
					goto _out
				}
			}
		case 26:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CLONE
				{
					(lex.p)++
					goto _out
				}
			}
		case 27:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CONST
				{
					(lex.p)++
					goto _out
				}
			}
		case 28:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CONTINUE
				{
					(lex.p)++
					goto _out
				}
			}
		case 29:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 30:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DEFAULT
				{
					(lex.p)++
					goto _out
				}
			}
		case 31:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DO
				{
					(lex.p)++
					goto _out
				}
			}
		case 32:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ECHO
				{
					(lex.p)++
					goto _out
				}
			}
		case 33:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ELSE
				{
					(lex.p)++
					goto _out
				}
			}
		case 34:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ELSEIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 35:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EMPTY
				{
					(lex.p)++
					goto _out
				}
			}
		case 36:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDDECLARE
				{
					(lex.p)++
					goto _out
				}
			}
		case 37:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDFOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 38:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDFOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 39:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDIF
				{
					(lex.p)++
					goto _out
				}
			}
		case 40:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDSWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 41:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ENDWHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 42:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EVAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 43:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EXIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 44:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_EXTENDS
				{
					(lex.p)++
					goto _out
				}
			}
		case 45:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FINAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 46:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FINALLY
				{
					(lex.p)++
					goto _out
				}
			}
		case 47:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 48:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FOREACH
				{
					(lex.p)++
					goto _out
				}
			}
		case 49:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FUNCTION
				{
					(lex.p)++
					goto _out
				}
			}
		case 50:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FN
				{
					(lex.p)++
					goto _out
				}
			}
		case 51:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_GLOBAL
				{
					(lex.p)++
					goto _out
				}
			}
		case 52:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_GOTO
				{
					(lex.p)++
					goto _out
				}
			}
		case 53:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_IF
				{
					(lex.p)++
					goto _out
				}
			}
		case 54:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_ISSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 55:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_IMPLEMENTS
				{
					(lex.p)++
					goto _out
				}
			}
		case 56:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INSTANCEOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 57:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INSTEADOF
				{
					(lex.p)++
					goto _out
				}
			}
		case 58:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INTERFACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 59:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LIST
				{
					(lex.p)++
					goto _out
				}
			}
		case 60:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NAMESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		case 61:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PRIVATE
				{
					(lex.p)++
					goto _out
				}
			}
		case 62:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PUBLIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 63:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PRINT
				{
					(lex.p)++
					goto _out
				}
			}
		case 64:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_PROTECTED
				{
					(lex.p)++
					goto _out
				}
			}
		case 65:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_RETURN
				{
					(lex.p)++
					goto _out
				}
			}
		case 66:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_READONLY
				{
					(lex.p)++
					goto _out
				}
			}
		case 67:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_STATIC
				{
					(lex.p)++
					goto _out
				}
			}
		case 68:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_SWITCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 69:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_MATCH
				{
					(lex.p)++
					goto _out
				}
			}
		case 70:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_THROW
				{
					(lex.p)++
					goto _out
				}
			}
		case 71:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_TRAIT
				{
					(lex.p)++
					goto _out
				}
			}
		case 72:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_TRY
				{
					(lex.p)++
					goto _out
				}
			}
		case 73:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_UNSET
				{
					(lex.p)++
					goto _out
				}
			}
		case 74:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_USE
				{
					(lex.p)++
					goto _out
				}
			}
		case 75:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_VAR
				{
					(lex.p)++
					goto _out
				}
			}
		case 76:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_WHILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 78:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_YIELD
				{
					(lex.p)++
					goto _out
				}
			}
		case 79:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INCLUDE
				{
					(lex.p)++
					goto _out
				}
			}
		case 80:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_INCLUDE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 81:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_REQUIRE
				{
					(lex.p)++
					goto _out
				}
			}
		case 82:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_REQUIRE_ONCE
				{
					(lex.p)++
					goto _out
				}
			}
		case 83:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_CLASS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 84:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DIR
				{
					(lex.p)++
					goto _out
				}
			}
		case 85:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FILE
				{
					(lex.p)++
					goto _out
				}
			}
		case 86:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_FUNC_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 87:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LINE
				{
					(lex.p)++
					goto _out
				}
			}
		case 88:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NS_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 89:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_METHOD_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 90:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_TRAIT_C
				{
					(lex.p)++
					goto _out
				}
			}
		case 91:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_HALT_COMPILER
				lex.cs = 540
				{
					(lex.p)++
					goto _out
				}
			}
		case 92:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_NEW
				{
					(lex.p)++
					goto _out
				}
			}
		case 93:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LOGICAL_AND
				{
					(lex.p)++
					goto _out
				}
			}
		case 94:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LOGICAL_OR
				{
					(lex.p)++
					goto _out
				}
			}
		case 95:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_LOGICAL_XOR
				{
					(lex.p)++
					goto _out
				}
			}
		case 124:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_SL
				{
					(lex.p)++
					goto _out
				}
			}
		case 143:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_STRING
				{
					(lex.p)++
					goto _out
				}
			}
		case 149:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('"'))
				lex.cs = 514
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr15:
		// line internal/php8/scanner.rl:361
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_CONSTANT_ENCAPSED_STRING
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr19:
		// line internal/php8/scanner.rl:384
		(lex.p) = (lex.te) - 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st131
	tr23:
		// line internal/php8/scanner.rl:347
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr38:
		// line internal/php8/scanner.rl:318
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ARRAY_CAST
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr45:
		// line internal/php8/scanner.rl:324
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING_CAST
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr49:
		// line internal/php8/scanner.rl:319
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOL_CAST
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr57:
		// line internal/php8/scanner.rl:321
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOUBLE_CAST
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr64:
		// line internal/php8/scanner.rl:322
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INT_CAST
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr73:
		// line internal/php8/scanner.rl:323
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_CAST
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr77:
		// line internal/php8/scanner.rl:320
		lex.te = (lex.p) + 1
		{
			lex.error(fmt.Sprintf("The (real) cast has been removed, use (float) instead"))
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr86:
		// line internal/php8/scanner.rl:325
		lex.te = (lex.p) + 1
		{
			lex.error(fmt.Sprintf("The (unset) cast is no longer supported"))
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr87:
		// line internal/php8/scanner.rl:286
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ELLIPSIS
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr90:
		// line internal/php8/scanner.rl:165
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr100:
		// line internal/php8/scanner.rl:334
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
				isDocComment = true
			}

			if isDocComment {
				lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
			} else {
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st131
	tr101:
		// line internal/php8/scanner.rl:176
		(lex.p) = (lex.te) - 1
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 131
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr114:
		lex.cs = 131
		// line internal/php8/scanner.rl:163
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 124
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr130:
		lex.cs = 131
		// line internal/php8/scanner.rl:359
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			lex.cs = 490
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr131:
		lex.cs = 131
		// line internal/php8/scanner.rl:162
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 124
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr134:
		// line internal/php8/scanner.rl:356
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr137:
		// line internal/php8/scanner.rl:267
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_YIELD
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr148:
		// line internal/php8/scanner.rl:266
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_YIELD_FROM
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr149:
		// line internal/php8/scanner.rl:204
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_FULLY_QUALIFIED
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr200:
		// line internal/php8/scanner.rl:384
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st131
	tr211:
		// line internal/php8/scanner.rl:347
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr249:
		lex.cs = 131
		// line internal/php8/scanner.rl:381
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('`'))
			lex.cs = 508
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr250:
		// line internal/php8/scanner.rl:353
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('{'))
			lex.call(131, 131)
			goto _out
		}
		goto st131
	tr252:
		// line internal/php8/scanner.rl:354
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('}'))
			lex.ret(1)
			goto _out
		}
		goto st131
	tr253:
		// line internal/php8/scanner.rl:161
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st131
	tr255:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:161
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st131
	tr259:
		// line internal/php8/scanner.rl:384
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st131
	tr260:
		// line internal/php8/scanner.rl:347
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr262:
		// line internal/php8/scanner.rl:304
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr263:
		// line internal/php8/scanner.rl:305
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_IDENTICAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr264:
		lex.cs = 131
		// line internal/php8/scanner.rl:382
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('"'))
			lex.cs = 514
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr265:
		// line internal/php8/scanner.rl:331
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st131
	tr267:
		// line internal/php8/scanner.rl:285
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ATTRIBUTE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr268:
		// line internal/php8/scanner.rl:327
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st131
	tr271:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:327
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetStr("?>")
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st131
	tr276:
		// line internal/php8/scanner.rl:355
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr277:
		// line internal/php8/scanner.rl:299
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MOD_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr278:
		// line internal/php8/scanner.rl:288
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOLEAN_AND
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr279:
		// line internal/php8/scanner.rl:290
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_AND_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr281:
		// line internal/php8/scanner.rl:293
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MUL_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr282:
		// line internal/php8/scanner.rl:312
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_POW
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr283:
		// line internal/php8/scanner.rl:294
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_POW_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr284:
		// line internal/php8/scanner.rl:301
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INC
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr285:
		// line internal/php8/scanner.rl:296
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PLUS_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr286:
		// line internal/php8/scanner.rl:300
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DEC
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr287:
		// line internal/php8/scanner.rl:297
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_MINUS_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr288:
		lex.cs = 131
		// line internal/php8/scanner.rl:358
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			lex.cs = 490
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr290:
		// line internal/php8/scanner.rl:292
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_CONCAT_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr291:
		// line internal/php8/scanner.rl:165
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr294:
		// line internal/php8/scanner.rl:295
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DIV_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr295:
		// line internal/php8/scanner.rl:176
		lex.te = (lex.p)
		(lex.p)--
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.Replace(string(lex.data[lex.ts:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 131
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr300:
		// line internal/php8/scanner.rl:166
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, 2, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 131
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr301:
		// line internal/php8/scanner.rl:191
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.Replace(string(lex.data[lex.ts+2:lex.te]), "_", "", -1)
			_, err := strconv.ParseInt(s, 16, 0)

			if err == nil {
				lex.setTokenPosition(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 131
					goto _out
				}
			}

			lex.setTokenPosition(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr302:
		// line internal/php8/scanner.rl:287
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_PAAMAYIM_NEKUDOTAYIM
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr303:
		lex.cs = 131
		// line internal/php8/scanner.rl:163
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 124
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr305:
		lex.cs = 131
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:163
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 124
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr308:
		// line internal/php8/scanner.rl:304
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr309:
		// line internal/php8/scanner.rl:313
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr310:
		// line internal/php8/scanner.rl:308
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SL_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr311:
		lex.cs = 131
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:367
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.heredocLabel = lex.data[lblStart:lblEnd]
			lex.setTokenPosition(tkn)
			tok = token.T_START_HEREDOC

			if lex.isHeredocEnd(lex.p + 1) {
				lex.cs = 520
			} else if lex.data[lblStart-1] == '\'' {
				lex.cs = 497
			} else {
				lex.cs = 501
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr312:
		// line internal/php8/scanner.rl:311
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_SMALLER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr313:
		// line internal/php8/scanner.rl:303
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SPACESHIP
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr315:
		// line internal/php8/scanner.rl:302
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOUBLE_ARROW
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr316:
		// line internal/php8/scanner.rl:306
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr317:
		// line internal/php8/scanner.rl:307
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_IDENTICAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr318:
		// line internal/php8/scanner.rl:310
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_IS_GREATER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr320:
		// line internal/php8/scanner.rl:314
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SR
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr321:
		// line internal/php8/scanner.rl:309
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_SR_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr325:
		lex.cs = 131
		// line internal/php8/scanner.rl:162
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 124
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr327:
		lex.cs = 131
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:162
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 124
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr328:
		// line internal/php8/scanner.rl:315
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_COALESCE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr329:
		// line internal/php8/scanner.rl:316
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_COALESCE_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr330:
		// line internal/php8/scanner.rl:356
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr336:
		// line internal/php8/scanner.rl:203
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_QUALIFIED
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr412:
		// line internal/php8/scanner.rl:222
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ELSE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr432:
		// line internal/php8/scanner.rl:226
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENDFOR
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr461:
		// line internal/php8/scanner.rl:234
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_FINAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr465:
		// line internal/php8/scanner.rl:236
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_FOR
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr497:
		// line internal/php8/scanner.rl:268
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_INCLUDE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr540:
		// line internal/php8/scanner.rl:249
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAMESPACE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr542:
		// line internal/php8/scanner.rl:202
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_RELATIVE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr578:
		// line internal/php8/scanner.rl:270
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_REQUIRE
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr624:
		// line internal/php8/scanner.rl:267
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_YIELD
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr625:
		// line internal/php8/scanner.rl:205
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NS_SEPARATOR
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr626:
		// line internal/php8/scanner.rl:204
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NAME_FULLY_QUALIFIED
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr628:
		// line internal/php8/scanner.rl:298
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_XOR_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr702:
		// line internal/php8/scanner.rl:291
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	tr703:
		// line internal/php8/scanner.rl:289
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_BOOLEAN_OR
			{
				(lex.p)++
				lex.cs = 131
				goto _out
			}
		}
		goto st131
	st131:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof131
		}
	st_case_131:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:3675
		switch lex.data[(lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr202
		case 32:
			goto tr201
		case 33:
			goto st135
		case 34:
			goto tr204
		case 35:
			goto st138
		case 36:
			goto st141
		case 37:
			goto st143
		case 38:
			goto st144
		case 39:
			goto tr209
		case 40:
			goto tr210
		case 42:
			goto st147
		case 43:
			goto st149
		case 45:
			goto st150
		case 46:
			goto tr215
		case 47:
			goto tr216
		case 48:
			goto tr217
		case 58:
			goto st160
		case 59:
			goto tr219
		case 60:
			goto st164
		case 61:
			goto st168
		case 62:
			goto st170
		case 63:
			goto tr223
		case 64:
			goto tr211
		case 65:
			goto tr224
		case 66:
			goto tr225
		case 67:
			goto tr226
		case 68:
			goto tr227
		case 69:
			goto tr228
		case 70:
			goto tr229
		case 71:
			goto tr230
		case 73:
			goto tr232
		case 76:
			goto tr233
		case 77:
			goto tr234
		case 78:
			goto tr235
		case 79:
			goto tr236
		case 80:
			goto tr237
		case 82:
			goto tr238
		case 83:
			goto tr239
		case 84:
			goto tr240
		case 85:
			goto tr241
		case 86:
			goto tr242
		case 87:
			goto tr243
		case 88:
			goto tr244
		case 89:
			goto tr245
		case 92:
			goto st421
		case 94:
			goto st423
		case 95:
			goto tr248
		case 96:
			goto tr249
		case 97:
			goto tr224
		case 98:
			goto tr225
		case 99:
			goto tr226
		case 100:
			goto tr227
		case 101:
			goto tr228
		case 102:
			goto tr229
		case 103:
			goto tr230
		case 105:
			goto tr232
		case 108:
			goto tr233
		case 109:
			goto tr234
		case 110:
			goto tr235
		case 111:
			goto tr236
		case 112:
			goto tr237
		case 114:
			goto tr238
		case 115:
			goto tr239
		case 116:
			goto tr240
		case 117:
			goto tr241
		case 118:
			goto tr242
		case 119:
			goto tr243
		case 120:
			goto tr244
		case 121:
			goto tr245
		case 123:
			goto tr250
		case 124:
			goto st489
		case 125:
			goto tr252
		case 126:
			goto tr211
		case 127:
			goto tr200
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr201
				}
			default:
				goto tr200
			}
		case lex.data[(lex.p)] > 31:
			switch {
			case lex.data[(lex.p)] < 49:
				if 41 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 44 {
					goto tr211
				}
			case lex.data[(lex.p)] > 57:
				if 91 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 93 {
					goto tr211
				}
			default:
				goto tr102
			}
		default:
			goto tr200
		}
		goto tr231
	tr201:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st132
	tr256:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		goto st132
	st132:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		// line internal/php8/scanner.go:3872
		switch lex.data[(lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr254
		case 32:
			goto tr201
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr201
		}
		goto tr253
	tr11:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st133
	tr257:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st133
	st133:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof133
		}
	st_case_133:
		// line internal/php8/scanner.go:3922
		switch lex.data[(lex.p)] {
		case 10:
			goto tr257
		case 13:
			goto tr258
		case 32:
			goto tr256
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr256
		}
		goto tr255
	tr254:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st6
	tr258:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		// line internal/php8/scanner.go:3966
		if lex.data[(lex.p)] == 10 {
			goto tr11
		}
		goto tr10
	tr202:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st134
	st134:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof134
		}
	st_case_134:
		// line internal/php8/scanner.go:3988
		if lex.data[(lex.p)] == 10 {
			goto tr11
		}
		goto tr259
	st135:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof135
		}
	st_case_135:
		if lex.data[(lex.p)] == 61 {
			goto st136
		}
		goto tr260
	st136:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		if lex.data[(lex.p)] == 61 {
			goto tr263
		}
		goto tr262
	tr204:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:382
		lex.act = 149
		goto st137
	st137:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof137
		}
	st_case_137:
		// line internal/php8/scanner.go:4023
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
	tr14:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st7
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		// line internal/php8/scanner.go:4056
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 123:
			goto st10
		}
		goto st7
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto st8
		case 92:
			goto st9
		case 96:
			goto st7
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st7
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st7
			}
		default:
			goto st7
		}
		goto tr12
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		}
		goto st7
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		case 34:
			goto tr15
		case 36:
			goto tr12
		case 92:
			goto st9
		}
		goto st7
	st138:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof138
		}
	st_case_138:
		if lex.data[(lex.p)] == 91 {
			goto tr267
		}
		goto st139
	tr270:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st139
	tr272:
		// line internal/php8/scanner.rl:54

		goto st139
	tr274:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st139
	st139:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof139
		}
	st_case_139:
		// line internal/php8/scanner.go:4178
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr269
		case 525:
			goto tr270
		}
		if 512 <= _widec && _widec <= 767 {
			goto st139
		}
		goto tr268
	tr269:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st140
	tr273:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st140
	st140:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof140
		}
	st_case_140:
		// line internal/php8/scanner.go:4258
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotPhpCloseToken() && lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr273
		case 525:
			goto tr274
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr272
		}
		goto tr271
	st141:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof141
		}
	st_case_141:
		if lex.data[(lex.p)] == 96 {
			goto tr260
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr260
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr260
			}
		default:
			goto tr260
		}
		goto st142
	st142:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof142
		}
	st_case_142:
		if lex.data[(lex.p)] == 96 {
			goto tr276
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr276
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr276
				}
			case lex.data[(lex.p)] >= 91:
				goto tr276
			}
		default:
			goto tr276
		}
		goto st142
	st143:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof143
		}
	st_case_143:
		if lex.data[(lex.p)] == 61 {
			goto tr277
		}
		goto tr260
	st144:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof144
		}
	st_case_144:
		switch lex.data[(lex.p)] {
		case 38:
			goto tr278
		case 61:
			goto tr279
		}
		goto tr260
	tr209:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st145
	st145:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof145
		}
	st_case_145:
		// line internal/php8/scanner.go:4385
		switch lex.data[(lex.p)] {
		case 10:
			goto tr21
		case 13:
			goto tr21
		case 39:
			goto tr15
		case 92:
			goto st12
		}
		goto st11
	tr21:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st11
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		// line internal/php8/scanner.go:4414
		switch lex.data[(lex.p)] {
		case 10:
			goto tr21
		case 13:
			goto tr21
		case 39:
			goto tr15
		case 92:
			goto st12
		}
		goto st11
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr21
		case 13:
			goto tr21
		}
		goto st11
	tr210:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st146
	st146:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof146
		}
	st_case_146:
		// line internal/php8/scanner.go:4448
		switch lex.data[(lex.p)] {
		case 9:
			goto st13
		case 32:
			goto st13
		case 65:
			goto st14
		case 66:
			goto st19
		case 68:
			goto st31
		case 70:
			goto st37
		case 73:
			goto st41
		case 79:
			goto st48
		case 82:
			goto st54
		case 83:
			goto st58
		case 85:
			goto st63
		case 97:
			goto st14
		case 98:
			goto st19
		case 100:
			goto st31
		case 102:
			goto st37
		case 105:
			goto st41
		case 111:
			goto st48
		case 114:
			goto st54
		case 115:
			goto st58
		case 117:
			goto st63
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st13
		}
		goto tr260
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch lex.data[(lex.p)] {
		case 9:
			goto st13
		case 32:
			goto st13
		case 65:
			goto st14
		case 66:
			goto st19
		case 68:
			goto st31
		case 70:
			goto st37
		case 73:
			goto st41
		case 79:
			goto st48
		case 82:
			goto st54
		case 83:
			goto st58
		case 85:
			goto st63
		case 97:
			goto st14
		case 98:
			goto st19
		case 100:
			goto st31
		case 102:
			goto st37
		case 105:
			goto st41
		case 111:
			goto st48
		case 114:
			goto st54
		case 115:
			goto st58
		case 117:
			goto st63
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st13
		}
		goto tr23
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch lex.data[(lex.p)] {
		case 82:
			goto st15
		case 114:
			goto st15
		}
		goto tr23
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch lex.data[(lex.p)] {
		case 82:
			goto st16
		case 114:
			goto st16
		}
		goto tr23
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch lex.data[(lex.p)] {
		case 65:
			goto st17
		case 97:
			goto st17
		}
		goto tr23
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch lex.data[(lex.p)] {
		case 89:
			goto st18
		case 121:
			goto st18
		}
		goto tr23
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch lex.data[(lex.p)] {
		case 9:
			goto st18
		case 32:
			goto st18
		case 41:
			goto tr38
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st18
		}
		goto tr23
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch lex.data[(lex.p)] {
		case 73:
			goto st20
		case 79:
			goto st25
		case 105:
			goto st20
		case 111:
			goto st25
		}
		goto tr23
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch lex.data[(lex.p)] {
		case 78:
			goto st21
		case 110:
			goto st21
		}
		goto tr23
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch lex.data[(lex.p)] {
		case 65:
			goto st22
		case 97:
			goto st22
		}
		goto tr23
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch lex.data[(lex.p)] {
		case 82:
			goto st23
		case 114:
			goto st23
		}
		goto tr23
	st23:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch lex.data[(lex.p)] {
		case 89:
			goto st24
		case 121:
			goto st24
		}
		goto tr23
	st24:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch lex.data[(lex.p)] {
		case 9:
			goto st24
		case 32:
			goto st24
		case 41:
			goto tr45
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st24
		}
		goto tr23
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch lex.data[(lex.p)] {
		case 79:
			goto st26
		case 111:
			goto st26
		}
		goto tr23
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch lex.data[(lex.p)] {
		case 76:
			goto st27
		case 108:
			goto st27
		}
		goto tr23
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch lex.data[(lex.p)] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 41:
			goto tr49
		case 69:
			goto st29
		case 101:
			goto st29
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st28
		}
		goto tr23
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch lex.data[(lex.p)] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 41:
			goto tr49
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st28
		}
		goto tr23
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch lex.data[(lex.p)] {
		case 65:
			goto st30
		case 97:
			goto st30
		}
		goto tr23
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch lex.data[(lex.p)] {
		case 78:
			goto st28
		case 110:
			goto st28
		}
		goto tr23
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch lex.data[(lex.p)] {
		case 79:
			goto st32
		case 111:
			goto st32
		}
		goto tr23
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch lex.data[(lex.p)] {
		case 85:
			goto st33
		case 117:
			goto st33
		}
		goto tr23
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch lex.data[(lex.p)] {
		case 66:
			goto st34
		case 98:
			goto st34
		}
		goto tr23
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		switch lex.data[(lex.p)] {
		case 76:
			goto st35
		case 108:
			goto st35
		}
		goto tr23
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		switch lex.data[(lex.p)] {
		case 69:
			goto st36
		case 101:
			goto st36
		}
		goto tr23
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch lex.data[(lex.p)] {
		case 9:
			goto st36
		case 32:
			goto st36
		case 41:
			goto tr57
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st36
		}
		goto tr23
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch lex.data[(lex.p)] {
		case 76:
			goto st38
		case 108:
			goto st38
		}
		goto tr23
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		switch lex.data[(lex.p)] {
		case 79:
			goto st39
		case 111:
			goto st39
		}
		goto tr23
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch lex.data[(lex.p)] {
		case 65:
			goto st40
		case 97:
			goto st40
		}
		goto tr23
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		switch lex.data[(lex.p)] {
		case 84:
			goto st36
		case 116:
			goto st36
		}
		goto tr23
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		switch lex.data[(lex.p)] {
		case 78:
			goto st42
		case 110:
			goto st42
		}
		goto tr23
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		switch lex.data[(lex.p)] {
		case 84:
			goto st43
		case 116:
			goto st43
		}
		goto tr23
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch lex.data[(lex.p)] {
		case 9:
			goto st44
		case 32:
			goto st44
		case 41:
			goto tr64
		case 69:
			goto st45
		case 101:
			goto st45
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st44
		}
		goto tr23
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		switch lex.data[(lex.p)] {
		case 9:
			goto st44
		case 32:
			goto st44
		case 41:
			goto tr64
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st44
		}
		goto tr23
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 71:
			goto st46
		case 103:
			goto st46
		}
		goto tr23
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch lex.data[(lex.p)] {
		case 69:
			goto st47
		case 101:
			goto st47
		}
		goto tr23
	st47:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch lex.data[(lex.p)] {
		case 82:
			goto st44
		case 114:
			goto st44
		}
		goto tr23
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch lex.data[(lex.p)] {
		case 66:
			goto st49
		case 98:
			goto st49
		}
		goto tr23
	st49:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch lex.data[(lex.p)] {
		case 74:
			goto st50
		case 106:
			goto st50
		}
		goto tr23
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch lex.data[(lex.p)] {
		case 69:
			goto st51
		case 101:
			goto st51
		}
		goto tr23
	st51:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch lex.data[(lex.p)] {
		case 67:
			goto st52
		case 99:
			goto st52
		}
		goto tr23
	st52:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch lex.data[(lex.p)] {
		case 84:
			goto st53
		case 116:
			goto st53
		}
		goto tr23
	st53:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch lex.data[(lex.p)] {
		case 9:
			goto st53
		case 32:
			goto st53
		case 41:
			goto tr73
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st53
		}
		goto tr23
	st54:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch lex.data[(lex.p)] {
		case 69:
			goto st55
		case 101:
			goto st55
		}
		goto tr23
	st55:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch lex.data[(lex.p)] {
		case 65:
			goto st56
		case 97:
			goto st56
		}
		goto tr23
	st56:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch lex.data[(lex.p)] {
		case 76:
			goto st57
		case 108:
			goto st57
		}
		goto tr23
	st57:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch lex.data[(lex.p)] {
		case 9:
			goto st57
		case 32:
			goto st57
		case 41:
			goto tr77
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st57
		}
		goto tr23
	st58:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch lex.data[(lex.p)] {
		case 84:
			goto st59
		case 116:
			goto st59
		}
		goto tr23
	st59:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch lex.data[(lex.p)] {
		case 82:
			goto st60
		case 114:
			goto st60
		}
		goto tr23
	st60:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch lex.data[(lex.p)] {
		case 73:
			goto st61
		case 105:
			goto st61
		}
		goto tr23
	st61:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch lex.data[(lex.p)] {
		case 78:
			goto st62
		case 110:
			goto st62
		}
		goto tr23
	st62:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch lex.data[(lex.p)] {
		case 71:
			goto st24
		case 103:
			goto st24
		}
		goto tr23
	st63:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch lex.data[(lex.p)] {
		case 78:
			goto st64
		case 110:
			goto st64
		}
		goto tr23
	st64:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch lex.data[(lex.p)] {
		case 83:
			goto st65
		case 115:
			goto st65
		}
		goto tr23
	st65:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch lex.data[(lex.p)] {
		case 69:
			goto st66
		case 101:
			goto st66
		}
		goto tr23
	st66:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch lex.data[(lex.p)] {
		case 84:
			goto st67
		case 116:
			goto st67
		}
		goto tr23
	st67:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch lex.data[(lex.p)] {
		case 9:
			goto st67
		case 32:
			goto st67
		case 41:
			goto tr86
		}
		if 11 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st67
		}
		goto tr23
	st147:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof147
		}
	st_case_147:
		switch lex.data[(lex.p)] {
		case 42:
			goto st148
		case 61:
			goto tr281
		}
		goto tr260
	st148:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof148
		}
	st_case_148:
		if lex.data[(lex.p)] == 61 {
			goto tr283
		}
		goto tr282
	st149:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof149
		}
	st_case_149:
		switch lex.data[(lex.p)] {
		case 43:
			goto tr284
		case 61:
			goto tr285
		}
		goto tr260
	st150:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof150
		}
	st_case_150:
		switch lex.data[(lex.p)] {
		case 45:
			goto tr286
		case 61:
			goto tr287
		case 62:
			goto tr288
		}
		goto tr260
	tr215:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st151
	st151:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof151
		}
	st_case_151:
		// line internal/php8/scanner.go:5313
		switch lex.data[(lex.p)] {
		case 46:
			goto st68
		case 61:
			goto tr290
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr91
		}
		goto tr260
	st68:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		if lex.data[(lex.p)] == 46 {
			goto tr87
		}
		goto tr23
	tr91:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:165
		lex.act = 10
		goto st152
	st152:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof152
		}
	st_case_152:
		// line internal/php8/scanner.go:5345
		switch lex.data[(lex.p)] {
		case 69:
			goto st69
		case 95:
			goto st71
		case 101:
			goto st69
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr91
		}
		goto tr291
	st69:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch lex.data[(lex.p)] {
		case 43:
			goto st70
		case 45:
			goto st70
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr89
		}
		goto tr12
	st70:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr89
		}
		goto tr12
	tr89:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:165
		lex.act = 10
		goto st153
	st153:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof153
		}
	st_case_153:
		// line internal/php8/scanner.go:5394
		if lex.data[(lex.p)] == 95 {
			goto st70
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr89
		}
		goto tr291
	st71:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr91
		}
		goto tr90
	tr216:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st154
	st154:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof154
		}
	st_case_154:
		// line internal/php8/scanner.go:5421
		switch lex.data[(lex.p)] {
		case 42:
			goto st72
		case 47:
			goto st139
		case 61:
			goto tr294
		}
		goto tr260
	tr94:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st72
	tr96:
		// line internal/php8/scanner.rl:54

		goto st72
	tr98:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st72
	st72:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		// line internal/php8/scanner.go:5466
		switch lex.data[(lex.p)] {
		case 10:
			goto tr93
		case 13:
			goto tr94
		case 42:
			goto st74
		}
		goto st72
	tr93:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st73
	tr97:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st73
	st73:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		// line internal/php8/scanner.go:5507
		switch lex.data[(lex.p)] {
		case 10:
			goto tr97
		case 13:
			goto tr98
		case 42:
			goto tr99
		}
		goto tr96
	tr99:
		// line internal/php8/scanner.rl:54

		goto st74
	st74:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		// line internal/php8/scanner.go:5526
		switch lex.data[(lex.p)] {
		case 10:
			goto tr93
		case 13:
			goto tr94
		case 42:
			goto st74
		case 47:
			goto tr100
		}
		goto st72
	tr217:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:176
		lex.act = 12
		goto st155
	st155:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof155
		}
	st_case_155:
		// line internal/php8/scanner.go:5550
		switch lex.data[(lex.p)] {
		case 46:
			goto tr296
		case 69:
			goto st69
		case 95:
			goto st75
		case 98:
			goto st76
		case 101:
			goto st69
		case 120:
			goto st77
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr102
		}
		goto tr295
	tr296:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:165
		lex.act = 10
		goto st156
	st156:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof156
		}
	st_case_156:
		// line internal/php8/scanner.go:5581
		switch lex.data[(lex.p)] {
		case 69:
			goto st69
		case 101:
			goto st69
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr91
		}
		goto tr291
	tr102:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:176
		lex.act = 12
		goto st157
	st157:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof157
		}
	st_case_157:
		// line internal/php8/scanner.go:5604
		switch lex.data[(lex.p)] {
		case 46:
			goto tr296
		case 69:
			goto st69
		case 95:
			goto st75
		case 101:
			goto st69
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr102
		}
		goto tr295
	st75:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr102
		}
		goto tr101
	st76:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr103
		}
		goto tr12
	tr103:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:166
		lex.act = 11
		goto st158
	st158:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof158
		}
	st_case_158:
		// line internal/php8/scanner.go:5649
		if lex.data[(lex.p)] == 95 {
			goto st76
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr103
		}
		goto tr300
	st77:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr104
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr104
			}
		default:
			goto tr104
		}
		goto tr12
	tr104:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:191
		lex.act = 13
		goto st159
	st159:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof159
		}
	st_case_159:
		// line internal/php8/scanner.go:5687
		if lex.data[(lex.p)] == 95 {
			goto st77
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr104
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr104
			}
		default:
			goto tr104
		}
		goto tr301
	st160:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof160
		}
	st_case_160:
		if lex.data[(lex.p)] == 58 {
			goto tr302
		}
		goto tr260
	tr219:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st161
	st161:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof161
		}
	st_case_161:
		// line internal/php8/scanner.go:5723
		switch lex.data[(lex.p)] {
		case 10:
			goto tr106
		case 13:
			goto tr107
		case 32:
			goto st78
		case 63:
			goto st81
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st78
		}
		goto tr260
	tr109:
		// line internal/php8/scanner.rl:54

		goto st78
	st78:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		// line internal/php8/scanner.go:5747
		switch lex.data[(lex.p)] {
		case 10:
			goto tr106
		case 13:
			goto tr107
		case 32:
			goto st78
		case 63:
			goto st81
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st78
		}
		goto tr23
	tr106:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st79
	tr110:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st79
	st79:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		// line internal/php8/scanner.go:5793
		switch lex.data[(lex.p)] {
		case 10:
			goto tr110
		case 13:
			goto tr111
		case 32:
			goto tr109
		case 63:
			goto tr112
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr109
		}
		goto tr23
	tr107:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st80
	tr111:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st80
	st80:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		// line internal/php8/scanner.go:5839
		if lex.data[(lex.p)] == 10 {
			goto tr106
		}
		goto tr23
	tr112:
		// line internal/php8/scanner.rl:54

		goto st81
	st81:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		// line internal/php8/scanner.go:5853
		if lex.data[(lex.p)] == 62 {
			goto tr113
		}
		goto tr23
	tr113:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st162
	st162:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof162
		}
	st_case_162:
		// line internal/php8/scanner.go:5868
		switch lex.data[(lex.p)] {
		case 10:
			goto tr115
		case 13:
			goto tr304
		}
		goto tr303
	tr115:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st163
	st163:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof163
		}
	st_case_163:
		// line internal/php8/scanner.go:5893
		goto tr305
	tr304:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st82
	st82:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		// line internal/php8/scanner.go:5912
		if lex.data[(lex.p)] == 10 {
			goto tr115
		}
		goto tr114
	st164:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof164
		}
	st_case_164:
		switch lex.data[(lex.p)] {
		case 60:
			goto tr306
		case 61:
			goto st167
		case 62:
			goto tr308
		}
		goto tr260
	tr306:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:313
		lex.act = 124
		goto st165
	st165:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof165
		}
	st_case_165:
		// line internal/php8/scanner.go:5943
		switch lex.data[(lex.p)] {
		case 60:
			goto st83
		case 61:
			goto tr310
		}
		goto tr309
	st83:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch lex.data[(lex.p)] {
		case 9:
			goto st83
		case 32:
			goto st83
		case 34:
			goto st84
		case 39:
			goto st88
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr119
	st84:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr120
	tr120:
		// line internal/php8/scanner.rl:35
		lblStart = lex.p
		goto st85
	st85:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		// line internal/php8/scanner.go:6011
		switch lex.data[(lex.p)] {
		case 34:
			goto tr121
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr12
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr12
				}
			case lex.data[(lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st85
	tr121:
		// line internal/php8/scanner.rl:36
		lblEnd = lex.p
		goto st86
	st86:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		// line internal/php8/scanner.go:6045
		switch lex.data[(lex.p)] {
		case 10:
			goto tr123
		case 13:
			goto tr124
		}
		goto tr12
	tr123:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st166
	tr127:
		// line internal/php8/scanner.rl:36
		lblEnd = lex.p
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st166
	st166:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof166
		}
	st_case_166:
		// line internal/php8/scanner.go:6084
		goto tr311
	tr124:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st87
	tr128:
		// line internal/php8/scanner.rl:36
		lblEnd = lex.p
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st87
	st87:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		// line internal/php8/scanner.go:6117
		if lex.data[(lex.p)] == 10 {
			goto tr123
		}
		goto tr12
	st88:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr125
	tr125:
		// line internal/php8/scanner.rl:35
		lblStart = lex.p
		goto st89
	st89:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		// line internal/php8/scanner.go:6152
		switch lex.data[(lex.p)] {
		case 39:
			goto tr121
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr12
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr12
				}
			case lex.data[(lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st89
	tr119:
		// line internal/php8/scanner.rl:35
		lblStart = lex.p
		goto st90
	st90:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		// line internal/php8/scanner.go:6186
		switch lex.data[(lex.p)] {
		case 10:
			goto tr127
		case 13:
			goto tr128
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr12
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr12
				}
			case lex.data[(lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st90
	st167:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof167
		}
	st_case_167:
		if lex.data[(lex.p)] == 62 {
			goto tr313
		}
		goto tr312
	st168:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof168
		}
	st_case_168:
		switch lex.data[(lex.p)] {
		case 61:
			goto st169
		case 62:
			goto tr315
		}
		goto tr260
	st169:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof169
		}
	st_case_169:
		if lex.data[(lex.p)] == 61 {
			goto tr317
		}
		goto tr316
	st170:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof170
		}
	st_case_170:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr318
		case 62:
			goto st171
		}
		goto tr260
	st171:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof171
		}
	st_case_171:
		if lex.data[(lex.p)] == 61 {
			goto tr321
		}
		goto tr320
	tr223:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st172
	st172:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof172
		}
	st_case_172:
		// line internal/php8/scanner.go:6274
		switch lex.data[(lex.p)] {
		case 45:
			goto st91
		case 62:
			goto tr323
		case 63:
			goto st175
		}
		goto tr260
	st91:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		if lex.data[(lex.p)] == 62 {
			goto tr130
		}
		goto tr23
	tr323:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st173
	st173:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof173
		}
	st_case_173:
		// line internal/php8/scanner.go:6303
		switch lex.data[(lex.p)] {
		case 10:
			goto tr132
		case 13:
			goto tr326
		}
		goto tr325
	tr132:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st174
	st174:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof174
		}
	st_case_174:
		// line internal/php8/scanner.go:6328
		goto tr327
	tr326:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st92
	st92:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		// line internal/php8/scanner.go:6347
		if lex.data[(lex.p)] == 10 {
			goto tr132
		}
		goto tr131
	st175:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof175
		}
	st_case_175:
		if lex.data[(lex.p)] == 61 {
			goto tr329
		}
		goto tr328
	tr224:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st176
	st176:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof176
		}
	st_case_176:
		// line internal/php8/scanner.go:6373
		switch lex.data[(lex.p)] {
		case 66:
			goto tr331
		case 78:
			goto tr332
		case 82:
			goto tr333
		case 83:
			goto tr334
		case 92:
			goto st93
		case 96:
			goto tr330
		case 98:
			goto tr331
		case 110:
			goto tr332
		case 114:
			goto tr333
		case 115:
			goto tr334
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr231:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st177
	tr334:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:209
		lex.act = 20
		goto st177
	tr342:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:207
		lex.act = 18
		goto st177
	tr343:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:282
		lex.act = 93
		goto st177
	tr346:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:208
		lex.act = 19
		goto st177
	tr351:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:210
		lex.act = 21
		goto st177
	tr363:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:211
		lex.act = 22
		goto st177
	tr364:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:212
		lex.act = 23
		goto st177
	tr366:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:213
		lex.act = 24
		goto st177
	tr373:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:238
		lex.act = 49
		goto st177
	tr377:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:214
		lex.act = 25
		goto st177
	tr379:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:215
		lex.act = 26
		goto st177
	tr383:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:216
		lex.act = 27
		goto st177
	tr387:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:217
		lex.act = 28
		goto st177
	tr390:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:220
		lex.act = 31
		goto st177
	tr396:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:218
		lex.act = 29
		goto st177
	tr400:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:219
		lex.act = 30
		goto st177
	tr401:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:232
		lex.act = 43
		goto st177
	tr409:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:221
		lex.act = 32
		goto st177
	tr414:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:223
		lex.act = 34
		goto st177
	tr417:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:224
		lex.act = 35
		goto st177
	tr429:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:225
		lex.act = 36
		goto st177
	tr436:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:227
		lex.act = 38
		goto st177
	tr437:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:228
		lex.act = 39
		goto st177
	tr442:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:229
		lex.act = 40
		goto st177
	tr446:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:230
		lex.act = 41
		goto st177
	tr448:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:231
		lex.act = 42
		goto st177
	tr454:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:233
		lex.act = 44
		goto st177
	tr456:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:239
		lex.act = 50
		goto st177
	tr463:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:235
		lex.act = 46
		goto st177
	tr469:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:237
		lex.act = 48
		goto st177
	tr475:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:240
		lex.act = 51
		goto st177
	tr477:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:241
		lex.act = 52
		goto st177
	tr478:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:242
		lex.act = 53
		goto st177
	tr489:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:244
		lex.act = 55
		goto st177
	tr502:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:269
		lex.act = 80
		goto st177
	tr510:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:245
		lex.act = 56
		goto st177
	tr514:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:246
		lex.act = 57
		goto st177
	tr520:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:247
		lex.act = 58
		goto st177
	tr523:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:243
		lex.act = 54
		goto st177
	tr526:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:248
		lex.act = 59
		goto st177
	tr530:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:258
		lex.act = 69
		goto st177
	tr543:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:281
		lex.act = 92
		goto st177
	tr544:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:283
		lex.act = 94
		goto st177
	tr551:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:252
		lex.act = 63
		goto st177
	tr554:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:250
		lex.act = 61
		goto st177
	tr560:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:253
		lex.act = 64
		goto st177
	tr564:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:251
		lex.act = 62
		goto st177
	tr573:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:255
		lex.act = 66
		goto st177
	tr583:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:271
		lex.act = 82
		goto st177
	tr586:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:254
		lex.act = 65
		goto st177
	tr592:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:256
		lex.act = 67
		goto st177
	tr596:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:257
		lex.act = 68
		goto st177
	tr601:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:259
		lex.act = 70
		goto st177
	tr603:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:261
		lex.act = 72
		goto st177
	tr605:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:260
		lex.act = 71
		goto st177
	tr610:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:262
		lex.act = 73
		goto st177
	tr611:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:263
		lex.act = 74
		goto st177
	tr613:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:264
		lex.act = 75
		goto st177
	tr617:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:265
		lex.act = 76
		goto st177
	tr619:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:284
		lex.act = 95
		goto st177
	tr643:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:272
		lex.act = 83
		goto st177
	tr647:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:273
		lex.act = 84
		goto st177
	tr653:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:274
		lex.act = 85
		goto st177
	tr661:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:275
		lex.act = 86
		goto st177
	tr673:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:280
		lex.act = 91
		goto st177
	tr678:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:276
		lex.act = 87
		goto st177
	tr685:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:278
		lex.act = 89
		goto st177
	tr695:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:277
		lex.act = 88
		goto st177
	tr701:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:279
		lex.act = 90
		goto st177
	st177:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof177
		}
	st_case_177:
		// line internal/php8/scanner.go:6909
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 96:
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr12
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr12
				}
			case lex.data[(lex.p)] >= 91:
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr231
	st93:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr133
	tr133:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:203
		lex.act = 15
		goto st178
	st178:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof178
		}
	st_case_178:
		// line internal/php8/scanner.go:6967
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 96:
			goto tr336
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr336
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr336
				}
			case lex.data[(lex.p)] >= 91:
				goto tr336
			}
		default:
			goto tr336
		}
		goto tr133
	tr331:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st179
	st179:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof179
		}
	st_case_179:
		// line internal/php8/scanner.go:7004
		switch lex.data[(lex.p)] {
		case 83:
			goto tr337
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr337
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr337:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st180
	st180:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof180
		}
	st_case_180:
		// line internal/php8/scanner.go:7045
		switch lex.data[(lex.p)] {
		case 84:
			goto tr338
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr338
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr338:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st181
	st181:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof181
		}
	st_case_181:
		// line internal/php8/scanner.go:7086
		switch lex.data[(lex.p)] {
		case 82:
			goto tr339
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr339
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr339:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st182
	st182:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof182
		}
	st_case_182:
		// line internal/php8/scanner.go:7127
		switch lex.data[(lex.p)] {
		case 65:
			goto tr340
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr340
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr340:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st183
	st183:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof183
		}
	st_case_183:
		// line internal/php8/scanner.go:7168
		switch lex.data[(lex.p)] {
		case 67:
			goto tr341
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr341
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr341:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st184
	st184:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof184
		}
	st_case_184:
		// line internal/php8/scanner.go:7209
		switch lex.data[(lex.p)] {
		case 84:
			goto tr342
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr342
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr332:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st185
	st185:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof185
		}
	st_case_185:
		// line internal/php8/scanner.go:7250
		switch lex.data[(lex.p)] {
		case 68:
			goto tr343
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr343
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr333:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st186
	st186:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof186
		}
	st_case_186:
		// line internal/php8/scanner.go:7291
		switch lex.data[(lex.p)] {
		case 82:
			goto tr344
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr344
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr344:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st187
	st187:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof187
		}
	st_case_187:
		// line internal/php8/scanner.go:7332
		switch lex.data[(lex.p)] {
		case 65:
			goto tr345
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr345
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr345:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st188
	st188:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof188
		}
	st_case_188:
		// line internal/php8/scanner.go:7373
		switch lex.data[(lex.p)] {
		case 89:
			goto tr346
		case 92:
			goto st93
		case 96:
			goto tr330
		case 121:
			goto tr346
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr225:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st189
	st189:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof189
		}
	st_case_189:
		// line internal/php8/scanner.go:7414
		switch lex.data[(lex.p)] {
		case 34:
			goto st7
		case 60:
			goto st94
		case 82:
			goto tr348
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr348
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	st94:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		if lex.data[(lex.p)] == 60 {
			goto st95
		}
		goto tr134
	st95:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		if lex.data[(lex.p)] == 60 {
			goto st83
		}
		goto tr134
	tr348:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st190
	st190:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof190
		}
	st_case_190:
		// line internal/php8/scanner.go:7477
		switch lex.data[(lex.p)] {
		case 69:
			goto tr349
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr349
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr349:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st191
	st191:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof191
		}
	st_case_191:
		// line internal/php8/scanner.go:7518
		switch lex.data[(lex.p)] {
		case 65:
			goto tr350
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr350
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr350:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st192
	st192:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof192
		}
	st_case_192:
		// line internal/php8/scanner.go:7559
		switch lex.data[(lex.p)] {
		case 75:
			goto tr351
		case 92:
			goto st93
		case 96:
			goto tr330
		case 107:
			goto tr351
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr226:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st193
	st193:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof193
		}
	st_case_193:
		// line internal/php8/scanner.go:7600
		switch lex.data[(lex.p)] {
		case 65:
			goto tr352
		case 70:
			goto tr353
		case 76:
			goto tr354
		case 79:
			goto tr355
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr352
		case 102:
			goto tr353
		case 108:
			goto tr354
		case 111:
			goto tr355
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr352:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st194
	st194:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof194
		}
	st_case_194:
		// line internal/php8/scanner.go:7653
		switch lex.data[(lex.p)] {
		case 76:
			goto tr356
		case 83:
			goto tr357
		case 84:
			goto tr358
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr356
		case 115:
			goto tr357
		case 116:
			goto tr358
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr356:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st195
	st195:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof195
		}
	st_case_195:
		// line internal/php8/scanner.go:7702
		switch lex.data[(lex.p)] {
		case 76:
			goto tr359
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr359
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr359:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st196
	st196:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof196
		}
	st_case_196:
		// line internal/php8/scanner.go:7743
		switch lex.data[(lex.p)] {
		case 65:
			goto tr360
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr360
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr360:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st197
	st197:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof197
		}
	st_case_197:
		// line internal/php8/scanner.go:7784
		switch lex.data[(lex.p)] {
		case 66:
			goto tr361
		case 92:
			goto st93
		case 96:
			goto tr330
		case 98:
			goto tr361
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr361:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st198
	st198:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof198
		}
	st_case_198:
		// line internal/php8/scanner.go:7825
		switch lex.data[(lex.p)] {
		case 76:
			goto tr362
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr362
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr362:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st199
	st199:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof199
		}
	st_case_199:
		// line internal/php8/scanner.go:7866
		switch lex.data[(lex.p)] {
		case 69:
			goto tr363
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr363
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr357:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st200
	st200:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof200
		}
	st_case_200:
		// line internal/php8/scanner.go:7907
		switch lex.data[(lex.p)] {
		case 69:
			goto tr364
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr364
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr358:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st201
	st201:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof201
		}
	st_case_201:
		// line internal/php8/scanner.go:7948
		switch lex.data[(lex.p)] {
		case 67:
			goto tr365
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr365
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr365:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st202
	st202:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof202
		}
	st_case_202:
		// line internal/php8/scanner.go:7989
		switch lex.data[(lex.p)] {
		case 72:
			goto tr366
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr366
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr353:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st203
	st203:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof203
		}
	st_case_203:
		// line internal/php8/scanner.go:8030
		switch lex.data[(lex.p)] {
		case 85:
			goto tr367
		case 92:
			goto st93
		case 96:
			goto tr330
		case 117:
			goto tr367
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr367:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st204
	st204:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof204
		}
	st_case_204:
		// line internal/php8/scanner.go:8071
		switch lex.data[(lex.p)] {
		case 78:
			goto tr368
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr368
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr368:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st205
	st205:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof205
		}
	st_case_205:
		// line internal/php8/scanner.go:8112
		switch lex.data[(lex.p)] {
		case 67:
			goto tr369
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr369
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr369:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st206
	st206:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof206
		}
	st_case_206:
		// line internal/php8/scanner.go:8153
		switch lex.data[(lex.p)] {
		case 84:
			goto tr370
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr370
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr370:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st207
	st207:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof207
		}
	st_case_207:
		// line internal/php8/scanner.go:8194
		switch lex.data[(lex.p)] {
		case 73:
			goto tr371
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr371
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr371:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st208
	st208:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof208
		}
	st_case_208:
		// line internal/php8/scanner.go:8235
		switch lex.data[(lex.p)] {
		case 79:
			goto tr372
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr372
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr372:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st209
	st209:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof209
		}
	st_case_209:
		// line internal/php8/scanner.go:8276
		switch lex.data[(lex.p)] {
		case 78:
			goto tr373
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr373
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr354:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st210
	st210:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof210
		}
	st_case_210:
		// line internal/php8/scanner.go:8317
		switch lex.data[(lex.p)] {
		case 65:
			goto tr374
		case 79:
			goto tr375
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr374
		case 111:
			goto tr375
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr374:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st211
	st211:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof211
		}
	st_case_211:
		// line internal/php8/scanner.go:8362
		switch lex.data[(lex.p)] {
		case 83:
			goto tr376
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr376
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr376:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st212
	st212:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof212
		}
	st_case_212:
		// line internal/php8/scanner.go:8403
		switch lex.data[(lex.p)] {
		case 83:
			goto tr377
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr377
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr375:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st213
	st213:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof213
		}
	st_case_213:
		// line internal/php8/scanner.go:8444
		switch lex.data[(lex.p)] {
		case 78:
			goto tr378
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr378
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr378:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st214
	st214:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof214
		}
	st_case_214:
		// line internal/php8/scanner.go:8485
		switch lex.data[(lex.p)] {
		case 69:
			goto tr379
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr379
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr355:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st215
	st215:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof215
		}
	st_case_215:
		// line internal/php8/scanner.go:8526
		switch lex.data[(lex.p)] {
		case 78:
			goto tr380
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr380
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr380:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st216
	st216:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof216
		}
	st_case_216:
		// line internal/php8/scanner.go:8567
		switch lex.data[(lex.p)] {
		case 83:
			goto tr381
		case 84:
			goto tr382
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr381
		case 116:
			goto tr382
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr381:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st217
	st217:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof217
		}
	st_case_217:
		// line internal/php8/scanner.go:8612
		switch lex.data[(lex.p)] {
		case 84:
			goto tr383
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr383
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr382:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st218
	st218:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof218
		}
	st_case_218:
		// line internal/php8/scanner.go:8653
		switch lex.data[(lex.p)] {
		case 73:
			goto tr384
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr384
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr384:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st219
	st219:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof219
		}
	st_case_219:
		// line internal/php8/scanner.go:8694
		switch lex.data[(lex.p)] {
		case 78:
			goto tr385
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr385
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr385:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st220
	st220:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof220
		}
	st_case_220:
		// line internal/php8/scanner.go:8735
		switch lex.data[(lex.p)] {
		case 85:
			goto tr386
		case 92:
			goto st93
		case 96:
			goto tr330
		case 117:
			goto tr386
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr386:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st221
	st221:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof221
		}
	st_case_221:
		// line internal/php8/scanner.go:8776
		switch lex.data[(lex.p)] {
		case 69:
			goto tr387
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr387
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr227:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st222
	st222:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof222
		}
	st_case_222:
		// line internal/php8/scanner.go:8817
		switch lex.data[(lex.p)] {
		case 69:
			goto tr388
		case 73:
			goto tr389
		case 79:
			goto tr390
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr388
		case 105:
			goto tr389
		case 111:
			goto tr390
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr388:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st223
	st223:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof223
		}
	st_case_223:
		// line internal/php8/scanner.go:8866
		switch lex.data[(lex.p)] {
		case 67:
			goto tr391
		case 70:
			goto tr392
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr391
		case 102:
			goto tr392
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr391:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st224
	st224:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof224
		}
	st_case_224:
		// line internal/php8/scanner.go:8911
		switch lex.data[(lex.p)] {
		case 76:
			goto tr393
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr393
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr393:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st225
	st225:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof225
		}
	st_case_225:
		// line internal/php8/scanner.go:8952
		switch lex.data[(lex.p)] {
		case 65:
			goto tr394
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr394
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr394:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st226
	st226:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof226
		}
	st_case_226:
		// line internal/php8/scanner.go:8993
		switch lex.data[(lex.p)] {
		case 82:
			goto tr395
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr395
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr395:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st227
	st227:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof227
		}
	st_case_227:
		// line internal/php8/scanner.go:9034
		switch lex.data[(lex.p)] {
		case 69:
			goto tr396
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr396
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr392:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st228
	st228:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof228
		}
	st_case_228:
		// line internal/php8/scanner.go:9075
		switch lex.data[(lex.p)] {
		case 65:
			goto tr397
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr397
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr397:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st229
	st229:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof229
		}
	st_case_229:
		// line internal/php8/scanner.go:9116
		switch lex.data[(lex.p)] {
		case 85:
			goto tr398
		case 92:
			goto st93
		case 96:
			goto tr330
		case 117:
			goto tr398
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr398:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st230
	st230:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof230
		}
	st_case_230:
		// line internal/php8/scanner.go:9157
		switch lex.data[(lex.p)] {
		case 76:
			goto tr399
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr399
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr399:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st231
	st231:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof231
		}
	st_case_231:
		// line internal/php8/scanner.go:9198
		switch lex.data[(lex.p)] {
		case 84:
			goto tr400
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr400
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr389:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st232
	st232:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof232
		}
	st_case_232:
		// line internal/php8/scanner.go:9239
		switch lex.data[(lex.p)] {
		case 69:
			goto tr401
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr401
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr228:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st233
	st233:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof233
		}
	st_case_233:
		// line internal/php8/scanner.go:9280
		switch lex.data[(lex.p)] {
		case 67:
			goto tr402
		case 76:
			goto tr403
		case 77:
			goto tr404
		case 78:
			goto tr405
		case 86:
			goto tr406
		case 88:
			goto tr407
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr402
		case 108:
			goto tr403
		case 109:
			goto tr404
		case 110:
			goto tr405
		case 118:
			goto tr406
		case 120:
			goto tr407
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr402:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st234
	st234:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof234
		}
	st_case_234:
		// line internal/php8/scanner.go:9341
		switch lex.data[(lex.p)] {
		case 72:
			goto tr408
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr408
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr408:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st235
	st235:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof235
		}
	st_case_235:
		// line internal/php8/scanner.go:9382
		switch lex.data[(lex.p)] {
		case 79:
			goto tr409
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr409
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr403:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st236
	st236:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof236
		}
	st_case_236:
		// line internal/php8/scanner.go:9423
		switch lex.data[(lex.p)] {
		case 83:
			goto tr410
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr410
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr410:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st237
	st237:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof237
		}
	st_case_237:
		// line internal/php8/scanner.go:9464
		switch lex.data[(lex.p)] {
		case 69:
			goto tr411
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr411
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr411:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:222
		lex.act = 33
		goto st238
	st238:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof238
		}
	st_case_238:
		// line internal/php8/scanner.go:9505
		switch lex.data[(lex.p)] {
		case 73:
			goto tr413
		case 92:
			goto st93
		case 96:
			goto tr412
		case 105:
			goto tr413
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr412
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr412
				}
			case lex.data[(lex.p)] >= 91:
				goto tr412
			}
		default:
			goto tr412
		}
		goto tr231
	tr413:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st239
	st239:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof239
		}
	st_case_239:
		// line internal/php8/scanner.go:9546
		switch lex.data[(lex.p)] {
		case 70:
			goto tr414
		case 92:
			goto st93
		case 96:
			goto tr330
		case 102:
			goto tr414
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr404:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st240
	st240:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof240
		}
	st_case_240:
		// line internal/php8/scanner.go:9587
		switch lex.data[(lex.p)] {
		case 80:
			goto tr415
		case 92:
			goto st93
		case 96:
			goto tr330
		case 112:
			goto tr415
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr415:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st241
	st241:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof241
		}
	st_case_241:
		// line internal/php8/scanner.go:9628
		switch lex.data[(lex.p)] {
		case 84:
			goto tr416
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr416
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr416:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st242
	st242:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof242
		}
	st_case_242:
		// line internal/php8/scanner.go:9669
		switch lex.data[(lex.p)] {
		case 89:
			goto tr417
		case 92:
			goto st93
		case 96:
			goto tr330
		case 121:
			goto tr417
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr405:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st243
	st243:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof243
		}
	st_case_243:
		// line internal/php8/scanner.go:9710
		switch lex.data[(lex.p)] {
		case 68:
			goto tr418
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr418
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr418:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st244
	st244:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof244
		}
	st_case_244:
		// line internal/php8/scanner.go:9751
		switch lex.data[(lex.p)] {
		case 68:
			goto tr419
		case 70:
			goto tr420
		case 73:
			goto tr421
		case 83:
			goto tr422
		case 87:
			goto tr423
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr419
		case 102:
			goto tr420
		case 105:
			goto tr421
		case 115:
			goto tr422
		case 119:
			goto tr423
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr419:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st245
	st245:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof245
		}
	st_case_245:
		// line internal/php8/scanner.go:9808
		switch lex.data[(lex.p)] {
		case 69:
			goto tr424
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr424
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr424:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st246
	st246:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof246
		}
	st_case_246:
		// line internal/php8/scanner.go:9849
		switch lex.data[(lex.p)] {
		case 67:
			goto tr425
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr425
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr425:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st247
	st247:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof247
		}
	st_case_247:
		// line internal/php8/scanner.go:9890
		switch lex.data[(lex.p)] {
		case 76:
			goto tr426
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr426
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr426:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st248
	st248:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof248
		}
	st_case_248:
		// line internal/php8/scanner.go:9931
		switch lex.data[(lex.p)] {
		case 65:
			goto tr427
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr427
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr427:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st249
	st249:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof249
		}
	st_case_249:
		// line internal/php8/scanner.go:9972
		switch lex.data[(lex.p)] {
		case 82:
			goto tr428
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr428
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr428:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st250
	st250:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof250
		}
	st_case_250:
		// line internal/php8/scanner.go:10013
		switch lex.data[(lex.p)] {
		case 69:
			goto tr429
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr429
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr420:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st251
	st251:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof251
		}
	st_case_251:
		// line internal/php8/scanner.go:10054
		switch lex.data[(lex.p)] {
		case 79:
			goto tr430
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr430
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr430:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st252
	st252:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof252
		}
	st_case_252:
		// line internal/php8/scanner.go:10095
		switch lex.data[(lex.p)] {
		case 82:
			goto tr431
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr431
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr431:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:226
		lex.act = 37
		goto st253
	st253:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof253
		}
	st_case_253:
		// line internal/php8/scanner.go:10136
		switch lex.data[(lex.p)] {
		case 69:
			goto tr433
		case 92:
			goto st93
		case 96:
			goto tr432
		case 101:
			goto tr433
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr432
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr432
				}
			case lex.data[(lex.p)] >= 91:
				goto tr432
			}
		default:
			goto tr432
		}
		goto tr231
	tr433:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st254
	st254:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof254
		}
	st_case_254:
		// line internal/php8/scanner.go:10177
		switch lex.data[(lex.p)] {
		case 65:
			goto tr434
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr434
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr434:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st255
	st255:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof255
		}
	st_case_255:
		// line internal/php8/scanner.go:10218
		switch lex.data[(lex.p)] {
		case 67:
			goto tr435
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr435
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr435:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st256
	st256:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof256
		}
	st_case_256:
		// line internal/php8/scanner.go:10259
		switch lex.data[(lex.p)] {
		case 72:
			goto tr436
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr436
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr421:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st257
	st257:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof257
		}
	st_case_257:
		// line internal/php8/scanner.go:10300
		switch lex.data[(lex.p)] {
		case 70:
			goto tr437
		case 92:
			goto st93
		case 96:
			goto tr330
		case 102:
			goto tr437
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr422:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st258
	st258:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof258
		}
	st_case_258:
		// line internal/php8/scanner.go:10341
		switch lex.data[(lex.p)] {
		case 87:
			goto tr438
		case 92:
			goto st93
		case 96:
			goto tr330
		case 119:
			goto tr438
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr438:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st259
	st259:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof259
		}
	st_case_259:
		// line internal/php8/scanner.go:10382
		switch lex.data[(lex.p)] {
		case 73:
			goto tr439
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr439
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr439:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st260
	st260:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof260
		}
	st_case_260:
		// line internal/php8/scanner.go:10423
		switch lex.data[(lex.p)] {
		case 84:
			goto tr440
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr440
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr440:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st261
	st261:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof261
		}
	st_case_261:
		// line internal/php8/scanner.go:10464
		switch lex.data[(lex.p)] {
		case 67:
			goto tr441
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr441
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr441:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st262
	st262:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof262
		}
	st_case_262:
		// line internal/php8/scanner.go:10505
		switch lex.data[(lex.p)] {
		case 72:
			goto tr442
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr442
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr423:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st263
	st263:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof263
		}
	st_case_263:
		// line internal/php8/scanner.go:10546
		switch lex.data[(lex.p)] {
		case 72:
			goto tr443
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr443
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr443:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st264
	st264:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof264
		}
	st_case_264:
		// line internal/php8/scanner.go:10587
		switch lex.data[(lex.p)] {
		case 73:
			goto tr444
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr444
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr444:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st265
	st265:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof265
		}
	st_case_265:
		// line internal/php8/scanner.go:10628
		switch lex.data[(lex.p)] {
		case 76:
			goto tr445
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr445
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr445:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st266
	st266:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof266
		}
	st_case_266:
		// line internal/php8/scanner.go:10669
		switch lex.data[(lex.p)] {
		case 69:
			goto tr446
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr446
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr406:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st267
	st267:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof267
		}
	st_case_267:
		// line internal/php8/scanner.go:10710
		switch lex.data[(lex.p)] {
		case 65:
			goto tr447
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr447
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr447:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st268
	st268:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof268
		}
	st_case_268:
		// line internal/php8/scanner.go:10751
		switch lex.data[(lex.p)] {
		case 76:
			goto tr448
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr448
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr407:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st269
	st269:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof269
		}
	st_case_269:
		// line internal/php8/scanner.go:10792
		switch lex.data[(lex.p)] {
		case 73:
			goto tr449
		case 84:
			goto tr450
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr449
		case 116:
			goto tr450
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr449:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st270
	st270:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof270
		}
	st_case_270:
		// line internal/php8/scanner.go:10837
		switch lex.data[(lex.p)] {
		case 84:
			goto tr401
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr401
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr450:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st271
	st271:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof271
		}
	st_case_271:
		// line internal/php8/scanner.go:10878
		switch lex.data[(lex.p)] {
		case 69:
			goto tr451
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr451
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr451:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st272
	st272:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof272
		}
	st_case_272:
		// line internal/php8/scanner.go:10919
		switch lex.data[(lex.p)] {
		case 78:
			goto tr452
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr452
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr452:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st273
	st273:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof273
		}
	st_case_273:
		// line internal/php8/scanner.go:10960
		switch lex.data[(lex.p)] {
		case 68:
			goto tr453
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr453
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr453:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st274
	st274:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof274
		}
	st_case_274:
		// line internal/php8/scanner.go:11001
		switch lex.data[(lex.p)] {
		case 83:
			goto tr454
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr454
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr229:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st275
	st275:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof275
		}
	st_case_275:
		// line internal/php8/scanner.go:11042
		switch lex.data[(lex.p)] {
		case 73:
			goto tr455
		case 78:
			goto tr456
		case 79:
			goto tr457
		case 85:
			goto tr367
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr455
		case 110:
			goto tr456
		case 111:
			goto tr457
		case 117:
			goto tr367
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr455:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st276
	st276:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof276
		}
	st_case_276:
		// line internal/php8/scanner.go:11095
		switch lex.data[(lex.p)] {
		case 78:
			goto tr458
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr458
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr458:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st277
	st277:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof277
		}
	st_case_277:
		// line internal/php8/scanner.go:11136
		switch lex.data[(lex.p)] {
		case 65:
			goto tr459
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr459
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr459:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st278
	st278:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof278
		}
	st_case_278:
		// line internal/php8/scanner.go:11177
		switch lex.data[(lex.p)] {
		case 76:
			goto tr460
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr460
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr460:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:234
		lex.act = 45
		goto st279
	st279:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof279
		}
	st_case_279:
		// line internal/php8/scanner.go:11218
		switch lex.data[(lex.p)] {
		case 76:
			goto tr462
		case 92:
			goto st93
		case 96:
			goto tr461
		case 108:
			goto tr462
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr461
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr461
				}
			case lex.data[(lex.p)] >= 91:
				goto tr461
			}
		default:
			goto tr461
		}
		goto tr231
	tr462:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st280
	st280:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof280
		}
	st_case_280:
		// line internal/php8/scanner.go:11259
		switch lex.data[(lex.p)] {
		case 89:
			goto tr463
		case 92:
			goto st93
		case 96:
			goto tr330
		case 121:
			goto tr463
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr457:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st281
	st281:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof281
		}
	st_case_281:
		// line internal/php8/scanner.go:11300
		switch lex.data[(lex.p)] {
		case 82:
			goto tr464
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr464
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr464:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:236
		lex.act = 47
		goto st282
	st282:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof282
		}
	st_case_282:
		// line internal/php8/scanner.go:11341
		switch lex.data[(lex.p)] {
		case 69:
			goto tr466
		case 92:
			goto st93
		case 96:
			goto tr465
		case 101:
			goto tr466
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr465
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr465
				}
			case lex.data[(lex.p)] >= 91:
				goto tr465
			}
		default:
			goto tr465
		}
		goto tr231
	tr466:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st283
	st283:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof283
		}
	st_case_283:
		// line internal/php8/scanner.go:11382
		switch lex.data[(lex.p)] {
		case 65:
			goto tr467
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr467
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr467:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st284
	st284:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof284
		}
	st_case_284:
		// line internal/php8/scanner.go:11423
		switch lex.data[(lex.p)] {
		case 67:
			goto tr468
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr468
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr468:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st285
	st285:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof285
		}
	st_case_285:
		// line internal/php8/scanner.go:11464
		switch lex.data[(lex.p)] {
		case 72:
			goto tr469
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr469
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr230:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st286
	st286:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof286
		}
	st_case_286:
		// line internal/php8/scanner.go:11505
		switch lex.data[(lex.p)] {
		case 76:
			goto tr470
		case 79:
			goto tr471
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr470
		case 111:
			goto tr471
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr470:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st287
	st287:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof287
		}
	st_case_287:
		// line internal/php8/scanner.go:11550
		switch lex.data[(lex.p)] {
		case 79:
			goto tr472
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr472
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr472:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st288
	st288:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof288
		}
	st_case_288:
		// line internal/php8/scanner.go:11591
		switch lex.data[(lex.p)] {
		case 66:
			goto tr473
		case 92:
			goto st93
		case 96:
			goto tr330
		case 98:
			goto tr473
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr473:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st289
	st289:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof289
		}
	st_case_289:
		// line internal/php8/scanner.go:11632
		switch lex.data[(lex.p)] {
		case 65:
			goto tr474
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr474
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr474:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st290
	st290:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof290
		}
	st_case_290:
		// line internal/php8/scanner.go:11673
		switch lex.data[(lex.p)] {
		case 76:
			goto tr475
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr475
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr471:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st291
	st291:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof291
		}
	st_case_291:
		// line internal/php8/scanner.go:11714
		switch lex.data[(lex.p)] {
		case 84:
			goto tr476
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr476
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr476:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st292
	st292:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof292
		}
	st_case_292:
		// line internal/php8/scanner.go:11755
		switch lex.data[(lex.p)] {
		case 79:
			goto tr477
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr477
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr232:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st293
	st293:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof293
		}
	st_case_293:
		// line internal/php8/scanner.go:11796
		switch lex.data[(lex.p)] {
		case 70:
			goto tr478
		case 77:
			goto tr479
		case 78:
			goto tr480
		case 83:
			goto tr481
		case 92:
			goto st93
		case 96:
			goto tr330
		case 102:
			goto tr478
		case 109:
			goto tr479
		case 110:
			goto tr480
		case 115:
			goto tr481
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr479:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st294
	st294:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof294
		}
	st_case_294:
		// line internal/php8/scanner.go:11849
		switch lex.data[(lex.p)] {
		case 80:
			goto tr482
		case 92:
			goto st93
		case 96:
			goto tr330
		case 112:
			goto tr482
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr482:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st295
	st295:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof295
		}
	st_case_295:
		// line internal/php8/scanner.go:11890
		switch lex.data[(lex.p)] {
		case 76:
			goto tr483
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr483
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr483:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st296
	st296:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof296
		}
	st_case_296:
		// line internal/php8/scanner.go:11931
		switch lex.data[(lex.p)] {
		case 69:
			goto tr484
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr484
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr484:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st297
	st297:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof297
		}
	st_case_297:
		// line internal/php8/scanner.go:11972
		switch lex.data[(lex.p)] {
		case 77:
			goto tr485
		case 92:
			goto st93
		case 96:
			goto tr330
		case 109:
			goto tr485
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr485:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st298
	st298:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof298
		}
	st_case_298:
		// line internal/php8/scanner.go:12013
		switch lex.data[(lex.p)] {
		case 69:
			goto tr486
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr486
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr486:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st299
	st299:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof299
		}
	st_case_299:
		// line internal/php8/scanner.go:12054
		switch lex.data[(lex.p)] {
		case 78:
			goto tr487
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr487
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr487:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st300
	st300:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof300
		}
	st_case_300:
		// line internal/php8/scanner.go:12095
		switch lex.data[(lex.p)] {
		case 84:
			goto tr488
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr488
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr488:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st301
	st301:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof301
		}
	st_case_301:
		// line internal/php8/scanner.go:12136
		switch lex.data[(lex.p)] {
		case 83:
			goto tr489
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr489
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr480:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st302
	st302:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof302
		}
	st_case_302:
		// line internal/php8/scanner.go:12177
		switch lex.data[(lex.p)] {
		case 67:
			goto tr490
		case 83:
			goto tr491
		case 84:
			goto tr492
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr490
		case 115:
			goto tr491
		case 116:
			goto tr492
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr490:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st303
	st303:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof303
		}
	st_case_303:
		// line internal/php8/scanner.go:12226
		switch lex.data[(lex.p)] {
		case 76:
			goto tr493
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr493
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr493:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st304
	st304:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof304
		}
	st_case_304:
		// line internal/php8/scanner.go:12267
		switch lex.data[(lex.p)] {
		case 85:
			goto tr494
		case 92:
			goto st93
		case 96:
			goto tr330
		case 117:
			goto tr494
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr494:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st305
	st305:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof305
		}
	st_case_305:
		// line internal/php8/scanner.go:12308
		switch lex.data[(lex.p)] {
		case 68:
			goto tr495
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr495
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr495:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st306
	st306:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof306
		}
	st_case_306:
		// line internal/php8/scanner.go:12349
		switch lex.data[(lex.p)] {
		case 69:
			goto tr496
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr496
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr496:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:268
		lex.act = 79
		goto st307
	st307:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof307
		}
	st_case_307:
		// line internal/php8/scanner.go:12390
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr498
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr497
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr497
				}
			case lex.data[(lex.p)] >= 91:
				goto tr497
			}
		default:
			goto tr497
		}
		goto tr231
	tr498:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st308
	st308:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof308
		}
	st_case_308:
		// line internal/php8/scanner.go:12427
		switch lex.data[(lex.p)] {
		case 79:
			goto tr499
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr499
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr499:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st309
	st309:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof309
		}
	st_case_309:
		// line internal/php8/scanner.go:12468
		switch lex.data[(lex.p)] {
		case 78:
			goto tr500
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr500
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr500:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st310
	st310:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof310
		}
	st_case_310:
		// line internal/php8/scanner.go:12509
		switch lex.data[(lex.p)] {
		case 67:
			goto tr501
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr501
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr501:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st311
	st311:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof311
		}
	st_case_311:
		// line internal/php8/scanner.go:12550
		switch lex.data[(lex.p)] {
		case 69:
			goto tr502
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr502
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr491:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st312
	st312:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof312
		}
	st_case_312:
		// line internal/php8/scanner.go:12591
		switch lex.data[(lex.p)] {
		case 84:
			goto tr503
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr503
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr503:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st313
	st313:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof313
		}
	st_case_313:
		// line internal/php8/scanner.go:12632
		switch lex.data[(lex.p)] {
		case 65:
			goto tr504
		case 69:
			goto tr505
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr504
		case 101:
			goto tr505
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr504:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st314
	st314:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof314
		}
	st_case_314:
		// line internal/php8/scanner.go:12677
		switch lex.data[(lex.p)] {
		case 78:
			goto tr506
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr506
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr506:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st315
	st315:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof315
		}
	st_case_315:
		// line internal/php8/scanner.go:12718
		switch lex.data[(lex.p)] {
		case 67:
			goto tr507
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr507
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr507:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st316
	st316:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof316
		}
	st_case_316:
		// line internal/php8/scanner.go:12759
		switch lex.data[(lex.p)] {
		case 69:
			goto tr508
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr508
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr508:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st317
	st317:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof317
		}
	st_case_317:
		// line internal/php8/scanner.go:12800
		switch lex.data[(lex.p)] {
		case 79:
			goto tr509
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr509
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr509:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st318
	st318:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof318
		}
	st_case_318:
		// line internal/php8/scanner.go:12841
		switch lex.data[(lex.p)] {
		case 70:
			goto tr510
		case 92:
			goto st93
		case 96:
			goto tr330
		case 102:
			goto tr510
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr505:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st319
	st319:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof319
		}
	st_case_319:
		// line internal/php8/scanner.go:12882
		switch lex.data[(lex.p)] {
		case 65:
			goto tr511
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr511
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr511:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st320
	st320:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof320
		}
	st_case_320:
		// line internal/php8/scanner.go:12923
		switch lex.data[(lex.p)] {
		case 68:
			goto tr512
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr512
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr512:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st321
	st321:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof321
		}
	st_case_321:
		// line internal/php8/scanner.go:12964
		switch lex.data[(lex.p)] {
		case 79:
			goto tr513
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr513
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr513:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st322
	st322:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof322
		}
	st_case_322:
		// line internal/php8/scanner.go:13005
		switch lex.data[(lex.p)] {
		case 70:
			goto tr514
		case 92:
			goto st93
		case 96:
			goto tr330
		case 102:
			goto tr514
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr492:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st323
	st323:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof323
		}
	st_case_323:
		// line internal/php8/scanner.go:13046
		switch lex.data[(lex.p)] {
		case 69:
			goto tr515
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr515
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr515:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st324
	st324:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof324
		}
	st_case_324:
		// line internal/php8/scanner.go:13087
		switch lex.data[(lex.p)] {
		case 82:
			goto tr516
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr516
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr516:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st325
	st325:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof325
		}
	st_case_325:
		// line internal/php8/scanner.go:13128
		switch lex.data[(lex.p)] {
		case 70:
			goto tr517
		case 92:
			goto st93
		case 96:
			goto tr330
		case 102:
			goto tr517
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr517:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st326
	st326:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof326
		}
	st_case_326:
		// line internal/php8/scanner.go:13169
		switch lex.data[(lex.p)] {
		case 65:
			goto tr518
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr518
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr518:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st327
	st327:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof327
		}
	st_case_327:
		// line internal/php8/scanner.go:13210
		switch lex.data[(lex.p)] {
		case 67:
			goto tr519
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr519
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr519:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st328
	st328:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof328
		}
	st_case_328:
		// line internal/php8/scanner.go:13251
		switch lex.data[(lex.p)] {
		case 69:
			goto tr520
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr520
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr481:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st329
	st329:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof329
		}
	st_case_329:
		// line internal/php8/scanner.go:13292
		switch lex.data[(lex.p)] {
		case 83:
			goto tr521
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr521
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr521:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st330
	st330:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof330
		}
	st_case_330:
		// line internal/php8/scanner.go:13333
		switch lex.data[(lex.p)] {
		case 69:
			goto tr522
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr522
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr522:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st331
	st331:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof331
		}
	st_case_331:
		// line internal/php8/scanner.go:13374
		switch lex.data[(lex.p)] {
		case 84:
			goto tr523
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr523
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr233:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st332
	st332:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof332
		}
	st_case_332:
		// line internal/php8/scanner.go:13415
		switch lex.data[(lex.p)] {
		case 73:
			goto tr524
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr524
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr524:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st333
	st333:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof333
		}
	st_case_333:
		// line internal/php8/scanner.go:13456
		switch lex.data[(lex.p)] {
		case 83:
			goto tr525
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr525
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr525:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st334
	st334:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof334
		}
	st_case_334:
		// line internal/php8/scanner.go:13497
		switch lex.data[(lex.p)] {
		case 84:
			goto tr526
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr526
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr234:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st335
	st335:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof335
		}
	st_case_335:
		// line internal/php8/scanner.go:13538
		switch lex.data[(lex.p)] {
		case 65:
			goto tr527
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr527
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr527:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st336
	st336:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof336
		}
	st_case_336:
		// line internal/php8/scanner.go:13579
		switch lex.data[(lex.p)] {
		case 84:
			goto tr528
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr528
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr528:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st337
	st337:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof337
		}
	st_case_337:
		// line internal/php8/scanner.go:13620
		switch lex.data[(lex.p)] {
		case 67:
			goto tr529
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr529
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr529:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st338
	st338:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof338
		}
	st_case_338:
		// line internal/php8/scanner.go:13661
		switch lex.data[(lex.p)] {
		case 72:
			goto tr530
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr530
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr235:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st339
	st339:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof339
		}
	st_case_339:
		// line internal/php8/scanner.go:13702
		switch lex.data[(lex.p)] {
		case 65:
			goto tr531
		case 69:
			goto tr532
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr531
		case 101:
			goto tr532
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr531:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st340
	st340:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof340
		}
	st_case_340:
		// line internal/php8/scanner.go:13747
		switch lex.data[(lex.p)] {
		case 77:
			goto tr533
		case 92:
			goto st93
		case 96:
			goto tr330
		case 109:
			goto tr533
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr533:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st341
	st341:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof341
		}
	st_case_341:
		// line internal/php8/scanner.go:13788
		switch lex.data[(lex.p)] {
		case 69:
			goto tr534
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr534
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr534:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st342
	st342:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof342
		}
	st_case_342:
		// line internal/php8/scanner.go:13829
		switch lex.data[(lex.p)] {
		case 83:
			goto tr535
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr535
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr535:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st343
	st343:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof343
		}
	st_case_343:
		// line internal/php8/scanner.go:13870
		switch lex.data[(lex.p)] {
		case 80:
			goto tr536
		case 92:
			goto st93
		case 96:
			goto tr330
		case 112:
			goto tr536
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr536:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st344
	st344:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof344
		}
	st_case_344:
		// line internal/php8/scanner.go:13911
		switch lex.data[(lex.p)] {
		case 65:
			goto tr537
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr537
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr537:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st345
	st345:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof345
		}
	st_case_345:
		// line internal/php8/scanner.go:13952
		switch lex.data[(lex.p)] {
		case 67:
			goto tr538
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr538
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr538:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st346
	st346:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof346
		}
	st_case_346:
		// line internal/php8/scanner.go:13993
		switch lex.data[(lex.p)] {
		case 69:
			goto tr539
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr539
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr539:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:249
		lex.act = 60
		goto st347
	st347:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof347
		}
	st_case_347:
		// line internal/php8/scanner.go:14034
		switch lex.data[(lex.p)] {
		case 92:
			goto st96
		case 96:
			goto tr540
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr540
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr540
				}
			case lex.data[(lex.p)] >= 91:
				goto tr540
			}
		default:
			goto tr540
		}
		goto tr231
	st96:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof96
		}
	st_case_96:
		if lex.data[(lex.p)] == 96 {
			goto tr12
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr12
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr136
	tr136:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:202
		lex.act = 14
		goto st348
	st348:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof348
		}
	st_case_348:
		// line internal/php8/scanner.go:14092
		switch lex.data[(lex.p)] {
		case 92:
			goto st96
		case 96:
			goto tr542
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr542
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr542
				}
			case lex.data[(lex.p)] >= 91:
				goto tr542
			}
		default:
			goto tr542
		}
		goto tr136
	tr532:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st349
	st349:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof349
		}
	st_case_349:
		// line internal/php8/scanner.go:14129
		switch lex.data[(lex.p)] {
		case 87:
			goto tr543
		case 92:
			goto st93
		case 96:
			goto tr330
		case 119:
			goto tr543
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr236:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st350
	st350:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof350
		}
	st_case_350:
		// line internal/php8/scanner.go:14170
		switch lex.data[(lex.p)] {
		case 82:
			goto tr544
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr544
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr237:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st351
	st351:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof351
		}
	st_case_351:
		// line internal/php8/scanner.go:14211
		switch lex.data[(lex.p)] {
		case 82:
			goto tr545
		case 85:
			goto tr546
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr545
		case 117:
			goto tr546
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr545:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st352
	st352:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof352
		}
	st_case_352:
		// line internal/php8/scanner.go:14256
		switch lex.data[(lex.p)] {
		case 73:
			goto tr547
		case 79:
			goto tr548
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr547
		case 111:
			goto tr548
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr547:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st353
	st353:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof353
		}
	st_case_353:
		// line internal/php8/scanner.go:14301
		switch lex.data[(lex.p)] {
		case 78:
			goto tr549
		case 86:
			goto tr550
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr549
		case 118:
			goto tr550
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr549:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st354
	st354:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof354
		}
	st_case_354:
		// line internal/php8/scanner.go:14346
		switch lex.data[(lex.p)] {
		case 84:
			goto tr551
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr551
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr550:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st355
	st355:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof355
		}
	st_case_355:
		// line internal/php8/scanner.go:14387
		switch lex.data[(lex.p)] {
		case 65:
			goto tr552
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr552
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr552:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st356
	st356:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof356
		}
	st_case_356:
		// line internal/php8/scanner.go:14428
		switch lex.data[(lex.p)] {
		case 84:
			goto tr553
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr553
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr553:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st357
	st357:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof357
		}
	st_case_357:
		// line internal/php8/scanner.go:14469
		switch lex.data[(lex.p)] {
		case 69:
			goto tr554
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr554
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr548:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st358
	st358:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof358
		}
	st_case_358:
		// line internal/php8/scanner.go:14510
		switch lex.data[(lex.p)] {
		case 84:
			goto tr555
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr555
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr555:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st359
	st359:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof359
		}
	st_case_359:
		// line internal/php8/scanner.go:14551
		switch lex.data[(lex.p)] {
		case 69:
			goto tr556
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr556
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr556:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st360
	st360:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof360
		}
	st_case_360:
		// line internal/php8/scanner.go:14592
		switch lex.data[(lex.p)] {
		case 67:
			goto tr557
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr557
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr557:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st361
	st361:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof361
		}
	st_case_361:
		// line internal/php8/scanner.go:14633
		switch lex.data[(lex.p)] {
		case 84:
			goto tr558
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr558
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr558:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st362
	st362:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof362
		}
	st_case_362:
		// line internal/php8/scanner.go:14674
		switch lex.data[(lex.p)] {
		case 69:
			goto tr559
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr559
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr559:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st363
	st363:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof363
		}
	st_case_363:
		// line internal/php8/scanner.go:14715
		switch lex.data[(lex.p)] {
		case 68:
			goto tr560
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr560
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr546:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st364
	st364:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof364
		}
	st_case_364:
		// line internal/php8/scanner.go:14756
		switch lex.data[(lex.p)] {
		case 66:
			goto tr561
		case 92:
			goto st93
		case 96:
			goto tr330
		case 98:
			goto tr561
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr561:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st365
	st365:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof365
		}
	st_case_365:
		// line internal/php8/scanner.go:14797
		switch lex.data[(lex.p)] {
		case 76:
			goto tr562
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr562
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr562:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st366
	st366:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof366
		}
	st_case_366:
		// line internal/php8/scanner.go:14838
		switch lex.data[(lex.p)] {
		case 73:
			goto tr563
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr563
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr563:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st367
	st367:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof367
		}
	st_case_367:
		// line internal/php8/scanner.go:14879
		switch lex.data[(lex.p)] {
		case 67:
			goto tr564
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr564
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr238:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st368
	st368:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof368
		}
	st_case_368:
		// line internal/php8/scanner.go:14920
		switch lex.data[(lex.p)] {
		case 69:
			goto tr565
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr565
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr565:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st369
	st369:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof369
		}
	st_case_369:
		// line internal/php8/scanner.go:14961
		switch lex.data[(lex.p)] {
		case 65:
			goto tr566
		case 81:
			goto tr567
		case 84:
			goto tr568
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr566
		case 113:
			goto tr567
		case 116:
			goto tr568
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr566:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st370
	st370:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof370
		}
	st_case_370:
		// line internal/php8/scanner.go:15010
		switch lex.data[(lex.p)] {
		case 68:
			goto tr569
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr569
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr569:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st371
	st371:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof371
		}
	st_case_371:
		// line internal/php8/scanner.go:15051
		switch lex.data[(lex.p)] {
		case 79:
			goto tr570
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr570
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr570:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st372
	st372:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof372
		}
	st_case_372:
		// line internal/php8/scanner.go:15092
		switch lex.data[(lex.p)] {
		case 78:
			goto tr571
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr571
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr571:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st373
	st373:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof373
		}
	st_case_373:
		// line internal/php8/scanner.go:15133
		switch lex.data[(lex.p)] {
		case 76:
			goto tr572
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr572
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr572:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st374
	st374:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof374
		}
	st_case_374:
		// line internal/php8/scanner.go:15174
		switch lex.data[(lex.p)] {
		case 89:
			goto tr573
		case 92:
			goto st93
		case 96:
			goto tr330
		case 121:
			goto tr573
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr567:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st375
	st375:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof375
		}
	st_case_375:
		// line internal/php8/scanner.go:15215
		switch lex.data[(lex.p)] {
		case 85:
			goto tr574
		case 92:
			goto st93
		case 96:
			goto tr330
		case 117:
			goto tr574
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr574:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st376
	st376:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof376
		}
	st_case_376:
		// line internal/php8/scanner.go:15256
		switch lex.data[(lex.p)] {
		case 73:
			goto tr575
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr575
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr575:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st377
	st377:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof377
		}
	st_case_377:
		// line internal/php8/scanner.go:15297
		switch lex.data[(lex.p)] {
		case 82:
			goto tr576
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr576
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr576:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st378
	st378:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof378
		}
	st_case_378:
		// line internal/php8/scanner.go:15338
		switch lex.data[(lex.p)] {
		case 69:
			goto tr577
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr577
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr577:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:270
		lex.act = 81
		goto st379
	st379:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof379
		}
	st_case_379:
		// line internal/php8/scanner.go:15379
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr579
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr578
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr578
				}
			case lex.data[(lex.p)] >= 91:
				goto tr578
			}
		default:
			goto tr578
		}
		goto tr231
	tr579:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st380
	st380:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof380
		}
	st_case_380:
		// line internal/php8/scanner.go:15416
		switch lex.data[(lex.p)] {
		case 79:
			goto tr580
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr580
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr580:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st381
	st381:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof381
		}
	st_case_381:
		// line internal/php8/scanner.go:15457
		switch lex.data[(lex.p)] {
		case 78:
			goto tr581
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr581
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr581:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st382
	st382:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof382
		}
	st_case_382:
		// line internal/php8/scanner.go:15498
		switch lex.data[(lex.p)] {
		case 67:
			goto tr582
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr582
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr582:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st383
	st383:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof383
		}
	st_case_383:
		// line internal/php8/scanner.go:15539
		switch lex.data[(lex.p)] {
		case 69:
			goto tr583
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr583
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr568:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st384
	st384:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof384
		}
	st_case_384:
		// line internal/php8/scanner.go:15580
		switch lex.data[(lex.p)] {
		case 85:
			goto tr584
		case 92:
			goto st93
		case 96:
			goto tr330
		case 117:
			goto tr584
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr584:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st385
	st385:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof385
		}
	st_case_385:
		// line internal/php8/scanner.go:15621
		switch lex.data[(lex.p)] {
		case 82:
			goto tr585
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr585
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr585:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st386
	st386:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof386
		}
	st_case_386:
		// line internal/php8/scanner.go:15662
		switch lex.data[(lex.p)] {
		case 78:
			goto tr586
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr586
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr239:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st387
	st387:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof387
		}
	st_case_387:
		// line internal/php8/scanner.go:15703
		switch lex.data[(lex.p)] {
		case 84:
			goto tr587
		case 87:
			goto tr588
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr587
		case 119:
			goto tr588
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr587:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st388
	st388:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof388
		}
	st_case_388:
		// line internal/php8/scanner.go:15748
		switch lex.data[(lex.p)] {
		case 65:
			goto tr589
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr589
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr589:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st389
	st389:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof389
		}
	st_case_389:
		// line internal/php8/scanner.go:15789
		switch lex.data[(lex.p)] {
		case 84:
			goto tr590
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr590
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr590:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st390
	st390:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof390
		}
	st_case_390:
		// line internal/php8/scanner.go:15830
		switch lex.data[(lex.p)] {
		case 73:
			goto tr591
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr591
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr591:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st391
	st391:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof391
		}
	st_case_391:
		// line internal/php8/scanner.go:15871
		switch lex.data[(lex.p)] {
		case 67:
			goto tr592
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr592
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr588:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st392
	st392:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof392
		}
	st_case_392:
		// line internal/php8/scanner.go:15912
		switch lex.data[(lex.p)] {
		case 73:
			goto tr593
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr593
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr593:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st393
	st393:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof393
		}
	st_case_393:
		// line internal/php8/scanner.go:15953
		switch lex.data[(lex.p)] {
		case 84:
			goto tr594
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr594
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr594:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st394
	st394:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof394
		}
	st_case_394:
		// line internal/php8/scanner.go:15994
		switch lex.data[(lex.p)] {
		case 67:
			goto tr595
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr595
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr595:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st395
	st395:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof395
		}
	st_case_395:
		// line internal/php8/scanner.go:16035
		switch lex.data[(lex.p)] {
		case 72:
			goto tr596
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr596
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr240:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st396
	st396:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof396
		}
	st_case_396:
		// line internal/php8/scanner.go:16076
		switch lex.data[(lex.p)] {
		case 72:
			goto tr597
		case 82:
			goto tr598
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr597
		case 114:
			goto tr598
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr597:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st397
	st397:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof397
		}
	st_case_397:
		// line internal/php8/scanner.go:16121
		switch lex.data[(lex.p)] {
		case 82:
			goto tr599
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr599
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr599:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st398
	st398:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof398
		}
	st_case_398:
		// line internal/php8/scanner.go:16162
		switch lex.data[(lex.p)] {
		case 79:
			goto tr600
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr600
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr600:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st399
	st399:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof399
		}
	st_case_399:
		// line internal/php8/scanner.go:16203
		switch lex.data[(lex.p)] {
		case 87:
			goto tr601
		case 92:
			goto st93
		case 96:
			goto tr330
		case 119:
			goto tr601
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr598:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st400
	st400:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof400
		}
	st_case_400:
		// line internal/php8/scanner.go:16244
		switch lex.data[(lex.p)] {
		case 65:
			goto tr602
		case 89:
			goto tr603
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr602
		case 121:
			goto tr603
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr602:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st401
	st401:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof401
		}
	st_case_401:
		// line internal/php8/scanner.go:16289
		switch lex.data[(lex.p)] {
		case 73:
			goto tr604
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr604
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr604:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st402
	st402:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof402
		}
	st_case_402:
		// line internal/php8/scanner.go:16330
		switch lex.data[(lex.p)] {
		case 84:
			goto tr605
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr605
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr241:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st403
	st403:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof403
		}
	st_case_403:
		// line internal/php8/scanner.go:16371
		switch lex.data[(lex.p)] {
		case 78:
			goto tr606
		case 83:
			goto tr607
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr606
		case 115:
			goto tr607
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr606:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st404
	st404:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof404
		}
	st_case_404:
		// line internal/php8/scanner.go:16416
		switch lex.data[(lex.p)] {
		case 83:
			goto tr608
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr608
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr608:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st405
	st405:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof405
		}
	st_case_405:
		// line internal/php8/scanner.go:16457
		switch lex.data[(lex.p)] {
		case 69:
			goto tr609
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr609
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr609:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st406
	st406:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof406
		}
	st_case_406:
		// line internal/php8/scanner.go:16498
		switch lex.data[(lex.p)] {
		case 84:
			goto tr610
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr610
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr607:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st407
	st407:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof407
		}
	st_case_407:
		// line internal/php8/scanner.go:16539
		switch lex.data[(lex.p)] {
		case 69:
			goto tr611
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr611
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr242:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st408
	st408:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof408
		}
	st_case_408:
		// line internal/php8/scanner.go:16580
		switch lex.data[(lex.p)] {
		case 65:
			goto tr612
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr612
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr612:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st409
	st409:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof409
		}
	st_case_409:
		// line internal/php8/scanner.go:16621
		switch lex.data[(lex.p)] {
		case 82:
			goto tr613
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr613
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr243:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st410
	st410:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof410
		}
	st_case_410:
		// line internal/php8/scanner.go:16662
		switch lex.data[(lex.p)] {
		case 72:
			goto tr614
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr614
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr614:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st411
	st411:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof411
		}
	st_case_411:
		// line internal/php8/scanner.go:16703
		switch lex.data[(lex.p)] {
		case 73:
			goto tr615
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr615
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr615:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st412
	st412:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof412
		}
	st_case_412:
		// line internal/php8/scanner.go:16744
		switch lex.data[(lex.p)] {
		case 76:
			goto tr616
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr616
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr616:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st413
	st413:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof413
		}
	st_case_413:
		// line internal/php8/scanner.go:16785
		switch lex.data[(lex.p)] {
		case 69:
			goto tr617
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr617
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr244:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st414
	st414:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof414
		}
	st_case_414:
		// line internal/php8/scanner.go:16826
		switch lex.data[(lex.p)] {
		case 79:
			goto tr618
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr618
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr618:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st415
	st415:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof415
		}
	st_case_415:
		// line internal/php8/scanner.go:16867
		switch lex.data[(lex.p)] {
		case 82:
			goto tr619
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr619
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr245:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st416
	st416:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof416
		}
	st_case_416:
		// line internal/php8/scanner.go:16908
		switch lex.data[(lex.p)] {
		case 73:
			goto tr620
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr620
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr620:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st417
	st417:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof417
		}
	st_case_417:
		// line internal/php8/scanner.go:16949
		switch lex.data[(lex.p)] {
		case 69:
			goto tr621
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr621
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr621:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st418
	st418:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof418
		}
	st_case_418:
		// line internal/php8/scanner.go:16990
		switch lex.data[(lex.p)] {
		case 76:
			goto tr622
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr622
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr622:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st419
	st419:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof419
		}
	st_case_419:
		// line internal/php8/scanner.go:17031
		switch lex.data[(lex.p)] {
		case 68:
			goto tr623
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr623
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr623:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:267
		lex.act = 78
		goto st420
	st420:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof420
		}
	st_case_420:
		// line internal/php8/scanner.go:17072
		switch lex.data[(lex.p)] {
		case 10:
			goto tr139
		case 13:
			goto tr140
		case 32:
			goto st97
		case 92:
			goto st93
		case 96:
			goto tr624
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st97
				}
			default:
				goto tr624
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 91:
				if 58 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 64 {
					goto tr624
				}
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr624
				}
			default:
				goto tr624
			}
		default:
			goto tr624
		}
		goto tr231
	tr142:
		// line internal/php8/scanner.rl:54

		goto st97
	st97:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		// line internal/php8/scanner.go:17121
		switch lex.data[(lex.p)] {
		case 10:
			goto tr139
		case 13:
			goto tr140
		case 32:
			goto st97
		case 70:
			goto st100
		case 102:
			goto st100
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st97
		}
		goto tr137
	tr139:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st98
	tr143:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st98
	st98:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		// line internal/php8/scanner.go:17169
		switch lex.data[(lex.p)] {
		case 10:
			goto tr143
		case 13:
			goto tr144
		case 32:
			goto tr142
		case 70:
			goto tr145
		case 102:
			goto tr145
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr142
		}
		goto tr137
	tr140:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st99
	tr144:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st99
	st99:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		// line internal/php8/scanner.go:17217
		if lex.data[(lex.p)] == 10 {
			goto tr139
		}
		goto tr137
	tr145:
		// line internal/php8/scanner.rl:54

		goto st100
	st100:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		// line internal/php8/scanner.go:17231
		switch lex.data[(lex.p)] {
		case 82:
			goto st101
		case 114:
			goto st101
		}
		goto tr137
	st101:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		switch lex.data[(lex.p)] {
		case 79:
			goto st102
		case 111:
			goto st102
		}
		goto tr137
	st102:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr148
		case 109:
			goto tr148
		}
		goto tr137
	st421:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof421
		}
	st_case_421:
		if lex.data[(lex.p)] == 96 {
			goto tr625
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr625
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr625
			}
		default:
			goto tr625
		}
		goto tr150
	tr150:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st422
	st422:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof422
		}
	st_case_422:
		// line internal/php8/scanner.go:17294
		switch lex.data[(lex.p)] {
		case 92:
			goto st103
		case 96:
			goto tr626
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr626
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr626
				}
			case lex.data[(lex.p)] >= 91:
				goto tr626
			}
		default:
			goto tr626
		}
		goto tr150
	st103:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		if lex.data[(lex.p)] == 96 {
			goto tr149
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr149
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr149
			}
		default:
			goto tr149
		}
		goto tr150
	st423:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof423
		}
	st_case_423:
		if lex.data[(lex.p)] == 61 {
			goto tr628
		}
		goto tr260
	tr248:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st424
	st424:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof424
		}
	st_case_424:
		// line internal/php8/scanner.go:17361
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr629
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr629:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st425
	st425:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof425
		}
	st_case_425:
		// line internal/php8/scanner.go:17398
		switch lex.data[(lex.p)] {
		case 67:
			goto tr630
		case 68:
			goto tr631
		case 70:
			goto tr632
		case 72:
			goto tr633
		case 76:
			goto tr634
		case 77:
			goto tr635
		case 78:
			goto tr636
		case 84:
			goto tr637
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr630
		case 100:
			goto tr631
		case 102:
			goto tr632
		case 104:
			goto tr633
		case 108:
			goto tr634
		case 109:
			goto tr635
		case 110:
			goto tr636
		case 116:
			goto tr637
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr630:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st426
	st426:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof426
		}
	st_case_426:
		// line internal/php8/scanner.go:17467
		switch lex.data[(lex.p)] {
		case 76:
			goto tr638
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr638
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr638:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st427
	st427:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof427
		}
	st_case_427:
		// line internal/php8/scanner.go:17508
		switch lex.data[(lex.p)] {
		case 65:
			goto tr639
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr639
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr639:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st428
	st428:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof428
		}
	st_case_428:
		// line internal/php8/scanner.go:17549
		switch lex.data[(lex.p)] {
		case 83:
			goto tr640
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr640
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr640:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st429
	st429:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof429
		}
	st_case_429:
		// line internal/php8/scanner.go:17590
		switch lex.data[(lex.p)] {
		case 83:
			goto tr641
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr641
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr641:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st430
	st430:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof430
		}
	st_case_430:
		// line internal/php8/scanner.go:17631
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr642
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr642:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st431
	st431:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof431
		}
	st_case_431:
		// line internal/php8/scanner.go:17668
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr643
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr631:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st432
	st432:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof432
		}
	st_case_432:
		// line internal/php8/scanner.go:17705
		switch lex.data[(lex.p)] {
		case 73:
			goto tr644
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr644
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr644:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st433
	st433:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof433
		}
	st_case_433:
		// line internal/php8/scanner.go:17746
		switch lex.data[(lex.p)] {
		case 82:
			goto tr645
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr645
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr645:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st434
	st434:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof434
		}
	st_case_434:
		// line internal/php8/scanner.go:17787
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr646
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr646:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st435
	st435:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof435
		}
	st_case_435:
		// line internal/php8/scanner.go:17824
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr647
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr632:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st436
	st436:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof436
		}
	st_case_436:
		// line internal/php8/scanner.go:17861
		switch lex.data[(lex.p)] {
		case 73:
			goto tr648
		case 85:
			goto tr649
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr648
		case 117:
			goto tr649
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr648:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st437
	st437:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof437
		}
	st_case_437:
		// line internal/php8/scanner.go:17906
		switch lex.data[(lex.p)] {
		case 76:
			goto tr650
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr650
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr650:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st438
	st438:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof438
		}
	st_case_438:
		// line internal/php8/scanner.go:17947
		switch lex.data[(lex.p)] {
		case 69:
			goto tr651
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr651
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr651:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st439
	st439:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof439
		}
	st_case_439:
		// line internal/php8/scanner.go:17988
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr652
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr652:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st440
	st440:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof440
		}
	st_case_440:
		// line internal/php8/scanner.go:18025
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr653
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr649:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st441
	st441:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof441
		}
	st_case_441:
		// line internal/php8/scanner.go:18062
		switch lex.data[(lex.p)] {
		case 78:
			goto tr654
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr654
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr654:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st442
	st442:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof442
		}
	st_case_442:
		// line internal/php8/scanner.go:18103
		switch lex.data[(lex.p)] {
		case 67:
			goto tr655
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr655
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr655:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st443
	st443:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof443
		}
	st_case_443:
		// line internal/php8/scanner.go:18144
		switch lex.data[(lex.p)] {
		case 84:
			goto tr656
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr656
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr656:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st444
	st444:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof444
		}
	st_case_444:
		// line internal/php8/scanner.go:18185
		switch lex.data[(lex.p)] {
		case 73:
			goto tr657
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr657
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr657:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st445
	st445:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof445
		}
	st_case_445:
		// line internal/php8/scanner.go:18226
		switch lex.data[(lex.p)] {
		case 79:
			goto tr658
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr658
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr658:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st446
	st446:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof446
		}
	st_case_446:
		// line internal/php8/scanner.go:18267
		switch lex.data[(lex.p)] {
		case 78:
			goto tr659
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr659
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr659:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st447
	st447:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof447
		}
	st_case_447:
		// line internal/php8/scanner.go:18308
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr660
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr660:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st448
	st448:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof448
		}
	st_case_448:
		// line internal/php8/scanner.go:18345
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr661
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr633:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st449
	st449:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof449
		}
	st_case_449:
		// line internal/php8/scanner.go:18382
		switch lex.data[(lex.p)] {
		case 65:
			goto tr662
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr662
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr662:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st450
	st450:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof450
		}
	st_case_450:
		// line internal/php8/scanner.go:18423
		switch lex.data[(lex.p)] {
		case 76:
			goto tr663
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr663
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr663:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st451
	st451:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof451
		}
	st_case_451:
		// line internal/php8/scanner.go:18464
		switch lex.data[(lex.p)] {
		case 84:
			goto tr664
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr664
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr664:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st452
	st452:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof452
		}
	st_case_452:
		// line internal/php8/scanner.go:18505
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr665
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr665:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st453
	st453:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof453
		}
	st_case_453:
		// line internal/php8/scanner.go:18542
		switch lex.data[(lex.p)] {
		case 67:
			goto tr666
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr666
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr666:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st454
	st454:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof454
		}
	st_case_454:
		// line internal/php8/scanner.go:18583
		switch lex.data[(lex.p)] {
		case 79:
			goto tr667
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr667
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr667:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st455
	st455:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof455
		}
	st_case_455:
		// line internal/php8/scanner.go:18624
		switch lex.data[(lex.p)] {
		case 77:
			goto tr668
		case 92:
			goto st93
		case 96:
			goto tr330
		case 109:
			goto tr668
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr668:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st456
	st456:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof456
		}
	st_case_456:
		// line internal/php8/scanner.go:18665
		switch lex.data[(lex.p)] {
		case 80:
			goto tr669
		case 92:
			goto st93
		case 96:
			goto tr330
		case 112:
			goto tr669
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr669:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st457
	st457:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof457
		}
	st_case_457:
		// line internal/php8/scanner.go:18706
		switch lex.data[(lex.p)] {
		case 73:
			goto tr670
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr670
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr670:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st458
	st458:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof458
		}
	st_case_458:
		// line internal/php8/scanner.go:18747
		switch lex.data[(lex.p)] {
		case 76:
			goto tr671
		case 92:
			goto st93
		case 96:
			goto tr330
		case 108:
			goto tr671
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr671:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st459
	st459:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof459
		}
	st_case_459:
		// line internal/php8/scanner.go:18788
		switch lex.data[(lex.p)] {
		case 69:
			goto tr672
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr672
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr672:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st460
	st460:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof460
		}
	st_case_460:
		// line internal/php8/scanner.go:18829
		switch lex.data[(lex.p)] {
		case 82:
			goto tr673
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr673
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr634:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st461
	st461:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof461
		}
	st_case_461:
		// line internal/php8/scanner.go:18870
		switch lex.data[(lex.p)] {
		case 73:
			goto tr674
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr674
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr674:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st462
	st462:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof462
		}
	st_case_462:
		// line internal/php8/scanner.go:18911
		switch lex.data[(lex.p)] {
		case 78:
			goto tr675
		case 92:
			goto st93
		case 96:
			goto tr330
		case 110:
			goto tr675
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr675:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st463
	st463:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof463
		}
	st_case_463:
		// line internal/php8/scanner.go:18952
		switch lex.data[(lex.p)] {
		case 69:
			goto tr676
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr676
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr676:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st464
	st464:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof464
		}
	st_case_464:
		// line internal/php8/scanner.go:18993
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr677
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr677:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st465
	st465:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof465
		}
	st_case_465:
		// line internal/php8/scanner.go:19030
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr678
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr635:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st466
	st466:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof466
		}
	st_case_466:
		// line internal/php8/scanner.go:19067
		switch lex.data[(lex.p)] {
		case 69:
			goto tr679
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr679
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr679:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st467
	st467:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof467
		}
	st_case_467:
		// line internal/php8/scanner.go:19108
		switch lex.data[(lex.p)] {
		case 84:
			goto tr680
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr680
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr680:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st468
	st468:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof468
		}
	st_case_468:
		// line internal/php8/scanner.go:19149
		switch lex.data[(lex.p)] {
		case 72:
			goto tr681
		case 92:
			goto st93
		case 96:
			goto tr330
		case 104:
			goto tr681
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr681:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st469
	st469:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof469
		}
	st_case_469:
		// line internal/php8/scanner.go:19190
		switch lex.data[(lex.p)] {
		case 79:
			goto tr682
		case 92:
			goto st93
		case 96:
			goto tr330
		case 111:
			goto tr682
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr682:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st470
	st470:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof470
		}
	st_case_470:
		// line internal/php8/scanner.go:19231
		switch lex.data[(lex.p)] {
		case 68:
			goto tr683
		case 92:
			goto st93
		case 96:
			goto tr330
		case 100:
			goto tr683
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr683:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st471
	st471:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof471
		}
	st_case_471:
		// line internal/php8/scanner.go:19272
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr684
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr684:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st472
	st472:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof472
		}
	st_case_472:
		// line internal/php8/scanner.go:19309
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr685
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr636:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st473
	st473:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof473
		}
	st_case_473:
		// line internal/php8/scanner.go:19346
		switch lex.data[(lex.p)] {
		case 65:
			goto tr686
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr686
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr686:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st474
	st474:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof474
		}
	st_case_474:
		// line internal/php8/scanner.go:19387
		switch lex.data[(lex.p)] {
		case 77:
			goto tr687
		case 92:
			goto st93
		case 96:
			goto tr330
		case 109:
			goto tr687
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr687:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st475
	st475:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof475
		}
	st_case_475:
		// line internal/php8/scanner.go:19428
		switch lex.data[(lex.p)] {
		case 69:
			goto tr688
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr688
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr688:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st476
	st476:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof476
		}
	st_case_476:
		// line internal/php8/scanner.go:19469
		switch lex.data[(lex.p)] {
		case 83:
			goto tr689
		case 92:
			goto st93
		case 96:
			goto tr330
		case 115:
			goto tr689
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr689:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st477
	st477:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof477
		}
	st_case_477:
		// line internal/php8/scanner.go:19510
		switch lex.data[(lex.p)] {
		case 80:
			goto tr690
		case 92:
			goto st93
		case 96:
			goto tr330
		case 112:
			goto tr690
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr690:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st478
	st478:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof478
		}
	st_case_478:
		// line internal/php8/scanner.go:19551
		switch lex.data[(lex.p)] {
		case 65:
			goto tr691
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr691
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr691:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st479
	st479:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof479
		}
	st_case_479:
		// line internal/php8/scanner.go:19592
		switch lex.data[(lex.p)] {
		case 67:
			goto tr692
		case 92:
			goto st93
		case 96:
			goto tr330
		case 99:
			goto tr692
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr692:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st480
	st480:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof480
		}
	st_case_480:
		// line internal/php8/scanner.go:19633
		switch lex.data[(lex.p)] {
		case 69:
			goto tr693
		case 92:
			goto st93
		case 96:
			goto tr330
		case 101:
			goto tr693
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr693:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st481
	st481:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof481
		}
	st_case_481:
		// line internal/php8/scanner.go:19674
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr694
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr694:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st482
	st482:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof482
		}
	st_case_482:
		// line internal/php8/scanner.go:19711
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr695
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr637:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st483
	st483:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof483
		}
	st_case_483:
		// line internal/php8/scanner.go:19748
		switch lex.data[(lex.p)] {
		case 82:
			goto tr696
		case 92:
			goto st93
		case 96:
			goto tr330
		case 114:
			goto tr696
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr696:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st484
	st484:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof484
		}
	st_case_484:
		// line internal/php8/scanner.go:19789
		switch lex.data[(lex.p)] {
		case 65:
			goto tr697
		case 92:
			goto st93
		case 96:
			goto tr330
		case 97:
			goto tr697
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr697:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st485
	st485:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof485
		}
	st_case_485:
		// line internal/php8/scanner.go:19830
		switch lex.data[(lex.p)] {
		case 73:
			goto tr698
		case 92:
			goto st93
		case 96:
			goto tr330
		case 105:
			goto tr698
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr698:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st486
	st486:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof486
		}
	st_case_486:
		// line internal/php8/scanner.go:19871
		switch lex.data[(lex.p)] {
		case 84:
			goto tr699
		case 92:
			goto st93
		case 96:
			goto tr330
		case 116:
			goto tr699
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr699:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st487
	st487:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof487
		}
	st_case_487:
		// line internal/php8/scanner.go:19912
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr700
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	tr700:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:356
		lex.act = 143
		goto st488
	st488:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof488
		}
	st_case_488:
		// line internal/php8/scanner.go:19949
		switch lex.data[(lex.p)] {
		case 92:
			goto st93
		case 95:
			goto tr701
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr330
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 96:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr330
				}
			case lex.data[(lex.p)] >= 91:
				goto tr330
			}
		default:
			goto tr330
		}
		goto tr231
	st489:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof489
		}
	st_case_489:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr702
		case 124:
			goto tr703
		}
		goto tr260
	tr151:
		// line internal/php8/scanner.rl:391
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st490
	tr153:
		// line internal/php8/scanner.rl:395
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				goto st131
			}
		}
		goto st490
	tr154:
		// line internal/php8/scanner.rl:393
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 490
				goto _out
			}
		}
		goto st490
	tr704:
		// line internal/php8/scanner.rl:395
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				goto st131
			}
		}
		goto st490
	tr710:
		// line internal/php8/scanner.rl:391
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st490
	tr712:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:391
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st490
	tr716:
		// line internal/php8/scanner.rl:395
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				goto st131
			}
		}
		goto st490
	tr717:
		// line internal/php8/scanner.rl:392
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 490
				goto _out
			}
		}
		goto st490
	tr719:
		lex.cs = 490
		// line internal/php8/scanner.rl:394
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			lex.cs = 131
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st490:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof490
		}
	st_case_490:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:20049
		switch lex.data[(lex.p)] {
		case 10:
			goto tr152
		case 13:
			goto tr706
		case 32:
			goto tr705
		case 45:
			goto st494
		case 63:
			goto tr708
		case 96:
			goto tr704
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr705
				}
			default:
				goto tr704
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr704
				}
			case lex.data[(lex.p)] >= 91:
				goto tr704
			}
		default:
			goto tr704
		}
		goto st496
	tr705:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st491
	tr713:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		goto st491
	st491:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof491
		}
	st_case_491:
		// line internal/php8/scanner.go:20104
		switch lex.data[(lex.p)] {
		case 10:
			goto tr152
		case 13:
			goto tr711
		case 32:
			goto tr705
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr705
		}
		goto tr710
	tr152:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st492
	tr714:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st492
	st492:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof492
		}
	st_case_492:
		// line internal/php8/scanner.go:20154
		switch lex.data[(lex.p)] {
		case 10:
			goto tr714
		case 13:
			goto tr715
		case 32:
			goto tr713
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr713
		}
		goto tr712
	tr711:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st104
	tr715:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st104
	st104:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		// line internal/php8/scanner.go:20198
		if lex.data[(lex.p)] == 10 {
			goto tr152
		}
		goto tr151
	tr706:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st493
	st493:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof493
		}
	st_case_493:
		// line internal/php8/scanner.go:20220
		if lex.data[(lex.p)] == 10 {
			goto tr152
		}
		goto tr716
	st494:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof494
		}
	st_case_494:
		if lex.data[(lex.p)] == 62 {
			goto tr717
		}
		goto tr716
	tr708:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st495
	st495:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof495
		}
	st_case_495:
		// line internal/php8/scanner.go:20244
		if lex.data[(lex.p)] == 45 {
			goto st105
		}
		goto tr716
	st105:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		if lex.data[(lex.p)] == 62 {
			goto tr154
		}
		goto tr153
	st496:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof496
		}
	st_case_496:
		if lex.data[(lex.p)] == 96 {
			goto tr719
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr719
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr719
				}
			case lex.data[(lex.p)] >= 91:
				goto tr719
			}
		default:
			goto tr719
		}
		goto st496
	tr723:
		lex.cs = 497
		// line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 156:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				lex.cs = 520
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr724:
		lex.cs = 497
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:399
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.cs = 520
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr728:
		lex.cs = 497
		// line internal/php8/scanner.rl:399
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.cs = 520
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st497:
		// line NONE:1
		lex.ts = 0

		// line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof497
		}
	st_case_497:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:20341
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr721
		case 1037:
			goto tr722
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr720
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	tr720:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:399
		lex.act = 156
		goto st498
	tr725:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:399
		lex.act = 156
		goto st498
	st498:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof498
		}
	st_case_498:
		// line internal/php8/scanner.go:20415
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr721
		case 1037:
			goto tr722
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr720
		}
		goto tr723
	tr721:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st499
	tr726:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st499
	st499:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof499
		}
	st_case_499:
		// line internal/php8/scanner.go:20495
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr726
		case 1037:
			goto tr727
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr725
		}
		goto tr724
	tr722:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st500
	tr727:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st500
	st500:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof500
		}
	st_case_500:
		// line internal/php8/scanner.go:20575
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			default:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) {
					_widec += 256
				}
			}
		default:
			_widec = 768 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) {
				_widec += 256
			}
		}
		switch _widec {
		case 1034:
			goto tr721
		case 1037:
			goto tr722
		}
		if 1024 <= _widec && _widec <= 1279 {
			goto tr720
		}
		goto tr728
	tr155:
		// line internal/php8/scanner.rl:408
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(501, 131)
			goto _out
		}
		goto st501
	tr736:
		// line internal/php8/scanner.rl:410
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 501
					lex.top++
					goto st522
				}
			}
		}
		goto st501
	tr737:
		// line internal/php8/scanner.rl:409
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(501, 538)
			goto _out
		}
		goto st501
	tr738:
		lex.cs = 501
		// line NONE:1
		switch lex.act {
		case 157:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(501, 131)
				goto _out
			}
		case 158:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(501, 538)
				goto _out
			}
		case 160:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE

				if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
					lex.cs = 520
				}
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr739:
		lex.cs = 501
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:411
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 520
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr743:
		lex.cs = 501
		// line internal/php8/scanner.rl:411
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE

			if len(lex.data) > lex.p+1 && lex.data[lex.p+1] != '$' && lex.data[lex.p+1] != '{' {
				lex.cs = 520
			}
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st501:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof501
		}
	st_case_501:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:20707
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto st502
		case 1403:
			goto st106
		case 1546:
			goto tr732
		case 1549:
			goto tr733
		case 1572:
			goto st506
		case 1659:
			goto st507
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr731
		}
		goto st0
	st502:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof502
		}
	st_case_502:
		if lex.data[(lex.p)] == 123 {
			goto tr737
		}
		goto tr736
	st106:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		if lex.data[(lex.p)] == 36 {
			goto tr155
		}
		goto st0
	tr731:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:411
		lex.act = 160
		goto st503
	tr740:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:411
		lex.act = 160
		goto st503
	tr744:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:409
		lex.act = 158
		goto st503
	tr745:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:408
		lex.act = 157
		goto st503
	st503:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof503
		}
	st_case_503:
		// line internal/php8/scanner.go:20817
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr732
		case 1549:
			goto tr733
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr731
		}
		goto tr738
	tr732:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st504
	tr741:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st504
	st504:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof504
		}
	st_case_504:
		// line internal/php8/scanner.go:20897
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr741
		case 1549:
			goto tr742
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr740
		}
		goto tr739
	tr733:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st505
	tr742:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st505
	st505:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof505
		}
	st_case_505:
		// line internal/php8/scanner.go:20977
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1546:
			goto tr732
		case 1549:
			goto tr733
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr731
		}
		goto tr743
	st506:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof506
		}
	st_case_506:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1403:
			goto tr737
		case 1546:
			goto tr732
		case 1549:
			goto tr733
		case 1659:
			goto tr744
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr731
		}
		goto tr736
	st507:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof507
		}
	st_case_507:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1280 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotHeredocEnd(lex.p) && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1316:
			goto tr155
		case 1546:
			goto tr732
		case 1549:
			goto tr733
		case 1572:
			goto tr745
		}
		if 1536 <= _widec && _widec <= 1791 {
			goto tr731
		}
		goto tr743
	tr157:
		// line internal/php8/scanner.rl:425
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(2)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 508
					lex.top++
					goto st522
				}
			}
		}
		goto st508
	tr158:
		// line internal/php8/scanner.rl:424
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(508, 538)
			goto _out
		}
		goto st508
	tr159:
		// line internal/php8/scanner.rl:423
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(508, 131)
			goto _out
		}
		goto st508
	tr747:
		lex.cs = 508
		// line internal/php8/scanner.rl:426
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('`'))
			lex.cs = 131
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr755:
		lex.cs = 508
		// line NONE:1
		switch lex.act {
		case 161:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(508, 131)
				goto _out
			}
		case 162:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(508, 538)
				goto _out
			}
		case 163:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(2)
				{
					lex.growCallStack()
					{
						lex.stack[lex.top] = 508
						lex.top++
						goto st522
					}
				}
			}
		case 164:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('`'))
				lex.cs = 131
				{
					(lex.p)++
					goto _out
				}
			}
		case 165:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr756:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:427
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 508
				goto _out
			}
		}
		goto st508
	tr760:
		// line internal/php8/scanner.rl:427
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 508
				goto _out
			}
		}
		goto st508
	st508:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof508
		}
	st_case_508:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:21220
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto st107
		case 1888:
			goto tr747
		case 1915:
			goto st108
		case 2058:
			goto tr750
		case 2061:
			goto tr751
		case 2084:
			goto st512
		case 2144:
			goto tr753
		case 2171:
			goto st513
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr749
		}
		goto st0
	st107:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch lex.data[(lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr158
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr157
	st108:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		if lex.data[(lex.p)] == 36 {
			goto tr159
		}
		goto st0
	tr749:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:427
		lex.act = 165
		goto st509
	tr753:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:426
		lex.act = 164
		goto st509
	tr757:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:427
		lex.act = 165
		goto st509
	tr761:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:425
		lex.act = 163
		goto st509
	tr762:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:424
		lex.act = 162
		goto st509
	tr763:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:423
		lex.act = 161
		goto st509
	st509:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof509
		}
	st_case_509:
		// line internal/php8/scanner.go:21363
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr750
		case 2061:
			goto tr751
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr749
		}
		goto tr755
	tr750:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st510
	tr758:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st510
	st510:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof510
		}
	st_case_510:
		// line internal/php8/scanner.go:21443
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr758
		case 2061:
			goto tr759
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr757
		}
		goto tr756
	tr751:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st511
	tr759:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st511
	st511:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof511
		}
	st_case_511:
		// line internal/php8/scanner.go:21523
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2058:
			goto tr750
		case 2061:
			goto tr751
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr749
		}
		goto tr760
	st512:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof512
		}
	st_case_512:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1887:
			goto tr157
		case 1915:
			goto tr158
		case 2058:
			goto tr750
		case 2061:
			goto tr751
		case 2143:
			goto tr761
		case 2171:
			goto tr762
		}
		switch {
		case _widec < 2113:
			switch {
			case _widec < 1889:
				if 1857 <= _widec && _widec <= 1882 {
					goto tr157
				}
			case _widec > 1914:
				switch {
				case _widec > 2047:
					if 2048 <= _widec && _widec <= 2112 {
						goto tr749
					}
				case _widec >= 1920:
					goto tr157
				}
			default:
				goto tr157
			}
		case _widec > 2138:
			switch {
			case _widec < 2145:
				if 2139 <= _widec && _widec <= 2144 {
					goto tr749
				}
			case _widec > 2170:
				switch {
				case _widec > 2175:
					if 2176 <= _widec && _widec <= 2303 {
						goto tr761
					}
				case _widec >= 2172:
					goto tr749
				}
			default:
				goto tr761
			}
		default:
			goto tr761
		}
		goto tr760
	st513:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof513
		}
	st_case_513:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('`') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('`') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 1792 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('`') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 1828:
			goto tr159
		case 2058:
			goto tr750
		case 2061:
			goto tr751
		case 2084:
			goto tr763
		}
		if 2048 <= _widec && _widec <= 2303 {
			goto tr749
		}
		goto tr760
	tr160:
		// line internal/php8/scanner.rl:437
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(2)
			{
				lex.growCallStack()
				{
					lex.stack[lex.top] = 514
					lex.top++
					goto st522
				}
			}
		}
		goto st514
	tr161:
		// line internal/php8/scanner.rl:436
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_DOLLAR_OPEN_CURLY_BRACES
			lex.call(514, 538)
			goto _out
		}
		goto st514
	tr162:
		// line internal/php8/scanner.rl:435
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_CURLY_OPEN
			lex.call(514, 131)
			goto _out
		}
		goto st514
	tr764:
		lex.cs = 514
		// line internal/php8/scanner.rl:438
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('"'))
			lex.cs = 131
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr773:
		lex.cs = 514
		// line NONE:1
		switch lex.act {
		case 166:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(1)
				lex.setTokenPosition(tkn)
				tok = token.T_CURLY_OPEN
				lex.call(514, 131)
				goto _out
			}
		case 167:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.T_DOLLAR_OPEN_CURLY_BRACES
				lex.call(514, 538)
				goto _out
			}
		case 168:
			{
				(lex.p) = (lex.te) - 1
				lex.ungetCnt(2)
				{
					lex.growCallStack()
					{
						lex.stack[lex.top] = 514
						lex.top++
						goto st522
					}
				}
			}
		case 169:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPosition(tkn)
				tok = token.ID(int('"'))
				lex.cs = 131
				{
					(lex.p)++
					goto _out
				}
			}
		case 170:
			{
				(lex.p) = (lex.te) - 1

				lex.setTokenPosition(tkn)
				tok = token.T_ENCAPSED_AND_WHITESPACE
				{
					(lex.p)++
					goto _out
				}
			}
		}

		goto _again
	tr774:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:439
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 514
				goto _out
			}
		}
		goto st514
	tr778:
		// line internal/php8/scanner.rl:439
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			{
				(lex.p)++
				lex.cs = 514
				goto _out
			}
		}
		goto st514
	st514:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof514
		}
	st_case_514:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:21807
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2338:
			goto tr764
		case 2340:
			goto st109
		case 2427:
			goto st110
		case 2570:
			goto tr768
		case 2573:
			goto tr769
		case 2594:
			goto tr770
		case 2596:
			goto st518
		case 2683:
			goto st519
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr767
		}
		goto st0
	st109:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch lex.data[(lex.p)] {
		case 96:
			goto st0
		case 123:
			goto tr161
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr160
	st110:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		if lex.data[(lex.p)] == 36 {
			goto tr162
		}
		goto st0
	tr767:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:439
		lex.act = 170
		goto st515
	tr770:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:438
		lex.act = 169
		goto st515
	tr775:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:439
		lex.act = 170
		goto st515
	tr779:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:437
		lex.act = 168
		goto st515
	tr780:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:436
		lex.act = 167
		goto st515
	tr781:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:435
		lex.act = 166
		goto st515
	st515:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof515
		}
	st_case_515:
		// line internal/php8/scanner.go:21950
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr768
		case 2573:
			goto tr769
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr767
		}
		goto tr773
	tr768:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st516
	tr776:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st516
	st516:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof516
		}
	st_case_516:
		// line internal/php8/scanner.go:22030
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr776
		case 2573:
			goto tr777
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr775
		}
		goto tr774
	tr769:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st517
	tr777:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st517
	st517:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof517
		}
	st_case_517:
		// line internal/php8/scanner.go:22110
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2570:
			goto tr768
		case 2573:
			goto tr769
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr767
		}
		goto tr778
	st518:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof518
		}
	st_case_518:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2399:
			goto tr160
		case 2427:
			goto tr161
		case 2570:
			goto tr768
		case 2573:
			goto tr769
		case 2655:
			goto tr779
		case 2683:
			goto tr780
		}
		switch {
		case _widec < 2625:
			switch {
			case _widec < 2401:
				if 2369 <= _widec && _widec <= 2394 {
					goto tr160
				}
			case _widec > 2426:
				switch {
				case _widec > 2559:
					if 2560 <= _widec && _widec <= 2624 {
						goto tr767
					}
				case _widec >= 2432:
					goto tr160
				}
			default:
				goto tr160
			}
		case _widec > 2650:
			switch {
			case _widec < 2657:
				if 2651 <= _widec && _widec <= 2656 {
					goto tr767
				}
			case _widec > 2682:
				switch {
				case _widec > 2687:
					if 2688 <= _widec && _widec <= 2815 {
						goto tr779
					}
				case _widec >= 2684:
					goto tr767
				}
			default:
				goto tr779
			}
		default:
			goto tr779
		}
		goto tr778
	st519:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof519
		}
	st_case_519:
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			default:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotStringEnd('"') && lex.isNotStringVar() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotStringEnd('"') && lex.isNotStringVar() {
					_widec += 256
				}
			}
		default:
			_widec = 2304 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotStringEnd('"') && lex.isNotStringVar() {
				_widec += 256
			}
		}
		switch _widec {
		case 2340:
			goto tr162
		case 2570:
			goto tr768
		case 2573:
			goto tr769
		case 2596:
			goto tr781
		}
		if 2560 <= _widec && _widec <= 2815 {
			goto tr767
		}
		goto tr778
	tr783:
		lex.cs = 520
		// line internal/php8/scanner.rl:447
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_END_HEREDOC
			lex.cs = 131
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st520:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof520
		}
	st_case_520:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:22337
		if lex.data[(lex.p)] == 96 {
			goto st0
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st0
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st0
			}
		default:
			goto st0
		}
		goto st521
	st521:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof521
		}
	st_case_521:
		if lex.data[(lex.p)] == 96 {
			goto tr783
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr783
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr783
				}
			case lex.data[(lex.p)] >= 91:
				goto tr783
			}
		default:
			goto tr783
		}
		goto st521
	tr163:
		// line internal/php8/scanner.rl:467
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st522
	tr164:
		// line internal/php8/scanner.rl:463
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 522
				goto _out
			}
		}
		goto st522
	tr166:
		// line internal/php8/scanner.rl:464
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_NULLSAFE_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 522
				goto _out
			}
		}
		goto st522
	tr784:
		// line internal/php8/scanner.rl:467
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st522
	tr789:
		// line internal/php8/scanner.rl:466
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('['))
			lex.call(522, 528)
			goto _out
		}
		goto st522
	tr790:
		// line internal/php8/scanner.rl:467
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				lex.top--
				lex.cs = lex.stack[lex.top]
				goto _again
			}
		}
		goto st522
	tr792:
		// line internal/php8/scanner.rl:462
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 522
				goto _out
			}
		}
		goto st522
	tr795:
		// line internal/php8/scanner.rl:465
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 522
				goto _out
			}
		}
		goto st522
	st522:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof522
		}
	st_case_522:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:22434
		switch lex.data[(lex.p)] {
		case 36:
			goto st523
		case 45:
			goto tr786
		case 63:
			goto tr787
		case 91:
			goto tr789
		case 96:
			goto tr784
		}
		switch {
		case lex.data[(lex.p)] < 92:
			if lex.data[(lex.p)] <= 64 {
				goto tr784
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr784
			}
		default:
			goto tr784
		}
		goto st527
	st523:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof523
		}
	st_case_523:
		if lex.data[(lex.p)] == 96 {
			goto tr790
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr790
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr790
			}
		default:
			goto tr790
		}
		goto st524
	st524:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof524
		}
	st_case_524:
		if lex.data[(lex.p)] == 96 {
			goto tr792
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr792
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr792
				}
			case lex.data[(lex.p)] >= 91:
				goto tr792
			}
		default:
			goto tr792
		}
		goto st524
	tr786:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st525
	st525:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof525
		}
	st_case_525:
		// line internal/php8/scanner.go:22517
		if lex.data[(lex.p)] == 62 {
			goto st111
		}
		goto tr790
	st111:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof111
		}
	st_case_111:
		if lex.data[(lex.p)] == 96 {
			goto tr163
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr163
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr163
			}
		default:
			goto tr163
		}
		goto tr164
	tr787:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st526
	st526:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof526
		}
	st_case_526:
		// line internal/php8/scanner.go:22553
		if lex.data[(lex.p)] == 45 {
			goto st112
		}
		goto tr790
	st112:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		if lex.data[(lex.p)] == 62 {
			goto st113
		}
		goto tr163
	st113:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		if lex.data[(lex.p)] == 96 {
			goto tr163
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr163
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr163
			}
		default:
			goto tr163
		}
		goto tr166
	st527:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof527
		}
	st_case_527:
		if lex.data[(lex.p)] == 96 {
			goto tr795
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr795
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr795
				}
			case lex.data[(lex.p)] >= 91:
				goto tr795
			}
		default:
			goto tr795
		}
		goto st527
	tr167:
		// line internal/php8/scanner.rl:471
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 528
				goto _out
			}
		}
		goto st528
	tr796:
		// line internal/php8/scanner.rl:477
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st528
	tr797:
		// line internal/php8/scanner.rl:474
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st528
	tr800:
		// line internal/php8/scanner.rl:475
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 528
				goto _out
			}
		}
		goto st528
	tr804:
		// line internal/php8/scanner.rl:476
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(']'))
			lex.ret(2)
			goto _out
		}
		goto st528
	tr805:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:474
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_ENCAPSED_AND_WHITESPACE
			lex.ret(2)
			goto _out
		}
		goto st528
	tr806:
		// line internal/php8/scanner.rl:477
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st528
	tr807:
		// line internal/php8/scanner.rl:475
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 528
				goto _out
			}
		}
		goto st528
	tr809:
		// line internal/php8/scanner.rl:472
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_VARIABLE
			{
				(lex.p)++
				lex.cs = 528
				goto _out
			}
		}
		goto st528
	tr810:
		// line internal/php8/scanner.rl:471
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_NUM_STRING
			{
				(lex.p)++
				lex.cs = 528
				goto _out
			}
		}
		goto st528
	tr814:
		// line internal/php8/scanner.rl:473
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPosition(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 528
				goto _out
			}
		}
		goto st528
	st528:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof528
		}
	st_case_528:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:22694
		switch lex.data[(lex.p)] {
		case 10:
			goto tr798
		case 13:
			goto tr799
		case 32:
			goto tr797
		case 33:
			goto tr800
		case 35:
			goto tr797
		case 36:
			goto st531
		case 39:
			goto tr797
		case 48:
			goto tr802
		case 92:
			goto tr797
		case 93:
			goto tr804
		case 96:
			goto tr796
		case 124:
			goto tr800
		case 126:
			goto tr800
		}
		switch {
		case lex.data[(lex.p)] < 37:
			switch {
			case lex.data[(lex.p)] < 9:
				if lex.data[(lex.p)] <= 8 {
					goto tr796
				}
			case lex.data[(lex.p)] > 12:
				if 14 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 34 {
					goto tr796
				}
			default:
				goto tr797
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 58:
				if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
					goto tr168
				}
			case lex.data[(lex.p)] > 64:
				switch {
				case lex.data[(lex.p)] > 94:
					if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
						goto tr796
					}
				case lex.data[(lex.p)] >= 91:
					goto tr800
				}
			default:
				goto tr800
			}
		default:
			goto tr800
		}
		goto st537
	tr798:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st529
	st529:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof529
		}
	st_case_529:
		// line internal/php8/scanner.go:22776
		goto tr805
	tr799:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st530
	st530:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof530
		}
	st_case_530:
		// line internal/php8/scanner.go:22795
		if lex.data[(lex.p)] == 10 {
			goto tr798
		}
		goto tr806
	st531:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof531
		}
	st_case_531:
		if lex.data[(lex.p)] == 96 {
			goto tr807
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr807
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr807
			}
		default:
			goto tr807
		}
		goto st532
	st532:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof532
		}
	st_case_532:
		if lex.data[(lex.p)] == 96 {
			goto tr809
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr809
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr809
				}
			case lex.data[(lex.p)] >= 91:
				goto tr809
			}
		default:
			goto tr809
		}
		goto st532
	tr802:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st533
	st533:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof533
		}
	st_case_533:
		// line internal/php8/scanner.go:22857
		switch lex.data[(lex.p)] {
		case 95:
			goto st114
		case 98:
			goto st115
		case 120:
			goto st116
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr168
		}
		goto tr810
	tr168:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st534
	st534:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof534
		}
	st_case_534:
		// line internal/php8/scanner.go:22880
		if lex.data[(lex.p)] == 95 {
			goto st114
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr168
		}
		goto tr810
	st114:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof114
		}
	st_case_114:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr168
		}
		goto tr167
	st115:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr169
		}
		goto tr167
	tr169:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st535
	st535:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof535
		}
	st_case_535:
		// line internal/php8/scanner.go:22916
		if lex.data[(lex.p)] == 95 {
			goto st115
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr169
		}
		goto tr810
	st116:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof116
		}
	st_case_116:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr170
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr170
			}
		default:
			goto tr170
		}
		goto tr167
	tr170:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st536
	st536:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof536
		}
	st_case_536:
		// line internal/php8/scanner.go:22952
		if lex.data[(lex.p)] == 95 {
			goto st116
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr170
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr170
			}
		default:
			goto tr170
		}
		goto tr810
	st537:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof537
		}
	st_case_537:
		if lex.data[(lex.p)] == 96 {
			goto tr814
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr814
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr814
				}
			case lex.data[(lex.p)] >= 91:
				goto tr814
			}
		default:
			goto tr814
		}
		goto st537
	tr171:
		lex.cs = 538
		// line internal/php8/scanner.rl:485
		(lex.p) = (lex.te) - 1
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	tr173:
		lex.cs = 538
		// line internal/php8/scanner.rl:484
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.setTokenPosition(tkn)
			tok = token.T_STRING_VARNAME
			lex.cs = 131
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr815:
		lex.cs = 538
		// line internal/php8/scanner.rl:485
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	tr817:
		lex.cs = 538
		// line internal/php8/scanner.rl:485
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	st538:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof538
		}
	st_case_538:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23031
		if lex.data[(lex.p)] == 96 {
			goto tr815
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto tr815
			}
		case lex.data[(lex.p)] > 94:
			if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto tr815
			}
		default:
			goto tr815
		}
		goto tr816
	tr816:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st539
	st539:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof539
		}
	st_case_539:
		// line internal/php8/scanner.go:23058
		switch lex.data[(lex.p)] {
		case 91:
			goto tr173
		case 96:
			goto tr817
		case 125:
			goto tr173
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr817
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr817
				}
			case lex.data[(lex.p)] >= 92:
				goto tr817
			}
		default:
			goto tr817
		}
		goto st117
	st117:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		switch lex.data[(lex.p)] {
		case 91:
			goto tr173
		case 96:
			goto tr171
		case 125:
			goto tr173
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr171
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr171
				}
			case lex.data[(lex.p)] >= 92:
				goto tr171
			}
		default:
			goto tr171
		}
		goto st117
	tr174:
		// line internal/php8/scanner.rl:489
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st540
	tr818:
		lex.cs = 540
		// line internal/php8/scanner.rl:491
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	tr821:
		lex.cs = 540
		// line internal/php8/scanner.rl:490
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int('('))
			lex.cs = 544
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr822:
		// line internal/php8/scanner.rl:489
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st540
	tr824:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:489
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st540
	tr828:
		lex.cs = 540
		// line internal/php8/scanner.rl:491
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	st540:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof540
		}
	st_case_540:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23165
		switch lex.data[(lex.p)] {
		case 10:
			goto tr175
		case 13:
			goto tr820
		case 32:
			goto tr819
		case 40:
			goto tr821
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr819
		}
		goto tr818
	tr819:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st541
	tr825:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		goto st541
	st541:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof541
		}
	st_case_541:
		// line internal/php8/scanner.go:23197
		switch lex.data[(lex.p)] {
		case 10:
			goto tr175
		case 13:
			goto tr823
		case 32:
			goto tr819
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr819
		}
		goto tr822
	tr175:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st542
	tr826:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st542
	st542:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof542
		}
	st_case_542:
		// line internal/php8/scanner.go:23247
		switch lex.data[(lex.p)] {
		case 10:
			goto tr826
		case 13:
			goto tr827
		case 32:
			goto tr825
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr825
		}
		goto tr824
	tr823:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st118
	tr827:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st118
	st118:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof118
		}
	st_case_118:
		// line internal/php8/scanner.go:23291
		if lex.data[(lex.p)] == 10 {
			goto tr175
		}
		goto tr174
	tr820:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st543
	st543:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof543
		}
	st_case_543:
		// line internal/php8/scanner.go:23313
		if lex.data[(lex.p)] == 10 {
			goto tr175
		}
		goto tr828
	tr176:
		// line internal/php8/scanner.rl:495
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st544
	tr829:
		lex.cs = 544
		// line internal/php8/scanner.rl:497
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	tr832:
		lex.cs = 544
		// line internal/php8/scanner.rl:496
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(')'))
			lex.cs = 548
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr833:
		// line internal/php8/scanner.rl:495
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st544
	tr835:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:495
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st544
	tr839:
		lex.cs = 544
		// line internal/php8/scanner.rl:497
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	st544:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof544
		}
	st_case_544:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23367
		switch lex.data[(lex.p)] {
		case 10:
			goto tr177
		case 13:
			goto tr831
		case 32:
			goto tr830
		case 41:
			goto tr832
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr830
		}
		goto tr829
	tr830:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st545
	tr836:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		goto st545
	st545:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof545
		}
	st_case_545:
		// line internal/php8/scanner.go:23399
		switch lex.data[(lex.p)] {
		case 10:
			goto tr177
		case 13:
			goto tr834
		case 32:
			goto tr830
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr830
		}
		goto tr833
	tr177:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st546
	tr837:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st546
	st546:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof546
		}
	st_case_546:
		// line internal/php8/scanner.go:23449
		switch lex.data[(lex.p)] {
		case 10:
			goto tr837
		case 13:
			goto tr838
		case 32:
			goto tr836
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr836
		}
		goto tr835
	tr834:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st119
	tr838:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st119
	st119:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof119
		}
	st_case_119:
		// line internal/php8/scanner.go:23493
		if lex.data[(lex.p)] == 10 {
			goto tr177
		}
		goto tr176
	tr831:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st547
	st547:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof547
		}
	st_case_547:
		// line internal/php8/scanner.go:23515
		if lex.data[(lex.p)] == 10 {
			goto tr177
		}
		goto tr839
	tr178:
		// line internal/php8/scanner.rl:501
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st548
	tr840:
		lex.cs = 548
		// line internal/php8/scanner.rl:503
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	tr843:
		lex.cs = 548
		// line internal/php8/scanner.rl:502
		lex.te = (lex.p) + 1
		{
			lex.setTokenPosition(tkn)
			tok = token.ID(int(';'))
			lex.cs = 552
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr844:
		// line internal/php8/scanner.rl:501
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st548
	tr846:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:501
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st548
	tr850:
		lex.cs = 548
		// line internal/php8/scanner.rl:503
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			lex.cs = 131
		}
		goto _again
	st548:
		// line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof548
		}
	st_case_548:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23569
		switch lex.data[(lex.p)] {
		case 10:
			goto tr179
		case 13:
			goto tr842
		case 32:
			goto tr841
		case 59:
			goto tr843
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr841
		}
		goto tr840
	tr841:
		// line NONE:1
		lex.te = (lex.p) + 1

		goto st549
	tr847:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		goto st549
	st549:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof549
		}
	st_case_549:
		// line internal/php8/scanner.go:23601
		switch lex.data[(lex.p)] {
		case 10:
			goto tr179
		case 13:
			goto tr845
		case 32:
			goto tr841
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr841
		}
		goto tr844
	tr179:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st550
	tr848:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st550
	st550:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof550
		}
	st_case_550:
		// line internal/php8/scanner.go:23651
		switch lex.data[(lex.p)] {
		case 10:
			goto tr848
		case 13:
			goto tr849
		case 32:
			goto tr847
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr847
		}
		goto tr846
	tr845:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st120
	tr849:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st120
	st120:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof120
		}
	st_case_120:
		// line internal/php8/scanner.go:23695
		if lex.data[(lex.p)] == 10 {
			goto tr179
		}
		goto tr178
	tr842:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st551
	st551:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof551
		}
	st_case_551:
		// line internal/php8/scanner.go:23717
		if lex.data[(lex.p)] == 10 {
			goto tr179
		}
		goto tr850
	tr854:
		// line NONE:1
		switch lex.act {
		case 0:
			{
				{
					goto st0
				}
			}
		case 197:
			{
				(lex.p) = (lex.te) - 1
				lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
			}
		}

		goto st552
	tr855:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:507
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
		}
		goto st552
	tr859:
		// line internal/php8/scanner.rl:507
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_HALT_COMPILER, lex.ts, lex.te)
		}
		goto st552
	st552:
		// line NONE:1
		lex.ts = 0

		// line NONE:1
		lex.act = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof552
		}
	st_case_552:
		// line NONE:1
		lex.ts = (lex.p)

		// line internal/php8/scanner.go:23761
		switch lex.data[(lex.p)] {
		case 10:
			goto tr852
		case 13:
			goto tr853
		}
		goto tr851
	tr851:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:507
		lex.act = 197
		goto st553
	tr856:
		// line NONE:1
		lex.te = (lex.p) + 1

		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:507
		lex.act = 197
		goto st553
	st553:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof553
		}
	st_case_553:
		// line internal/php8/scanner.go:23790
		switch lex.data[(lex.p)] {
		case 10:
			goto tr852
		case 13:
			goto tr853
		}
		goto tr851
	tr852:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st554
	tr857:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st554
	st554:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof554
		}
	st_case_554:
		// line internal/php8/scanner.go:23829
		switch lex.data[(lex.p)] {
		case 10:
			goto tr857
		case 13:
			goto tr858
		}
		goto tr856
	tr853:
		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st555
	tr858:
		// line internal/php8/scanner.rl:54

		// line internal/php8/scanner.rl:38

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st555
	st555:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof555
		}
	st_case_555:
		// line internal/php8/scanner.go:23868
		switch lex.data[(lex.p)] {
		case 10:
			goto tr852
		case 13:
			goto tr853
		}
		goto tr851
	st_out:
	_test_eof121:
		lex.cs = 121
		goto _test_eof
	_test_eof122:
		lex.cs = 122
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof123:
		lex.cs = 123
		goto _test_eof
	_test_eof124:
		lex.cs = 124
		goto _test_eof
	_test_eof125:
		lex.cs = 125
		goto _test_eof
	_test_eof126:
		lex.cs = 126
		goto _test_eof
	_test_eof127:
		lex.cs = 127
		goto _test_eof
	_test_eof128:
		lex.cs = 128
		goto _test_eof
	_test_eof129:
		lex.cs = 129
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof130:
		lex.cs = 130
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof131:
		lex.cs = 131
		goto _test_eof
	_test_eof132:
		lex.cs = 132
		goto _test_eof
	_test_eof133:
		lex.cs = 133
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof134:
		lex.cs = 134
		goto _test_eof
	_test_eof135:
		lex.cs = 135
		goto _test_eof
	_test_eof136:
		lex.cs = 136
		goto _test_eof
	_test_eof137:
		lex.cs = 137
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof138:
		lex.cs = 138
		goto _test_eof
	_test_eof139:
		lex.cs = 139
		goto _test_eof
	_test_eof140:
		lex.cs = 140
		goto _test_eof
	_test_eof141:
		lex.cs = 141
		goto _test_eof
	_test_eof142:
		lex.cs = 142
		goto _test_eof
	_test_eof143:
		lex.cs = 143
		goto _test_eof
	_test_eof144:
		lex.cs = 144
		goto _test_eof
	_test_eof145:
		lex.cs = 145
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof146:
		lex.cs = 146
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof18:
		lex.cs = 18
		goto _test_eof
	_test_eof19:
		lex.cs = 19
		goto _test_eof
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof22:
		lex.cs = 22
		goto _test_eof
	_test_eof23:
		lex.cs = 23
		goto _test_eof
	_test_eof24:
		lex.cs = 24
		goto _test_eof
	_test_eof25:
		lex.cs = 25
		goto _test_eof
	_test_eof26:
		lex.cs = 26
		goto _test_eof
	_test_eof27:
		lex.cs = 27
		goto _test_eof
	_test_eof28:
		lex.cs = 28
		goto _test_eof
	_test_eof29:
		lex.cs = 29
		goto _test_eof
	_test_eof30:
		lex.cs = 30
		goto _test_eof
	_test_eof31:
		lex.cs = 31
		goto _test_eof
	_test_eof32:
		lex.cs = 32
		goto _test_eof
	_test_eof33:
		lex.cs = 33
		goto _test_eof
	_test_eof34:
		lex.cs = 34
		goto _test_eof
	_test_eof35:
		lex.cs = 35
		goto _test_eof
	_test_eof36:
		lex.cs = 36
		goto _test_eof
	_test_eof37:
		lex.cs = 37
		goto _test_eof
	_test_eof38:
		lex.cs = 38
		goto _test_eof
	_test_eof39:
		lex.cs = 39
		goto _test_eof
	_test_eof40:
		lex.cs = 40
		goto _test_eof
	_test_eof41:
		lex.cs = 41
		goto _test_eof
	_test_eof42:
		lex.cs = 42
		goto _test_eof
	_test_eof43:
		lex.cs = 43
		goto _test_eof
	_test_eof44:
		lex.cs = 44
		goto _test_eof
	_test_eof45:
		lex.cs = 45
		goto _test_eof
	_test_eof46:
		lex.cs = 46
		goto _test_eof
	_test_eof47:
		lex.cs = 47
		goto _test_eof
	_test_eof48:
		lex.cs = 48
		goto _test_eof
	_test_eof49:
		lex.cs = 49
		goto _test_eof
	_test_eof50:
		lex.cs = 50
		goto _test_eof
	_test_eof51:
		lex.cs = 51
		goto _test_eof
	_test_eof52:
		lex.cs = 52
		goto _test_eof
	_test_eof53:
		lex.cs = 53
		goto _test_eof
	_test_eof54:
		lex.cs = 54
		goto _test_eof
	_test_eof55:
		lex.cs = 55
		goto _test_eof
	_test_eof56:
		lex.cs = 56
		goto _test_eof
	_test_eof57:
		lex.cs = 57
		goto _test_eof
	_test_eof58:
		lex.cs = 58
		goto _test_eof
	_test_eof59:
		lex.cs = 59
		goto _test_eof
	_test_eof60:
		lex.cs = 60
		goto _test_eof
	_test_eof61:
		lex.cs = 61
		goto _test_eof
	_test_eof62:
		lex.cs = 62
		goto _test_eof
	_test_eof63:
		lex.cs = 63
		goto _test_eof
	_test_eof64:
		lex.cs = 64
		goto _test_eof
	_test_eof65:
		lex.cs = 65
		goto _test_eof
	_test_eof66:
		lex.cs = 66
		goto _test_eof
	_test_eof67:
		lex.cs = 67
		goto _test_eof
	_test_eof147:
		lex.cs = 147
		goto _test_eof
	_test_eof148:
		lex.cs = 148
		goto _test_eof
	_test_eof149:
		lex.cs = 149
		goto _test_eof
	_test_eof150:
		lex.cs = 150
		goto _test_eof
	_test_eof151:
		lex.cs = 151
		goto _test_eof
	_test_eof68:
		lex.cs = 68
		goto _test_eof
	_test_eof152:
		lex.cs = 152
		goto _test_eof
	_test_eof69:
		lex.cs = 69
		goto _test_eof
	_test_eof70:
		lex.cs = 70
		goto _test_eof
	_test_eof153:
		lex.cs = 153
		goto _test_eof
	_test_eof71:
		lex.cs = 71
		goto _test_eof
	_test_eof154:
		lex.cs = 154
		goto _test_eof
	_test_eof72:
		lex.cs = 72
		goto _test_eof
	_test_eof73:
		lex.cs = 73
		goto _test_eof
	_test_eof74:
		lex.cs = 74
		goto _test_eof
	_test_eof155:
		lex.cs = 155
		goto _test_eof
	_test_eof156:
		lex.cs = 156
		goto _test_eof
	_test_eof157:
		lex.cs = 157
		goto _test_eof
	_test_eof75:
		lex.cs = 75
		goto _test_eof
	_test_eof76:
		lex.cs = 76
		goto _test_eof
	_test_eof158:
		lex.cs = 158
		goto _test_eof
	_test_eof77:
		lex.cs = 77
		goto _test_eof
	_test_eof159:
		lex.cs = 159
		goto _test_eof
	_test_eof160:
		lex.cs = 160
		goto _test_eof
	_test_eof161:
		lex.cs = 161
		goto _test_eof
	_test_eof78:
		lex.cs = 78
		goto _test_eof
	_test_eof79:
		lex.cs = 79
		goto _test_eof
	_test_eof80:
		lex.cs = 80
		goto _test_eof
	_test_eof81:
		lex.cs = 81
		goto _test_eof
	_test_eof162:
		lex.cs = 162
		goto _test_eof
	_test_eof163:
		lex.cs = 163
		goto _test_eof
	_test_eof82:
		lex.cs = 82
		goto _test_eof
	_test_eof164:
		lex.cs = 164
		goto _test_eof
	_test_eof165:
		lex.cs = 165
		goto _test_eof
	_test_eof83:
		lex.cs = 83
		goto _test_eof
	_test_eof84:
		lex.cs = 84
		goto _test_eof
	_test_eof85:
		lex.cs = 85
		goto _test_eof
	_test_eof86:
		lex.cs = 86
		goto _test_eof
	_test_eof166:
		lex.cs = 166
		goto _test_eof
	_test_eof87:
		lex.cs = 87
		goto _test_eof
	_test_eof88:
		lex.cs = 88
		goto _test_eof
	_test_eof89:
		lex.cs = 89
		goto _test_eof
	_test_eof90:
		lex.cs = 90
		goto _test_eof
	_test_eof167:
		lex.cs = 167
		goto _test_eof
	_test_eof168:
		lex.cs = 168
		goto _test_eof
	_test_eof169:
		lex.cs = 169
		goto _test_eof
	_test_eof170:
		lex.cs = 170
		goto _test_eof
	_test_eof171:
		lex.cs = 171
		goto _test_eof
	_test_eof172:
		lex.cs = 172
		goto _test_eof
	_test_eof91:
		lex.cs = 91
		goto _test_eof
	_test_eof173:
		lex.cs = 173
		goto _test_eof
	_test_eof174:
		lex.cs = 174
		goto _test_eof
	_test_eof92:
		lex.cs = 92
		goto _test_eof
	_test_eof175:
		lex.cs = 175
		goto _test_eof
	_test_eof176:
		lex.cs = 176
		goto _test_eof
	_test_eof177:
		lex.cs = 177
		goto _test_eof
	_test_eof93:
		lex.cs = 93
		goto _test_eof
	_test_eof178:
		lex.cs = 178
		goto _test_eof
	_test_eof179:
		lex.cs = 179
		goto _test_eof
	_test_eof180:
		lex.cs = 180
		goto _test_eof
	_test_eof181:
		lex.cs = 181
		goto _test_eof
	_test_eof182:
		lex.cs = 182
		goto _test_eof
	_test_eof183:
		lex.cs = 183
		goto _test_eof
	_test_eof184:
		lex.cs = 184
		goto _test_eof
	_test_eof185:
		lex.cs = 185
		goto _test_eof
	_test_eof186:
		lex.cs = 186
		goto _test_eof
	_test_eof187:
		lex.cs = 187
		goto _test_eof
	_test_eof188:
		lex.cs = 188
		goto _test_eof
	_test_eof189:
		lex.cs = 189
		goto _test_eof
	_test_eof94:
		lex.cs = 94
		goto _test_eof
	_test_eof95:
		lex.cs = 95
		goto _test_eof
	_test_eof190:
		lex.cs = 190
		goto _test_eof
	_test_eof191:
		lex.cs = 191
		goto _test_eof
	_test_eof192:
		lex.cs = 192
		goto _test_eof
	_test_eof193:
		lex.cs = 193
		goto _test_eof
	_test_eof194:
		lex.cs = 194
		goto _test_eof
	_test_eof195:
		lex.cs = 195
		goto _test_eof
	_test_eof196:
		lex.cs = 196
		goto _test_eof
	_test_eof197:
		lex.cs = 197
		goto _test_eof
	_test_eof198:
		lex.cs = 198
		goto _test_eof
	_test_eof199:
		lex.cs = 199
		goto _test_eof
	_test_eof200:
		lex.cs = 200
		goto _test_eof
	_test_eof201:
		lex.cs = 201
		goto _test_eof
	_test_eof202:
		lex.cs = 202
		goto _test_eof
	_test_eof203:
		lex.cs = 203
		goto _test_eof
	_test_eof204:
		lex.cs = 204
		goto _test_eof
	_test_eof205:
		lex.cs = 205
		goto _test_eof
	_test_eof206:
		lex.cs = 206
		goto _test_eof
	_test_eof207:
		lex.cs = 207
		goto _test_eof
	_test_eof208:
		lex.cs = 208
		goto _test_eof
	_test_eof209:
		lex.cs = 209
		goto _test_eof
	_test_eof210:
		lex.cs = 210
		goto _test_eof
	_test_eof211:
		lex.cs = 211
		goto _test_eof
	_test_eof212:
		lex.cs = 212
		goto _test_eof
	_test_eof213:
		lex.cs = 213
		goto _test_eof
	_test_eof214:
		lex.cs = 214
		goto _test_eof
	_test_eof215:
		lex.cs = 215
		goto _test_eof
	_test_eof216:
		lex.cs = 216
		goto _test_eof
	_test_eof217:
		lex.cs = 217
		goto _test_eof
	_test_eof218:
		lex.cs = 218
		goto _test_eof
	_test_eof219:
		lex.cs = 219
		goto _test_eof
	_test_eof220:
		lex.cs = 220
		goto _test_eof
	_test_eof221:
		lex.cs = 221
		goto _test_eof
	_test_eof222:
		lex.cs = 222
		goto _test_eof
	_test_eof223:
		lex.cs = 223
		goto _test_eof
	_test_eof224:
		lex.cs = 224
		goto _test_eof
	_test_eof225:
		lex.cs = 225
		goto _test_eof
	_test_eof226:
		lex.cs = 226
		goto _test_eof
	_test_eof227:
		lex.cs = 227
		goto _test_eof
	_test_eof228:
		lex.cs = 228
		goto _test_eof
	_test_eof229:
		lex.cs = 229
		goto _test_eof
	_test_eof230:
		lex.cs = 230
		goto _test_eof
	_test_eof231:
		lex.cs = 231
		goto _test_eof
	_test_eof232:
		lex.cs = 232
		goto _test_eof
	_test_eof233:
		lex.cs = 233
		goto _test_eof
	_test_eof234:
		lex.cs = 234
		goto _test_eof
	_test_eof235:
		lex.cs = 235
		goto _test_eof
	_test_eof236:
		lex.cs = 236
		goto _test_eof
	_test_eof237:
		lex.cs = 237
		goto _test_eof
	_test_eof238:
		lex.cs = 238
		goto _test_eof
	_test_eof239:
		lex.cs = 239
		goto _test_eof
	_test_eof240:
		lex.cs = 240
		goto _test_eof
	_test_eof241:
		lex.cs = 241
		goto _test_eof
	_test_eof242:
		lex.cs = 242
		goto _test_eof
	_test_eof243:
		lex.cs = 243
		goto _test_eof
	_test_eof244:
		lex.cs = 244
		goto _test_eof
	_test_eof245:
		lex.cs = 245
		goto _test_eof
	_test_eof246:
		lex.cs = 246
		goto _test_eof
	_test_eof247:
		lex.cs = 247
		goto _test_eof
	_test_eof248:
		lex.cs = 248
		goto _test_eof
	_test_eof249:
		lex.cs = 249
		goto _test_eof
	_test_eof250:
		lex.cs = 250
		goto _test_eof
	_test_eof251:
		lex.cs = 251
		goto _test_eof
	_test_eof252:
		lex.cs = 252
		goto _test_eof
	_test_eof253:
		lex.cs = 253
		goto _test_eof
	_test_eof254:
		lex.cs = 254
		goto _test_eof
	_test_eof255:
		lex.cs = 255
		goto _test_eof
	_test_eof256:
		lex.cs = 256
		goto _test_eof
	_test_eof257:
		lex.cs = 257
		goto _test_eof
	_test_eof258:
		lex.cs = 258
		goto _test_eof
	_test_eof259:
		lex.cs = 259
		goto _test_eof
	_test_eof260:
		lex.cs = 260
		goto _test_eof
	_test_eof261:
		lex.cs = 261
		goto _test_eof
	_test_eof262:
		lex.cs = 262
		goto _test_eof
	_test_eof263:
		lex.cs = 263
		goto _test_eof
	_test_eof264:
		lex.cs = 264
		goto _test_eof
	_test_eof265:
		lex.cs = 265
		goto _test_eof
	_test_eof266:
		lex.cs = 266
		goto _test_eof
	_test_eof267:
		lex.cs = 267
		goto _test_eof
	_test_eof268:
		lex.cs = 268
		goto _test_eof
	_test_eof269:
		lex.cs = 269
		goto _test_eof
	_test_eof270:
		lex.cs = 270
		goto _test_eof
	_test_eof271:
		lex.cs = 271
		goto _test_eof
	_test_eof272:
		lex.cs = 272
		goto _test_eof
	_test_eof273:
		lex.cs = 273
		goto _test_eof
	_test_eof274:
		lex.cs = 274
		goto _test_eof
	_test_eof275:
		lex.cs = 275
		goto _test_eof
	_test_eof276:
		lex.cs = 276
		goto _test_eof
	_test_eof277:
		lex.cs = 277
		goto _test_eof
	_test_eof278:
		lex.cs = 278
		goto _test_eof
	_test_eof279:
		lex.cs = 279
		goto _test_eof
	_test_eof280:
		lex.cs = 280
		goto _test_eof
	_test_eof281:
		lex.cs = 281
		goto _test_eof
	_test_eof282:
		lex.cs = 282
		goto _test_eof
	_test_eof283:
		lex.cs = 283
		goto _test_eof
	_test_eof284:
		lex.cs = 284
		goto _test_eof
	_test_eof285:
		lex.cs = 285
		goto _test_eof
	_test_eof286:
		lex.cs = 286
		goto _test_eof
	_test_eof287:
		lex.cs = 287
		goto _test_eof
	_test_eof288:
		lex.cs = 288
		goto _test_eof
	_test_eof289:
		lex.cs = 289
		goto _test_eof
	_test_eof290:
		lex.cs = 290
		goto _test_eof
	_test_eof291:
		lex.cs = 291
		goto _test_eof
	_test_eof292:
		lex.cs = 292
		goto _test_eof
	_test_eof293:
		lex.cs = 293
		goto _test_eof
	_test_eof294:
		lex.cs = 294
		goto _test_eof
	_test_eof295:
		lex.cs = 295
		goto _test_eof
	_test_eof296:
		lex.cs = 296
		goto _test_eof
	_test_eof297:
		lex.cs = 297
		goto _test_eof
	_test_eof298:
		lex.cs = 298
		goto _test_eof
	_test_eof299:
		lex.cs = 299
		goto _test_eof
	_test_eof300:
		lex.cs = 300
		goto _test_eof
	_test_eof301:
		lex.cs = 301
		goto _test_eof
	_test_eof302:
		lex.cs = 302
		goto _test_eof
	_test_eof303:
		lex.cs = 303
		goto _test_eof
	_test_eof304:
		lex.cs = 304
		goto _test_eof
	_test_eof305:
		lex.cs = 305
		goto _test_eof
	_test_eof306:
		lex.cs = 306
		goto _test_eof
	_test_eof307:
		lex.cs = 307
		goto _test_eof
	_test_eof308:
		lex.cs = 308
		goto _test_eof
	_test_eof309:
		lex.cs = 309
		goto _test_eof
	_test_eof310:
		lex.cs = 310
		goto _test_eof
	_test_eof311:
		lex.cs = 311
		goto _test_eof
	_test_eof312:
		lex.cs = 312
		goto _test_eof
	_test_eof313:
		lex.cs = 313
		goto _test_eof
	_test_eof314:
		lex.cs = 314
		goto _test_eof
	_test_eof315:
		lex.cs = 315
		goto _test_eof
	_test_eof316:
		lex.cs = 316
		goto _test_eof
	_test_eof317:
		lex.cs = 317
		goto _test_eof
	_test_eof318:
		lex.cs = 318
		goto _test_eof
	_test_eof319:
		lex.cs = 319
		goto _test_eof
	_test_eof320:
		lex.cs = 320
		goto _test_eof
	_test_eof321:
		lex.cs = 321
		goto _test_eof
	_test_eof322:
		lex.cs = 322
		goto _test_eof
	_test_eof323:
		lex.cs = 323
		goto _test_eof
	_test_eof324:
		lex.cs = 324
		goto _test_eof
	_test_eof325:
		lex.cs = 325
		goto _test_eof
	_test_eof326:
		lex.cs = 326
		goto _test_eof
	_test_eof327:
		lex.cs = 327
		goto _test_eof
	_test_eof328:
		lex.cs = 328
		goto _test_eof
	_test_eof329:
		lex.cs = 329
		goto _test_eof
	_test_eof330:
		lex.cs = 330
		goto _test_eof
	_test_eof331:
		lex.cs = 331
		goto _test_eof
	_test_eof332:
		lex.cs = 332
		goto _test_eof
	_test_eof333:
		lex.cs = 333
		goto _test_eof
	_test_eof334:
		lex.cs = 334
		goto _test_eof
	_test_eof335:
		lex.cs = 335
		goto _test_eof
	_test_eof336:
		lex.cs = 336
		goto _test_eof
	_test_eof337:
		lex.cs = 337
		goto _test_eof
	_test_eof338:
		lex.cs = 338
		goto _test_eof
	_test_eof339:
		lex.cs = 339
		goto _test_eof
	_test_eof340:
		lex.cs = 340
		goto _test_eof
	_test_eof341:
		lex.cs = 341
		goto _test_eof
	_test_eof342:
		lex.cs = 342
		goto _test_eof
	_test_eof343:
		lex.cs = 343
		goto _test_eof
	_test_eof344:
		lex.cs = 344
		goto _test_eof
	_test_eof345:
		lex.cs = 345
		goto _test_eof
	_test_eof346:
		lex.cs = 346
		goto _test_eof
	_test_eof347:
		lex.cs = 347
		goto _test_eof
	_test_eof96:
		lex.cs = 96
		goto _test_eof
	_test_eof348:
		lex.cs = 348
		goto _test_eof
	_test_eof349:
		lex.cs = 349
		goto _test_eof
	_test_eof350:
		lex.cs = 350
		goto _test_eof
	_test_eof351:
		lex.cs = 351
		goto _test_eof
	_test_eof352:
		lex.cs = 352
		goto _test_eof
	_test_eof353:
		lex.cs = 353
		goto _test_eof
	_test_eof354:
		lex.cs = 354
		goto _test_eof
	_test_eof355:
		lex.cs = 355
		goto _test_eof
	_test_eof356:
		lex.cs = 356
		goto _test_eof
	_test_eof357:
		lex.cs = 357
		goto _test_eof
	_test_eof358:
		lex.cs = 358
		goto _test_eof
	_test_eof359:
		lex.cs = 359
		goto _test_eof
	_test_eof360:
		lex.cs = 360
		goto _test_eof
	_test_eof361:
		lex.cs = 361
		goto _test_eof
	_test_eof362:
		lex.cs = 362
		goto _test_eof
	_test_eof363:
		lex.cs = 363
		goto _test_eof
	_test_eof364:
		lex.cs = 364
		goto _test_eof
	_test_eof365:
		lex.cs = 365
		goto _test_eof
	_test_eof366:
		lex.cs = 366
		goto _test_eof
	_test_eof367:
		lex.cs = 367
		goto _test_eof
	_test_eof368:
		lex.cs = 368
		goto _test_eof
	_test_eof369:
		lex.cs = 369
		goto _test_eof
	_test_eof370:
		lex.cs = 370
		goto _test_eof
	_test_eof371:
		lex.cs = 371
		goto _test_eof
	_test_eof372:
		lex.cs = 372
		goto _test_eof
	_test_eof373:
		lex.cs = 373
		goto _test_eof
	_test_eof374:
		lex.cs = 374
		goto _test_eof
	_test_eof375:
		lex.cs = 375
		goto _test_eof
	_test_eof376:
		lex.cs = 376
		goto _test_eof
	_test_eof377:
		lex.cs = 377
		goto _test_eof
	_test_eof378:
		lex.cs = 378
		goto _test_eof
	_test_eof379:
		lex.cs = 379
		goto _test_eof
	_test_eof380:
		lex.cs = 380
		goto _test_eof
	_test_eof381:
		lex.cs = 381
		goto _test_eof
	_test_eof382:
		lex.cs = 382
		goto _test_eof
	_test_eof383:
		lex.cs = 383
		goto _test_eof
	_test_eof384:
		lex.cs = 384
		goto _test_eof
	_test_eof385:
		lex.cs = 385
		goto _test_eof
	_test_eof386:
		lex.cs = 386
		goto _test_eof
	_test_eof387:
		lex.cs = 387
		goto _test_eof
	_test_eof388:
		lex.cs = 388
		goto _test_eof
	_test_eof389:
		lex.cs = 389
		goto _test_eof
	_test_eof390:
		lex.cs = 390
		goto _test_eof
	_test_eof391:
		lex.cs = 391
		goto _test_eof
	_test_eof392:
		lex.cs = 392
		goto _test_eof
	_test_eof393:
		lex.cs = 393
		goto _test_eof
	_test_eof394:
		lex.cs = 394
		goto _test_eof
	_test_eof395:
		lex.cs = 395
		goto _test_eof
	_test_eof396:
		lex.cs = 396
		goto _test_eof
	_test_eof397:
		lex.cs = 397
		goto _test_eof
	_test_eof398:
		lex.cs = 398
		goto _test_eof
	_test_eof399:
		lex.cs = 399
		goto _test_eof
	_test_eof400:
		lex.cs = 400
		goto _test_eof
	_test_eof401:
		lex.cs = 401
		goto _test_eof
	_test_eof402:
		lex.cs = 402
		goto _test_eof
	_test_eof403:
		lex.cs = 403
		goto _test_eof
	_test_eof404:
		lex.cs = 404
		goto _test_eof
	_test_eof405:
		lex.cs = 405
		goto _test_eof
	_test_eof406:
		lex.cs = 406
		goto _test_eof
	_test_eof407:
		lex.cs = 407
		goto _test_eof
	_test_eof408:
		lex.cs = 408
		goto _test_eof
	_test_eof409:
		lex.cs = 409
		goto _test_eof
	_test_eof410:
		lex.cs = 410
		goto _test_eof
	_test_eof411:
		lex.cs = 411
		goto _test_eof
	_test_eof412:
		lex.cs = 412
		goto _test_eof
	_test_eof413:
		lex.cs = 413
		goto _test_eof
	_test_eof414:
		lex.cs = 414
		goto _test_eof
	_test_eof415:
		lex.cs = 415
		goto _test_eof
	_test_eof416:
		lex.cs = 416
		goto _test_eof
	_test_eof417:
		lex.cs = 417
		goto _test_eof
	_test_eof418:
		lex.cs = 418
		goto _test_eof
	_test_eof419:
		lex.cs = 419
		goto _test_eof
	_test_eof420:
		lex.cs = 420
		goto _test_eof
	_test_eof97:
		lex.cs = 97
		goto _test_eof
	_test_eof98:
		lex.cs = 98
		goto _test_eof
	_test_eof99:
		lex.cs = 99
		goto _test_eof
	_test_eof100:
		lex.cs = 100
		goto _test_eof
	_test_eof101:
		lex.cs = 101
		goto _test_eof
	_test_eof102:
		lex.cs = 102
		goto _test_eof
	_test_eof421:
		lex.cs = 421
		goto _test_eof
	_test_eof422:
		lex.cs = 422
		goto _test_eof
	_test_eof103:
		lex.cs = 103
		goto _test_eof
	_test_eof423:
		lex.cs = 423
		goto _test_eof
	_test_eof424:
		lex.cs = 424
		goto _test_eof
	_test_eof425:
		lex.cs = 425
		goto _test_eof
	_test_eof426:
		lex.cs = 426
		goto _test_eof
	_test_eof427:
		lex.cs = 427
		goto _test_eof
	_test_eof428:
		lex.cs = 428
		goto _test_eof
	_test_eof429:
		lex.cs = 429
		goto _test_eof
	_test_eof430:
		lex.cs = 430
		goto _test_eof
	_test_eof431:
		lex.cs = 431
		goto _test_eof
	_test_eof432:
		lex.cs = 432
		goto _test_eof
	_test_eof433:
		lex.cs = 433
		goto _test_eof
	_test_eof434:
		lex.cs = 434
		goto _test_eof
	_test_eof435:
		lex.cs = 435
		goto _test_eof
	_test_eof436:
		lex.cs = 436
		goto _test_eof
	_test_eof437:
		lex.cs = 437
		goto _test_eof
	_test_eof438:
		lex.cs = 438
		goto _test_eof
	_test_eof439:
		lex.cs = 439
		goto _test_eof
	_test_eof440:
		lex.cs = 440
		goto _test_eof
	_test_eof441:
		lex.cs = 441
		goto _test_eof
	_test_eof442:
		lex.cs = 442
		goto _test_eof
	_test_eof443:
		lex.cs = 443
		goto _test_eof
	_test_eof444:
		lex.cs = 444
		goto _test_eof
	_test_eof445:
		lex.cs = 445
		goto _test_eof
	_test_eof446:
		lex.cs = 446
		goto _test_eof
	_test_eof447:
		lex.cs = 447
		goto _test_eof
	_test_eof448:
		lex.cs = 448
		goto _test_eof
	_test_eof449:
		lex.cs = 449
		goto _test_eof
	_test_eof450:
		lex.cs = 450
		goto _test_eof
	_test_eof451:
		lex.cs = 451
		goto _test_eof
	_test_eof452:
		lex.cs = 452
		goto _test_eof
	_test_eof453:
		lex.cs = 453
		goto _test_eof
	_test_eof454:
		lex.cs = 454
		goto _test_eof
	_test_eof455:
		lex.cs = 455
		goto _test_eof
	_test_eof456:
		lex.cs = 456
		goto _test_eof
	_test_eof457:
		lex.cs = 457
		goto _test_eof
	_test_eof458:
		lex.cs = 458
		goto _test_eof
	_test_eof459:
		lex.cs = 459
		goto _test_eof
	_test_eof460:
		lex.cs = 460
		goto _test_eof
	_test_eof461:
		lex.cs = 461
		goto _test_eof
	_test_eof462:
		lex.cs = 462
		goto _test_eof
	_test_eof463:
		lex.cs = 463
		goto _test_eof
	_test_eof464:
		lex.cs = 464
		goto _test_eof
	_test_eof465:
		lex.cs = 465
		goto _test_eof
	_test_eof466:
		lex.cs = 466
		goto _test_eof
	_test_eof467:
		lex.cs = 467
		goto _test_eof
	_test_eof468:
		lex.cs = 468
		goto _test_eof
	_test_eof469:
		lex.cs = 469
		goto _test_eof
	_test_eof470:
		lex.cs = 470
		goto _test_eof
	_test_eof471:
		lex.cs = 471
		goto _test_eof
	_test_eof472:
		lex.cs = 472
		goto _test_eof
	_test_eof473:
		lex.cs = 473
		goto _test_eof
	_test_eof474:
		lex.cs = 474
		goto _test_eof
	_test_eof475:
		lex.cs = 475
		goto _test_eof
	_test_eof476:
		lex.cs = 476
		goto _test_eof
	_test_eof477:
		lex.cs = 477
		goto _test_eof
	_test_eof478:
		lex.cs = 478
		goto _test_eof
	_test_eof479:
		lex.cs = 479
		goto _test_eof
	_test_eof480:
		lex.cs = 480
		goto _test_eof
	_test_eof481:
		lex.cs = 481
		goto _test_eof
	_test_eof482:
		lex.cs = 482
		goto _test_eof
	_test_eof483:
		lex.cs = 483
		goto _test_eof
	_test_eof484:
		lex.cs = 484
		goto _test_eof
	_test_eof485:
		lex.cs = 485
		goto _test_eof
	_test_eof486:
		lex.cs = 486
		goto _test_eof
	_test_eof487:
		lex.cs = 487
		goto _test_eof
	_test_eof488:
		lex.cs = 488
		goto _test_eof
	_test_eof489:
		lex.cs = 489
		goto _test_eof
	_test_eof490:
		lex.cs = 490
		goto _test_eof
	_test_eof491:
		lex.cs = 491
		goto _test_eof
	_test_eof492:
		lex.cs = 492
		goto _test_eof
	_test_eof104:
		lex.cs = 104
		goto _test_eof
	_test_eof493:
		lex.cs = 493
		goto _test_eof
	_test_eof494:
		lex.cs = 494
		goto _test_eof
	_test_eof495:
		lex.cs = 495
		goto _test_eof
	_test_eof105:
		lex.cs = 105
		goto _test_eof
	_test_eof496:
		lex.cs = 496
		goto _test_eof
	_test_eof497:
		lex.cs = 497
		goto _test_eof
	_test_eof498:
		lex.cs = 498
		goto _test_eof
	_test_eof499:
		lex.cs = 499
		goto _test_eof
	_test_eof500:
		lex.cs = 500
		goto _test_eof
	_test_eof501:
		lex.cs = 501
		goto _test_eof
	_test_eof502:
		lex.cs = 502
		goto _test_eof
	_test_eof106:
		lex.cs = 106
		goto _test_eof
	_test_eof503:
		lex.cs = 503
		goto _test_eof
	_test_eof504:
		lex.cs = 504
		goto _test_eof
	_test_eof505:
		lex.cs = 505
		goto _test_eof
	_test_eof506:
		lex.cs = 506
		goto _test_eof
	_test_eof507:
		lex.cs = 507
		goto _test_eof
	_test_eof508:
		lex.cs = 508
		goto _test_eof
	_test_eof107:
		lex.cs = 107
		goto _test_eof
	_test_eof108:
		lex.cs = 108
		goto _test_eof
	_test_eof509:
		lex.cs = 509
		goto _test_eof
	_test_eof510:
		lex.cs = 510
		goto _test_eof
	_test_eof511:
		lex.cs = 511
		goto _test_eof
	_test_eof512:
		lex.cs = 512
		goto _test_eof
	_test_eof513:
		lex.cs = 513
		goto _test_eof
	_test_eof514:
		lex.cs = 514
		goto _test_eof
	_test_eof109:
		lex.cs = 109
		goto _test_eof
	_test_eof110:
		lex.cs = 110
		goto _test_eof
	_test_eof515:
		lex.cs = 515
		goto _test_eof
	_test_eof516:
		lex.cs = 516
		goto _test_eof
	_test_eof517:
		lex.cs = 517
		goto _test_eof
	_test_eof518:
		lex.cs = 518
		goto _test_eof
	_test_eof519:
		lex.cs = 519
		goto _test_eof
	_test_eof520:
		lex.cs = 520
		goto _test_eof
	_test_eof521:
		lex.cs = 521
		goto _test_eof
	_test_eof522:
		lex.cs = 522
		goto _test_eof
	_test_eof523:
		lex.cs = 523
		goto _test_eof
	_test_eof524:
		lex.cs = 524
		goto _test_eof
	_test_eof525:
		lex.cs = 525
		goto _test_eof
	_test_eof111:
		lex.cs = 111
		goto _test_eof
	_test_eof526:
		lex.cs = 526
		goto _test_eof
	_test_eof112:
		lex.cs = 112
		goto _test_eof
	_test_eof113:
		lex.cs = 113
		goto _test_eof
	_test_eof527:
		lex.cs = 527
		goto _test_eof
	_test_eof528:
		lex.cs = 528
		goto _test_eof
	_test_eof529:
		lex.cs = 529
		goto _test_eof
	_test_eof530:
		lex.cs = 530
		goto _test_eof
	_test_eof531:
		lex.cs = 531
		goto _test_eof
	_test_eof532:
		lex.cs = 532
		goto _test_eof
	_test_eof533:
		lex.cs = 533
		goto _test_eof
	_test_eof534:
		lex.cs = 534
		goto _test_eof
	_test_eof114:
		lex.cs = 114
		goto _test_eof
	_test_eof115:
		lex.cs = 115
		goto _test_eof
	_test_eof535:
		lex.cs = 535
		goto _test_eof
	_test_eof116:
		lex.cs = 116
		goto _test_eof
	_test_eof536:
		lex.cs = 536
		goto _test_eof
	_test_eof537:
		lex.cs = 537
		goto _test_eof
	_test_eof538:
		lex.cs = 538
		goto _test_eof
	_test_eof539:
		lex.cs = 539
		goto _test_eof
	_test_eof117:
		lex.cs = 117
		goto _test_eof
	_test_eof540:
		lex.cs = 540
		goto _test_eof
	_test_eof541:
		lex.cs = 541
		goto _test_eof
	_test_eof542:
		lex.cs = 542
		goto _test_eof
	_test_eof118:
		lex.cs = 118
		goto _test_eof
	_test_eof543:
		lex.cs = 543
		goto _test_eof
	_test_eof544:
		lex.cs = 544
		goto _test_eof
	_test_eof545:
		lex.cs = 545
		goto _test_eof
	_test_eof546:
		lex.cs = 546
		goto _test_eof
	_test_eof119:
		lex.cs = 119
		goto _test_eof
	_test_eof547:
		lex.cs = 547
		goto _test_eof
	_test_eof548:
		lex.cs = 548
		goto _test_eof
	_test_eof549:
		lex.cs = 549
		goto _test_eof
	_test_eof550:
		lex.cs = 550
		goto _test_eof
	_test_eof120:
		lex.cs = 120
		goto _test_eof
	_test_eof551:
		lex.cs = 551
		goto _test_eof
	_test_eof552:
		lex.cs = 552
		goto _test_eof
	_test_eof553:
		lex.cs = 553
		goto _test_eof
	_test_eof554:
		lex.cs = 554
		goto _test_eof
	_test_eof555:
		lex.cs = 555
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 122:
				goto tr182
			case 1:
				goto tr0
			case 123:
				goto tr183
			case 125:
				goto tr188
			case 126:
				goto tr190
			case 127:
				goto tr188
			case 128:
				goto tr188
			case 129:
				goto tr196
			case 2:
				goto tr4
			case 3:
				goto tr4
			case 4:
				goto tr4
			case 130:
				goto tr199
			case 5:
				goto tr4
			case 132:
				goto tr253
			case 133:
				goto tr255
			case 6:
				goto tr10
			case 134:
				goto tr259
			case 135:
				goto tr260
			case 136:
				goto tr262
			case 137:
				goto tr264
			case 7:
				goto tr12
			case 8:
				goto tr12
			case 9:
				goto tr12
			case 10:
				goto tr12
			case 138:
				goto tr265
			case 139:
				goto tr268
			case 140:
				goto tr271
			case 141:
				goto tr260
			case 142:
				goto tr276
			case 143:
				goto tr260
			case 144:
				goto tr260
			case 145:
				goto tr259
			case 11:
				goto tr19
			case 12:
				goto tr19
			case 146:
				goto tr260
			case 13:
				goto tr23
			case 14:
				goto tr23
			case 15:
				goto tr23
			case 16:
				goto tr23
			case 17:
				goto tr23
			case 18:
				goto tr23
			case 19:
				goto tr23
			case 20:
				goto tr23
			case 21:
				goto tr23
			case 22:
				goto tr23
			case 23:
				goto tr23
			case 24:
				goto tr23
			case 25:
				goto tr23
			case 26:
				goto tr23
			case 27:
				goto tr23
			case 28:
				goto tr23
			case 29:
				goto tr23
			case 30:
				goto tr23
			case 31:
				goto tr23
			case 32:
				goto tr23
			case 33:
				goto tr23
			case 34:
				goto tr23
			case 35:
				goto tr23
			case 36:
				goto tr23
			case 37:
				goto tr23
			case 38:
				goto tr23
			case 39:
				goto tr23
			case 40:
				goto tr23
			case 41:
				goto tr23
			case 42:
				goto tr23
			case 43:
				goto tr23
			case 44:
				goto tr23
			case 45:
				goto tr23
			case 46:
				goto tr23
			case 47:
				goto tr23
			case 48:
				goto tr23
			case 49:
				goto tr23
			case 50:
				goto tr23
			case 51:
				goto tr23
			case 52:
				goto tr23
			case 53:
				goto tr23
			case 54:
				goto tr23
			case 55:
				goto tr23
			case 56:
				goto tr23
			case 57:
				goto tr23
			case 58:
				goto tr23
			case 59:
				goto tr23
			case 60:
				goto tr23
			case 61:
				goto tr23
			case 62:
				goto tr23
			case 63:
				goto tr23
			case 64:
				goto tr23
			case 65:
				goto tr23
			case 66:
				goto tr23
			case 67:
				goto tr23
			case 147:
				goto tr260
			case 148:
				goto tr282
			case 149:
				goto tr260
			case 150:
				goto tr260
			case 151:
				goto tr260
			case 68:
				goto tr23
			case 152:
				goto tr291
			case 69:
				goto tr12
			case 70:
				goto tr12
			case 153:
				goto tr291
			case 71:
				goto tr90
			case 154:
				goto tr260
			case 72:
				goto tr23
			case 73:
				goto tr23
			case 74:
				goto tr23
			case 155:
				goto tr295
			case 156:
				goto tr291
			case 157:
				goto tr295
			case 75:
				goto tr101
			case 76:
				goto tr12
			case 158:
				goto tr300
			case 77:
				goto tr12
			case 159:
				goto tr301
			case 160:
				goto tr260
			case 161:
				goto tr260
			case 78:
				goto tr23
			case 79:
				goto tr23
			case 80:
				goto tr23
			case 81:
				goto tr23
			case 162:
				goto tr303
			case 163:
				goto tr305
			case 82:
				goto tr114
			case 164:
				goto tr260
			case 165:
				goto tr309
			case 83:
				goto tr12
			case 84:
				goto tr12
			case 85:
				goto tr12
			case 86:
				goto tr12
			case 166:
				goto tr311
			case 87:
				goto tr12
			case 88:
				goto tr12
			case 89:
				goto tr12
			case 90:
				goto tr12
			case 167:
				goto tr312
			case 168:
				goto tr260
			case 169:
				goto tr316
			case 170:
				goto tr260
			case 171:
				goto tr320
			case 172:
				goto tr260
			case 91:
				goto tr23
			case 173:
				goto tr325
			case 174:
				goto tr327
			case 92:
				goto tr131
			case 175:
				goto tr328
			case 176:
				goto tr330
			case 177:
				goto tr12
			case 93:
				goto tr12
			case 178:
				goto tr336
			case 179:
				goto tr330
			case 180:
				goto tr330
			case 181:
				goto tr330
			case 182:
				goto tr330
			case 183:
				goto tr330
			case 184:
				goto tr330
			case 185:
				goto tr330
			case 186:
				goto tr330
			case 187:
				goto tr330
			case 188:
				goto tr330
			case 189:
				goto tr330
			case 94:
				goto tr134
			case 95:
				goto tr134
			case 190:
				goto tr330
			case 191:
				goto tr330
			case 192:
				goto tr330
			case 193:
				goto tr330
			case 194:
				goto tr330
			case 195:
				goto tr330
			case 196:
				goto tr330
			case 197:
				goto tr330
			case 198:
				goto tr330
			case 199:
				goto tr330
			case 200:
				goto tr330
			case 201:
				goto tr330
			case 202:
				goto tr330
			case 203:
				goto tr330
			case 204:
				goto tr330
			case 205:
				goto tr330
			case 206:
				goto tr330
			case 207:
				goto tr330
			case 208:
				goto tr330
			case 209:
				goto tr330
			case 210:
				goto tr330
			case 211:
				goto tr330
			case 212:
				goto tr330
			case 213:
				goto tr330
			case 214:
				goto tr330
			case 215:
				goto tr330
			case 216:
				goto tr330
			case 217:
				goto tr330
			case 218:
				goto tr330
			case 219:
				goto tr330
			case 220:
				goto tr330
			case 221:
				goto tr330
			case 222:
				goto tr330
			case 223:
				goto tr330
			case 224:
				goto tr330
			case 225:
				goto tr330
			case 226:
				goto tr330
			case 227:
				goto tr330
			case 228:
				goto tr330
			case 229:
				goto tr330
			case 230:
				goto tr330
			case 231:
				goto tr330
			case 232:
				goto tr330
			case 233:
				goto tr330
			case 234:
				goto tr330
			case 235:
				goto tr330
			case 236:
				goto tr330
			case 237:
				goto tr330
			case 238:
				goto tr412
			case 239:
				goto tr330
			case 240:
				goto tr330
			case 241:
				goto tr330
			case 242:
				goto tr330
			case 243:
				goto tr330
			case 244:
				goto tr330
			case 245:
				goto tr330
			case 246:
				goto tr330
			case 247:
				goto tr330
			case 248:
				goto tr330
			case 249:
				goto tr330
			case 250:
				goto tr330
			case 251:
				goto tr330
			case 252:
				goto tr330
			case 253:
				goto tr432
			case 254:
				goto tr330
			case 255:
				goto tr330
			case 256:
				goto tr330
			case 257:
				goto tr330
			case 258:
				goto tr330
			case 259:
				goto tr330
			case 260:
				goto tr330
			case 261:
				goto tr330
			case 262:
				goto tr330
			case 263:
				goto tr330
			case 264:
				goto tr330
			case 265:
				goto tr330
			case 266:
				goto tr330
			case 267:
				goto tr330
			case 268:
				goto tr330
			case 269:
				goto tr330
			case 270:
				goto tr330
			case 271:
				goto tr330
			case 272:
				goto tr330
			case 273:
				goto tr330
			case 274:
				goto tr330
			case 275:
				goto tr330
			case 276:
				goto tr330
			case 277:
				goto tr330
			case 278:
				goto tr330
			case 279:
				goto tr461
			case 280:
				goto tr330
			case 281:
				goto tr330
			case 282:
				goto tr465
			case 283:
				goto tr330
			case 284:
				goto tr330
			case 285:
				goto tr330
			case 286:
				goto tr330
			case 287:
				goto tr330
			case 288:
				goto tr330
			case 289:
				goto tr330
			case 290:
				goto tr330
			case 291:
				goto tr330
			case 292:
				goto tr330
			case 293:
				goto tr330
			case 294:
				goto tr330
			case 295:
				goto tr330
			case 296:
				goto tr330
			case 297:
				goto tr330
			case 298:
				goto tr330
			case 299:
				goto tr330
			case 300:
				goto tr330
			case 301:
				goto tr330
			case 302:
				goto tr330
			case 303:
				goto tr330
			case 304:
				goto tr330
			case 305:
				goto tr330
			case 306:
				goto tr330
			case 307:
				goto tr497
			case 308:
				goto tr330
			case 309:
				goto tr330
			case 310:
				goto tr330
			case 311:
				goto tr330
			case 312:
				goto tr330
			case 313:
				goto tr330
			case 314:
				goto tr330
			case 315:
				goto tr330
			case 316:
				goto tr330
			case 317:
				goto tr330
			case 318:
				goto tr330
			case 319:
				goto tr330
			case 320:
				goto tr330
			case 321:
				goto tr330
			case 322:
				goto tr330
			case 323:
				goto tr330
			case 324:
				goto tr330
			case 325:
				goto tr330
			case 326:
				goto tr330
			case 327:
				goto tr330
			case 328:
				goto tr330
			case 329:
				goto tr330
			case 330:
				goto tr330
			case 331:
				goto tr330
			case 332:
				goto tr330
			case 333:
				goto tr330
			case 334:
				goto tr330
			case 335:
				goto tr330
			case 336:
				goto tr330
			case 337:
				goto tr330
			case 338:
				goto tr330
			case 339:
				goto tr330
			case 340:
				goto tr330
			case 341:
				goto tr330
			case 342:
				goto tr330
			case 343:
				goto tr330
			case 344:
				goto tr330
			case 345:
				goto tr330
			case 346:
				goto tr330
			case 347:
				goto tr540
			case 96:
				goto tr12
			case 348:
				goto tr542
			case 349:
				goto tr330
			case 350:
				goto tr330
			case 351:
				goto tr330
			case 352:
				goto tr330
			case 353:
				goto tr330
			case 354:
				goto tr330
			case 355:
				goto tr330
			case 356:
				goto tr330
			case 357:
				goto tr330
			case 358:
				goto tr330
			case 359:
				goto tr330
			case 360:
				goto tr330
			case 361:
				goto tr330
			case 362:
				goto tr330
			case 363:
				goto tr330
			case 364:
				goto tr330
			case 365:
				goto tr330
			case 366:
				goto tr330
			case 367:
				goto tr330
			case 368:
				goto tr330
			case 369:
				goto tr330
			case 370:
				goto tr330
			case 371:
				goto tr330
			case 372:
				goto tr330
			case 373:
				goto tr330
			case 374:
				goto tr330
			case 375:
				goto tr330
			case 376:
				goto tr330
			case 377:
				goto tr330
			case 378:
				goto tr330
			case 379:
				goto tr578
			case 380:
				goto tr330
			case 381:
				goto tr330
			case 382:
				goto tr330
			case 383:
				goto tr330
			case 384:
				goto tr330
			case 385:
				goto tr330
			case 386:
				goto tr330
			case 387:
				goto tr330
			case 388:
				goto tr330
			case 389:
				goto tr330
			case 390:
				goto tr330
			case 391:
				goto tr330
			case 392:
				goto tr330
			case 393:
				goto tr330
			case 394:
				goto tr330
			case 395:
				goto tr330
			case 396:
				goto tr330
			case 397:
				goto tr330
			case 398:
				goto tr330
			case 399:
				goto tr330
			case 400:
				goto tr330
			case 401:
				goto tr330
			case 402:
				goto tr330
			case 403:
				goto tr330
			case 404:
				goto tr330
			case 405:
				goto tr330
			case 406:
				goto tr330
			case 407:
				goto tr330
			case 408:
				goto tr330
			case 409:
				goto tr330
			case 410:
				goto tr330
			case 411:
				goto tr330
			case 412:
				goto tr330
			case 413:
				goto tr330
			case 414:
				goto tr330
			case 415:
				goto tr330
			case 416:
				goto tr330
			case 417:
				goto tr330
			case 418:
				goto tr330
			case 419:
				goto tr330
			case 420:
				goto tr624
			case 97:
				goto tr137
			case 98:
				goto tr137
			case 99:
				goto tr137
			case 100:
				goto tr137
			case 101:
				goto tr137
			case 102:
				goto tr137
			case 421:
				goto tr625
			case 422:
				goto tr626
			case 103:
				goto tr149
			case 423:
				goto tr260
			case 424:
				goto tr330
			case 425:
				goto tr330
			case 426:
				goto tr330
			case 427:
				goto tr330
			case 428:
				goto tr330
			case 429:
				goto tr330
			case 430:
				goto tr330
			case 431:
				goto tr330
			case 432:
				goto tr330
			case 433:
				goto tr330
			case 434:
				goto tr330
			case 435:
				goto tr330
			case 436:
				goto tr330
			case 437:
				goto tr330
			case 438:
				goto tr330
			case 439:
				goto tr330
			case 440:
				goto tr330
			case 441:
				goto tr330
			case 442:
				goto tr330
			case 443:
				goto tr330
			case 444:
				goto tr330
			case 445:
				goto tr330
			case 446:
				goto tr330
			case 447:
				goto tr330
			case 448:
				goto tr330
			case 449:
				goto tr330
			case 450:
				goto tr330
			case 451:
				goto tr330
			case 452:
				goto tr330
			case 453:
				goto tr330
			case 454:
				goto tr330
			case 455:
				goto tr330
			case 456:
				goto tr330
			case 457:
				goto tr330
			case 458:
				goto tr330
			case 459:
				goto tr330
			case 460:
				goto tr330
			case 461:
				goto tr330
			case 462:
				goto tr330
			case 463:
				goto tr330
			case 464:
				goto tr330
			case 465:
				goto tr330
			case 466:
				goto tr330
			case 467:
				goto tr330
			case 468:
				goto tr330
			case 469:
				goto tr330
			case 470:
				goto tr330
			case 471:
				goto tr330
			case 472:
				goto tr330
			case 473:
				goto tr330
			case 474:
				goto tr330
			case 475:
				goto tr330
			case 476:
				goto tr330
			case 477:
				goto tr330
			case 478:
				goto tr330
			case 479:
				goto tr330
			case 480:
				goto tr330
			case 481:
				goto tr330
			case 482:
				goto tr330
			case 483:
				goto tr330
			case 484:
				goto tr330
			case 485:
				goto tr330
			case 486:
				goto tr330
			case 487:
				goto tr330
			case 488:
				goto tr330
			case 489:
				goto tr260
			case 491:
				goto tr710
			case 492:
				goto tr712
			case 104:
				goto tr151
			case 493:
				goto tr716
			case 494:
				goto tr716
			case 495:
				goto tr716
			case 105:
				goto tr153
			case 496:
				goto tr719
			case 498:
				goto tr723
			case 499:
				goto tr724
			case 500:
				goto tr728
			case 502:
				goto tr736
			case 503:
				goto tr738
			case 504:
				goto tr739
			case 505:
				goto tr743
			case 506:
				goto tr736
			case 507:
				goto tr743
			case 509:
				goto tr755
			case 510:
				goto tr756
			case 511:
				goto tr760
			case 512:
				goto tr760
			case 513:
				goto tr760
			case 515:
				goto tr773
			case 516:
				goto tr774
			case 517:
				goto tr778
			case 518:
				goto tr778
			case 519:
				goto tr778
			case 521:
				goto tr783
			case 523:
				goto tr790
			case 524:
				goto tr792
			case 525:
				goto tr790
			case 111:
				goto tr163
			case 526:
				goto tr790
			case 112:
				goto tr163
			case 113:
				goto tr163
			case 527:
				goto tr795
			case 529:
				goto tr805
			case 530:
				goto tr806
			case 531:
				goto tr807
			case 532:
				goto tr809
			case 533:
				goto tr810
			case 534:
				goto tr810
			case 114:
				goto tr167
			case 115:
				goto tr167
			case 535:
				goto tr810
			case 116:
				goto tr167
			case 536:
				goto tr810
			case 537:
				goto tr814
			case 539:
				goto tr817
			case 117:
				goto tr171
			case 541:
				goto tr822
			case 542:
				goto tr824
			case 118:
				goto tr174
			case 543:
				goto tr828
			case 545:
				goto tr833
			case 546:
				goto tr835
			case 119:
				goto tr176
			case 547:
				goto tr839
			case 549:
				goto tr844
			case 550:
				goto tr846
			case 120:
				goto tr178
			case 551:
				goto tr850
			case 553:
				goto tr854
			case 554:
				goto tr855
			case 555:
				goto tr859
			}
		}

	_out:
		{
		}
	}

	// line internal/php8/scanner.rl:511

	tkn.Value = lex.data[lex.ts:lex.te]
	tkn.ID = token.ID(tok)

	return tkn
}
