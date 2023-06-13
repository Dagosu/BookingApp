import React from 'react';
import { useQuery, gql } from '@apollo/client';
import { Link } from 'react-router-dom';
import '../style/RecommendedFlightList.css';

const RECOMMENDED_FLIGHTS_QUERY = gql`
  query ($in: RecommendFlightInput!) {
    recommendFlight(in: $in) {
      flights {
        id
        departure
        arrival
        bookableSeats
        airline
        price
      }
    }
  }
`;

function RecommendedFlightList() {
  const userId = localStorage.getItem('userId');
  const { loading, error, data } = useQuery(RECOMMENDED_FLIGHTS_QUERY, {
    variables: {
      in: {
        userId: userId
      }
    },
  });

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  return (
    <div className="recommended-flight-list-container">
      <h2 className="recommended-flight-list-title">Recommended Flights for You</h2>
      <div className="recommended-flight-list">
      {data && data.recommendFlight && data.recommendFlight.flights.map((flight) => (
        <Link to={`/flight/${flight.id}`} key={flight.id} className="recommended-flight-item-link">
        <div className="recommended-flight-item">
          <p>Departure: {flight.departure}</p>
          <p>Arrival: {flight.arrival}</p>
          <p>Bookable Seats: {flight.bookableSeats}</p>
          <p>Airline: {flight.airline}</p>
          <p>Price: ${flight.price}</p>
        </div>
      </Link>
      ))}
      </div>
    </div>
  );
}

export default RecommendedFlightList;
