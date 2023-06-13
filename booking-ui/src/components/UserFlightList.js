import React, { useContext } from 'react';
import { useQuery, gql } from '@apollo/client';
import { AuthContext } from '../contexts/AuthContext';
import { Link } from 'react-router-dom';
import '../style/UserFlightList.css'; 

const GET_PURCHASED_FLIGHTS = gql`
  query GetPurchasedFlights($in: GetPurchasedFlightsInput!) {
    getPurchasedFlights(in: $in) {
      flights {
        id
        departure
        arrival
        bookableSeats
      }
    }
  }
`;

const GET_FAVORITED_FLIGHTS = gql`
  query GetFavoritedFlights($in: GetFavoritedFlightsInput!) {
    getFavoritedFlights(in: $in) {
      flights {
        id
        departure
        arrival
        bookableSeats
      }
    }
  }
`;

function UserFlightList({ type }) {
  const { userId } = useContext(AuthContext);
  const query = type === 'purchased' ? GET_PURCHASED_FLIGHTS : GET_FAVORITED_FLIGHTS;

  const { loading, error, data } = useQuery(query, {
    variables: {
      in: {
        userId: userId
      }
    },
    fetchPolicy: 'network-only' // This ensures the query is sent to the server on each execution
  });

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  const flightList = data ? (type === 'purchased' ? data.getPurchasedFlights : data.getFavoritedFlights) : null;

  return (
    <div className="user-flight-list-container">
      <div className="user-flight-list">
        {flightList && flightList.flights.map((flight) => (
          <Link to={`/flight/${flight.id}`} key={flight.id} className="user-flight-item-link">
            <div className="user-flight-item">
              <p>Departure: {flight.departure}</p>
              <p>Arrival: {flight.arrival}</p>
              <p>Bookable Seats: {flight.bookableSeats}</p>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
}

export default UserFlightList;
