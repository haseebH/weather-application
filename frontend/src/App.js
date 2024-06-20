import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import TemperatureDashboard from './components/TemperatureDashboard';
import LoginPage from './components/Login';
import './App.css';

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/" element={<LoginPage />} />
          <Route path="/dashboard" element={<TemperatureDashboard />} />
        </Routes>
      </div>
    </Router >
  );
}

export default App;
