package Server

import (
	"context"
	"fmt"
	"timesheet/GetInfo"
	"timesheet/ProtoApi"
)

type GRPCServer struct {
	ProtoApi.SheduleServiceServer
}

func (G GRPCServer) mustEmbedUnimplementedSheduleServiceServer() {
	//TODO implement me
	panic("implement me")
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
		fmt.Println(i.Type)
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
