import React, { useState } from 'react';
import axios from 'axios';

const SignupForm = () => {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    const user = {
      email,
      username,
      password,
      role,
    };

    let url = '';
    if (role === 'player') {
      url = 'http://localhost:8080/player';
    } else if (role === 'org') {
      url = 'http://localhost:8080/organizer';
    }

    try {
      const response = await axios.post(url, user, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      console.log(response.data);
      // Handle successful signup here, e.g., redirect or show a message
    } catch (error) {
      console.error('Signup failed:', error.response?.data || error.message);
      // Handle signup failure here
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Email:
        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
      </label>
      <label>
        Username:
        <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} required />
      </label>
      <label>
        Password:
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
      </label>
      <label>
        Role:
        <select value={role} onChange={(e) => setRole(e.target.value)} required>
          <option value="">Select Role</option>
          <option value="player">Player</option>
          <option value="org">Organizer</option>
        </select>
      </label>
      <button type="submit">Sign Up</button>
    </form>
  );
};

export default SignupForm;
