package main
 
import (
	"fmt"
	"strings"
)

func main(){

str:="abcabc"
substr:="c"

allSubStrings:=subStrings(str,len(str))
count:=0

for _,v := range allSubStrings{

	if strings.Contains(v,substr) == true{
		//fmt.Println(v)
		count++
	}
}
fmt.Println(count)

//fmt.Println(allSubStrings)

}

func subStrings(str string, n int)([]string) {

	slice:=[]string{}
	//starting point of string
	for len:=1;len<=n;len++{

		//fmt.Printf("len %d",len)
		//ending point of sring

		for i:=0;i<=n-len;i++{
			//fmt.Printf("i= %d",i)
			j:=i+len-1
			var substr string
			//fmt.Printf("j= %d",j)
			for k:=i;k<=j;k++{
				
					//fmt.Printf("%c",str[k])
					substr=substr+string(str[k])	
						
			}
			slice=append(slice,substr)
			//fmt.Println()
		}

	}
	return slice
	//fmt.Println(slice)
}
