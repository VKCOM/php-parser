package printer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/z7zmey/php-parser/internal/php8"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/conf"
	"github.com/z7zmey/php-parser/pkg/version"
	"github.com/z7zmey/php-parser/pkg/visitor/printer"
	"gotest.tools/assert"
)

func ExamplePrinterPHP8() {
	src := `<?php

namespace Foo;

abstract class Bar extends Baz
{
    public function greet()
    {
        echo "Hello";
        // some comment
    }
}
	`

	// parsePHP8

	config := conf.Config{
		Version: &version.Version{
			Major: 8,
			Minor: 0,
		},
	}
	lexer := php8.NewLexer([]byte(src), config)
	php8parser := php8.NewParser(lexer, config)
	php8parser.Parse()

	rootNode := php8parser.GetRootNode()

	// change namespace

	parts := &rootNode.(*ast.Root).Stmts[0].(*ast.StmtNamespace).Name.(*ast.Name).Parts
	*parts = append(*parts, &ast.NamePart{Value: []byte("Quuz")})

	// print

	p := printer.NewPrinter(os.Stdout)
	rootNode.Accept(p)

	// Output:
	// <?php
	//
	// namespace Foo\Quuz;
	//
	// abstract class Bar extends Baz
	// {
	//     public function greet()
	//     {
	//         echo "Hello";
	//         // some comment
	//     }
	// }
}

func parsePHP8(src string) ast.Vertex {
	config := conf.Config{
		Version: &version.Version{
			Major: 8,
			Minor: 1,
		},
	}
	lexer := php8.NewLexer([]byte(src), config)
	php8parser := php8.NewParser(lexer, config)
	php8parser.Parse()

	return php8parser.GetRootNode()
}

func printPHP8(n ast.Vertex) string {
	o := bytes.NewBufferString("")

	p := printer.NewPrinter(o)
	n.Accept(p)

	return o.String()
}

// test node

func TestParseAndPrintRootPHP8(t *testing.T) {

	src := ` <div>Hello</div> 
	<?php
	$a;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIdentifierPHP8(t *testing.T) {
	src := `<? ;
	/* Foo */
	Foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintParameterTMPPHP8(t *testing.T) {

	src := `<?php
	function foo ( foo & ... $foo = null ) {}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintParameterPHP8(t *testing.T) {

	src := `<?php
	function & foo (
		? int $a , & $b = null
		, \ Foo ...$c
	) : namespace  \ Bar \  baz \ quuz{
		;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNullablePHP8(t *testing.T) {

	src := `<?php
	function & foo ( ? int $a ) {
		/* do nothing */
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArgumentPHP8(t *testing.T) {
	src := `<?php
	foo ( $a , $b
		, ... $c ,
	) ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test name

func TestParseAndPrintNamesPHP8(t *testing.T) {
	src := `<?php
	foo ( ) ;
	\ foo ( ) ;
	namespace \ foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test scalar

func TestParseAndPrintMagicConstantPHP8(t *testing.T) {
	src := `<?php
	__CLASS__     ;
	__DIR__       ;
	__FILE__      ;
	__FUNCTION__  ;
	__LINE__      ;
	__NAMESPACE__ ;
	__METHOD__    ;
	__TRAIT__     ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNumberPHP8(t *testing.T) {
	src := `<?php
	// LNumber
	1234567890123456789 ;

	// DNumber
	12345678901234567890 ;
	0. ;
	.2 ;
	0.2 ;

	// binary LNumber
	0b0111111111111111111111111111111111111111111111111111111111111111 ;

	// binary DNumber
	0b1111111111111111111111111111111111111111111111111111111111111111 ;

	// HLNumber
	0x007111111111111111 ;

	// HDNumber
	0x8111111111111111 ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStringPHP8(t *testing.T) {
	src := `<?php
	'Hello' ;
	"Hello {$world } " ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintHeredocPHP8(t *testing.T) {
	src := `<?php
	foo(<<<EAP
test
EAP
, 'test'
);

<<<EAP
test
EAP;

<<<'EAP'
test
EAP;

<<<"EAP"
test
EAP;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test assign

func TestParseAndPrintAssignPHP8(t *testing.T) {
	src := `<?php
	$a = $b ;
	$a = & $b ;
	$a &= $b ;
	$a |= $b ;
	$a ^= $b ;
	$a ??= $b ;
	$a .= $b ;
	$a /= $b ;
	$a -= $b ;
	$a %= $b ;
	$a *= $b ;
	$a += $b ;
	$a **= $b ;
	$a <<= $b ;
	$a >>= $b ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test binary

func TestParseAndPrintBinaryPHP8(t *testing.T) {
	src := `<?php
	$a & $b ;
	$a | $b ;
	$a ^ $b ;
	$a && $b ;
	$a || $b ;
	$a ?? $b ;
	$a . $b ;
	$a / $b ;
	$a == $b ;
	$a >= $b ;
	$a > $b ;
	$a === $b ;
	$a and $b ;
	$a or $b ;
	$a xor $b ;
	$a - $b ;
	$a % $b ;
	$a * $b ;
	$a != $b ;
	$a <> $b ;
	$a !== $b ;
	$a + $b ;
	$a ** $b ;
	$a << $b ;
	$a >> $b ;
	$a <= $b ;
	$a < $b ;
	$a <=> $b ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test cast

func TestParseAndPrintCastPHP8(t *testing.T) {
	src := `<?php
	(  array     ) $a ;
	(  bool      ) $a ;
	(  boolean   ) $a ;
	// (  real      ) $a ; real cast was removed in PHP 8
	(  double    ) $a ;
	(  float     ) $a ;
	(  int       ) $a ;
	(  integer   ) $a ;
	(  object    ) $a ;
	(  string    ) $a ;
	(  binary    ) $a ;
	// (  unset     ) $a ; unset cast was removed in PHP 8
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test expr

func TestParseAndPrintArrayDimFetchPHP8(t *testing.T) {
	src := `<?php
	FOO [ ] ;
	FOO [ 1 ] ;
	$a [ ] ;
	$a [ 1 ] ;
	$a { 1 } ;
	new $a [ ] ;
	new $a [ 1 ] ;
	new $a { 1 } ;
	"$a[1]test" ;
	"${ a [ 1 ] }test" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArrayItemPHP8(t *testing.T) {
	src := `<?php
	$foo = [
		$world ,
		& $world ,
		'Hello' => $world ,
		... $unpack
	] ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArrayPHP8(t *testing.T) {
	src := `<?php
	array ( /* empty array */ ) ;
	array ( 0 , 2 => 2 ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBitwiseNotPHP8(t *testing.T) {
	src := `<?php
	~ $var ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBooleanNotPHP8(t *testing.T) {
	src := `<?php
	! $var ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassConstFetchPHP8(t *testing.T) {
	src := `<?php
	$var :: CONST ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClonePHP8(t *testing.T) {
	src := `<?php
	clone $var ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClosureUsePHP8(t *testing.T) {
	src := `<?php
	$a = function ( ) use ( $a , & $b ) {
		// do nothing
	} ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClosurePHP8(t *testing.T) {
	src := `<?php
	$a  = static function & ( ) : void {
		// do nothing
	} ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArrowFunctionPHP8(t *testing.T) {
	src := `<?php
	$a = static fn & ( $b ) : void => $c ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintConstFetchPHP8(t *testing.T) {
	src := `<?php
	null ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEmptyPHP8(t *testing.T) {
	src := `<?php
	empty ( $a ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintErrorSuppressPHP8(t *testing.T) {
	src := `<?php
	@ foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEvalPHP8(t *testing.T) {
	src := `<?php
	eval ( " " ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintExitPHP8(t *testing.T) {
	src := `<?php
	exit ;
	exit ( ) ;
	exit (1) ;
	exit ( 1 );
	die ;
	die ( ) ;
	die (1) ;
	die ( 1 );
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunctionCallPHP8(t *testing.T) {
	src := `<?php
	foo ( ) ;
	$var ( $a , ... $b , $c ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIncludePHP8(t *testing.T) {

	src := `<?php
	include 'foo' ;
	include_once 'bar' ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInstanceOfPHP8(t *testing.T) {
	src := `<?php
	$a instanceof Foo ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIssetPHP8(t *testing.T) {
	src := `<?php
	isset ( $a , $b [ 2 ] , ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintListPHP8(t *testing.T) {
	src := `<?php
	list( , $var , ) = $b ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintMethodCallPHP8(t *testing.T) {
	src := `<?php
	$a -> bar ( $arg , ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNewPHP8(t *testing.T) {
	src := `<?php

	new Foo ;

	new Foo ( $a, $b ) ;

	new class ( $c ) extends Foo implements Bar , Baz {

	} ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIncDecPHP8(t *testing.T) {
	src := `<?php
	++ $a ;
	-- $a ;
	$a ++ ;
	$a -- ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPrintPHP8(t *testing.T) {
	src := `<?php
	print $a ;
	print ( $a ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPropertyFetchPHP8(t *testing.T) {
	src := `<?php
	$a -> b ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintReferencePHP8(t *testing.T) {
	src := `<?php
	$a = & $b ;
	$a = [ & $b ] ;
	$a = [ $b => & $c ] ;

	$a = function ( ) use ( & $b ) {
		// do nothing
	} ;

	foreach ( $a as & $b ) {
		// do nothing
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintRequirePHP8(t *testing.T) {

	src := `<?php
	require __DIR__ . '/folder' ;
	require_once $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShellExecPHP8(t *testing.T) {
	src := "<?php ` {$v} cmd ` ; "

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShortArrayPHP8(t *testing.T) {
	src := `<?php
	$a = [ ] ;
	$a = [ 0 ] ;
	$a = [
		1 => & $b , // one
		$c       , /* two */ 
	] ;
	$a = [0, 1, 2] ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShortListPHP8(t *testing.T) {
	src := `<?php
	[ 
		/* skip */,
		$b 
		/* skip */,
	] = $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticCallPHP8(t *testing.T) {
	src := `<?php
	Foo :: bar ( $a , $b ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticPropertyFetchPHP8(t *testing.T) {
	src := `<?php
	Foo :: $bar ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTernaryPHP8(t *testing.T) {
	src := `<?php
	$a ? $b : $c ;
	$a ? : $c ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUnaryPHP8(t *testing.T) {
	src := `<?php
	- $a ;
	+ $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintVariablePHP8(t *testing.T) {
	src := `<?php
	$ /* variable variable comment */ $var ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintYieldPHP8(t *testing.T) {
	src := `<?php
	yield $a ;
	yield $k => $v ;
	yield from $a ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// test stmt

func TestParseAndPrintAltIfPHP8(t *testing.T) {
	src := `<?php
	if ( 1 ) :
		// do nothing
	elseif ( 2 ) :
	elseif ( 3 ) :
		;
	else :
		;
	endif ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltForPHP8(t *testing.T) {
	src := `<?php
	for ( $a ; $b ; $c ) :
	endfor ;
	
	for ( ; ; ) :
	endfor ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltForeachPHP8(t *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) :
		echo $v ;
	endforeach ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltSwitchPHP8(t *testing.T) {
	src := `<?php

	switch ( $a ) : 
	case 1 : 
		;
	case 2 : ;
	case 3 :
		 ;
	default :
		;
	endswitch ;
	

	switch ( $a ) : ;
	case 1 ; ;
	default ; ;
	endswitch ;

	switch ( $a ) :
	endswitch ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAltWhilePHP8(t *testing.T) {
	src := `<?php

	while ( $a ) :
		// do nothing
	endwhile ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintBreakPHP8(t *testing.T) {
	src := `<?php

	break ;
	break 1 ;
	break ( 2 ) ;
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassMethodPHP8(t *testing.T) {
	src := `<?php
	class Foo {
		/**
		 * abstract method
		 */
		public static function & greet ( ? Foo $a ) : void ;
		
		function greet ( string $a )
		{
			return 'hello' ;
		}
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassPHP8(t *testing.T) {
	src := `<?php
	abstract final class Foo extends Bar implements Baz , Quuz {
		
	}

	new class ( $c, $a ) extends Foo implements Bar , Baz {

	} ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassConstListPHP8(t *testing.T) {
	src := `<?php
	class Foo {
		public const FOO = 'f' , BAR = 'b' ;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintConstListPHP8(t *testing.T) {
	src := `<?php
	const FOO = 1 , BAR = 2 ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintContinuePHP8(t *testing.T) {
	src := `<?php

	continue ;
	continue 1 ;
	continue ( 2 ) ;
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDeclarePHP8(t *testing.T) {
	src := `<?php
	declare ( FOO = 'bar' , BAR = "foo" ) ;
	declare ( FOO = 'bar' ) $a ;
	declare ( FOO = 'bar' ) { }

	declare ( FOO = 'bar' ) : enddeclare ;
	declare ( FOO = 'bar' ) :
		;
	enddeclare ;
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintDoWhilePHP8(t *testing.T) {
	src := `<?php
	do {
		;
	} while ( $a ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintEchoPHP8(t *testing.T) {
	src := `<?php
	echo '' ;
	echo $a , ' ' , PHP_EOL;

	?>

	<?= $a, $b ?>
	<?= $c ; 

`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIfExpressionPHP8(t *testing.T) {
	src := `<?php
	$a ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintForPHP8(t *testing.T) {
	src := `<?php
	for ( $i = 0 ; $i < 3 ; $i ++ ) 
		echo $i . PHP_EOL;
	
	for ( ; ; ) {

	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintForeachPHP8(t *testing.T) {
	src := `<?php
	foreach ( $a as $k => & $v ) {
		;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunctionPHP8(t *testing.T) {

	src := `<?php
	function & foo ( ) : void {
		;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGlobalPHP8(t *testing.T) {
	src := `<?php
	global $a , $b ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGotoPHP8(t *testing.T) {
	src := `<?php
	goto Foo ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGroupUsePHP8(t *testing.T) {
	src := `<?php
	use function Foo \ { Bar as Baz , Quuz , } ;
	use Foo \ { function Bar as Baz , const Quuz } ;
	use \ Foo \ { function Bar as Baz , const Quuz , } ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintHaltCompilerPHP8(t *testing.T) {
	src := `<?php
	__halt_compiler ( ) ;
	this text is ignored by parser
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintIfElseIfElsePHP8(t *testing.T) {
	src := `<?php
	if ( 1 ) ;
	elseif ( 2 ) {
		;
	}
	else if ( 3 ) $a;
	else { }`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInlineHtmlPHP8(t *testing.T) {
	src := `<?php
	$a;?>test<? `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintShebangPHP8(t *testing.T) {
	src := `#!/usr/bin/env php
	<?php
	$a;?>test<? `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInterfacePHP8(t *testing.T) {
	src := `<?php
	interface Foo extends Bar , Baz {

	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintGotoLabelPHP8(t *testing.T) {
	src := `<?php
	Foo : $b ; `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNamespacePHP8(t *testing.T) {
	src := `<?php
	namespace Foo \ Bar ; 
	namespace Baz {

	}
	namespace {
		
	}
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNopPHP8(t *testing.T) {
	src := `<?php `

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintPropertyListPHP8(t *testing.T) {
	src := `<?php
	class Foo {
		var $a = '' , $b = null ;
		private $c ;
		public static Bar $d ;
		
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintReturnPHP8(t *testing.T) {
	src := `<?php
	class Foo {
		function bar ( )
		{
			return null ;
		}
	}

	function foo ( )
	{
		return $a ;
	}

	function bar ( )
	{
		return ;
	}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStaticVarPHP8(t *testing.T) {
	src := `<?php
	static $a , $b = foo ( ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintStmtListPHP8(t *testing.T) {
	src := `<?php
	{
		;
	}
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintSwitchPHP8(t *testing.T) {
	src := `<?php

	switch ( $a ) {
		case 1 : ;
		default : ;
	}
	switch ( $a ) { ;
		case 1 ; ;
		default ; ;
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintThrowPHP8(t *testing.T) {
	src := `<?php
	throw new \ Exception ( "msg" ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTraitUsePHP8(t *testing.T) {
	src := `<?php
	class foo {
		use \ foo , bar ;
		use foo , \ bar { }
		use \ foo , \ bar {
			foo :: a as b ;
			bar :: a insteadof foo ;
			foo :: c as public ;
			foo :: d as public e;
			f as g ;
		}
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTraitPHP8(t *testing.T) {
	src := `<?php
	trait foo {
		function bar ( ) { }
	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTryCatchFinallyPHP8(t *testing.T) {
	src := `<?php

	try {

	} catch ( \ Exception | \ Foo \ Bar $e) {

	} finally {

	}`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUnsetPHP8(t *testing.T) {
	src := `<?php
	unset ( $a ) ;
	unset ( $a , $b , ) ;`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintUseListPHP8(t *testing.T) {
	src := `<?php
	use Foo ;
	use \ Foo as Bar ;
	use function \ Foo as Bar ;
	use const Foo as Bar, baz ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintWhilePHP8(t *testing.T) {
	src := `<?php
	while ( $a ) echo '' ;
	while ( $a ) { }
	while ( $a ) ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

// other

func TestParseAndPrintParenthesesPHP8(t *testing.T) {
	src := `<?php
	$b = (($a));
	$b = ( ($a) );
	$b = ( ( $a ) );
	$b = ( ($a ));
	$b = (( $a) );

	( $a + $b ) * 2 ;
	( $ $foo . 'foo' ) :: { $bar . 'bar' } ( ) ;
	( $ $foo ) [ 'bar' ] ;
	$ { $a . 'b' } -> call ( ) ;
	$a -> { $b . 'b' } ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString1PHP8(t *testing.T) {
	src := `<?php
	// "test $foo" ;
	"test $foo[1]" ;
	"test $foo[-1]" ;
	"test $foo[112345678901234567890] " ;
	"test $foo[-112345678901234567890] " ;
	"test $foo[a]" ;
	"test $foo[$bar]" ;
	"test $foo->bar" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString2PHP8(t *testing.T) {
	src := `<?php
	"test ${ foo }" ;
	"test ${ foo . 'bar' }" ;
	"test ${ foo [ ] }" ;
	"test ${ foo [ $b ] }" ;
	"test ${ foo [ 1 ] }" ;
	"test ${ foo [ 'expr' . $bar ] }" ;
	"test ${ $foo }" ;
	"test ${ $foo -> bar }" ;
	"test ${ $foo -> bar ( ) }" ;
	"test ${ $a . '' }" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString3PHP8(t *testing.T) {
	src := `<?php
	"test ${foo}" ;
	"test ${foo[0]}";
	"test ${foo::$bar}";
	"test ${foo }" ;
	"test ${foo . 'bar' }" ;
	"test ${foo [ ] }" ;
	"test ${foo [ $b ] }" ;
	"test ${foo [ 1 ] }" ;
	"test ${foo [ 'expr' . $bar ] }" ;
	"test ${$foo }" ;
	"test ${$foo -> bar }" ;
	"test ${$foo -> bar ( ) }" ;
	"test ${$a . '' }" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintComplexString4PHP8(t *testing.T) {
	src := `<?php
	"test {$foo }" ;
	"test {$foo [ ] }" ;
	"test {$foo [ 1 ] }" ;
	"test {$foo -> bar }" ;
	"test {$foo -> bar ( ) }" ;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNullsafePHP8(t *testing.T) {
	src := `<?php
$a?->method();
$a?->prop;
"$a?->prop_inside_string";
(f())?->prop_for_expr;
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintNamedMethods(t *testing.T) {
	src := `<?php
foo($a, name: $b, ...$b);
foo(name: $b);
foo(name: $b, c: 10);
$foo(a: $a);
$foo(a: $a, $c, ...$b);
$foo->bar(some: $a);
$foo->bar(some: $a, $b, ...$c);
foo::bar(b: $b);
foo::bar($a, b: $b, ...$c);
$foo::bar(c: $a);
$foo::bar($b, c: $a, ...$b);
new foo(a: $a);
new foo(a: $a, $c, ...$b);
new class (name: $a) {};
new class (name: $a, $b, ...$c) {};
	`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintMatchPHP8(t *testing.T) {
	src := `<?php
echo match($a) {};

echo match($a) {
	default => 100,
};
echo match($a) {
	default => 100
};
echo match($a) {
	default, => 100
};
echo match($a) {
	default, => 100,
};

echo match($a) {
	10 => 100
};
echo match($a) {
	10 => 100,
};
echo match($a) {
	10, => 100
};
echo match($a) {
	10, => 100,
};

echo match($a) {
	10, 20 => 100
};
echo match($a) {
	10, 20 => 100,
};
echo match($a) {
	10, 20, => 100
};
echo match($a) {
	10, 20, => 100,
};

echo match($a) {
	10, 20 => 100,
	30, 40 => 101
};
echo match($a) {
	10, 20 => 100,
	30, 40 => 101,
};
echo match($a) {
	10, 20, => 100,
	30, 40, => 101
};

echo match($a) {
	10, 20 => 100,
	30, 40 => 101,
	default => 102,
};
	`

	actual := printPHP8(parsePHP8(src))

	assert.Equal(t, src, actual)
}

func TestParseAndPrintUnionTypesPHP8(t *testing.T) {
	src := `<?php
function f(int|string $a) {}
function f(int|string|float $a) {}
function f(int|string|float|bool $a) {}

function f(int|\Foo $a) {}
function f(\Foo|int $a) {}
function f(\Foo|\Boo $a) {}
function f(\Namespaced\Foo|\Boo $a) {}
// function f(static|\Boo $a) {} // static cannot be used here.

function f(): int|string {}
function f(): int|\Boo {}
function f(): \Foo|\Boo {}
function f(): \Namespaced\Foo|\Boo {}

function f(): \Foo|static {}
function f(): \Foo|void {}

function f() {
	$_ = function(int|string $a) {};
	$_ = function(int|string|float $a) {};
	$_ = function(int|string|float|bool $a) {};
	$_ = function(int|\Foo $a) {};
	$_ = function(\Foo|int $a) {};
	$_ = function(\Foo|\Boo $a) {};
	// $_ = function(\Foo|static $a) {}; // static cannot be used here.
	$_ = function(\Namespaced\Foo|\Boo $a) {};

	$_ = function(): int|string {};
	$_ = function(): int|\Boo {};
	$_ = function(): \Foo|\Boo {};
	$_ = function(): \Namespaced\Foo|\Boo {};

	$_ = function(): \Foo|static {}; // static can be used here.
	$_ = function(): \Foo|void {};

	$_ = fn(int|string $a) => $a;
	$_ = fn(int|string|float $a) => $a;
	$_ = fn(int|string|float|bool $a) => $a;
	$_ = fn(int|\Foo $a) => $a;
	$_ = fn(\Foo|int $a) => $a;
	$_ = fn(\Foo|\Boo $a) => $a;
	// $_ = fn(\Foo|static $a) => $a; // static cannot be used here.
	$_ = fn(\Namespaced\Foo|\Boo $a) => $a;

	$_ = fn(): int|string => null;
	$_ = fn(): int|\Boo => null;
	$_ = fn(): \Foo|\Boo => null;
	$_ = fn(): \Namespaced\Foo|\Boo => null;
	$_ = fn(): \Foo|static => null;
	$_ = fn(): \Foo|void => null;
}

class Foo {
	public int|string $a;
	public int|string|float $a;
	public int|\Foo $a;
	public \Foo|\Boo $a;
	public \Namespaced\Foo|\Boo $a;
	public \Namespaced\Foo|\Boo|int $a;
	// public static $a; // static cannot be used here.

	public function f(int|\Foo $a) {}
	public function f(\Foo|int $a) {}
	public function f(\Foo|\Boo $a) {}
	public function f(\Namespaced\Foo|\Boo $a) {}
	// public function f(static|\Boo $a) {} // static cannot be used here.

	public function f(int|\Foo $a = null) {}
	public function f(\Foo|int $a = null) {}
	public function f(\Foo|\Boo $a = null) {}
	public function f(\Namespaced\Foo|\Boo $a = null) {}

	public function f(): int|string {}
	public function f(): int|\Boo {}
	public function f(): \Foo|\Boo {}
	public function f(): \Namespaced\Foo|\Boo {}
	public function f(): int|static {} // static can be used here.
}
	`

	actual := printPHP8(parsePHP8(src))

	assert.Equal(t, src, actual)
}

func TestParseAndPrintTryCatchWithoutVariablePHP8(t *testing.T) {
	src := `<?php
try {} catch (Exception) {}
try {} catch (Exception|Exception2) {}
try {} catch (Exception|Exception2|Exception3) {}

try {} catch (Exception) {} catch (Exception2) {}
try {} catch (Exception|Exception2) {} catch (Exception3|Exception4) {}
try {} catch (Exception|Exception2|Exception3) {} catch (Exception4) {}

// With variable.
try {} catch (Exception $a) {}
try {} catch (Exception|Exception2 $a) {}
try {} catch (Exception|Exception2|Exception3 $ax) {}

try {} catch (Exception $a) {} catch (Exception2 $a) {}
try {} catch (Exception|Exception2 $a) {} catch (Exception3|Exception4 $a) {}
try {} catch (Exception|Exception2|Exception3 $a) {} catch (Exception4 $a) {}
	`

	actual := printPHP8(parsePHP8(src))

	assert.Equal(t, src, actual)
}

func TestParseAndPrintParametersWithTrailingCommaPHP8(t *testing.T) {
	src := `<?php
function f(
  $a,
  $b,
) {}

function f(
  $a,
) {}


class Foo {
  public function f(
    $a,
    $b,
  ) {}

  public function f(
    $a,
  ) {}
}

function f() {
  $_ = function(
    $a,
    $b,
  ) {};

  $_ = function(
    $a,
  ) {};

  $_ = fn(
    $a,
    $b,
  ) => 10;

  $_ = fn(
    $a,
  ) => 10;
}

// Without comma.
function f(
  $a,
  $b
) {}

function f(
  $a
) {}


class Foo {
  public function f(
    $a,
    $b
  ) {}

  public function f(
    $a
  ) {}
}

function f() {
  $_ = function(
    $a,
    $b
  ) {};

  $_ = function(
    $a
  ) {};

  $_ = fn(
    $a,
    $b
  ) => 10;

  $_ = fn(
    $a
  ) => 10;
}
	`

	actual := printPHP8(parsePHP8(src))

	assert.Equal(t, src, actual)
}

func TestParseAndPrintThrowExprAndStmtPHP8(t *testing.T) {
	src := `<?php
$a ??= throw new InvalidArgumentException();
$a && throw new InvalidArgumentException();
$a || throw new InvalidArgumentException();

$a = $b ?: throw new InvalidArgumentException();

$a = !empty($b)
    ? reset($b)
    : throw new InvalidArgumentException();

throw $this->createNotFoundException();

throw $exception = new Exception();

// As stmt.
throw new InvalidArgumentException();
	`

	actual := printPHP8(parsePHP8(src))

	assert.Equal(t, src, actual)
}

func TestParseAndPrintPropertyInConstructorPHP8(t *testing.T) {
	src := `<?php
class Point {
    public function __construct(
        protected float $x = 0.0,
        protected float $y = 0.0,
        public float $z = 0.0,
    ) {}
}

class Point2 {
    public function __construct(
        private float $x,
        protected float $y,
        public float $z,
    ) {}
}

class Point3 {
    public function __construct(
        protected $x,
        private $y,
        public $z,
    ) {}
}

class Point4 {
    public function __construct(
        public &$x,
        private &$y,
        protected &$z,
    ) {}
}


class Point4 {
    public function __construct(
        public float &$x = 0.0,
        protected float &$y = 0.0,
        private float &$z = 0.0,
    ) {}
}

class Point5 {
	public function __construct(
		public $x,
		public float $y,
		protected float &$z,
		private float $a = 0.0,
		protected float &$b = 0.0,
	) {}
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClass2PHP8(t *testing.T) {
	src := `<?php
class Foo {}
abstract class Foo {}
final class Foo {}
abstract final class Foo {}

abstract final class Foo extends Boo {}
abstract class Foo extends Boo implements Goo {}
final class Foo extends Boo implements Goo, Doo {}
class Foo implements Goo {}
class Foo implements Goo, Doo {}

class Foo {
	public function f() {}
}

abstract final class Foo extends Boo implements Goo, Doo  {
	public function f() {}
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTrait2PHP8(t *testing.T) {
	src := `<?php
trait Foo {}

trait Foo {
	public function f() {}
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInterface2PHP8(t *testing.T) {
	src := `<?php
interface Foo {}
interface Foo extends Boo {}
interface Foo extends Boo, Goo {}

interface Foo {
	public function f() {}
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassWithAttributesPHP8(t *testing.T) {
	src := `<?php
#[SimpleAttribute]
class Foo {}

#[AttributeWithArgs(100, SomeClass::class)]
class Foo extends Boo {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
class Foo implements Goo {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
abstract class Foo implements Goo, Doo {}

#[
	FirstLineAttribute(100, SomeClass::class), 
	SecondLineAttribute()
]
class Foo {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
class Foo extends Boo implements Goo {}

#[\WithSlash]
class Foo {}

#[\F\Q\N]
class Foo {}

#[\WithSlash]
final class Foo extends Boo implements Goo, Doo {}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintTraitWithAttributesPHP8(t *testing.T) {
	src := `<?php
#[SimpleAttribute]
trait Foo {}

#[AttributeWithArgs(100, SomeClass::class)]
trait Foo {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
trait Foo {}

#[
	FirstLineAttribute(100, SomeClass::class), 
	SecondLineAttribute()
]
trait Foo {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
trait Foo {}

#[\WithSlash]
trait Foo {}

#[\F\Q\N]
trait Foo {}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintInterfaceWithAttributesPHP8(t *testing.T) {
	src := `<?php
#[SimpleAttribute]
interface Foo {}

#[AttributeWithArgs(100, SomeClass::class)]
interface Foo extends Boo {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
interface Foo {}

#[
	FirstLineAttribute(100, SomeClass::class), 
	SecondLineAttribute()
]
interface Foo {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
interface Foo extends Boo, Goo {}

#[\WithSlash]
interface Foo {}

#[\F\Q\N]
interface Foo {}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintParamsWithAttributesPHP8(t *testing.T) {
	src := `<?php
class Foo {
	public function __construct(#[SimpleAttribute] public string|int $a) {}
	public function __construct(
		#[SimpleAttribute] public string|int $a,
		#[AttributeWithArgs(100, SomeClass::class)] public string|int $b,
		#[
			FirstLineAttribute(100, SomeClass::class), 
			SecondLineAttribute()
		]
		$c,
		#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
		#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
		private \Foo|string $d = 100,

		#[
			FirstAttribute(100, SomeClass::class), 
			SecondAttribute(),
		]
		#[
			SecondGroupFirstAttribute(SomeClass::class), 
			SecondGroupSecondAttribute
		]
		$e = 100,

		#[\F\Q\N] string|int $f = 100,
	) {}
}

function f(#[SimpleAttribute] $a) {}

function f(
	#[SimpleAttribute] string|int $a,
	#[AttributeWithArgs(100, SomeClass::class)] string|int $b,
	#[
		FirstLineAttribute(100, SomeClass::class), 
		SecondLineAttribute()
	]
	$c,
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
	\Foo|string $d = 100,

	#[
		FirstAttribute(100, SomeClass::class), 
		SecondAttribute(),
	]
	#[
		SecondGroupFirstAttribute(SomeClass::class), 
		SecondGroupSecondAttribute
	]
	$e = 100,

	#[\F\Q\N] string|int $f = 100,
) {}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintFunctionWithAttributesPHP8(t *testing.T) {
	src := `<?php
#[SimpleAttribute]
function &f(): void {}

#[AttributeWithArgs(100, SomeClass::class)]
function f() {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
function &f(): void {}

#[
	FirstLineAttribute(100, SomeClass::class), 
	SecondLineAttribute()
]
function f() {}

#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
function &f(): int|string {}

#[\WithSlash]
function f() {}

#[\F\Q\N]
function &f(): \Foo|string {}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClosureWithAttributesPHP8(t *testing.T) {
	src := `<?php
function f() {
	$_ =
		#[SimpleAttribute]
		function &(): void {};

	$_ = 
		#[AttributeWithArgs(100, SomeClass::class)]
		function () {};

	$_ = 
		#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
		function &(): void {};

	$_ = 
		#[
			FirstLineAttribute(100, SomeClass::class), 
			SecondLineAttribute()
		]
		function () {};

	$_ = 
		#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
		#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
		function &(): int|string {};

	$_ = 
		#[\WithSlash]
		function () {};

	$_ = 
		#[\F\Q\N]
		function &(): \Foo|string {};
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintArrowFunctionWithAttributesPHP8(t *testing.T) {
	src := `<?php
function f() {
	$_ =
		#[SimpleAttribute]
		fn &(): void => 1;

	$_ = 
		#[AttributeWithArgs(100, SomeClass::class)]
		fn () => 1;

	$_ = 
		#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
		fn &(): void => 1;

	$_ = 
		#[
			FirstLineAttribute(100, SomeClass::class), 
			SecondLineAttribute()
		]
		fn () => 1;

	$_ = 
		#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
		#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
		fn &(): int|string => 1;

	$_ = 
		#[\WithSlash]
		fn () => 1;

	$_ = 
		#[\F\Q\N]
		fn &(): \Foo|string => 1;
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintAnonClassWithAttributesPHP8(t *testing.T) {
	src := `<?php
$_ = new
	#[SimpleAttribute]
	class {};

$_ = new
	#[AttributeWithArgs(100, SomeClass::class)]
	class extends Boo {};

$_ = new
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	class implements Goo {};

$_ = new
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	class implements Goo, Doo {};

$_ = new
	#[
		FirstLineAttribute(100, SomeClass::class), 
		SecondLineAttribute()
	]
	class {};

$_ = new
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
	class extends Boo implements Goo {};

$_ = new
	#[\WithSlash]
	class {};

$_ = new
	#[\F\Q\N]
	class {};

$_ = new
	#[\WithSlash]
	class extends Boo implements Goo, Doo {};
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassPropertyWithAttributesPHP8(t *testing.T) {
	src := `<?php
class Foo {
	#[SimpleAttribute]
	public $a;
	
	#[AttributeWithArgs(100, SomeClass::class)]
	protected $a = null;
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	var $a = null, $b;
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	private ?Foo $a = null;
	
	#[
		FirstLineAttribute(100, SomeClass::class), 
		SecondLineAttribute()
	]
	public int|string $a = 10, $b, $c, $d;
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
	var $a = 100;
	
	#[\WithSlash]
	public $a = null;
	
	#[\F\Q\N]
	protected $a = null;
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassConstWithAttributesPHP8(t *testing.T) {
	src := `<?php
class Foo {
	#[SimpleAttribute]
	public const A = 100;
	
	#[AttributeWithArgs(100, SomeClass::class)]
	private const A = 10, B = 10, C = 10;
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	protected const A = "a";
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	public const A = 100;
	
	#[
		FirstLineAttribute(100, SomeClass::class), 
		SecondLineAttribute()
	]
	public const A = 100;
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
	public const A = 100;
	
	#[\WithSlash]
	protected const A = 100;
	
	#[\F\Q\N]
	public const A = 100;
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}

func TestParseAndPrintClassMethodWithAttributesPHP8(t *testing.T) {
	src := `<?php
class Foo {
	#[SimpleAttribute]
	function &f(): void {}
	
	#[AttributeWithArgs(100, SomeClass::class)]
	public function f() {}
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	protected function &f(): void {}
	
	#[
		FirstLineAttribute(100, SomeClass::class), 
		SecondLineAttribute()
	]
	public function f() {}
	
	#[FirstAttribute(100, SomeClass::class), SecondAttribute()]
	#[SecondGroupFirstAttribute(SomeClass::class), SecondGroupSecondAttribute]
	protected function &f(): int|string {}
	
	#[\WithSlash]
	public function f() {}
	
	#[\F\Q\N]
	protected function &f(): \Foo|string {}
}
`

	actual := printPHP8(parsePHP8(src))

	if src != actual {
		t.Errorf("\nexpected: %s\ngot: %s\n", src, actual)
	}
}
