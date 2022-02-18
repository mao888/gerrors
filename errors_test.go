package gerrors

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {

	//fmt.Printf("%+v\n", err1)
	//err1 = Wrap(err1, "exec0 wrap")
	//err1 = Wrap(err1, "exec1 wrap")/
	//fmt.Printf("%s\n", wrap2().Error())
	//fmt.Printf("%#+v\n", err1)
	err := wrap2().Error()
	//fmt.Println(fmt.Sprintf("%+v \n", err))
	//fmt.Printf("%#+v\n", err)
	fmt.Printf("%s\n", err)
	//glog.Error(context.TODO(), err)
	fmt.Printf("----------------------------------------------------------- \n")
	//glog.Errorf(context.TODO(), "%s\n", err)

}

func wrap2() error {
	if err := wrap1(); err != nil {
		return Wrap(err, "exec2 wrap")
	}
	return nil
}

func wrap1() error {
	if err := wrap0(); err != nil {
		return Wrap(err, "exec1 wrap")
	}
	return nil
}

func wrap0() error {
	if err := openfile(); err != nil {
		return Wrap(err, "exec0 wrap")
	}
	return nil
}

func openfile() error {
	if _, err := os.Open("1"); err != nil {
		//glog.Error(context.TODO(), err)
		return err
	}
	return nil
}
