package main

import(
	"math/rand"
	"time"
	"os"
	"fmt"
	"bufio"
	"strconv"
)

var (
	uids []int64
	maxUid int64
	minUid int64
	filename string
)


func main(){
	maxUid = 1000000000
	minUid = 0
	filename = "./uid.txt"
	
	var file *os.File
	var err error
	if checkFileIsExist(filename){
		file, err = os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)  
	}else{
		file, err = os.Create(filename)
	}
    if err != nil {  
        fmt.Println(err)  
        return 
    }  
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		uid,err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil{
			fmt.Println(err)
			return 
		}
		uids = append(uids, uid)
	}
   
    writer := bufio.NewWriter(file)  
	uids=make([]int64,0)
	var uid int64
	uid=createUser()
	writer.WriteString(strconv.FormatInt(uid,10))  
    writer.WriteString("\n")  
    writer.Flush()  
	uids = append(uids, uid)

	return 
}

func findUid(uid int64) bool {
	for _, id := range uids {
		if uid == id{
			return true
		}
	}
	return false
}

func createUser() int64{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var uid int64
	uid = r.Int63n(maxUid-minUid)+minUid
	for findUid(uid) {
		uid = r.Int63n(maxUid-minUid)+minUid
	}
	return uid
}

func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}