import React from 'react';
import './App.css';
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';
import { BrowserRouter as Router, Route, Routes, Navigate, useLocation } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import Login from './components/Login';
import Home from './components/Home';
import FlightDetail from "./components/FlightDetail";
import ProtectedRoute from "./components/ProtectedRoute";
import UserFlightList from './components/UserFlightList';

// Create an Apollo client
const client = new ApolloClient({
  uri: 'http://localhost:8080/query', 
  cache: new InMemoryCache()
});


function UserFlightListWithKey({ type }) {
  const location = useLocation();
  return <UserFlightList key={location.pathname} type={type} />;
} 

function App() {
  return (
    <ApolloProvider client={client}>
      <AuthProvider>
        <Router>
          <Routes>
            <Route path="/" element={<Navigate to="/login" />} />
            <Route path="/login" element={<Login />} />
            <Route path="/home" element={<ProtectedRoute><Home /></ProtectedRoute>} />
            <Route path="/flight/:id" element={<ProtectedRoute><FlightDetail /></ProtectedRoute>} />
            <Route path="/purchased" element={<UserFlightListWithKey type="purchased" />} />
            <Route path="/favorites" element={<UserFlightListWithKey type="favorites" />} />
          </Routes>
        </Router>
      </AuthProvider>
    </ApolloProvider>
  );
}



export default App;
