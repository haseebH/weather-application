const axios = require('axios');

const getTemperatureData = async (req, res) => {
  const { location, timeframe } = req.params;

  const token = req.cookies.token;
  if (!token) {
    return res.status(401).json({ error: 'Unauthorized' });
  }

  try {
    const response = await axios.get(`${process.env.DATA_URL}/temperature/api/v1/temperature/${location}/${timeframe}`, {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });

    res.json(response.data);
  } catch (error) {
    res.status(500).json({ error: 'Error fetching temperature data' });
  }
};

module.exports = { getTemperatureData };
