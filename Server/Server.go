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
		fmt.Println(temp.Id, " ", temp.SubjectItem)
		results = append(results, &temp)
	}
	return &ProtoApi.SubjectsFromDb{Subjects: results}, nil
}

func (G GRPCServer) GetTutorsFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.TutorsFromDb, error) {
	//TODO implement me
	return &ProtoApi.TutorsFromDb{}, nil
}

func (G GRPCServer) GetAuditoriumFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.AuditoriumsFromDb, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) GetTypeFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.TypesFromDb, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) GetGroupFromDbRPC(ctx context.Context, wrap *ProtoApi.Wrap) (*ProtoApi.GroupsFromDb, error) {
	//TODO implement me
	panic("implement me")
}
