package Server

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"timesheet/AuthPack"
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
	var ArrayOfFilterFunction [7]FilterPack.ArrayStructString
	ArrayOfFilterFunction = FilterPack.GetDaysOfWeekForStrings(FilterPack.FilterFunctionWithGroup(Filter.Filter, Filter.Value))
	var SortedAdditionalArray = FilterPack.AdditionSubjectToEachDayForStrings(ArrayOfFilterFunction)
	var results ProtoApi.SheduleArrayByWeek
	for i := 0; i < 7; i++ {
		var tArray ProtoApi.SheduleArrayByDay
		for j := 0; j < 8; j++ {
			var t ProtoApi.StringSheduleObject
			t.Auditorium = SortedAdditionalArray[i].SubjectsOfThisDay[j].Auditorium
			t.Tutor = SortedAdditionalArray[i].SubjectsOfThisDay[j].SecondName + " " + SortedAdditionalArray[i].SubjectsOfThisDay[j].FirstName + " " + SortedAdditionalArray[i].SubjectsOfThisDay[j].LastName
			t.Type = SortedAdditionalArray[i].SubjectsOfThisDay[j].Type
			t.Subject = SortedAdditionalArray[i].SubjectsOfThisDay[j].Subject
			t.Number = SortedAdditionalArray[i].SubjectsOfThisDay[j].Number
			t.Group = SortedAdditionalArray[i].SubjectsOfThisDay[j].Group
			t.Dates = &ProtoApi.DateTime{
				Year:    int32(SortedAdditionalArray[i].SubjectsOfThisDay[j].Dates.Year()),
				Month:   int32(SortedAdditionalArray[i].SubjectsOfThisDay[j].Dates.Month()),
				Day:     int32(SortedAdditionalArray[i].SubjectsOfThisDay[j].Dates.Day()),
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			}
			tArray.ObjectsString = append(tArray.ObjectsString, &t)
		}
		results.DaysOfObjects = append(results.DaysOfObjects, &tArray)
	}
	return &results, nil
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

func (G GRPCServer) CheckAuth(ctx context.Context, AuthCtx *ProtoApi.AuthorizationContext) (*ProtoApi.AuthorizationResult, error) {
	fmt.Println(AuthCtx.Login, " ", AuthCtx.Pass)
	DefStruct, RoleResult := AuthPack.Authorization(AuthCtx.Login, AuthCtx.Pass)
	return &ProtoApi.AuthorizationResult{IsAccept: strconv.FormatBool(DefStruct), Role: RoleResult}, nil
}

func (G GRPCServer) GetSheduleFromDbForTutor(ctx context.Context, Filter *ProtoApi.Filter) (*ProtoApi.SheduleArrayByWeek, error) {
	var ArrayOfFilterFunction [7]FilterPack.ArrayStructString
	ArrayOfFilterFunction = FilterPack.GetDaysOfWeekForStrings(FilterPack.FilterFunction(Filter.Filter, Filter.Value))
	var SortedAdditionalArray = FilterPack.AdditionSubjectToEachDayForStrings(ArrayOfFilterFunction)
	var results ProtoApi.SheduleArrayByWeek
	for i := 0; i < 7; i++ {
		var tArray ProtoApi.SheduleArrayByDay
		for j := 0; j < 8; j++ {
			var t ProtoApi.StringSheduleObject
			t.Auditorium = SortedAdditionalArray[i].SubjectsOfThisDay[j].Auditorium
			t.Tutor = SortedAdditionalArray[i].SubjectsOfThisDay[j].SecondName + " " + SortedAdditionalArray[i].SubjectsOfThisDay[j].FirstName + " " + SortedAdditionalArray[i].SubjectsOfThisDay[j].LastName
			t.Type = SortedAdditionalArray[i].SubjectsOfThisDay[j].Type
			t.Subject = SortedAdditionalArray[i].SubjectsOfThisDay[j].Subject
			t.Number = SortedAdditionalArray[i].SubjectsOfThisDay[j].Number
			t.Group = SortedAdditionalArray[i].SubjectsOfThisDay[j].Group
			t.Dates = &ProtoApi.DateTime{
				Year:    int32(SortedAdditionalArray[i].SubjectsOfThisDay[j].Dates.Year()),
				Month:   int32(SortedAdditionalArray[i].SubjectsOfThisDay[j].Dates.Month()),
				Day:     int32(SortedAdditionalArray[i].SubjectsOfThisDay[j].Dates.Day()),
				Hours:   0,
				Minutes: 0,
				Seconds: 0,
			}
			tArray.ObjectsString = append(tArray.ObjectsString, &t)
		}
		results.DaysOfObjects = append(results.DaysOfObjects, &tArray)
	}
	return &results, nil
}
