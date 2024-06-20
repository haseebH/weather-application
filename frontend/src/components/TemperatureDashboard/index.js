import React, { useEffect, useState } from 'react';
import { Line } from 'react-chartjs-2';
import axios from 'axios';
import Cookies from 'js-cookie';
import 'chart.js/auto';
import './index.css';

const TemperatureDashboard = () => {
  const [data, setData] = useState([]);
  const [timeframe, setTimeframe] = useState('month');
  const [labels, setLabels] = useState([]);
  const [temperatures, setTemperatures] = useState([]);

  let user = {name: Cookies.get('name'), location: Cookies.get('location')};

  useEffect(() => {
    fetchTemperatureData(user.location, timeframe);
  }, [timeframe]);

  const fetchTemperatureData = async (location, timeframe) => {
    try {
      const response = await axios.get(`${process.env.REACT_APP_API_URL}/api/temperature/${location}/${timeframe}`, {
        withCredentials: true
      });

      const data = response.data;

      if (data) {
        const labels = data.map(entry => entry.date);
        const temperatures = data.map(entry => entry.value);

        setLabels(labels);
        setTemperatures(temperatures);
      }
    } catch (error) {
      console.error('Error fetching temperature data:', error);
    }
  };

  const handleTimeframeChange = (event) => {
    setTimeframe(event.target.value);
  };

  const chartData = {
    labels,
    datasets: [
      {
        label: 'Air Temperature',
        data: temperatures,
        fill: false,
        backgroundColor: 'rgba(75,192,192,0.4)',
        borderColor: 'rgba(75,192,192,1)',
      },
    ],
  };

  return (
    <div className="dashboard-container">
      <h1>Temperature Dashboard</h1>
      <p>Welcome, {user.name} from {user.location}!</p>
      <div className="controls">
        <label htmlFor="timeframe">Select Timeframe:</label>
        <select id="timeframe" value={timeframe} onChange={handleTimeframeChange}>
          <option value="month">Past Month</option>
          <option value="year">Past Year</option>
          <option value="3years">Past 3 Years</option>
        </select>
      </div>
      <div className="chart-container">
        <Line data={chartData} />
      </div>
    </div>
  );
};

export default TemperatureDashboard;
