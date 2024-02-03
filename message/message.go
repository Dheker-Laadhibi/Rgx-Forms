package message
import(
	//regular exepression (validate some inputs)
	"regexp"
	"strings"
)
/* 
.+: Matches one or more of any character (except for a newline).
@: Matches the at symbol.
.+: Matches one or more of any character.
\\.: Matches a literal dot (escaped with a backslash, as . has a special meaning in regular expressions). */
//algorithme@gmail.com

var rxEmail= regexp.MustCompile(".+@.+\\..+")
type Message struct{
	Email string
	Content string
	Errors map[string]string
}