const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');
const cookieParser = require('cookie-parser');
const temperatureRoutes = require('./routes/temperatureRoutes');
const authRoutes = require('./routes/authRoutes');
require('dotenv').config();

const app = express();
const port = process.env.PORT || 5000;

app.use(cors({
  origin: 'http://localhost:8082',
  credentials: true
}));
app.use(bodyParser.json());
app.use(cookieParser());

app.use('/api/temperature', temperatureRoutes);
app.use('/api/auth', authRoutes);

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
