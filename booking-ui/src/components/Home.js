import React, { useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';
import FlightList from './FlightList';
import '../style//Home.css'; // import your CSS

function Home() {
  const { isLoggedIn } = useContext(AuthContext);

  if (!isLoggedIn) {
    return <Navigate to="/login" />;
  }

  return (
    <div className="home-container">
      <h1 className="home-title">Flight List</h1>
      <FlightList /> {/* Include the FlightList component */}
    </div>
  );
}

export default Home;
