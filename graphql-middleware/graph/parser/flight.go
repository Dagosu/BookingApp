package parser

import (
	dt "github.com/Dagosu/BookingApp/datatypes"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/model"
)

func ParseFlight(f *dt.Flight) *model.Flight {
	if f == nil {
		return nil
	}

	return &model.Flight{
		ID:            f.GetId(),
		Departure:     StrRefer(f.GetDeparture()),
		DepartureTime: ParsePbTimestamp(f.GetDepartureTime()),
		Arrival:       StrRefer(f.GetArrival()),
		ArrivalTime:   ParsePbTimestamp(f.GetArrivalTime()),
		TotalSeats:    IntRefer(f.GetTotalSeats()),
		BookableSeats: IntRefer(f.GetBookableSeats()),
		Airline:       StrRefer(f.GetAirline()),
		Price:         FloatRefer(float64(f.GetPrice())),
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

func FlightStreamToChan(c chan *model.FlightListResponse, o *dt.FlightListResponse, flights *[]*model.Flight, ready *bool) {
	// log.Println("FlightList received msg, pushing to channel", f)

	opt := model.OperationType(o.OperationType.String())

	fl := ParseFlight(o.Flight)
	if !*ready {
		if fl != nil {
			*flights = append(*flights, fl)
		}

		*ready = o.OperationType == dt.OperationType_READY

		if *ready {
			optInsert := model.OperationTypeInsert

			c <- &model.FlightListResponse{
				OperationType: &optInsert,
				Flights:       *flights,
			}

			*flights = []*model.Flight{}

			// send a separate READY message
			c <- &model.FlightListResponse{
				OperationType: &opt,
			}
		}

		return
	}

	flightsList := []*model.Flight{fl}

	c <- &model.FlightListResponse{
		OperationType: &opt,
		Flights:       flightsList,
	}
}
