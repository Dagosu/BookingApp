import React, { useContext } from 'react';
import { gql, useLazyQuery } from '@apollo/client';
import { useNavigate, Navigate } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';
import styles from '../style/Login.module.css'; 

const CHECK_CREDENTIALS = gql`
  query CheckCredentials($in: CheckCredentialsInput!) {
    checkCredentials(in: $in) {
      authorized
      userId
    }
  }
`;

function Login() {
  const { isLoggedIn, setIsLoggedIn, setUserId } = useContext(AuthContext);
  const navigate = useNavigate();
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [loginError, setLoginError] = React.useState('');

  const [checkCredentials, { loading }] = useLazyQuery(CHECK_CREDENTIALS, {
    onError: (err) => {
      // Handle login errors
      if (err.message === 'rpc error: code = Unknown desc = Invalid credentials!') {
        setLoginError('Invalid username or password. Please try again.');
      } else {
        setLoginError(err.message);
      }
    },
    onCompleted: (data) => {
      if (data?.checkCredentials?.authorized) {
        // Redirect the user to the home page if they are authorized
        setIsLoggedIn(true);
        setUserId(data.checkCredentials.userId);
        navigate('/home');
      } else {
        // Set an error message if the user is not authorized
        setLoginError('Invalid username or password. Please try again.');
      }
    },
  });
  

  const handleLogin = (event) => {
    event.preventDefault();
    checkCredentials({ variables: { in: { email, password } } });
  };

  // Return early if the user is already logged in
  if (isLoggedIn) {
    return <Navigate to="/home" />;
  }

  if (loading) return 'Loading...';

  return (
    <div className={styles.container}>
      <form className={styles.form} onSubmit={handleLogin}>
        <input
          className={styles.input}
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
          required
        />

        <input
          className={styles.input}
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Password"
          required
        />

        {loginError && <div className={styles.error}>{loginError}</div>}

        <button className={styles.button} type="submit">Log in</button>
      </form>
    </div>
  );
}

export default Login;
