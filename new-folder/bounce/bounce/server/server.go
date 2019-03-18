package server

import (
	"database/sql"
	"log"
	"relay/reports/bounce/bounce/models"
	//"github.com/wizelineacademy/GoWorkshop/bounce/models"
	"relay/reports/bounce/proto/bounce"

	"golang.org/x/net/context"
)

type Env struct {
	db *sql.DB
}
type Server struct{}

// I have to add server struct here as well....
//
// func init() {
// 	db, err := models.NewDB()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	env := &Env{db: db}
// 	bounces, err := models.BounceGetData(env.db)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// bounces, err := models.BounceGetData(env.db)
// 	// bounces, err := models.DB.BounceGetData()
// 	response := new(bounce.BounceGetDataResponse)
// 	fmt.Println("this is Bonces", bounces)
// 	fmt.Println("this is Bonces", response)
// 	return
// }
//
// type BounceGetDataResponse struct {
// 	Reason                 string   `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
// 	BounceClassName        string   `protobuf:"bytes,2,opt,name=bounce_class_name,json=bounceClassName,proto3" json:"bounce_class_name,omitempty"`
// 	BounceClassDescription string   `protobuf:"bytes,3,opt,name=bounce_class_description,json=bounceClassDescription,proto3" json:"bounce_class_description,omitempty"`
// 	OunceCategoryId        int32    `protobuf:"varint,4,opt,name=ounce_category_id,json=ounceCategoryId,proto3" json:"ounce_category_id,omitempty"`
// 	BounceCategoryName     string   `protobuf:"bytes,5,opt,name=bounce_category_name,json=bounceCategoryName,proto3" json:"bounce_category_name,omitempty"`
// 	ClassificationId       int32    `protobuf:"varint,6,opt,name=classification_id,json=classificationId,proto3" json:"classification_id,omitempty"`
// 	CountInbandBounce      int32    `protobuf:"varint,7,opt,name=count_inband_bounce,json=countInbandBounce,proto3" json:"count_inband_bounce,omitempty"`
// 	CountOutofbandBounce   int32    `protobuf:"varint,8,opt,name=count_outofband_bounce,json=countOutofbandBounce,proto3" json:"count_outofband_bounce,omitempty"`
// 	CountBounce            int32    `protobuf:"varint,9,opt,name=count_bounce,json=countBounce,proto3" json:"count_bounce,omitempty"`
// 	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
// 	XXX_unrecognized       []byte   `json:"-"`
// 	XXX_sizecache          int32    `json:"-"`
// }

func (s *Server) GetData(ctx context.Context, in *bounce.BounceGetData) (*bounce.BounceGetDataResponse, error) {
	db, err := models.NewDB()
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}
	bounces, err := models.BounceGetData(env.db)
	response := new(bounce.BounceGetDataResponse)
	for i := 0; i < len(bounces); i++ {
		response.AccountID = bounces[i].AccountID
		response.BouncesReason = bounces[i].BonuceReason
		response.BounceClassName = bounces[i].BounceClassName
		response.ClassificationID = bounces[i].ClassificationID
		response.BounceType = bounces[i].BounceType
		response.BounceClassDesc = bounces[i].BounceClassDesc
		response.Count = bounces[i].Count
	}

	return response, err
}
