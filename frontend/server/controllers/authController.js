const jwt = require('jsonwebtoken');
const axios = require('axios')
require('dotenv').config();
const cookie = require('cookie');

const login = async (req, res) => {
  const { email, password } = req.body;

  try {
    const response = await axios.post(`${process.env.LOGIN_URL}/rbac/api/v1/login`, { email, password });

    if (response.status === 200 && response.data) {
      const user = response.data;

      res.json(user);
    } else {
      res.status(401).json({ error: 'Invalid credentials' });
    }
  } catch (error) {
    res.status(500).json({ error: 'Error during login' });
  }
};

module.exports = { login };
