import React, { useContext } from 'react';
import { Route, Navigate } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';

function ProtectedRoute({ children }) {
    const { isLoggedIn } = useContext(AuthContext);
  
    if (!isLoggedIn) {
      return <Navigate to="/login" />;
    }
  
    return children;
}  
  

export default ProtectedRoute;
