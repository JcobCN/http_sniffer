package file

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var configFile = "./windoc"

func WriteConfigInit(){

	configFileState := false

	_, err := os.Stat(configFile)
	fmt.Println(err)
	if err == nil{
		configFileState = true
	} else if err != nil{
		if os.IsExist(err){
			configFileState = true
		}
	}

	if configFileState == false{
		fp, err := os.Create(configFile)
		defer fp.Close()

		if err != nil{
			fmt.Println(err.Error())
			return
		}else{
			_, err = fp.Write([]byte("write=1\n"))
			if err != nil{
				fmt.Println("write error")
				return
			}
		}
	}
}

func WriteConfigRead() (canWrite bool,err error)  {

	fp, err := os.Open(configFile)
	defer fp.Close()
	if err != nil{
		fmt.Println(err.Error())
		return canWrite, err
	}

	b := make([]byte, 128)
	n, _ :=fp.Read(b)
	fmt.Println(n, string(b))

	writeStr := string(b)
	writeValueArry := strings.Split(writeStr,"=")
	if strings.Contains(writeValueArry[1], "1") {
		return true, err
	}

	return canWrite, err
}

func ModifiedWriteValue(value int){
	fp, err := os.OpenFile(configFile, os.O_RDWR, 0)
	defer fp.Close()
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fp.Seek(0,os.SEEK_SET)

	b := make([]byte,128)
	n,_ := fp.Read(b)
	fmt.Println(n, string(b))

	re,_ := regexp.Compile(`write=\d`)
	replaceStr := fmt.Sprintf("write=%d", value)
	reb := re.ReplaceAll(b, []byte(replaceStr))
	fmt.Println("reb:", string(reb))

	fp.Seek(0, os.SEEK_SET)

	fp.Write(reb)

}
