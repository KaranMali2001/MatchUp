import React, { useState } from 'react';
import LoginForm from './LoginForm';
import SignupForm from './SignupForm';

const HomePage = () => {
  const [showLoginForm, setShowLoginForm] = useState(false);
  const [showSignupForm, setShowSignupForm] = useState(false);

  const toggleLoginForm = () => {
    setShowLoginForm(!showLoginForm);
    setShowSignupForm(false);
  };

  const toggleSignupForm = () => {
    setShowSignupForm(!showSignupForm);
    setShowLoginForm(false);
  };

  return (
    <div className="home-page">
      <h1>Welcome to Our Platform</h1>
      {!showLoginForm &&!showSignupForm && (
        <>
          <button onClick={toggleLoginForm}>Login</button>
          <button onClick={toggleSignupForm}>Sign Up</button>
        </>
      )}
      {showLoginForm && (
        <>
          <LoginForm />
          <button onClick={toggleLoginForm}>Back to Home</button>
        </>
      )}
      {showSignupForm && (
        <>
          <SignupForm />
          <button onClick={toggleSignupForm}>Back to Home</button>
        </>
      )}
    </div>
  );
};

export default HomePage;
