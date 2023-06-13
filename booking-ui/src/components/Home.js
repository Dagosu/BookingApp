import React, { useContext } from 'react';
import { useQuery, gql } from '@apollo/client';
import { Navigate, Link } from 'react-router-dom';
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
  const { userId } = useContext(AuthContext);
  
  const { loading, error, data } = useQuery(RECOMMENDED_FLIGHTS_QUERY, {
    variables: {
      in: {
        userId: userId
      }
    },
    fetchPolicy: 'network-only' 
  });

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  return (
    <div className="home-container">
      <h1 className="home-title">Flight List</h1>
      <div className="button-group">
        <Link to="/purchased" className="button">Purchased Flights</Link>
        <Link to="/favorites" className="button">Favorite Flights</Link>
      </div>
      
      <h2 className="home-title">Recommended Flights</h2>
      <RecommendedFlightList flights={data.recommendFlight.flights} />
      
      <h2 className="home-title">All Flights</h2>
      <FlightList />
    </div>
  );
}

export default Home;
