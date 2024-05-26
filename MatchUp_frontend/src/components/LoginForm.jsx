import React, { useState } from 'react';
import axios from 'axios';

const LoginForm = ({ onLoginSuccess }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post('http://localhost:8080/login', { username, password }, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      });

      // Handle successful login here
      onLoginSuccess(response.data);
    } catch (error) {
      console.error('Login failed:', error.response?.data || error.message);
      // Handle login failure here
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Username:
        <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} required />
      </label>
      <label>
        Password:
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
      </label>
      <button type="submit">Log In</button>
    </form>
  );
};

export default LoginForm;
