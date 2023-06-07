package parser

import (
	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
)

func ParseMyObject(o *dt.MyObject) *model.MyObject {
	if o == nil {
		return nil
	}

	return &model.MyObject{
		ID:   o.GetId(),
		Name: StrRefer(o.GetName()),
		Date: ParsePbTimestamp(o.GetDate()),
	}
}

func ParseSort(in []*model.SortParamInput) []*dt.SortParam {
	out := make([]*dt.SortParam, 0, len(in))
	for _, i := range in {
		so := dt.ViewSortOrder_ASC
		if *i.Order == model.ViewSortOrderDesc {
			so = dt.ViewSortOrder_DESC
		}
		out = append(out, &dt.SortParam{
			Field: StrDerefer(i.Field),
			Order: so,
		})
	}

	return out
}

func ParseFilter(in []*model.FilterParamInput) []*dt.FilterParam {
	out := make([]*dt.FilterParam, 0, len(in))
	for _, i := range in {
		out = append(out, &dt.FilterParam{
			Condition: StrDerefer(i.Condition),
			Field:     StrDerefer(i.Field),
			Operator:  StrDerefer(i.Operator),
			Value:     StrDerefer(i.Value),
		})
	}

	return out
}

func ObjectsStreamToChan(c chan *model.TestListResponse, o *dt.TestListResponse, myObjects *[]*model.MyObject, ready *bool) {
	// log.Println("FlightList received msg, pushing to channel", f)

	opt := model.OperationType(o.OperationType.String())

	myO := ParseMyObject(o.Object)
	if !*ready {
		if myO != nil {
			*myObjects = append(*myObjects, myO)
		}

		*ready = o.OperationType == dt.OperationType_READY

		if *ready {
			optInsert := model.OperationTypeInsert

			c <- &model.TestListResponse{
				OperationType: &optInsert,
				Objects:       *myObjects,
			}

			*myObjects = []*model.MyObject{}

			// send a separate READY message
			c <- &model.TestListResponse{
				OperationType: &opt,
			}
		}

		return
	}

	myObjectList := []*model.MyObject{myO}

	c <- &model.TestListResponse{
		OperationType: &opt,
		Objects:       myObjectList,
	}
}
