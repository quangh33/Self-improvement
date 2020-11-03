package main

import (
	complexpb "complex"
	day_of_week_pb "day_of_week"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"simple"
)

func main() {
	//sm := doSimple()
	//fmt.Println(sm)
	//writeToFile("simple.bin", sm)
	//sm2 := &simplepb.Simple{}
	//readFromFile("simple.bin", sm2)
	//fmt.Println(sm2)
	//smAsString := toJson(sm2)
	//fmt.Println(smAsString)
	doEnum()

	cm := complexpb.ComplexMessage{OneDummy: &complexpb.DummyMessage{Name: "Quang", Id: 1}}
	oneDummy := cm.GetOneDummy()
	fmt.Println(oneDummy)
	fmt.Println(cm)
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0600); err != nil {
		log.Fatalln("Can't write to file", err)
	}

	return nil
}

func readFromFile(filename string, pb proto.Message) error {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Cannot unmarshal byte to proto buf messaeg ", err2)
		return err
	}
	return nil
}

func doSimple() *simplepb.Simple {
	sm := simplepb.Simple{
		Id: 1,
		Name: "Quang",
		List: []int32{1,2,3},
	}

	sm.Name = "Nick"
	return &sm
}

func toJson(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cannot convert to JSON", err)
		return ""
	}
	return out
}

func fromJson(in string, message proto.Message) error {
	err := jsonpb.UnmarshalString(in, message)
	if err != nil {
		log.Fatalln("Cannot convert JSON to pb", err)
		return err
	}
	return nil
}


func doEnum() {
	day := day_of_week_pb.DayOfWeek{
		Day: day_of_week_pb.Days_SUNDAY,
	}

	fmt.Println(&day)
}