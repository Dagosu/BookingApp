import React, { useContext, useState } from 'react';
import { useQuery, useMutation, gql } from '@apollo/client';
import { useParams } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';
import { Navigate } from 'react-router-dom';
import '../style/FlightDetail.css'; 
import '../style/Notification.css'; 

const GET_FLIGHT = gql`
  query ($in: GetFlightInput!) {
    getFlight(in: $in) {
      flight {
        id
        departure
        departureTime {
          seconds
        }
        arrival
        arrivalTime {
          seconds
        }
        totalSeats
        bookableSeats
        airline
        price
      }
    }
  }
`;

const PURCHASE_FLIGHT = gql`
  mutation ($in: PurchaseFlightInput!) {
    purchaseFlight(in: $in) {
      purchasedFlight {
        id
        departure
        arrival
        totalSeats
        airline
        price
      }
    }
  }
`;

const FAVORITE_FLIGHT = gql`
  mutation ($in: FavoriteFlightInput!) {
    favoriteFlight(in: $in) {
      favoritedFlight {
        id
        departure
        arrival
      }
    }
  }
`;

function FlightDetail() {
  const { isLoggedIn } = useContext(AuthContext);
  const { id } = useParams();
  const { userId } = useContext(AuthContext);
  
  const { loading, error, data } = useQuery(GET_FLIGHT, {
    variables: {
      in: {
        flightId: id,
      },
    },
  });

  const [purchaseFlight] = useMutation(PURCHASE_FLIGHT, {
    variables: { in: 
        { 
            userId, 
            flightId: id 
        } 
    },
  });

  const [favoriteFlight] = useMutation(FAVORITE_FLIGHT, {
    variables: { 
        in: { 
            userId, 
            flightId: id 
        } 
    },
  });

  const [notification, setNotification] = useState({ show: false, message: '' });

  const handlePurchase = () => {
    purchaseFlight().then(() => {
      setNotification({ show: true, message: 'Flight purchased successfully!' });
      setTimeout(() => setNotification({ show: false, message: '' }), 3000); // Hide after 3 seconds
    });
  };

  const handleFavorite = () => {
    favoriteFlight().then(() => {
      setNotification({ show: true, message: 'Flight added to favorites!' });
      setTimeout(() => setNotification({ show: false, message: '' }), 3000); // Hide after 3 seconds
    });
  };

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  const flight = data.getFlight.flight;

  return (
    <div className="flight-detail-container">
      <div className="flight-detail">
        <p className="flight-detail-text">Departure: {flight.departure} at {new Date(flight.departureTime.seconds * 1000).toLocaleString()}</p>
        <p className="flight-detail-text">Arrival: {flight.arrival} at {new Date(flight.arrivalTime.seconds * 1000).toLocaleString()}</p>
        <p className="flight-detail-text">Total Seats: {flight.totalSeats}</p>
        <p className="flight-detail-text">Bookable Seats: {flight.bookableSeats}</p>
        <p className="flight-detail-text">Airline: {flight.airline}</p>
        <p className="flight-detail-text">Price: ${flight.price}</p>
      </div>
      <div className="flight-detail-actions">
        <button className="flight-action-button" onClick={handlePurchase}>Purchase</button>
        <button className="flight-action-button" onClick={handleFavorite}>Favorite</button>
      </div>
      {notification.show && <div className="notification">{notification.message}</div>}
    </div>
  );
}

export default FlightDetail;
