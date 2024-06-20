const express = require('express');
const router = express.Router();
const { getTemperatureData } = require('../controllers/temperatureController');

router.get('/:location/:timeframe', getTemperatureData);

module.exports = router;
