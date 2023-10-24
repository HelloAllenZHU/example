/*
 * 本程序演示了Go调用Windows dll
 */
package main

/*
 #include "LibTest.h"
 #cgo LDFLAGS: -L${SRCDIR} -lTest
*/
import "C"

func main() {
    C.Add( 1, 2 )
}
