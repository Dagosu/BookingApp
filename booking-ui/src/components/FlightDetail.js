import React, { useContext, useState, useEffect } from 'react';
import { useQuery, useMutation, gql } from '@apollo/client';
import { useParams } from 'react-router-dom';
import { AuthContext } from '../contexts/AuthContext';
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
        reviews {
          userName
          text
        }
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

const WRITE_REVIEW = gql`
  mutation ($in: WriteReviewInput!) {
    writeReview(in: $in) {
      flight {
        reviews {
          userName
          text
        }
      }
    }
  }
`;

const CHECK_FLIGHT_PURCHASE = gql`
  query ($in: CheckFlightPurchaseInput!) {
    checkFlightPurchase(in: $in) {
      flight {
        id
      }
    }
  }
`;


function FlightDetail() {
  const { id } = useParams();
  const { userId } = useContext(AuthContext);
  const [reviews, setReviews] = useState([]);
  const [canWriteReview, setCanWriteReview] = useState(false);
  
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

  const [reviewText, setReviewText] = useState('');
  const [writeReview] = useMutation(WRITE_REVIEW, {
    variables: {
      in: {
        flightId: id,
        userId,
        text: reviewText  
      },
    },
  });

  const { loading: checkFlightLoading, error: checkFlightError, data: checkFlightData } = useQuery(CHECK_FLIGHT_PURCHASE, {
    variables: {
      in: {
        flightId: id,
        userId,
      },
    },
  });
  

  const [notification, setNotification] = useState({ show: false, message: '' });

  const handlePurchase = () => {
    purchaseFlight().then(() => {
      setNotification({ show: true, message: 'Flight purchased successfully!' });
      setTimeout(() => setNotification({ show: false, message: '' }), 3000); // Hide after 3 seconds
    }).catch((error) => {
      if (error.message === 'rpc error: code = Unknown desc = You already purchased this flight!') {
        setNotification({ show: true, message: 'You already purchased this flight!' });
        setTimeout(() => setNotification({ show: false, message: '' }), 3000); // Hide after 3 seconds
      } else {
        // Handle other errors
        console.error(error);
      }
    });
  };

  const handleFavorite = () => {
    favoriteFlight().then(() => {
      setNotification({ show: true, message: 'Flight added to favorites!' });
      setTimeout(() => setNotification({ show: false, message: '' }), 3000); // Hide after 3 seconds
    });
  };

  const handleReviewSubmit = () => {
    writeReview().then((response) => {
      setNotification({ show: true, message: 'Review added successfully!' });
      setTimeout(() => setNotification({ show: false, message: '' }), 3000); // Hide after 3 seconds
      setReviewText(''); // clear the review text
      // Update reviews state
      const updatedReviews = response.data.writeReview.flight.reviews; // Get all reviews from response
      setReviews(updatedReviews); // Set state to updated reviews
    });
  };
  

  useEffect(() => {
    if (!loading && !error && data) {
      setReviews(data.getFlight.flight.reviews || []);
    }
  }, [loading, data, error]);

  useEffect(() => {
    if (!checkFlightLoading && !checkFlightError && checkFlightData) {
      setCanWriteReview(checkFlightData.checkFlightPurchase?.flight !== null);
    }
  }, [checkFlightLoading, checkFlightData, checkFlightError]);  

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
      <div className="review-form">
      <textarea
      className="review-input"
      value={reviewText}
      onChange={(e) => setReviewText(e.target.value)}
      placeholder="Write your review here..."
      disabled={!canWriteReview}
      ></textarea>
      <button
        className="review-submit-button"
        onClick={handleReviewSubmit}
        disabled={!canWriteReview}
      > Submit Review </button>
      </div>
      <div className="reviews-container">
        <h2>Reviews:</h2>
        {reviews.map((review, index) => (
          <div className="review" key={index}>
            <p><strong>{review.userName}:</strong> {review.text}</p>
          </div>
        ))}
      </div>
    </div>
  );  
}

export default FlightDetail;
