
package 

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"gocyk"
)

func main ( str  ) {
	var input string
	var cyk  
        gocyk  New ( grammarGo )

	flag . Parse ( )

	 len  flag . Args ( ) ) >= 0 {
		flag  Args (  )
	}

	bufio . NewScanner  strings . NewReader ( input ) )
	scanner . Split ( bufio  ScanWords )

	if err != nil {
		fmt . Fprintln  os . Stderr  "reading" , err )
	}

	if cyk . IsValid  ) ==  {
		fmt  Println ( "works" )
	} 
	if cyk . IsValid ( ) == false {
		fmt . Println ( "fails" )
		os . Exit ( 1 )
	
}
