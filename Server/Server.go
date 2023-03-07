package Server

import (
	"context"
	"time"
	"timesheet/FilterPack"
	"timesheet/GetInfo"
	"timesheet/ProtoApi"
	"timesheet/SetInfo"
)

type GRPCServer struct {
	ProtoApi.SheduleServiceServer
}

func (G GRPCServer) GetSubjectFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.SubjectsFromDb, error) {
	DefStruct := GetInfo.GetSubjectFromDb()
	var results []*ProtoApi.StructSubject
	for _, i := range DefStruct {
		var temp = ProtoApi.StructSubject{SubjectItem: i.SubjectItem, Id: i.Id}
		results = append(results, &temp)
	}
	return &ProtoApi.SubjectsFromDb{Subjects: results}, nil
}

func (G GRPCServer) GetTutorsFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.TutorsFromDb, error) {
	DefStruct := GetInfo.GetTutorsFromDb()
	var results []*ProtoApi.StructTutor
	for _, i := range DefStruct {
		var temp = ProtoApi.StructTutor{Id: i.Id, SecondName: i.SecondName, FirstName: i.FirstName, LastName: i.LastName}
		results = append(results, &temp)
	}
	return &ProtoApi.TutorsFromDb{Tutors: results}, nil
}

func (G GRPCServer) GetAuditoriumFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.AuditoriumsFromDb, error) {
	DefStruct := GetInfo.GetAuditoriumFromDb()
	var results []*ProtoApi.StructClassroom
	for _, i := range DefStruct {
		var temp = ProtoApi.StructClassroom{Id: i.Id, Classroom: i.Classroom}
		results = append(results, &temp)
	}
	return &ProtoApi.AuditoriumsFromDb{Auditoriums: results}, nil
}

func (G GRPCServer) GetTypeFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.TypesFromDb, error) {
	DefStruct := GetInfo.GetTypeFromDb()
	var results []*ProtoApi.StructType
	for _, i := range DefStruct {
		var temp = ProtoApi.StructType{Id: i.Id, Type: i.Type}
		results = append(results, &temp)
	}
	return &ProtoApi.TypesFromDb{Types: results}, nil
}

func (G GRPCServer) GetGroupFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.GroupsFromDb, error) {
	DefStruct := GetInfo.GetGroupFromDb()
	var results []*ProtoApi.StructGroup
	for _, i := range DefStruct {
		var temp = ProtoApi.StructGroup{Id: i.Id, GroupName: i.GroupName}
		results = append(results, &temp)
	}
	return &ProtoApi.GroupsFromDb{Groups: results}, nil
}

func (G GRPCServer) GetSheduleFromDb(ctx context.Context, Filter *ProtoApi.Filter) (*ProtoApi.SheduleArrayByWeek, error) {
	var ArrayOfFilterFunction [7]FilterPack.ArrayStruct
	ArrayOfFilterFunction = FilterPack.GetDayOfWeek(FilterPack.FilterFunction(Filter.Filter, Filter.Value))
	var results *ProtoApi.SheduleArrayByWeek
// 	for j := 0; j < 7; j++ {
// 		for _, i := range ArrayOfFilterFunction {
// 			var t *ProtoApi.SheduleObject
// 			t.Auditorium = i.Arr[j].Auditorium
// 			t.Tutor = i.Arr[j].Tutor
// 			t.Type = i.Arr[j].Type
// 			t.Subject = i.Arr[j].Subject
// 			t.Number = i.Arr[j].Number
// 			t.Group = i.Arr[j].Group
// 			results.DaysOfObjects[j].Objects = append(results.DaysOfObjects[j].Objects, t)
// 		}
// 	}
	for i := 0; i < 7; i++ {
		for j := 0; j < 8; j++{
			var t *ProtoApi.SheduleObject
			t.Auditorium = ArrayOfFilterFunction[i].Arr[j].Auditorium
			t.Tutor = ArrayOfFilterFunction[i].Arr[j].Tutor
			t.Type = ArrayOfFilterFunction[i].Arr[j].Type
			t.Subject = ArrayOfFilterFunction[i].Arr[j].Subject
			t.Number = ArrayOfFilterFunction[i].Arr[j].Number
			t.Group = ArrayOfFilterFunction[i].Arr[j].Group
			t.Dates = ArrayOfFilterFunction[i].Arr[j].Dates
		}
		results.DaysOfObjects[i].Objects = append(results.DaysOfObjects[i].Objects, t)
	}
	return results, nil
}

func (G GRPCServer) AddShedule(ctx context.Context, SheduleArray *ProtoApi.AllSheduleArray) (*ProtoApi.Wrap, error) {
	var results []*SetInfo.ObjectPattern
	for _, i := range SheduleArray.Objects {
		var t SetInfo.ObjectPattern
		t.Auditorium = i.Auditorium
		t.Tutor = i.Tutor
		t.Type = i.Type
		t.Subject = i.Subject
		t.Number = i.Number
		t.Group = i.Group
		t.Dates = time.Date(int(i.Dates.Year), time.Month(i.Dates.Month), int(i.Dates.Day), int(i.Dates.Hours), int(i.Dates.Minutes), int(i.Dates.Seconds), 0, time.UTC)
		results = append(results, &t)
	}
	SetInfo.InsertionToDb(results)
	return &ProtoApi.Wrap{}, nil
}
