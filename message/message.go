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
func (msg *Message) validate()bool{
	//initializing a map named Errors
msg.Errors=make( map[string]string)
//if the email match the regular expressions
match:=rxEmail.Match([]byte(msg.Email))
if match ==false{
	msg.Errors["email"]="please entre a valid email adress "
}
// strings.TrimSpace to remove leading and trailing whitespaces from a string.
if strings.TrimSpace(msg.Content)==""{
	msg.Errors["content"]="please enter a message "

}
return len(msg.Errors)==0

}