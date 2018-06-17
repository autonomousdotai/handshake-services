const configs = require('./configs');
const express = require('express');
const logger = require('morgan');
const bodyParser = require('body-parser');
const Twit = require('twit');
const utils = require('./utils');

const T = new Twit({
	consumer_key:         configs.consumer_key,
	consumer_secret:      configs.consumer_secret,
	access_token:         configs.access_token,
	access_token_secret:  configs.access_token_secret,
	app_only_auth:        true,
	timeout_ms:           60*1000,  // optional HTTP request timeout to apply to all requests.
	strictSSL:            true,     // optional - requires SSL certificates to be valid.
})

app = express();
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

const router = express.Router();
router.get('/', function(req, res) {
	res.json(utils.responseSuccess('', 'Twitter REST Service'));
});

router.get('/:username?', function(req, res) {
	if (req.params.username) {
		try {
			T.get('followers/list', { screen_name: 'ninja_org', count: '10', cursor: -1 },  function (err, data, response) {
				let result = false;
				if (data && data.users && data.users.length > 0) {
					console.log(data.users.length);
					for (let u of data.users) {
						if(u.screen_name == req.params.username){
							result = true
							console.log('found', u);
							break;
						}
					}
				}
				res.json(utils.responseSuccess(result));
			})	
		} catch (e) {
			console.log('Error: ', e.message);
			res.json(utils.responseError('Error occured, please try again'));
		}
	}
	else{
		res.json(utils.responseError('Invalid params'));
	}
});

app.use('/', router);

app.listen(configs.port, function() {
	console.log('Twitter service listening on port ' + configs.port + '!')
});