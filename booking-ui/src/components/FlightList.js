import React, { useState } from 'react';
import { useSubscription } from '@apollo/client';
import { useForm } from 'react-hook-form';
import gql from 'graphql-tag';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import '../style/FlightList.css';
import { Link } from 'react-router-dom';

const FLIGHTS_SUBSCRIPTION = gql`
  subscription ($in: FlightListInput!) {
    flightList(in: $in) {
      flights {
        id
        departure
        departureTime {
          seconds
        }
        arrival
        arrivalTime {
          seconds
        }
        bookableSeats
      }
    }
  }
`;

function FlightList() {
  const { register, handleSubmit } = useForm();
  const [startTime, setStartTime] = useState(new Date());
  const [endTime, setEndTime] = useState(new Date());
  const [filter, setFilter] = useState({});
  const [searchText, setSearchText] = useState("");  
  const [status, setStatus] = useState({ scheduled: false, active: false, arrived: false });

  const { loading, error, data } = useSubscription(FLIGHTS_SUBSCRIPTION, {
    variables: {
      in: filter,
    },
  });

  const onSubmit = (data) => {
    let filterArray = [
      {
        condition: "and",
        field: "departure_time",
        operator: "gte",
        value: Math.floor(startTime.getTime() / 1000),
      },
      {
        condition: "and",
        field: "departure_time",
        operator: "lte",
        value: Math.floor(endTime.getTime() / 1000),
      }
    ];

    Object.keys(status).forEach(key => {
      if (status[key]) {
        filterArray.push({
          condition: "or",
          field: "status",
          operator: "eq",
          value: key
        });
      }
    });

    setFilter({
      filter: filterArray,
      query: searchText,  
    });
  };

  const clearDates = () => {
    setStartTime(null);
    setEndTime(null);
  }

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  return (
    <div className="flight-list-container">
      <form className="flight-filter-form" onSubmit={handleSubmit(onSubmit)}>
        <label className="flight-filter-label">
          Start Time:
          <DatePicker selected={startTime} onChange={date => setStartTime(date)} showTimeSelect dateFormat="Pp" className="flight-datepicker" />
        </label>
        <label className="flight-filter-label">
          End Time:
          <DatePicker selected={endTime} onChange={date => setEndTime(date)} showTimeSelect dateFormat="Pp" className="flight-datepicker" />
          <button onClick={clearDates} type="button" className="flight-filter-clear">Clear</button>
        </label>
        <label className="flight-filter-label"> {}
          Search:
          <input type="text" value={searchText} onChange={(e) => setSearchText(e.target.value)} className="flight-search-input" />
        </label>
        <label className="checkbox-container">
          <input type="checkbox" checked={status.scheduled} onChange={(e) => setStatus({ ...status, scheduled: e.target.checked })} />
          Scheduled
        </label>
        <label className="checkbox-container">
          <input type="checkbox" checked={status.active} onChange={(e) => setStatus({ ...status, active: e.target.checked })} />
          Active
        </label>
        <label className="checkbox-container">
          <input type="checkbox" checked={status.arrived} onChange={(e) => setStatus({ ...status, arrived: e.target.checked })} />
          Arrived
        </label>
        <input type="submit" value="Filter" className="flight-filter-submit" />
      </form>

      <div className="flight-list">
      {data && data.flightList && data.flightList.flights.map((flight) => (
        <Link to={`/flight/${flight.id}`} key={flight.id} className="flight-item-link">
        <div className="flight-item">
          <p>Departure: {flight.departure} at {new Date(flight.departureTime.seconds * 1000).toLocaleString()}</p>
          <p>Arrival: {flight.arrival} at {new Date(flight.arrivalTime.seconds * 1000).toLocaleString()}</p>
          <p>Bookable Seats: {flight.bookableSeats}</p>
        </div>
      </Link>
      ))}
      </div>
    </div>
  );
}

export default FlightList;
