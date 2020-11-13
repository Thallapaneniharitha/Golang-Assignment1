1.//JSON(Javascript Object Notation)
package main    
import ( 
    "fmt"
    "encoding/json"
)   
type Human struct{         
   Name string 
    Age int
    Address string 
}   
func main() { 
       
        human1 := Human{"Haritha", 24, "Rajampet"} 
    human_enc, err := json.Marshal(human1) 
       
    if err != nil { 
           
        fmt.Println(err) 
    } 
    fmt.Println(string(human_enc)) 
    human2 := []Human{ 
        {Name: "Saritha", Age: 23, Address: "New Delhi"},
        {Name: "sarada", Age: 25, Address: "Pune"}, 
        {Name: "hari", Age: 27, Address: "Bangalore"}, 
    } 
       
    // encoding into JSON format 
    human2_enc, err := json.Marshal(human2) 
        
        if err != nil { 
       
            fmt.Println(err) 
        } 
    fmt.Println() 
        fmt.Println(string(human2_enc)) 
} 
$go run main.go
{"Name":"Haritha","Age":24,"Address":"Rajampet"}
[{"Name":"Saritha","Age":23,"Address":"New Delhi"},{"Name":"sarada","Age":25,"Address":"Pune"},{"Name":"hari","Age":27,"Address":"Bangalore"}]