import React, { useContext } from 'react';
import { useQuery, gql } from '@apollo/client';
import { useNavigate, Link } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';
import FlightList from './FlightList';
import RecommendedFlightList from './RecommendedFlightList';
import '../style/Home.css'; 

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

function Home() {
  const { userId, setIsLoggedIn, setUserId } = useContext(AuthContext);
  const navigate = useNavigate();
  
  const { loading, error, data } = useQuery(RECOMMENDED_FLIGHTS_QUERY, {
    variables: {
      in: {
        userId: userId
      }
    },
    fetchPolicy: 'network-only' 
  });

  const logout = () => {
    setIsLoggedIn(false);
    setUserId(null);
    localStorage.removeItem('isLoggedIn');
    localStorage.removeItem('userId');
    navigate('/login');
  };

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  return (
    <div className="home-container">
      <RecommendedFlightList flights={data.recommendFlight.flights} />
      <h1 className="home-title">All Flights</h1>
      <FlightList />
      <div className="button-group">
        <Link to="/purchased" className="button">Purchased Flights</Link>
        <Link to="/favorites" className="button">Favorite Flights</Link>
        <button onClick={logout} className="button logout-button">Log Out</button> 
      </div>
    </div>
  );
}

export default Home;
